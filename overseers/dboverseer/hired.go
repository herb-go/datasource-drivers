package dboverseer

import (
	"github.com/herb-go/datasource/sql/db"
	"github.com/herb-go/worker"
)

var dbworker = db.New()
var Team = worker.GetWorkerTeam(&dbworker)

func GetDBByID(id string) *db.PlainDB {
	w := worker.FindWorker(id)
	if w == nil {
		return nil
	}
	c, ok := w.Interface.(**db.PlainDB)
	if ok == false || c == nil {
		return nil
	}
	return *c
}
