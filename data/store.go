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

func PutEvent(event map[string]interface{}) int {
	myDB := Database()
	defer myDB.Close()
	events := myDB.Use("events")

	docID, err := events.Insert(event)
	if err != nil {
		log.Fatal(err)
	}
	return docID
}

func GetEvent(docID int) map[string]interface{} {
	myDB := Database()
	defer myDB.Close()
	events := myDB.Use("events")

	readBack, err := events.Read(docID)
	if err != nil {
		log.Fatal(err)
	}
	return readBack
}

func GetEvents() map[int]struct{} {
	myDB := Database()
	defer myDB.Close()
	events := myDB.Use("events")

	query := "all"
	//map[string]interface{} { }

	queryResult := make(map[int]struct{})
	if err := db.EvalQuery(query, events, &queryResult); err != nil {
		log.Fatal(err)
	}
	return queryResult
}
