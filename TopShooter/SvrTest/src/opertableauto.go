//Warnning: This file is auto generate, don't modify it manual.
package app
import (
    "fmt"
	"github.com/go-redis/redis"
	"database/sql"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var client *redis.Client
var sqldb  *sql.DB

func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func init(){
	client = redis.NewClient(redisOptions())
	client.FlushDB()
	var err error
	sqldb, err = sql.Open("mysql", sqlOptions())
	check(err)
}


func Read_t_account(key string)(result map[string][string]){
    redisKey:= t_account+":"+key
    isExsit, _ := client.Exists(redisKey)
    if isExsit == int64(1) { //在redis中有数据,则直接返回redis的数据
        result["accid"]=client.HGet(redisKey, "accid")
        result["accountName"]=client.HGet(redisKey, "accountName")
        result["password"]=client.HGet(redisKey, "password")
        result["createtm"]=client.HGet(redisKey, "createtm")
        result["lastLoginTm"]=client.HGet(redisKey, "lastLoginTm")
        result["gameid"]=client.HGet(redisKey, "gameid")
        return
    }
    sql := fmt.Sprintf("select * from t_account where accid = %s", key)
"
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
	i := 0
	resultTmp := make(map[int]map[string]string)
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
}


func Read_t_role(key string)(result map[string][string]){
    redisKey:= t_role+":"+key
    isExsit, _ := client.Exists(redisKey)
    if isExsit == int64(1) { //在redis中有数据,则直接返回redis的数据
        result["roleid"]=client.HGet(redisKey, "roleid")
        result["accid"]=client.HGet(redisKey, "accid")
        result["nickName"]=client.HGet(redisKey, "nickName")
        result["sex"]=client.HGet(redisKey, "sex")
        result["templateId"]=client.HGet(redisKey, "templateId")
        result["createtm"]=client.HGet(redisKey, "createtm")
        result["lastsceneid"]=client.HGet(redisKey, "lastsceneid")
        result["lastposX"]=client.HGet(redisKey, "lastposX")
        result["lastposY"]=client.HGet(redisKey, "lastposY")
        result["handWeapon"]=client.HGet(redisKey, "handWeapon")
        result["bulletCount"]=client.HGet(redisKey, "bulletCount")
        result["weaponList"]=client.HGet(redisKey, "weaponList")
        result["level"]=client.HGet(redisKey, "level")
        result["gold"]=client.HGet(redisKey, "gold")
        return
    }
    sql := fmt.Sprintf("select * from t_role where roleid = %s", key)
"
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
	i := 0
	resultTmp := make(map[int]map[string]string)
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
}
