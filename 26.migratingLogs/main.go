package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"migratingLogsModule/models"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// constants
// TO DO: need to work on the go-rouitnes part
var waitGroup = &sync.WaitGroup{}
var mux =  &sync.Mutex{}
var logFileName = time.Now().String() + "logFile.json"
var counter int64 = 1

type logObj struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Seach_type   string             `json:"seach_type,omitempty" bson:"seach_type,omitempty"`
	Created_at   primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Linkedin_url string             `json:"linkedin_url,omitempty" bson:"linkedin_url,omitempty"`
	Logs         []interface{}      `json:"logs,omitempty" bson:"logs,omitempty"`
}

// function to handle error logs
func handleErrors(functionName string, err error) {
	if err != nil {
		log.Fatal("Function :: ", functionName, "ERROR ", err)
	}
}

func main() {
	startTime := time.Now()
	fmt.Println("Welcome to migratingLogs module")
	fmt.Println("The total logs documents are ")

	// fetching the total count of documents
	totalLogCount := getLogCount()
	fmt.Println(totalLogCount)

	// creating a log file
	logFile, fileErr := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	handleErrors("fileErr:LogFileOpen", fileErr)
	defer logFile.Close()

	// appendging th "[" at the start of file
	logFile.WriteString("[")

	// appending the logs to file
	writeLogsToFile(logFile)

	// appendging th "]" at the end of file
	logFile.WriteString("]")

	for{
		if(counter >= totalLogCount){
			break
		}
	}

	waitGroup.Wait()
	endTime := time.Now()
	fmt.Printf("The total time taken by program to move %v logs from DB to file is %v \n", totalLogCount, endTime.Sub(startTime))
	fmt.Printf("Counter ::::::::: %d", counter)
}

func getLogCount() int64 {
	opts := options.Count().SetHint("_id_")
	logCount, countErr := models.Collections.CountDocuments(context.Background(), bson.D{}, opts)
	handleErrors("CountErr:DBCountDoucments", countErr)

	return logCount
}

func writeLogsToFile(logFile *os.File) {
	// opts := options.Find().SetLimit(78)
	cursor, findErr := models.Collections.Find(context.Background(), bson.M{})
	handleErrors("FindERR:FatchAllRecords", findErr)
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		waitGroup.Add(1)
		go routineFuncToAppendLogs(cursor, logFile)
	}
}

// called by writeLogsToFile using go routine
func routineFuncToAppendLogs(cursor *mongo.Cursor, logFile *os.File) {
	defer waitGroup.Done()
	var logData logObj
	decodeErr := cursor.Decode(&logData)
	handleErrors("DecodeErr", decodeErr)

	logDataStr, err := json.Marshal(logData)
	handleErrors("MarshalErr", err)

	mux.Lock()
	// dataChannel  <- string(logDataStr)
	counter += 1
	logFile.WriteString(string(logDataStr) + ",")
	mux.Unlock()
}
