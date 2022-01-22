# Mongo Leaf
Mongo leaf is a simple tool between you and go-mongodb driver. in leaf we only use json format to use mongo just like mongo-shell if you have a dynamic scenario mongo leaf might be useful for you. I'm noob in go and programing sorry for bad explantion so I tryed to make anything simple for noobs like me.


# Contents
- [Create Connetion](#create-connetion)
- [Create Database](#create-database)
- [Create Colletion](#create-colletion)
- [example](#simple-example)
- [functions](#crud)
    - [Client](#client-functions)
        - [Connect](#Connect)
        - [Disconnect](#Disconnect)
        - [ListDatabaseNames](#ListDatabaseNames)
        - [ListDatabases](#ListDatabases)
    - [Database](#database-functions)
        - [Aggregate](#Aggregate)
        - [Drop](#Drop)
        - [CreateView](#CreateView)
        - [ListCollections](#ListCollections)
        - [RunCommand](#RunCommand)
        - [RunCommandCursor](#RunCommandCursor)
    - [Colletion](#colletion-functions)
        - [Find](#Find)
        - [Aggregate](#Aggregate)
        - [Update](#Update)
        - [Delete](#Delete)
        - [Insert](#Insert)
        - [FindOneAnd](#FindOneAnd)
        - [CountDocuments](#CountDocuments)
        - [Distinct](#Distinct)
        - [Drop](#Drop)
        - [EstimatedDocumentCount](#EstimatedDocumentCount)


# Install
first execute this command
```
go get github.com/aamirmousavi/mongoleaf
```
then in your go code import leaf package
```go
import (
   "github.com/aamirmousavi/mongoleaf"
)
```


# Create Connetion

for connect to mongodb 
```go
    client, err := mongoleaf.New("mongodb://localhost:27017")
    if err != nil {
        //deal with error
	}
```
put your mongo URL as an argument for make connetion with your database


# Create Database

in your mongodb client if you want to create or refrence to a database you would use DataBase Function just like this 
for making a refrence to your database
```go
    client, err := mongoleaf.New("mongodb://localhost:27017")
    if err != nil {
         //deal with error
    }
    myDataBase := client.DataBase("myDBName")
```
now you can esly access the colletions and function of each database you have

# Create Colletion

as you now in mongo each database might have many collations so when you define a database you can call or create a collation in that database just like this example
```go
myCollation := client.DataBase("DbName").Colletion("myCollation")
```
if database or colletion you calling is not exsists it will create them


# Simple Example
```go
package main

import (
	"fmt"

	"github.com/aamirmousavi/mongoleaf"
)

func main() {
	client, err := mongoleaf.New("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	siteDB := client.DataBase("website")
	products := siteDB.Colletion("products")

	products.InsertMany("", `{
		"title":"pro_1",
		"price":5
	}`, `{
		"title":"pro_2",
		"price":2
	}`, `{
		"title":"pro_3",
		"price":15
	}`)

	//first arg is filter that mongo filter rows and second is the update you want to set
	//and last is option if you don't have any option use empty string or empty json map: "" or "{}"
	result, err := products.UpdateMany(`{
		"price":{
			"$gt":10
		}
	}`, `{
		"$set":{
			"more_than_10":true
		}
	}`, "")

	fmt.Printf("update \n\terror: %v \n\tresult: %v \n", err, result)

	//if your filter is "{}" you can also use empty string as no filter
	//it will return all rows in a colletion
	rows, err := products.Find("", "")
	jsonRows := mongoleaf.JSONPretty(rows)
	fmt.Printf("find result \n\terror: %v \n\trows: %v \n", err, jsonRows)
}

```
OutPut
```
update 
        error: <nil> 
	result: map[MatchedCount:1 ModifiedCount:1 UpsertedCount:0 UpsertedID:<nil>] 
find result 
        error: <nil> 
        rows: [
        {
                "_id": "61e4bf46d16388a583426754",
                "price": 5,
                "title": "pro_1"
        },
        {
                "_id": "61e4bf46d16388a583426755",
                "price": 2,
                "title": "pro_2"
        },
        {
                "_id": "61e4bf46d16388a583426756",
                "more_than_10": true,
                "price": 15,
                "title": "pro_3"
        }
 ] 
```
# Functions

- # Client Functions

    you always able to use dirvers functions like start session or watch and extra like this example
    ```go
        client, err := mongoleaf.New("mongodb://localhost:27017")
        if err != nil {
            panic(err)
        }
        dbNames, err := client.Client.StartSession(&options.SessionOptions{})
    ```
    just add **``.Client``** and call your functions also this works for database and colletion not only client type
    - # Connect
        ```go
            client, err := mongoleaf.New("mongodb://localhost:27017")
            if err != nil {
                panic(err)
            }
            err = client.Connect()
        ```
    - # Disconnect
        ```go
            err = client.Disconnect()
        ```

    - # ListDatabaseNames
        ````go
            filter, option := `{}`, `{}` //empty filter return all 
            dbNames, err := client.ListDatabaseNames(filter, option)
        ````

    - # ListDatabases
        reutrn a array of maps with Name, Empty and SizeOnDisk keys
        ```go
            dataBases, err := client.ListDatabases(filter, option)
        ```

-  # Database Functions
    
    - # Aggregate
        ```go
            out, err := client.DataBase(dbName).Aggregate(pipline, option)
        ```
    - # Drop
        ```go
            err := client.DataBase(dbName).Drop()
        ```
    - # CreateCollection
        ```go
        err := client.DataBase(dbName).CreateCollection("colName")
        ``` 

    - # CreateView
        ```go
        err := client.DataBase(dbName).CreateView(viewName, viewOn, pipline, option)
        ```
    - # ListCollectionNames
        ```go
        colNames, err := client.DataBase(dbName).ListCollectionNames(filter, option)
        ```

    - # ListCollections
        ```go
        colNames, err := client.DataBase(dbName).ListCollections(filter, option)
        ```

    - # RunCommand
        ```go
        out, err := client.DataBase(dbName).RunCommand(Command, option)
        ```

    - # RunCommandCursor
        ```go
        res, err := client.DataBase(dbName).RunCommandCursor(Command, option)
        ```

- # Colletion Functions

    - # Find 
        for reading many recoard in a mongo colletion 
        ```go
        myCollation := client.DataBase("DbName").Colletion("myCollation")
        filter:=`{"title":"my title"}`  //just like mongoshell filter
        options:=`{"limit":2}`           //just like mongoshell option
        results, err := myCollation.Find(filter,options)
        ```
        it's retrun an array of maps if you want json  result you can parse any type to json with mongoleaf.JSON function
        ```go
        json := mongoleaf.JSON(results)
        ```
        for reading only one reacord 
        ```go
        results, err := myCollation.FindOne(`{"title":"my title"}`,``)
        ```
        as you see we don't have any option we should pass empty  `` "" `` or `` `{}` `` as option 
        also if you have empty filter you can use empty string `` "" `` or `` `{}` `` as filter

    - # Aggregate

        ```go
        results, err := myCollation.Aggregate(`[{},{}]`,``) 
        ```
        as you know in mongodb Aggregate has array of map style so keep it in mind
    - # Update 
        ```go
        result, err := myCollation.UpdateMany(`{"_id":"sd1d12ds12d..."}`,`{"title":"new title"}`,option) 
        ```
    - # Delete 
        for deleteing many recoards use DeleteMany fountion
        ```go
        result, err := client.DataBase("DbName").Colletion("myCollation").DeleteMany(`{"_id":"sd12d12d2123"}`,option)
        ```
        as you see we called it in defrent way you can call all CRUD fountions just like this code
        ```go
        result, err := myCollation.DeleteOne(`{"_id":"sd1d12ds12d..."}`,option) 
        ```
    - # Insert
        for inserting if you are using insert many you have a little difrent parameters
        ```go
        result, err := myCollation.InsertMany(option,`{"title":"the title 1"}`,`{"title":"the title 12"}`,`{"title":"the title 2"}`) 
        ```
        in isert many you can insert many rows you want just with adding a new  
        ```go 
        func (hand colletion) InsertMany(optionQuery , documents ...) (map[]interface{}, error) 
        ```
        for insetring one row 
        ```go
        result, err := myCollation.InsertOne(option,`{"title":"the title 1"}`) 
        ```

    - # FindOneAnd

    - # CountDocuments

    - # Distinct

    - # Drop

    - # EstimatedDocumentCount
