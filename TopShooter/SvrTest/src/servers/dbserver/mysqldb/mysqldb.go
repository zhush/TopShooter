package mysqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"libs/log"
	"servers/dbserver/common"
	"servers/dbserver/config"

	_ "github.com/go-sql-driver/mysql"
)

var DBMgr *MysqlDB

func init() {
	DBMgr = new(MysqlDB)
	DBMgr.init()
}

type MysqlDB struct {
	db *sql.DB
}

func (dbsvr *MysqlDB) init() {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		config.DBConf["DatabaseUser"].(string),
		config.DBConf["DatabasePwd"].(string),
		config.DBConf["DatabaseAddr"].(string),
		config.DBConf["Database"].(string))
	dbsvr.db, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Connect Database(%s) failed! error:%s", config.DBConf["Database"].(string), err.Error())
	}
}

//判断是否是数字串类型
func (dbsvr *MysqlDB) isString(val interface{}) bool {
	switch val.(type) {
	case string:
		return true
	default:
		return false
	}
}

func (dbsvr *MysqlDB) convertConditonStr(conditons []Params.KeyValue) string {
	conditionsStr := ""
	for i := 0; i < len(conditons); i++ {
		condition := conditons[i]
		if i > 0 {
			if dbsvr.isString(condition.Value) {
				conditionsStr = fmt.Sprintf("%s and %s = '%s'", conditionsStr, condition.Key, condition.Value)
			} else {
				conditionsStr = fmt.Sprintf("%s and  %s = %v", conditionsStr, condition.Key, condition.Value)
			}
		} else {
			if dbsvr.isString(condition.Value) {
				conditionsStr = fmt.Sprintf(" %s = '%s'", condition.Key, condition.Value)
			} else {
				conditionsStr = fmt.Sprintf(" %s = %v", condition.Key, condition.Value)
			}
		}
	}
	return conditionsStr
}

////////////////////////////////////////////////////////////////////////////////
func (dbsvr *MysqlDB) Add(param *Params.AddParam, result *Params.AddResult) error {
	if len(param.Values) == 0 {
		(*result).Result = -1
		(*result).ErrorMsg = "Invalid Add Params, param.Values length == 0"
		log.Error("Invalid Add Params, param.Values length == 0")
		return errors.New("Invalid Add Params, param.Values length == 0")
	}
	keys := ""
	values := ""
	for i := 0; i < len(param.Values); i++ {
		value := param.Values[i]
		if i > 0 {
			keys = fmt.Sprintf("%s,%s", keys, value.Key)
			if dbsvr.isString(value.Value) {
				values = fmt.Sprintf("%s,'%v'", values, value.Value)
			} else {
				values = fmt.Sprintf("%s, %v", values, value.Value)
			}
		} else {
			keys = value.Key
			values = fmt.Sprintf("%v", value.Value)
		}
	}
	sql := fmt.Sprintf("insert into %s (%s) values(%s)", param.TableName, keys, values)
	ret, err := dbsvr.db.Exec(sql)
	if err != nil {
		log.Error("exec:%s failed!", sql)
		(*result).Result = -1

		(*result).ErrorMsg = err.Error()
		return err
	}
	if LastInsertId, err2 := ret.LastInsertId(); nil == err2 {
		(*result).AutoIncrementId = LastInsertId
	}
	if RowsAffected, err3 := ret.RowsAffected(); nil != err3 {
		(*result).Result = -1
		(*result).ErrorMsg = err3.Error()
		return err3
	} else {
		if RowsAffected == 0 {
			(*result).Result = -1
			(*result).ErrorMsg = "No Add, Affected 0 Rows"
			return err3
		}
	}

	(*result).Result = 0
	(*result).ErrorMsg = ""
	return nil

}

//执行删除语句
func (dbsvr *MysqlDB) Del(param *Params.DelParam, result *Params.DelResult) error {

	if len(param.Conditions) == 0 {
		(*result).Result = -1
		(*result).ErrorMsg = "Invalid Del Params, param.Conditions length == 0"
		log.Error("Invalid Del Params, param.Conditions length == 0")
		return errors.New("Invalid Del Params, param.Conditions length == 0")
	}
	conditions := dbsvr.convertConditonStr(param.Conditions)
	sql := fmt.Sprintf("delete from  %s  where %s", param.TableName, conditions)
	ret, err := dbsvr.db.Exec(sql)
	if err != nil {
		log.Error("exec:%s failed!", sql)
		(*result).Result = -1

		(*result).ErrorMsg = err.Error()
		return err
	}

	if RowsAffected, err3 := ret.RowsAffected(); nil != err3 {
		(*result).Result = -1
		(*result).ErrorMsg = err3.Error()
		return err3
	} else {
		if RowsAffected == 0 {
			(*result).Result = -1
			(*result).ErrorMsg = "Del, RowsAffect 0 rows"
		}
	}
	(*result).Result = 0
	(*result).ErrorMsg = ""
	return nil
}

//更新语句
func (dbsvr *MysqlDB) Update(param *Params.UpdateParam, result *Params.UpdateResult) error {

	if len(param.Values) == 0 {
		(*result).Result = -1
		errmsg := "Invalid Update Params, param.Values length == 0"
		(*result).ErrorMsg = errmsg
		log.Error(errmsg)
		return errors.New(errmsg)
	}

	if len(param.Conditions) == 0 {
		(*result).Result = -1
		errmsg := "Invalid Update Params, param.Conditions length == 0"
		(*result).ErrorMsg = errmsg
		log.Error(errmsg)
		return errors.New(errmsg)
	}

	updateSets := ""
	for i := 0; i < len(param.Values); i++ {
		setValue := param.Values[i]
		if i > 0 {
			if dbsvr.isString(setValue.Value) {
				updateSets = fmt.Sprintf("%s, %s = '%s'", updateSets, setValue.Key, setValue.Value)
			} else {
				updateSets = fmt.Sprintf("%s,  %s = %v", updateSets, setValue.Key, setValue.Value)
			}
		} else {
			if dbsvr.isString(setValue.Value) {
				updateSets = fmt.Sprintf(" %s = '%s'", setValue.Key, setValue.Value)
			} else {
				updateSets = fmt.Sprintf(" %s = %v", setValue.Key, setValue.Value)
			}
		}
	}

	conditions := dbsvr.convertConditonStr(param.Conditions)

	sql := fmt.Sprintf("update from  %s  set %s where %s", param.TableName, updateSets, conditions)
	ret, err := dbsvr.db.Exec(sql)
	if err != nil {
		log.Error("exec:%s failed!", sql)
		(*result).Result = -1
		(*result).ErrorMsg = err.Error()
		return err
	}

	rowAffected := 0
	if RowsAffected, err3 := ret.RowsAffected(); nil != err3 {
		(*result).Result = -1
		(*result).ErrorMsg = err3.Error()
		return err3
	} else {
		rowAffected = int(RowsAffected)
	}

	if rowAffected == 0 {
		(*result).Result = -1
		errmsg := "RowsAffected == 0"
		(*result).ErrorMsg = errmsg
		log.Error(errmsg)
		return errors.New(errmsg)
	}

	(*result).Result = 0
	(*result).ErrorMsg = ""
	return nil
}

//查询语句
func (dbsvr *MysqlDB) Read(param *Params.ReadParam, result *Params.ReadResult) error {
	if len(param.Conditions) == 0 {
		(*result).Result = -1
		errmsg := "Invalid Read Params, param.Conditions length == 0"
		log.Error(errmsg)
		return errors.New(errmsg)
	}

	conditions := ""
	for i := 0; i < len(param.Conditions); i++ {
		condition := param.Conditions[i]
		if i > 0 {
			if dbsvr.isString(condition.Value) {
				conditions = fmt.Sprintf(" and %s = '%s'", condition.Key, condition.Value)
			} else {
				conditions = fmt.Sprintf(" and  %s = %v", condition.Key, condition.Value)
			}
		} else {
			if dbsvr.isString(condition.Value) {
				conditions = fmt.Sprintf(" %s = '%s'", condition.Key, condition.Value)
			} else {
				conditions = fmt.Sprintf(" %s = %v", condition.Key, condition.Value)
			}
		}
	}

	selectKeys := "*"
	if len(param.Keys) > 0 {
		selectKeys = " "
		for i := 0; i < len(param.Keys); i++ {
			key := param.Keys[i]
			if i > 0 {
				selectKeys = fmt.Sprintf(", %s ", key)
			} else {
				selectKeys = key
			}
		}
	}

	sql := fmt.Sprintf("select (%s) from %s where %s", selectKeys, param.TableName, conditions)

	rows, err := dbsvr.db.Query(sql)
	if err != nil {
		errmsg := fmt.Sprintf("Exec (%s) failed! error:", sql, err.Error())
		log.Error(errmsg)
		(*result).Result = -1
		return errors.New(errmsg)
	}

	//返回所有列
	cols, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	resultTmp := make(map[int]map[string]string)
	for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		resultTmp[i] = row
		i++
	}

	log.Debug("sql:%s", sql)
	log.Debug("ret:%v", resultTmp)
	(*result).Result = 0
	(*result).Rows = len(resultTmp)
	for i := 0; i < len(resultTmp); i++ {
		rowResult := Params.RowResult{}
		rowResult.Values = make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range resultTmp[i] {
			rowResult.Values[k] = string(v)
		}
		log.Debug("rowResult:%v", rowResult)
		(*result).RowValues = append((*result).RowValues, rowResult)
	}

	log.Debug("result:%v", (*result))

	return nil
}
