//Warnning: This file is auto generate, don't modify it manual.
package app
import (
    "fmt"
	"github.com/go-redis/redis"
	"database/sql"
	"encoding/json"
	"libs/log"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

var client *redis.Client
var sqldb  *sql.DB

func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}



func toString(a interface{}) string{
	
	if  v,p:=a.(int);p{
	 	return strconv.Itoa(v)
	}
	
	if v,p:=a.(float64);  p{
	 	return strconv.FormatFloat(v,'f', -1, 64)
	}
	
	if v,p:=a.(float32); p {
		return strconv.FormatFloat(float64(v),'f', -1, 32)
	}
	
	if v,p:=a.(int16); p { 
		return strconv.Itoa(int(v))
	}
	if v,p:=a.(uint); p { 
		return strconv.Itoa(int(v))
	}
	if v,p:=a.(int32); p { 
		return strconv.Itoa(int(v))
	}
	return "wrong"
}


func sqlValueStr(a interface{}) string{
	switch val:=a.(type){
		case string:
			return "'" + toString(val) + "'"  
		default:
			return toString(val)
	}
}

func init(){
	client = redis.NewClient(redisOptions())
	client.FlushDB()
	var err error
	sqldb, err = sql.Open("mysql", sqlOptions())
	check(err)
	
	registerAllOperateTableHandlers()
	
}


func Read_t_account(key string)(result map[string]string){
    redisKey:= "t_account:"+key
    isExsit, _ := client.Exists(redisKey).Result()
    if isExsit == int64(1) { //在redis中有数据,则直接返回redis的数据
        result["accid"]=client.HGet(redisKey, "accid").Val()
        result["accountName"]=client.HGet(redisKey, "accountName").Val()
        result["password"]=client.HGet(redisKey, "password").Val()
        result["createtm"]=client.HGet(redisKey, "createtm").Val()
        result["lastLoginTm"]=client.HGet(redisKey, "lastLoginTm").Val()
        result["gameid"]=client.HGet(redisKey, "gameid").Val()
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
func Add_t_account(contentJson string) (bool, int64){
	var contentMaps map[string]interface{}
	err := json.Unmarshal([]byte(contentJson), &contentMaps)
	if err == nil {
		return false, 0
	}
    tableKey, isOk := contentMaps["accid"]
	if isOk == false{
		return false, 0
	}
    redisKey := "t_account:"+tableKey.(string)

	isExsit, _ := client.Exists(redisKey).Result()
	if isExsit == int64(1) {
		return false, 0
	}
	
	//Write to Redis
	var fieldValue interface{}
	var keyIsExsit bool
	
    fieldValue, keyIsExsit = contentMaps["accid"]
    if keyIsExsit == true {
        client.HSet(redisKey, "accid", fieldValue)
    }else{
        client.HSet(redisKey, "accid", "")
    }

    fieldValue, keyIsExsit = contentMaps["accountName"]
    if keyIsExsit == true {
        client.HSet(redisKey, "accountName", fieldValue)
    }else{
        client.HSet(redisKey, "accountName", "")
    }

    fieldValue, keyIsExsit = contentMaps["password"]
    if keyIsExsit == true {
        client.HSet(redisKey, "password", fieldValue)
    }else{
        client.HSet(redisKey, "password", "")
    }

    fieldValue, keyIsExsit = contentMaps["createtm"]
    if keyIsExsit == true {
        client.HSet(redisKey, "createtm", fieldValue)
    }else{
        client.HSet(redisKey, "createtm", "")
    }

    fieldValue, keyIsExsit = contentMaps["lastLoginTm"]
    if keyIsExsit == true {
        client.HSet(redisKey, "lastLoginTm", fieldValue)
    }else{
        client.HSet(redisKey, "lastLoginTm", "")
    }

    fieldValue, keyIsExsit = contentMaps["gameid"]
    if keyIsExsit == true {
        client.HSet(redisKey, "gameid", fieldValue)
    }else{
        client.HSet(redisKey, "gameid", "")
    }

	//Write to Mysql!
	keys := "("
	values := "("
	for k, v := range(contentMaps){
		if keys != "("{
			keys = keys + ","
		}
		keys = keys + k
		if values != "("{
			values = values + ","
		}
		switch val:=v.(type){
			case string:
				values = values + "\"" + val + "\""
			default:
				values = values + toString(val)
		}
	}
	
	keys = keys + ")"
	values = values + ")"	
    tableNames := "t_account"
	sql := fmt.Sprintf("insert into %s  %s  values  %s ", tableNames, keys, values)
	ret1, err1 := sqldb.Exec(sql)
	if err1 != nil {
		log.Error("exec:%s failed!", sql)
		return false, 0
	}
	var lastInserId int64 = 0
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
func Update_t_account(key string, contentJson string)(bool, int64){
	var contentMaps map[string]interface{}
	err := json.Unmarshal([]byte(contentJson), &contentMaps)
	if err == nil {
		return false, 0
	}
    redisKey := "t_account:"+key
	isExsit, _ := client.Exists(redisKey).Result()
	if isExsit == int64(0) {
		return false, 0
	}
	//更新redis
	for k, v := range(contentMaps){
		client.HSet(redisKey, k, v)
	}
	
	//Write to Mysql!
	keyvalues := ""
	for k, v := range(contentMaps){
		if keyvalues != ""{
			keyvalues = keyvalues + ","
		}
		keyvalues = keyvalues + k + " = " + toString(v)
	}
    conditions := fmt.Sprintf("accid = %s",sqlValueStr(key))
    tableNames := "t_account"

	sql := fmt.Sprintf("update from  %s  set(%s) where (%s)", tableNames, keyvalues, conditions)
	ret1, err1 := sqldb.Exec(sql)
	if err1 != nil {
		return false, 0
	}
	var lastInserId int64 = 0
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
func Remove_t_account(key string)(bool, int64){
    redisKey := "t_account:"+key
	isExsit, _ := client.Exists(redisKey).Result()
	if isExsit == int64(0) {
		return false, 0
	}
	//删除redis
	client.Del(redisKey)
	
	//Delete from Mysql!
    conditions := fmt.Sprintf("accid = %s",sqlValueStr(key))
    tableNames := "t_account"

	sql := fmt.Sprintf("delete from  %s  where (%s)", tableNames, conditions)
	ret1, err1 := sqldb.Exec(sql)
	if err1 != nil {
		return false, 0
	}

	var err2 error
	var RowsAffected int64 = 0
	if RowsAffected, err2 = ret1.RowsAffected(); nil != err2 {
		return false, 0
	} else {
		if RowsAffected == 0 {
			return false, 0
		}
	}
	return true, RowsAffected
}


func Read_t_role(key string)(result map[string]string){
    redisKey:= "t_role:"+key
    isExsit, _ := client.Exists(redisKey).Result()
    if isExsit == int64(1) { //在redis中有数据,则直接返回redis的数据
        result["roleid"]=client.HGet(redisKey, "roleid").Val()
        result["accid"]=client.HGet(redisKey, "accid").Val()
        result["nickName"]=client.HGet(redisKey, "nickName").Val()
        result["sex"]=client.HGet(redisKey, "sex").Val()
        result["templateId"]=client.HGet(redisKey, "templateId").Val()
        result["createtm"]=client.HGet(redisKey, "createtm").Val()
        result["lastsceneid"]=client.HGet(redisKey, "lastsceneid").Val()
        result["lastposX"]=client.HGet(redisKey, "lastposX").Val()
        result["lastposY"]=client.HGet(redisKey, "lastposY").Val()
        result["handWeapon"]=client.HGet(redisKey, "handWeapon").Val()
        result["bulletCount"]=client.HGet(redisKey, "bulletCount").Val()
        result["weaponList"]=client.HGet(redisKey, "weaponList").Val()
        result["level"]=client.HGet(redisKey, "level").Val()
        result["gold"]=client.HGet(redisKey, "gold").Val()
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
func Add_t_role(contentJson string) (bool, int64){
	var contentMaps map[string]interface{}
	err := json.Unmarshal([]byte(contentJson), &contentMaps)
	if err == nil {
		return false, 0
	}
    tableKey, isOk := contentMaps["roleid"]
	if isOk == false{
		return false, 0
	}
    redisKey := "t_role:"+tableKey.(string)

	isExsit, _ := client.Exists(redisKey).Result()
	if isExsit == int64(1) {
		return false, 0
	}
	
	//Write to Redis
	var fieldValue interface{}
	var keyIsExsit bool
	
    fieldValue, keyIsExsit = contentMaps["roleid"]
    if keyIsExsit == true {
        client.HSet(redisKey, "roleid", fieldValue)
    }else{
        client.HSet(redisKey, "roleid", "")
    }

    fieldValue, keyIsExsit = contentMaps["accid"]
    if keyIsExsit == true {
        client.HSet(redisKey, "accid", fieldValue)
    }else{
        client.HSet(redisKey, "accid", "")
    }

    fieldValue, keyIsExsit = contentMaps["nickName"]
    if keyIsExsit == true {
        client.HSet(redisKey, "nickName", fieldValue)
    }else{
        client.HSet(redisKey, "nickName", "")
    }

    fieldValue, keyIsExsit = contentMaps["sex"]
    if keyIsExsit == true {
        client.HSet(redisKey, "sex", fieldValue)
    }else{
        client.HSet(redisKey, "sex", "")
    }

    fieldValue, keyIsExsit = contentMaps["templateId"]
    if keyIsExsit == true {
        client.HSet(redisKey, "templateId", fieldValue)
    }else{
        client.HSet(redisKey, "templateId", "")
    }

    fieldValue, keyIsExsit = contentMaps["createtm"]
    if keyIsExsit == true {
        client.HSet(redisKey, "createtm", fieldValue)
    }else{
        client.HSet(redisKey, "createtm", "")
    }

    fieldValue, keyIsExsit = contentMaps["lastsceneid"]
    if keyIsExsit == true {
        client.HSet(redisKey, "lastsceneid", fieldValue)
    }else{
        client.HSet(redisKey, "lastsceneid", "")
    }

    fieldValue, keyIsExsit = contentMaps["lastposX"]
    if keyIsExsit == true {
        client.HSet(redisKey, "lastposX", fieldValue)
    }else{
        client.HSet(redisKey, "lastposX", "")
    }

    fieldValue, keyIsExsit = contentMaps["lastposY"]
    if keyIsExsit == true {
        client.HSet(redisKey, "lastposY", fieldValue)
    }else{
        client.HSet(redisKey, "lastposY", "")
    }

    fieldValue, keyIsExsit = contentMaps["handWeapon"]
    if keyIsExsit == true {
        client.HSet(redisKey, "handWeapon", fieldValue)
    }else{
        client.HSet(redisKey, "handWeapon", "")
    }

    fieldValue, keyIsExsit = contentMaps["bulletCount"]
    if keyIsExsit == true {
        client.HSet(redisKey, "bulletCount", fieldValue)
    }else{
        client.HSet(redisKey, "bulletCount", "")
    }

    fieldValue, keyIsExsit = contentMaps["weaponList"]
    if keyIsExsit == true {
        client.HSet(redisKey, "weaponList", fieldValue)
    }else{
        client.HSet(redisKey, "weaponList", "")
    }

    fieldValue, keyIsExsit = contentMaps["level"]
    if keyIsExsit == true {
        client.HSet(redisKey, "level", fieldValue)
    }else{
        client.HSet(redisKey, "level", "")
    }

    fieldValue, keyIsExsit = contentMaps["gold"]
    if keyIsExsit == true {
        client.HSet(redisKey, "gold", fieldValue)
    }else{
        client.HSet(redisKey, "gold", "")
    }

	//Write to Mysql!
	keys := "("
	values := "("
	for k, v := range(contentMaps){
		if keys != "("{
			keys = keys + ","
		}
		keys = keys + k
		if values != "("{
			values = values + ","
		}
		switch val:=v.(type){
			case string:
				values = values + "\"" + val + "\""
			default:
				values = values + toString(val)
		}
	}
	
	keys = keys + ")"
	values = values + ")"	
    tableNames := "t_role"
	sql := fmt.Sprintf("insert into %s  %s  values  %s ", tableNames, keys, values)
	ret1, err1 := sqldb.Exec(sql)
	if err1 != nil {
		log.Error("exec:%s failed!", sql)
		return false, 0
	}
	var lastInserId int64 = 0
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
func Update_t_role(key string, contentJson string)(bool, int64){
	var contentMaps map[string]interface{}
	err := json.Unmarshal([]byte(contentJson), &contentMaps)
	if err == nil {
		return false, 0
	}
    redisKey := "t_role:"+key
	isExsit, _ := client.Exists(redisKey).Result()
	if isExsit == int64(0) {
		return false, 0
	}
	//更新redis
	for k, v := range(contentMaps){
		client.HSet(redisKey, k, v)
	}
	
	//Write to Mysql!
	keyvalues := ""
	for k, v := range(contentMaps){
		if keyvalues != ""{
			keyvalues = keyvalues + ","
		}
		keyvalues = keyvalues + k + " = " + toString(v)
	}
    conditions := fmt.Sprintf("roleid = %s",sqlValueStr(key))
    tableNames := "t_role"

	sql := fmt.Sprintf("update from  %s  set(%s) where (%s)", tableNames, keyvalues, conditions)
	ret1, err1 := sqldb.Exec(sql)
	if err1 != nil {
		return false, 0
	}
	var lastInserId int64 = 0
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
func Remove_t_role(key string)(bool, int64){
    redisKey := "t_role:"+key
	isExsit, _ := client.Exists(redisKey).Result()
	if isExsit == int64(0) {
		return false, 0
	}
	//删除redis
	client.Del(redisKey)
	
	//Delete from Mysql!
    conditions := fmt.Sprintf("roleid = %s",sqlValueStr(key))
    tableNames := "t_role"

	sql := fmt.Sprintf("delete from  %s  where (%s)", tableNames, conditions)
	ret1, err1 := sqldb.Exec(sql)
	if err1 != nil {
		return false, 0
	}

	var err2 error
	var RowsAffected int64 = 0
	if RowsAffected, err2 = ret1.RowsAffected(); nil != err2 {
		return false, 0
	} else {
		if RowsAffected == 0 {
			return false, 0
		}
	}
	return true, RowsAffected
}


func Read_t_testtable(key string)(result map[string]string){
    redisKey:= "t_testtable:"+key
    isExsit, _ := client.Exists(redisKey).Result()
    if isExsit == int64(1) { //在redis中有数据,则直接返回redis的数据
        result["field1"]=client.HGet(redisKey, "field1").Val()
        result["field2"]=client.HGet(redisKey, "field2").Val()
        result["field3"]=client.HGet(redisKey, "field3").Val()
        result["field4"]=client.HGet(redisKey, "field4").Val()
        result["field5"]=client.HGet(redisKey, "field5").Val()
        result["field6"]=client.HGet(redisKey, "field6").Val()
        result["field7"]=client.HGet(redisKey, "field7").Val()
        return
    }
    sql := fmt.Sprintf("select * from t_testtable where field1 = %s", key)

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
func Add_t_testtable(contentJson string) (bool, int64){
	var contentMaps map[string]interface{}
	err := json.Unmarshal([]byte(contentJson), &contentMaps)
	if err == nil {
		return false, 0
	}
    tableKey, isOk := contentMaps["field1"]
	if isOk == false{
		return false, 0
	}
    redisKey := "t_testtable:"+tableKey.(string)

	isExsit, _ := client.Exists(redisKey).Result()
	if isExsit == int64(1) {
		return false, 0
	}
	
	//Write to Redis
	var fieldValue interface{}
	var keyIsExsit bool
	
    fieldValue, keyIsExsit = contentMaps["field1"]
    if keyIsExsit == true {
        client.HSet(redisKey, "field1", fieldValue)
    }else{
        client.HSet(redisKey, "field1", "")
    }

    fieldValue, keyIsExsit = contentMaps["field2"]
    if keyIsExsit == true {
        client.HSet(redisKey, "field2", fieldValue)
    }else{
        client.HSet(redisKey, "field2", "")
    }

    fieldValue, keyIsExsit = contentMaps["field3"]
    if keyIsExsit == true {
        client.HSet(redisKey, "field3", fieldValue)
    }else{
        client.HSet(redisKey, "field3", "")
    }

    fieldValue, keyIsExsit = contentMaps["field4"]
    if keyIsExsit == true {
        client.HSet(redisKey, "field4", fieldValue)
    }else{
        client.HSet(redisKey, "field4", "")
    }

    fieldValue, keyIsExsit = contentMaps["field5"]
    if keyIsExsit == true {
        client.HSet(redisKey, "field5", fieldValue)
    }else{
        client.HSet(redisKey, "field5", "")
    }

    fieldValue, keyIsExsit = contentMaps["field6"]
    if keyIsExsit == true {
        client.HSet(redisKey, "field6", fieldValue)
    }else{
        client.HSet(redisKey, "field6", "")
    }

    fieldValue, keyIsExsit = contentMaps["field7"]
    if keyIsExsit == true {
        client.HSet(redisKey, "field7", fieldValue)
    }else{
        client.HSet(redisKey, "field7", "")
    }

	//Write to Mysql!
	keys := "("
	values := "("
	for k, v := range(contentMaps){
		if keys != "("{
			keys = keys + ","
		}
		keys = keys + k
		if values != "("{
			values = values + ","
		}
		switch val:=v.(type){
			case string:
				values = values + "\"" + val + "\""
			default:
				values = values + toString(val)
		}
	}
	
	keys = keys + ")"
	values = values + ")"	
    tableNames := "t_testtable"
	sql := fmt.Sprintf("insert into %s  %s  values  %s ", tableNames, keys, values)
	ret1, err1 := sqldb.Exec(sql)
	if err1 != nil {
		log.Error("exec:%s failed!", sql)
		return false, 0
	}
	var lastInserId int64 = 0
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
func Update_t_testtable(key string, contentJson string)(bool, int64){
	var contentMaps map[string]interface{}
	err := json.Unmarshal([]byte(contentJson), &contentMaps)
	if err == nil {
		return false, 0
	}
    redisKey := "t_testtable:"+key
	isExsit, _ := client.Exists(redisKey).Result()
	if isExsit == int64(0) {
		return false, 0
	}
	//更新redis
	for k, v := range(contentMaps){
		client.HSet(redisKey, k, v)
	}
	
	//Write to Mysql!
	keyvalues := ""
	for k, v := range(contentMaps){
		if keyvalues != ""{
			keyvalues = keyvalues + ","
		}
		keyvalues = keyvalues + k + " = " + toString(v)
	}
    conditions := fmt.Sprintf("field1 = %s",sqlValueStr(key))
    tableNames := "t_testtable"

	sql := fmt.Sprintf("update from  %s  set(%s) where (%s)", tableNames, keyvalues, conditions)
	ret1, err1 := sqldb.Exec(sql)
	if err1 != nil {
		return false, 0
	}
	var lastInserId int64 = 0
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
func Remove_t_testtable(key string)(bool, int64){
    redisKey := "t_testtable:"+key
	isExsit, _ := client.Exists(redisKey).Result()
	if isExsit == int64(0) {
		return false, 0
	}
	//删除redis
	client.Del(redisKey)
	
	//Delete from Mysql!
    conditions := fmt.Sprintf("field1 = %s",sqlValueStr(key))
    tableNames := "t_testtable"

	sql := fmt.Sprintf("delete from  %s  where (%s)", tableNames, conditions)
	ret1, err1 := sqldb.Exec(sql)
	if err1 != nil {
		return false, 0
	}

	var err2 error
	var RowsAffected int64 = 0
	if RowsAffected, err2 = ret1.RowsAffected(); nil != err2 {
		return false, 0
	} else {
		if RowsAffected == 0 {
			return false, 0
		}
	}
	return true, RowsAffected
}

	
func registerAllOperateTableHandlers(){	
	
	RegisterReadTableHandler("t_account", Read_t_account)
	RegisterAddTableHandler("t_account", Add_t_account)
	RegisterUpdateTableHandler("t_account", Update_t_account)
	RegisterRemoveTableHandler("t_account", Remove_t_account)
		
	RegisterReadTableHandler("t_role", Read_t_role)
	RegisterAddTableHandler("t_role", Add_t_role)
	RegisterUpdateTableHandler("t_role", Update_t_role)
	RegisterRemoveTableHandler("t_role", Remove_t_role)
		
	RegisterReadTableHandler("t_testtable", Read_t_testtable)
	RegisterAddTableHandler("t_testtable", Add_t_testtable)
	RegisterUpdateTableHandler("t_testtable", Update_t_testtable)
	RegisterRemoveTableHandler("t_testtable", Remove_t_testtable)
		
}
