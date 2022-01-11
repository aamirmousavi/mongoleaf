# Mongo Leaf
Mongo leaf is a simple tool between you and go.mongodb.org/mongo-driver. in leaf we only use json format to use mongo just like mongo-shell if you have a dynamic senario it's very usfull for you. I'm noob in go and programing sorry for bad explantion so I tryed to make anything simple for noobs like me.


## Contents
- [Create Connetion](#create-connetion)
- [Create Database](#crud)
- [Create Colletion](#crud)
- [functions](#crud)
    - [Client](#)
    - [Database](#)
    - [Colletion](#)
- [exaplme](#crud)

## Install
exute this command on your terminal
```
    go get github.com/aamirmousavi/Mongo-Leaf
```
then in your go code import leaf package
```go
import (
   mongoleaf "github.com/aamirmousavi/MongoLeaf"
)
```

## Create Connetion

The way of connecting to mongo database as a client you need to use New function just like this exmaple
```go
    branch, err := mongoleaf.New("mongodb://localhost:27017")
    if err != nil {
		fmt.Println("error ", err)
		return
	}
```
it's take a string parameter as URI of your mongodb and connect you to your database

## Create Database

in your mongodb client if you want to create or refrence to a database you would use DataBase Function just like this example
```go
    branch, err := mongoleaf.New("mongodb://localhost:27017")
    if err != nil {
        fmt.Println("error ", err)
        return
    }
    myDataBase := branch.DataBase("myDBName")
```
I think you undrastand now what branch means if you used original mongo driver for golang it's just like **mongo.Client** in original driver

## Create Colletion

as you now in mongo each database can have many collations so when you define a database you can call or create a collation in that database just like this example
```go
myCollation := branch.DataBase("DbName").Colletion("myCollation")
```
if database or colletion you calling is not exsists it will create them

## Simple Example



## Client Functions

Connect

Disconnect

ListDatabaseNames

ListDatabases

NumberSessionsInProgress

StartSession

## Database Functions

Aggregate

Drop

CreateView

ListCollectionNames

ListCollections

RunCommand

RunCommandCursor

## Colletion Functions

### Read 
for reading many recoard in a mongo colletion 
```go
myCollation := branch.DataBase("DbName").Colletion("myCollation")
filter:=`{"title":"my title"}`  //just like mongoshell filter
options:=`{"limit":2}`           //just like mongoshell option
results, err := myCollation.Find(filter,options)
```
it's retrun an array of maps if you want json string result you can parse any type to json with mongoleaf.JSON function
```go
json := mongoleaf.JSON(results)
```
for reading only one reacord 
```go
results, err := myCollation.FindOne(`{"title":"my title"}`,``)
```
as you see we don't have any option we should pass empty string `` "" `` or `` `{}` `` as option 


## Aggregate

```go
results, err := myCollation.Aggregate(`[{},{}]`,``) 
```
as you know in mongodb Aggregate has array of map style so keep it in mind
### Update 
```go
result, err := myCollation.UpdateMany(`{"_id":"sd1d12ds12d..."}`,`{"title":"new title"}`,optionString) 
```
### Delete 
for deleteing many recoards use DeleteMany fountion
```go
result, err := branch.DataBase("DbName").Colletion("myCollation").DeleteMany(`{"_id":"sd12d12d2123"}`,optionString)
```
as you see we called it in defrent way you can call all CRUD fountions just like this code
```go
result, err := myCollation.DeleteOne(`{"_id":"sd1d12ds12d..."}`,optionString) 
```
### Insert
for inserting if you are using insert many you have a little difrent parameters
```go
result, err := myCollation.InsertMany(optionString,`{"title":"the title 1"}`,`{"title":"the title 12"}`,`{"title":"the title 2"}`) 
```
in isert many you can insert many rows you want just with adding a new string 
```go 
func (hand colletion) InsertMany(optionQuery string, documents ...string) (map[string]interface{}, error) 
```
for insetring one row 
```go
result, err := myCollation.InsertOne(optionString,`{"title":"the title 1"}`) 
```

FindOneAndUpdateDistinct

FindOneAndReplace

FindOneAndDelete

CountDocuments

Distinct

Drop

EstimatedDocumentCount
