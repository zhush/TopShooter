package server

import (
	"servers/dbserver/common"
	"servers/dbserver/mysqldb"
)

type DBService struct {
}

func (dbsvr *DBService) Add(param *Params.AddParam, result *Params.AddResult) error {
	return mysqldb.DBMgr.Add(param, result)
}

func (dbsvr *DBService) Del(param *Params.DelParam, result *Params.DelResult) error {
	return mysqldb.DBMgr.Del(param, result)
}

func (dbsvr *DBService) Update(param *Params.UpdateParam, result *Params.UpdateResult) error {
	return mysqldb.DBMgr.Update(param, result)
}

func (dbsvr *DBService) Read(param *Params.ReadParam, result *Params.ReadResult) error {
	return mysqldb.DBMgr.Read(param, result)
}
