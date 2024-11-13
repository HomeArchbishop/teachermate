package model

import (
	"github.com/syndtr/goleveldb/leveldb/util"
)

func AddSubscription(lessonId, studentId, connId string) error {
	safeDB.mux.Lock()
	defer safeDB.mux.Unlock()

	key := keySubscription(lessonId, studentId)
	err := safeDB.db.Put([]byte(key), []byte(connId), nil)
	if err != nil {
		return err
	}

	return nil
}

func HasSubscription(lessonId, studentId string) (bool, error) {
	safeDB.mux.Lock()
	defer safeDB.mux.Unlock()

	key := keySubscription(lessonId, studentId)
	has, err := safeDB.db.Has([]byte(key), nil)
	if err != nil {
		return false, err
	}

	return has, nil
}

func RemoveSubscription(lessonId, studentId string) error {
	safeDB.mux.Lock()
	defer safeDB.mux.Unlock()

	key := keySubscription(lessonId, studentId)
	if err := safeDB.db.Delete([]byte(key), nil); err != nil {
		if err.Error() == "leveldb: not found" {
			return nil
		}
	}

	return nil
}

func RemoveAllSubscription() error {
	safeDB.mux.Lock()
	defer safeDB.mux.Unlock()

	prefix := prefixSubscription4All()
	iter := safeDB.db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)

	for iter.Next() {
		safeDB.db.Delete(iter.Key(), nil)
	}

	return nil
}

func GetSubscription4Lesson(lessonId string) ([]string, error) {
	safeDB.mux.Lock()
	defer safeDB.mux.Unlock()

	prefix := prefixSubscription4Lesson(lessonId)
	iter := safeDB.db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)

	var connIdList []string
	for iter.Next() {
		connIdList = append(connIdList, string(iter.Value()))
	}

	return connIdList, nil
}
