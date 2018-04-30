package data

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/docker/docker/pkg/homedir"
	//"github.com/fatih/structs"
)

type EventId int
type InvocationId int
type StatId int
type ResponseId int

func Database() *db.DB {
	dbPath := filepath.Join(homedir.Get(), ".eva", "data")

	// Open database (creates as necessary /w directories)
	myDB, err := db.OpenDB(dbPath)
	if err != nil {
		panic(err)
	}

	/* Create all needed collections.
	   ignore errors as we want to idompotently
	   create these if they do not exist. */
	myDB.Create("events")
	myDB.Create("responses")
	myDB.Create("invocations")
	myDB.Create("stats")

	return myDB
}

func PutRecordTable(record map[string]interface{}, tableName string) int {
	myDB := Database()
	defer myDB.Close()
	invocations := myDB.Use(tableName)

	docID, err := invocations.Insert(record)
	if err != nil {
		log.Fatal(err)
	}
	return docID
}

func PutInvocation(invocation InvocationLog) InvocationId {
	//invocationMap := structs.Map(invocation)
	fmt.Println("PutInvStatID: ", invocation.StatId)
	fmt.Println("PutInvResponseID: ", invocation.ResponseId)
	invocationMap := map[string]interface{}{
		"StatId":            string(invocation.StatId),
		"ResponseId":        string(invocation.ResponseId),
		"InvocationRequest": invocation.InvocationRequest,
	}
	id := InvocationId(PutRecordTable(invocationMap, "invocations"))
	fmt.Println("InvocationId: ", id)
	return id
}

func PutEvent(event map[string]interface{}) EventId {
	return EventId(PutRecordTable(event, "events"))
}

func PutStat(event map[string]interface{}) StatId {
	return StatId(PutRecordTable(event, "stats"))
}

func PutResponse(event map[string]interface{}) ResponseId {
	return ResponseId(PutRecordTable(event, "responses"))
}

func GetIdTable(docID int, tableName string) map[string]interface{} {
	myDB := Database()
	defer myDB.Close()
	events := myDB.Use(tableName)

	readBack, err := events.Read(docID)
	if err != nil {
		log.Fatal(err)
	}
	return readBack
}

func GetEvent(docId EventId) map[string]interface{} {
	return GetIdTable(int(docId), "events")
}

func GetEventJson(docId EventId) ([]byte, error) {
	result := GetIdTable(int(docId), "events")
	return json.MarshalIndent(result, "", " ")
}

func GetInvocation(docId InvocationId) *InvocationLog {
	invMap := GetIdTable(int(docId), "invocations")

	statId, _ := strconv.Atoi(invMap["StatId"].(string))
	responseId, _ := strconv.Atoi(invMap["ResponseId"].(string))
	return &InvocationLog{
		StatId:     StatId(statId),         //invMap["StatId"].(int)),
		ResponseId: ResponseId(responseId), //invMap["ResponseId"].(int)),
	}
}

func GetStat(docId StatId) map[string]interface{} {
	return GetIdTable(int(docId), "stats")
}

func GetResponseJson(docId ResponseId) ([]byte, error) {
	result := GetIdTable(int(docId), "responses")
	return json.MarshalIndent(result, "", " ")
}

func GetAllTable(tableName string) map[int]struct{} {
	myDB := Database()
	defer myDB.Close()
	table := myDB.Use(tableName)

	query := "all"
	queryResult := make(map[int]struct{})
	if err := db.EvalQuery(query, table, &queryResult); err != nil {
		log.Fatal(err)
	}
	return queryResult
}

func GetInvocations() map[int]struct{} {
	return GetAllTable("invocations")
}

func GetEvents() map[int]struct{} {
	return GetAllTable("events")
}

func GetStats() map[int]struct{} {
	return GetAllTable("stats")
}
