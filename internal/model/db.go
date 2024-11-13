package model

import (
	"log"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

type safeDBType struct {
	db  *leveldb.DB
	mux sync.Mutex
}

var safeDB *safeDBType

func InitDB() {
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		log.Fatal(err)
	}

	safeDB = &safeDBType{db: db}
}

func CloseDB() {
	safeDB.db.Close()
}

func keySubscription(lessonId, studentId string) string {
	return "subscription:" + lessonId + ":" + studentId
}

func prefixSubscription4Lesson(lessonId string) string {
	return "subscription:" + lessonId + ":"
}
