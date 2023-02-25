package api

//D:\golang\goAAA\goDeployVercel-master\api\index.go
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"users99"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	app *gin.Engine
)

var Db *gorm.DB
var err error

type users88 struct {
	gorm.Model
	USER_NAME     string
	USER_NAME_EN  string
	USER_POSITION string
}

// type costCenter struct {
// 	ID         int    `json:"id"`
// 	Lang       string `json:"lang"`
// 	COST_ID    string `json:"COST_ID"`
// 	COSTCENTER string `json:"COSTCENTER"`
// }

type ResultObject struct {
	Resultstatus string   `json:"resultstatus"`
	HeaderAr     []string `json:"headeAr"`
	ErrorMsg     string   `json:"ErrorMsg"`
	DataResult   []struct {
		Field1  string `json:"field1"`
		Field2  string `json:"field2"`
		Field3  string `json:"field1"`
		Field4  string `json:"field2"`
		Field5  string `json:"field1"`
		Field6  string `json:"field2"`
		Field7  string `json:"field1"`
		Field8  string `json:"field2"`
		Field9  string `json:"field2"`
		Field10 string `json:"field2"`
	} `json:"Dataresult"`
}
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type costCenter struct {
	ID         string `json:"id"`
	COST_ID    string `json:"title"`
	COSTCENTER string `json:"costcenter"`
}

type ProjectAsset struct {
	id                   int
	assetCode            string
	assetDesc            string
	sqlData              string
	sqlCountRec          string
	headerArray          string
	caption              string
	jsonDataReturnFormat string
}

var companyJSON = `{
	"name" : "GOLinuxCloud",
    "years_of_service" : "5",
    "nature_of_company" : "Online Academy",
    "no_of_staff" : "10"
}`

// CREATE ENDPOIND

func myRoute(r *gin.RouterGroup) {

	r.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from golang nutv99 in vercel")

	})

	r.GET("/books", listBooksHandler)
	r.GET("/getbypageno/:tablename/:pageno", getByPageNo)
	r.GET("/getbypagenoV2/:tablename/:pageno", getByPageNoV2)
	r.GET("/getbypagenoV3/:tablename/:pageno", getByPageNoV3)
	r.GET("/getbyid/:tablename/:id", getByID)

	r.POST("/readjson", readJSON)
	//	r.post("/create", createNew)
	r.GET("/costcenter", listCostCenter)
	// http.HandleFunc("/write-json-to-file", createNew)
}

func init() {

	// http.Get("/readjson",
	//    http.HandleFunc("/readjson", readJSON)
	// )

	// http.ListenAndServe("", nil)
	users99.FetchAll()
	app = gin.New()
	r := app.Group("/api")
	r.Use(cors.Default())
	myRoute(r)

	dsn := "lbg5pjees347lrun2wdl:pscale_pw_CgPWWfYLYTy3ziSn28WizLQ3fpt3dTTU24kgX82qxNA@tcp(ap-southeast.connect.psdb.cloud)/it_asset?tls=true"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Db.AutoMigrate(&Book2{})

	// Db.AutoMigrate(&Member{})

}

func createNew(w http.ResponseWriter, r *http.Request) {
	// Declared an empty map interface
	//var dataResult map[string]interface{}

	/* New จาก  chatgpt */
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the JSON data into a map[string]interface{}
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Print the JSON data to the console
	fmt.Println(string(jsonData))

	/*  จบ ส่วนจาก  ChatGPT  */

	// Unmarshal the JSON to the interface.its same as decode
	// err := json.Unmarshal([]byte(companyJSON), &dataResult)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	var tablename = "costCenter"
	var id = "999"
	var sql = "INSERT INTO " + tablename
	var sqlUpdate = "UPDATE " + tablename + " SET "
	var keyclause = "("
	var valueclause = " VALUES("
	// Print elements in map on the terminal the key and its value
	for key, value := range jsonData {
		fmt.Printf("%s : %v \n", key, value)
		keyclause = keyclause + string(key) + ","
		valueclause = valueclause + "'" + string(value) + "',"
		sqlUpdate = sqlUpdate + string(key) + "='" + string(value) + "',"
	}
	keyclause = strings.TrimSuffix(keyclause, ",")
	keyclause = keyclause + ")"
	valueclause = strings.TrimSuffix(valueclause, ",") + ")"

	sqlUpdate = strings.TrimSuffix(sqlUpdate, ",")

	sql = sql + keyclause + valueclause
	sqlUpdate = sqlUpdate + " WHERE id=" + id

	fmt.Println(sql)
	fmt.Println(sqlUpdate)

}

func getByPageNo(c *gin.Context) {

	dsn := "lbg5pjees347lrun2wdl:pscale_pw_CgPWWfYLYTy3ziSn28WizLQ3fpt3dTTU24kgX82qxNA@tcp(ap-southeast.connect.psdb.cloud)/it_asset?tls=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	lines := make([][]string, 0)
	var tablename = c.Param("tablename")
	var cc1 string = ""
	cc1 = c.Param("pageno")
	// var pageno int = strconv.Atoi(cc1)
	pageno, _ := strconv.ParseInt(cc1, 10, 64)
	var startrec = (pageno - 1) * 10
	startrec2 := strconv.Itoa(int(startrec))

	rows, err := db.Query("SELECT * FROM " + tablename + " LIMIT " + string(startrec2) + ",10")

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("Error getting columns from table", err)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// var names []string
	// var name string

	for rows.Next() {
		// read the row on the table
		// each column value will be stored in the slice
		err = rows.Scan(scanArgs...)

		fmt.Println("Error scanning rows from table", err)

		var value string
		var line []string

		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
				line = append(line, value)
			}
		}

		lines = append(lines, line)
	}

	c.String(http.StatusOK, "Read Books Success"+tablename)
	c.JSON(http.StatusOK, &lines)

	if err != nil {
		panic(err.Error())
	}

}

func getByPageNoV3(c *gin.Context) {

	var db *gorm.DB
	dsn := "lbg5pjees347lrun2wdl:pscale_pw_CgPWWfYLYTy3ziSn28WizLQ3fpt3dTTU24kgX82qxNA@tcp(ap-southeast.connect.psdb.cloud)/it_asset?tls=true"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()

	var costCenters []costCenter

	if result := db.Find(&costCenters); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error999": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &costCenters)

	// var tablename = c.Param("tablename")
	// var cc1 string = ""
	// cc1 = c.Param("pageno")
	// // var pageno int = strconv.Atoi(cc1)
	// pageno, _ := strconv.ParseInt(cc1, 10, 64)
	// var startrec = (pageno - 1) * 10
	// startrec2 := strconv.Itoa(int(startrec))

	// rows, err := db.Query("SELECT * FROM " + tablename + " LIMIT " + string(startrec2) + ",10")

	// // Get column names
	// columns, err := rows.Columns()
	// if err != nil {
	// 	fmt.Println("Error getting columns from table", err)
	// }

	// values := make([]sql.RawBytes, len(columns))

	// scanArgs := make([]interface{}, len(values))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }

	// // var names []string
	// // var name string

	// for rows.Next() {
	// 	// read the row on the table
	// 	// each column value will be stored in the slice
	// 	err = rows.Scan(scanArgs...)

	// 	fmt.Println("Error scanning rows from table", err)

	// 	var value string
	// 	var line []string

	// 	for _, col := range values {
	// 		// Here we can check if the value is nil (NULL value)
	// 		if col == nil {
	// 			value = "NULL"
	// 		} else {
	// 			value = string(col)
	// 			line = append(line, value)
	// 		}
	// 	}

	// 	lines = append(lines, line)
	// }

	// c.String(http.StatusOK, "Read Books Success"+tablename)
	// c.JSON(http.StatusOK, &lines)

	// if err != nil {
	// 	panic(err.Error())
	// }

}

func getByPageNoV2(c *gin.Context) {

	//dsn := "lbg5pjees347lrun2wdl:pscale_pw_CgPWWfYLYTy3ziSn28WizLQ3fpt3dTTU24kgX82qxNA@tcp(ap-southeast.connect.psdb.cloud)/it_asset?tls=true"
	dsn := "ddhousin:y4e2Q44rBw@tcp(https://lovetoshopmall.com:443)/it_asset?tls=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	lines := make([][]string, 0)
	var tablename = c.Param("tablename")
	var cc1 string = ""
	cc1 = c.Param("pageno")

	var strArray [30]string
	strArray[0] = "A"
	strArray[1] = "B"
	strArray[2] = "C"
	strArray[3] = "D"
	strArray[4] = "E"
	strArray[5] = "F"
	strArray[6] = "G"
	strArray[7] = "H"
	strArray[8] = "I"

	for i := 65; i < 90; i++ {
		strArray[i-65] = string(i)
	}

	var sqlData, sqlCountRec, headerArray, caption, jsonDataResultFormat, fieldList string
	//var dataAsset string

	err = db.QueryRow("SELECT sqlData,sqlCountRec,headerArray,caption,jsonDataReturnFormat,fieldList FROM ProjectAsset WHERE assetCode=?", tablename).Scan(&sqlData, &sqlCountRec, &headerArray, &caption, &jsonDataResultFormat, &fieldList)
	if err != nil {
		//fmt.Println(err.Error())
		c.JSON(200, err.Error())
		c.String(200, sqlData+"-"+tablename)
		return
	}

	fieldListArray := strings.Split(fieldList, ",")

	totalRec := 0
	err = db.QueryRow(sqlCountRec).Scan(&totalRec)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.String(200, sqlData+"-"+tablename)
		return
	}

	// var pageno int = strconv.Atoi(cc1)
	pageno, _ := strconv.ParseInt(cc1, 10, 64)
	var startrec = (pageno - 1) * 10
	var endrec = int(startrec) + 10

	endrec2 := strconv.Itoa(int(endrec))

	fmt.Println(endrec)
	orderClause := ""

	//sortbyfieldno =  c.Query("sortby")
	var sortbyfieldno = c.Query("sortby")
	var sortbytype = c.Query("sortbytype")
	sortno2, err := strconv.ParseInt(sortbyfieldno, 10, 0)

	if sortbyfieldno != "" {
		orderClause = " ORDER BY " + fieldListArray[sortno2] + " " + sortbytype
	} else {
		orderClause = " "
	}

	var wherefieldName = c.Query("fieldName")
	var searchText = c.Query("searchText")
	var whereClause = " "
	if searchText != "" {
		whereClause = " WHERE " + wherefieldName + " LIKE '" + "%" + searchText + "%' "
		startrec = 0

	} else {
		whereClause = " "
	}

	startrec2 := strconv.Itoa(int(startrec))

	fmt.Println((whereClause))
	query := sqlData + whereClause + orderClause + " LIMIT " + string(startrec2) + ",10"
	rows, err := db.Query(query)
	//c.String(200, query)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("Error getting columns from table", err)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// var names []string
	// var name string
	var line99 string = "["

	for rows.Next() {
		// read the row on the table
		// each column value will be stored in the slice
		err = rows.Scan(scanArgs...)

		fmt.Println("Error scanning rows from table", err)
		line99 = line99 + "{"

		var value string
		var line []string
		asciiNum := 34 // Uppercase A
		character := string(asciiNum)
		var colNo int8
		colNo = 0
		fieldName := ""
		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
				//fieldName = "field_" + strArray[colNo]
				fieldName = strconv.Itoa(int(colNo))

				line99 = line99 + character +
					fieldName + character + ":" + character + "" + character + ","

			} else {

				value = string(col)
				//fieldName = "field_" + strArray[colNo]
				//fieldName = string(colNo)
				fieldName = strconv.Itoa(int(colNo))

				//line = append(line, fieldName+":"+value)
				line99 = line99 + character +
					fieldName + character + ":" + character + value + character + ","
			}
			colNo++
		}
		fieldName = fieldName
		line99 = strings.TrimSuffix(line99, ",")
		line99 = line99 + "},"

		lines = append(lines, line)
	}

	//jsonData, err := json.Marshal(lines)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}
	line99 = strings.TrimSuffix(line99, ",")
	line99 = line99 + "]"

	//line99 = `[{"name":"John","age":30,"city":"New York"},{"name":"John","age":30,"city":"New York"}]`

	var data interface{}
	var data2 interface{}

	err = json.Unmarshal([]byte(line99), &data)
	if err != nil {
		panic(err)
	}

	//ll  = json.Marshal(line99)
	totalPage := 0
	totalPage = totalRec / 10

	totalRec2 := strconv.Itoa(int(totalRec))
	totalPage2 := strconv.Itoa(int(totalPage))
	pageno2 := strconv.Itoa(int(pageno))
	fmt.Println((totalPage2))

	jsonDataResultFormat = strings.Replace(jsonDataResultFormat, "@#status", "success", 10)
	jsonDataResultFormat = strings.Replace(jsonDataResultFormat, "@#pageno", string(pageno2), 10)
	jsonDataResultFormat = strings.Replace(jsonDataResultFormat, "@#dataResult", line99, 10)
	jsonDataResultFormat = strings.Replace(jsonDataResultFormat, "@#HeadCol", headerArray, 1)
	jsonDataResultFormat = strings.Replace(jsonDataResultFormat, "@#startRec", startrec2, 1)
	jsonDataResultFormat = strings.Replace(jsonDataResultFormat, "@#endRec", string(endrec2), 1)
	jsonDataResultFormat = strings.Replace(jsonDataResultFormat, "@#totalRec", string(totalRec2), 1)
	jsonDataResultFormat = strings.Replace(jsonDataResultFormat, "@#totalPage", string(totalPage2), 1)

	//c.String(http.StatusOK, jsonDataResultFormat)

	err = json.Unmarshal([]byte(jsonDataResultFormat), &data2)
	if err != nil {
		//c.JSON(http.StatusOK, error.Error())
		c.JSON(http.StatusOK, err.Error())
		//panic(err)
		return
	}
	c.JSON(http.StatusOK, data2)

	// c.JSON(http.StatusOK, jsonData)
	//c.Data(http.StatusOK, "application/json", jsonData)

	if err != nil {
		panic(err.Error())
	}

}

func getByID(c *gin.Context) {

	dsn := "lbg5pjees347lrun2wdl:pscale_pw_CgPWWfYLYTy3ziSn28WizLQ3fpt3dTTU24kgX82qxNA@tcp(ap-southeast.connect.psdb.cloud)/it_asset?tls=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	lines := make([][]string, 0)
	var tablename = c.Param("tablename")
	var cc1 string = ""
	cc1 = c.Param("id")

	rows, err := db.Query("SELECT * FROM " + tablename + " where id= " + cc1)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("Error getting columns from table", err)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// var names []string
	// var name string

	for rows.Next() {
		// read the row on the table
		// each column value will be stored in the slice
		err = rows.Scan(scanArgs...)

		fmt.Println("Error scanning rows from table", err)

		var value string
		var line []string

		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
				line = append(line, value)
			}
		}

		lines = append(lines, line)
	}

	c.String(http.StatusOK, "Read Books Success"+tablename)
	c.JSON(http.StatusOK, &lines)

	if err != nil {
		panic(err.Error())
	}

}

func fck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func listBooksHandler(c *gin.Context) {
	var books []Book
	if result := Db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &books)
}

func listCostCenter(c *gin.Context) {
	var costs []costCenter
	if result := Db.Find(&costs); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &costs)
}

// ADD THIS SCRIPT
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

func readJSON(c *gin.Context) {

	dsn := "lbg5pjees347lrun2wdl:pscale_pw_CgPWWfYLYTy3ziSn28WizLQ3fpt3dTTU24kgX82qxNA@tcp(ap-southeast.connect.psdb.cloud)/it_asset?tls=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Read the JSON data from the request body

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the JSON data into a map[string]interface{}
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Print the parsed data to the console
	fmt.Printf("%+v\n", data)

	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Print the JSON data to the console
	fmt.Println(string(jsonData))
	var tablename = "notable"
	var id = "999"

	for key, value := range data {
		fmt.Printf("%s : %v \n", key, value)
		if key == "id" || key == "tablename" {
			if key == "id" || key == "tablename" {
				id = value.(string)
			}
			if key == "tablename" {
				tablename = value.(string)
			}

		}
	}

	//var id = r.URL.Query().Get("id")

	var sql = "INSERT INTO " + tablename
	var sqlUpdate = "UPDATE " + tablename + " SET "
	var keyclause = "("
	var valueclause = " VALUES("
	//var stVal
	// Print elements in map on the terminal the key and its value
	for key, value := range data {
		fmt.Printf("%s : %v \n", key, value)
		if key == "id" || key == "tablename" {
			if key == "id" || key == "tablename" {
				id = value.(string)
			}
			if key == "tablename" {
				tablename = value.(string)
			}
		} else {
			keyclause = keyclause + string(key) + ","
			valueclause = valueclause + "'" + value.(string) + "',"
			sqlUpdate = sqlUpdate + string(key) + "='" + value.(string) + "',"
		}
	}

	keyclause = strings.TrimSuffix(keyclause, ",")
	keyclause = keyclause + ")"
	valueclause = strings.TrimSuffix(valueclause, ",") + ")"

	sqlUpdate = strings.TrimSuffix(sqlUpdate, ",")

	sql = sql + keyclause + valueclause
	sqlUpdate = sqlUpdate + " WHERE id=" + id

	//c.String(http.StatusOK, "Read Books Success"+sql+" and "+sqlUpdate)
	var query string = ""
	if id == "000" {
		query = sql
	} else {
		query = sqlUpdate
	}

	rows, err := db.Query(query)

	fmt.Println(rows)
	fmt.Println(sqlUpdate)

	// Get column names

	if err != nil {
		c.String(http.StatusOK, "error id is "+id+" query is "+query+"-"+err.Error())
	} else {
		c.String(http.StatusOK, "Success-"+query)
	}

	// Return a success response to the client
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Data parsed successfully"))
}

/*
ALTER TABLE `it_asset`.`books`
CHANGE COLUMN `id` `id` INT NOT NULL AUTO_INCREMENT ;
*/
