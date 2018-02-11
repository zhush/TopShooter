package main

import (
	"container/list"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"libs/log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//字段结构
type DataField struct {
	FieldName string
	FieldType string
}

//表结构
type TableInfo struct {
	TableName   string
	TableFields []DataField
}

func main() {
	fmt.Println("Start Generate Code")
	var dbhost string = "127.0.0.1:3306"
	var dbuser string = "root"
	var dbpass string = "12345678"
	var dbname string = "topshooter"
	data, err := ioutil.ReadFile("dbconfig.json")
	if data != nil && err == nil {
		var dbConf map[string]interface{}
		err := json.Unmarshal(data, &dbConf)
		if err == nil {
			dbhost = dbConf["host"].(string)
			dbuser = dbConf["user"].(string)
			dbpass = dbConf["pass"].(string)
			dbname = dbConf["database"].(string)
		}
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		dbuser,
		dbpass,
		dbhost,
		dbname)
	db, err1 := sql.Open("mysql", dns)
	if err1 != nil {
		log.Fatal("Connect Database(%s) failed! error:%s", dbname, err1.Error())
	}
	fmt.Println("Open Database:" + dbname + " succeed!")

	tableNames := QueryAllTableNames(db, dbname)
	fmt.Println("tableNames:", tableNames)
	tableList := list.New()
	for i := 0; i < len(tableNames); i++ {
		tableList.PushBack(&TableInfo{tableNames[i], QueryTableAllFieldNames(db, tableNames[i], dbname)})
	}

	WriteDatabaseFiles(tableList, "../../opertableauto.go")

	for e := tableList.Front(); e != nil; e = e.Next() {
		//fmt.Println(e.Value)
	}
	defer db.Close()

}

func QueryAllTableNames(db *sql.DB, dbname string) []string {
	sql := fmt.Sprintf("SELECT table_name name,TABLE_COMMENT value FROM INFORMATION_SCHEMA.TABLES WHERE table_type='base table' and table_schema = '%s' order by table_name asc", dbname)

	var result []string = make([]string, 0)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("Excute sql:(%s) Failed, error:(%s)\n", sql, err.Error())
		return result
	}
	//返回所有列
	cols, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//返回所有列
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

	for i := 0; i < len(resultTmp); i++ {
		result = append(result, resultTmp[i]["name"])
	}
	return result
}

//遍历数据库中表字段名称
func QueryTableAllFieldNames(db *sql.DB, tableName string, dbName string) (result []DataField) {
	sql := fmt.Sprintf("select COLUMN_NAME, DATA_TYPE from information_schema.COLUMNS where table_name = '%s' and table_schema = '%s' ", tableName, dbName)
	fmt.Println("TableName:", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("Excute sql:(%s) Failed, error:(%s)\n", sql, err.Error())
		return result
	}
	//返回所有列
	cols, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//返回所有列
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

	for i := 0; i < len(resultTmp); i++ {
		result = append(result, DataField{resultTmp[i]["COLUMN_NAME"], resultTmp[i]["DATA_TYPE"]})
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//将自动处理数据库(拉取到redis和回写到mysql和redis的处理)
func WriteDatabaseFiles(tableList *list.List, outputFile string) {
	content := "//Warnning: This file is auto generate, don't modify it manual.\n"
	content = content +
		`package app
import (
    "fmt"
	"github.com/go-redis/redis"
	"database/sql"
	"encoding/json"
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
`
	for e := tableList.Front(); e != nil; e = e.Next() {
		tableInfo := e.Value.(*TableInfo)
		content = content + GenerateTableReadFunction(tableInfo)
		content = content + GenerateTableAddFunction(tableInfo)
		content = content + GenerateTableUpdateFunction(tableInfo)
		content = content + GenerateTableRemoveFunction(tableInfo)
	}

	var f *os.File
	var err1 error
	if checkFileIsExist(outputFile) { //如果文件存在
		err1 = os.Remove(outputFile) //删除文件
		check(err1)
	}
	f, err1 = os.Create(outputFile) //创建文件
	check(err1)
	n, err1 := io.WriteString(f, content) //写入文件(字符串)
	check(err1)
	fmt.Printf("写入 %d 个字节n", n)
}

func ConvertCondFormatStr(keyType string) string {
	if keyType == "tinyint" || keyType == "smallint" || keyType == "mediumint" || keyType == "int" || keyType == "bigint" || keyType == "float" || keyType == "double" || keyType == "decimal" {
		return "%s"
	} else {
		return "'%s'"
	}
}

func GenerateTableReadFunction(tableInfo *TableInfo) string {

	tableName := tableInfo.TableName
	keyName := tableInfo.TableFields[0].FieldName
	keyType := tableInfo.TableFields[0].FieldType

	ret := "\n\n"
	ret = ret + "func Read_" + tableName + "(key string)(result map[string][string]){\n"
	ret = ret + "    redisKey:= " + tableName + "+\":\"+key\n"
	ret = ret + "    isExsit, _ := client.Exists(redisKey)\n"
	ret = ret + "    if isExsit == int64(1) { //在redis中有数据,则直接返回redis的数据\n"
	for i := 0; i < len(tableInfo.TableFields); i++ {
		tableField := tableInfo.TableFields[i]
		ret = ret + fmt.Sprintf("        result[\"%s\"]=client.HGet(redisKey, \"%s\")\n", tableField.FieldName, tableField.FieldName)
	}
	ret = ret + "        return\n"
	ret = ret + "    }\n"
	ret = ret + fmt.Sprintf("    sql := fmt.Sprintf(\"select * from %s where %s = %s\", key)\n", tableName, keyName, ConvertCondFormatStr(keyType))
	ret = ret + `
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
`
	return ret
}

func GenerateTableAddFunction(tableInfo *TableInfo) string {
	tableName := tableInfo.TableName
	keyName := tableInfo.TableFields[0].FieldName
	//keyType := tableInfo.TableFields[0].FieldType

	ret := "\n\n"
	ret = ret + "//添加表记录的方法，传入的是json字符串,如果插入成功，则返回true,否则返回false\n"
	ret = ret + "func Add_" + tableName + "(contentJson string) bool{"
	ret = ret + `
	var contentMaps map[string]interface{}
	err := json.Unmarshal(contentJson, &contentMaps)
	if err == nil {
		return false
	}
`
	ret = ret + "    tableKey, isOk := contentMaps[\"" + keyName + "\"]"
	ret = ret + `
	if isOk == false{
		return false
	}
`
	ret = ret + "    redisKey := " + tableName + "+\":\"+tableKey\n"
	ret = ret + `
	isExsit, _ := client.Exists(redisKey)
	if isExsit == true {
		return false
	}
	
	//Write to Redis
	var fieldValue string
	var isExsit bool
	`
	for i := 0; i < len(tableInfo.TableFields); i++ {
		tableField := tableInfo.TableFields[i]
		ret = ret + fmt.Sprintf("\n    fieldValue, isExsit = contentMaps[\"%s\"])\n", tableField.FieldName)
		ret = ret + "    if isExsit == true {\n"
		ret = ret + fmt.Sprintf("        client.HSet(redisKey, \"%s\", fieldValue)\n    }else{\n", tableField.FieldName)
		ret = ret + fmt.Sprintf("        client.HSet(redisKey, \"%s\", \"\")\n    }\n", tableField.FieldName)
	}

	ret = ret + `
	//Write to Mysql!
	sql := 
	
	
	`

	ret = ret + "\n    return true\n}\n"
	return ret
}

func GenerateTableUpdateFunction(tableInfo *TableInfo) string {
	return ""
}

func GenerateTableRemoveFunction(tableInfo *TableInfo) string {
	return ""
}
