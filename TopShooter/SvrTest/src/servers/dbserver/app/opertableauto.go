//Warnning: This file is auto generate, don't modify it manual.
package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

var client *redis.Client
var sqldb *sql.DB

func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func toString(a interface{}) string {

	if v, p := a.(int); p {
		return strconv.Itoa(v)
	}

	if v, p := a.(float64); p {
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	if v, p := a.(float32); p {
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	}

	if v, p := a.(int16); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(uint); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(int32); p {
		return strconv.Itoa(int(v))
	}
	return "wrong"
}

func sqlValueStr(a interface{}) string {
	switch vtype := a.(type) {
	case string:
		return "'" + toString(a) + "'"
	default:
		return toString(a)
	}
}

func init() {
	client = redis.NewClient(redisOptions())
	client.FlushDB()
	var err error
	sqldb, err = sql.Open("mysql", sqlOptions())
	check(err)
}

func Read_t_account(key string) (result map[string]string) {
	redisKey := "t_account:" + key
	isExsit, _ := client.Exists(redisKey)
	if isExsit == int64(1) { //在redis中有数据,则直接返回redis的数据
		result["accid"] = client.HGet(redisKey, "accid")
		result["accountName"] = client.HGet(redisKey, "accountName")
		result["password"] = client.HGet(redisKey, "password")
		result["createtm"] = client.HGet(redisKey, "createtm")
		result["lastLoginTm"] = client.HGet(redisKey, "lastLoginTm")
		result["gameid"] = client.HGet(redisKey, "gameid")
		return
	}
	sql := fmt.Sprintf("select * from t_account where accid = %s", key)

	rows, err := sqldb.Query(sql)
	check(err)
	//返回所有列
	cols, err1 := rows.Columns()
	check(err1)
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			result[key] = string(v)
		}
		break
	}
	return
}

//添加表记录的方法，传入的是json字符串,如果插入成功，则返回true,和自增的id, 否则返回false, 0
func Add_t_account(contentJson string) (bool, int64) {
	var contentMaps map[string]interface{}
	err := json.Unmarshal(contentJson, &contentMaps)
	if err == nil {
		return false, 0
	}
	tableKey, isOk := contentMaps["accid"]
	if isOk == false {
		return false, 0
	}
	redisKey := "t_account:" + tableKey

	isExsit, _ := client.Exists(redisKey)
	if isExsit == true {
		return false, 0
	}

	//Write to Redis
	var fieldValue string

	fieldValue, isExsit = contentMaps["accid"]
	if isExsit == true {
		client.HSet(redisKey, "accid", fieldValue)
	} else {
		client.HSet(redisKey, "accid", "")
	}

	fieldValue, isExsit = contentMaps["accountName"]
	if isExsit == true {
		client.HSet(redisKey, "accountName", fieldValue)
	} else {
		client.HSet(redisKey, "accountName", "")
	}

	fieldValue, isExsit = contentMaps["password"]
	if isExsit == true {
		client.HSet(redisKey, "password", fieldValue)
	} else {
		client.HSet(redisKey, "password", "")
	}

	fieldValue, isExsit = contentMaps["createtm"]
	if isExsit == true {
		client.HSet(redisKey, "createtm", fieldValue)
	} else {
		client.HSet(redisKey, "createtm", "")
	}

	fieldValue, isExsit = contentMaps["lastLoginTm"]
	if isExsit == true {
		client.HSet(redisKey, "lastLoginTm", fieldValue)
	} else {
		client.HSet(redisKey, "lastLoginTm", "")
	}

	fieldValue, isExsit = contentMaps["gameid"]
	if isExsit == true {
		client.HSet(redisKey, "gameid", fieldValue)
	} else {
		client.HSet(redisKey, "gameid", "")
	}

	//Write to Mysql!
	keys := "("
	values := "("
	for k, v := range contentMaps {
		if keys != "(" {
			keys = keys + ","
		}
		keys = keys + k
		if values != "(" {
			values = values + ","
		}
		switch vtype := v.(type) {
		case string:
			values = values + "\"" + v + "\""
		default:
			values = values + toString(v)
		}
	}

	keys = keys + ")"
	values = values + ")"
	tableNames := "t_account"
	sql := fmt.Sprintf("insert into %s  %s  values  %s ", tableNames, keys, values)
	ret1, err1 := dbsvr.db.Exec(sql)
	if err1 != nil {
		log.Error("exec:%s failed!", sql)
		return false, 0
	}
	lastInserId := 0
	if LastInsertId, err2 := ret1.LastInsertId(); nil == err2 {
		lastInserId = LastInsertId
	}
	if RowsAffected, err3 := ret1.RowsAffected(); nil != err3 {
		return false, 0
	} else {
		if RowsAffected == 0 {
			return false, 0
		}
	}
	return true, int64(lastInserId)
}

//添加表记录的方法，传入的是json字符串,如果插入成功，则返回true,和自增的id, 否则返回false, 0
func Update_t_account(key string, contentJson string) (bool, int) {
	var contentMaps map[string]interface{}
	err := json.Unmarshal(contentJson, &contentMaps)
	if err == nil {
		return false, 0
	}
	redisKey := "t_account:" + key
	isExsit, _ := client.Exists(redisKey)
	if isExsit == false {
		return false, 0
	}
	//更新redis
	for k, v := range contentMaps {
		client.HSet(redisKey, k, v)
	}

	//Write to Mysql!
	keyvalues := ""
	for k, v := range contentMaps {
		if keyvalues != "" {
			keyvalues = keyvalues + ","
		}
		keyvalues = keyvalues + k + " = " + toString(v)
	}
	conditoins := fmt.Sprintf("accid = %s", sqlValueStr(key))
	tableNames := "t_account"

	sql := fmt.Sprintf("update from  %s  set(%s) where (%s)", tableNames, keyvalues, conditions)
	ret1, err1 := dbsvr.db.Exec(sql)
	if err1 != nil {
		return false, 0
	}
	lastInserId := 0
	if LastInsertId, err2 := ret1.LastInsertId(); nil == err2 {
		lastInserId = LastInsertId
	}
	if RowsAffected, err3 := ret1.RowsAffected(); nil != err3 {
		return false, 0
	} else {
		if RowsAffected == 0 {
			return false, 0
		}
	}
	return true, int64(lastInserId)
}

//删除表记录的方法，传入的是key 字符串,如果删除成功，则返回true,和影响的行数, 否则返回false, 0
func Remove_t_account(key string) (bool, int) {
	redisKey := "t_account:" + key
	isExsit, _ := client.Exists(redisKey)
	if isExsit == false {
		return false, 0
	}
	//删除redis
	client.Del(redisKey)

	//Delete from Mysql!
	conditoins := fmt.Sprintf("accid = %s", sqlValueStr(key))
	tableNames := "t_account"

	sql := fmt.Sprintf("delete from  %s  where (%s)", tableNames, conditions)
	ret1, err1 := dbsvr.db.Exec(sql)
	if err1 != nil {
		return false, 0
	}

	var err2 error
	var RowsAffected int
	if RowsAffected, err2 = ret1.RowsAffected(); nil != err2 {
		return false, 0
	} else {
		if RowsAffected == 0 {
			return false, 0
		}
	}
	return true, RowsAffected
}

func Read_t_role(key string) (result map[string]string) {
	redisKey := "t_role:" + key
	isExsit, _ := client.Exists(redisKey)
	if isExsit == int64(1) { //在redis中有数据,则直接返回redis的数据
		result["roleid"] = client.HGet(redisKey, "roleid")
		result["accid"] = client.HGet(redisKey, "accid")
		result["nickName"] = client.HGet(redisKey, "nickName")
		result["sex"] = client.HGet(redisKey, "sex")
		result["templateId"] = client.HGet(redisKey, "templateId")
		result["createtm"] = client.HGet(redisKey, "createtm")
		result["lastsceneid"] = client.HGet(redisKey, "lastsceneid")
		result["lastposX"] = client.HGet(redisKey, "lastposX")
		result["lastposY"] = client.HGet(redisKey, "lastposY")
		result["handWeapon"] = client.HGet(redisKey, "handWeapon")
		result["bulletCount"] = client.HGet(redisKey, "bulletCount")
		result["weaponList"] = client.HGet(redisKey, "weaponList")
		result["level"] = client.HGet(redisKey, "level")
		result["gold"] = client.HGet(redisKey, "gold")
		return
	}
	sql := fmt.Sprintf("select * from t_role where roleid = %s", key)

	rows, err := sqldb.Query(sql)
	check(err)
	//返回所有列
	cols, err1 := rows.Columns()
	check(err1)
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			result[key] = string(v)
		}
		break
	}
	return
}

//添加表记录的方法，传入的是json字符串,如果插入成功，则返回true,和自增的id, 否则返回false, 0
func Add_t_role(contentJson string) (bool, int64) {
	var contentMaps map[string]interface{}
	err := json.Unmarshal(contentJson, &contentMaps)
	if err == nil {
		return false, 0
	}
	tableKey, isOk := contentMaps["roleid"]
	if isOk == false {
		return false, 0
	}
	redisKey := "t_role:" + tableKey

	isExsit, _ := client.Exists(redisKey)
	if isExsit == true {
		return false, 0
	}

	//Write to Redis
	var fieldValue string

	fieldValue, isExsit = contentMaps["roleid"]
	if isExsit == true {
		client.HSet(redisKey, "roleid", fieldValue)
	} else {
		client.HSet(redisKey, "roleid", "")
	}

	fieldValue, isExsit = contentMaps["accid"]
	if isExsit == true {
		client.HSet(redisKey, "accid", fieldValue)
	} else {
		client.HSet(redisKey, "accid", "")
	}

	fieldValue, isExsit = contentMaps["nickName"]
	if isExsit == true {
		client.HSet(redisKey, "nickName", fieldValue)
	} else {
		client.HSet(redisKey, "nickName", "")
	}

	fieldValue, isExsit = contentMaps["sex"]
	if isExsit == true {
		client.HSet(redisKey, "sex", fieldValue)
	} else {
		client.HSet(redisKey, "sex", "")
	}

	fieldValue, isExsit = contentMaps["templateId"]
	if isExsit == true {
		client.HSet(redisKey, "templateId", fieldValue)
	} else {
		client.HSet(redisKey, "templateId", "")
	}

	fieldValue, isExsit = contentMaps["createtm"]
	if isExsit == true {
		client.HSet(redisKey, "createtm", fieldValue)
	} else {
		client.HSet(redisKey, "createtm", "")
	}

	fieldValue, isExsit = contentMaps["lastsceneid"]
	if isExsit == true {
		client.HSet(redisKey, "lastsceneid", fieldValue)
	} else {
		client.HSet(redisKey, "lastsceneid", "")
	}

	fieldValue, isExsit = contentMaps["lastposX"]
	if isExsit == true {
		client.HSet(redisKey, "lastposX", fieldValue)
	} else {
		client.HSet(redisKey, "lastposX", "")
	}

	fieldValue, isExsit = contentMaps["lastposY"]
	if isExsit == true {
		client.HSet(redisKey, "lastposY", fieldValue)
	} else {
		client.HSet(redisKey, "lastposY", "")
	}

	fieldValue, isExsit = contentMaps["handWeapon"]
	if isExsit == true {
		client.HSet(redisKey, "handWeapon", fieldValue)
	} else {
		client.HSet(redisKey, "handWeapon", "")
	}

	fieldValue, isExsit = contentMaps["bulletCount"]
	if isExsit == true {
		client.HSet(redisKey, "bulletCount", fieldValue)
	} else {
		client.HSet(redisKey, "bulletCount", "")
	}

	fieldValue, isExsit = contentMaps["weaponList"]
	if isExsit == true {
		client.HSet(redisKey, "weaponList", fieldValue)
	} else {
		client.HSet(redisKey, "weaponList", "")
	}

	fieldValue, isExsit = contentMaps["level"]
	if isExsit == true {
		client.HSet(redisKey, "level", fieldValue)
	} else {
		client.HSet(redisKey, "level", "")
	}

	fieldValue, isExsit = contentMaps["gold"]
	if isExsit == true {
		client.HSet(redisKey, "gold", fieldValue)
	} else {
		client.HSet(redisKey, "gold", "")
	}

	//Write to Mysql!
	keys := "("
	values := "("
	for k, v := range contentMaps {
		if keys != "(" {
			keys = keys + ","
		}
		keys = keys + k
		if values != "(" {
			values = values + ","
		}
		switch vtype := v.(type) {
		case string:
			values = values + "\"" + v + "\""
		default:
			values = values + toString(v)
		}
	}

	keys = keys + ")"
	values = values + ")"
	tableNames := "t_role"
	sql := fmt.Sprintf("insert into %s  %s  values  %s ", tableNames, keys, values)
	ret1, err1 := dbsvr.db.Exec(sql)
	if err1 != nil {
		log.Error("exec:%s failed!", sql)
		return false, 0
	}
	lastInserId := 0
	if LastInsertId, err2 := ret1.LastInsertId(); nil == err2 {
		lastInserId = LastInsertId
	}
	if RowsAffected, err3 := ret1.RowsAffected(); nil != err3 {
		return false, 0
	} else {
		if RowsAffected == 0 {
			return false, 0
		}
	}
	return true, int64(lastInserId)
}

//添加表记录的方法，传入的是json字符串,如果插入成功，则返回true,和自增的id, 否则返回false, 0
func Update_t_role(key string, contentJson string) (bool, int) {
	var contentMaps map[string]interface{}
	err := json.Unmarshal(contentJson, &contentMaps)
	if err == nil {
		return false, 0
	}
	redisKey := "t_role:" + key
	isExsit, _ := client.Exists(redisKey)
	if isExsit == false {
		return false, 0
	}
	//更新redis
	for k, v := range contentMaps {
		client.HSet(redisKey, k, v)
	}

	//Write to Mysql!
	keyvalues := ""
	for k, v := range contentMaps {
		if keyvalues != "" {
			keyvalues = keyvalues + ","
		}
		keyvalues = keyvalues + k + " = " + toString(v)
	}
	conditoins := fmt.Sprintf("roleid = %s", sqlValueStr(key))
	tableNames := "t_role"

	sql := fmt.Sprintf("update from  %s  set(%s) where (%s)", tableNames, keyvalues, conditions)
	ret1, err1 := dbsvr.db.Exec(sql)
	if err1 != nil {
		return false, 0
	}
	lastInserId := 0
	if LastInsertId, err2 := ret1.LastInsertId(); nil == err2 {
		lastInserId = LastInsertId
	}
	if RowsAffected, err3 := ret1.RowsAffected(); nil != err3 {
		return false, 0
	} else {
		if RowsAffected == 0 {
			return false, 0
		}
	}
	return true, int64(lastInserId)
}

//删除表记录的方法，传入的是key 字符串,如果删除成功，则返回true,和影响的行数, 否则返回false, 0
func Remove_t_role(key string) (bool, int) {
	redisKey := "t_role:" + key
	isExsit, _ := client.Exists(redisKey)
	if isExsit == false {
		return false, 0
	}
	//删除redis
	client.Del(redisKey)

	//Delete from Mysql!
	conditoins := fmt.Sprintf("roleid = %s", sqlValueStr(key))
	tableNames := "t_role"

	sql := fmt.Sprintf("delete from  %s  where (%s)", tableNames, conditions)
	ret1, err1 := dbsvr.db.Exec(sql)
	if err1 != nil {
		return false, 0
	}

	var err2 error
	var RowsAffected int
	if RowsAffected, err2 = ret1.RowsAffected(); nil != err2 {
		return false, 0
	} else {
		if RowsAffected == 0 {
			return false, 0
		}
	}
	return true, RowsAffected
}
