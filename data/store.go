package data

import (
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/docker/docker/pkg/homedir"
	"log"
	"path/filepath"
)

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

func PutInvocation(invocation map[string]interface{}) int {
	return PutRecordTable(invocation, "invocations")
}

func PutEvent(event map[string]interface{}) int {
	return PutRecordTable(event, "events")
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

func GetEvent(docId int) map[string]interface{} {
	return GetIdTable(docId, "events")
}

func GetInvocation(docId int) map[string]interface{} {
	return GetIdTable(docId, "invocations")
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
