# Mongo Leaf
Mongo leaf is a simple tool between you and go.mongodb.org/mongo-driver. in leaf we only use json format to use mongo just like mongo-shell if you have a dynamic senario it's very usfull for you. I'm noob in go and programing sorry for bad explantion so I tryed to make anything simple for noobs like me.


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
exute this command on your terminal
```
    go get github.com/aamirmousavi/mongoleaf
```
then in your go code import leaf package
```go
import (
   mongoleaf "github.com/aamirmousavi/mongoleaf"
)
```


# Create Connetion

The way of connecting to mongo database as a client you need to use New function just like this exmaple
```go
    client, err := mongoleaf.New("mongodb://localhost:27017")
    if err != nil {
		fmt.Println("error ", err)
		return
	}
```
it's take a  parameter as URI of your mongodb and connect you to your database


# Create Database

in your mongodb client if you want to create or refrence to a database you would use DataBase Function just like this example
```go
    client, err := mongoleaf.New("mongodb://localhost:27017")
    if err != nil {
        fmt.Println("error ", err)
        return
    }
    myDataBase := client.DataBase("myDBName")
```
I think you undrastand now what client means if you used original mongo driver for golang it's just like **mongo.Client** in original driver


# Create Colletion

as you now in mongo each database can have many collations so when you define a database you can call or create a collation in that database just like this example
```go
myCollation := client.DataBase("DbName").Colletion("myCollation")
```
if database or colletion you calling is not exsists it will create them


# Simple Example

# Functions

- # Client Functions

    you alway able to use dirvers functions like start session or watch and extra like this example
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
