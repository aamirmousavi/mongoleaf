# Mongo Leaf
Mongo leaf is a simple tool between you and go-mongodb driver. in leaf we only use json format to use mongo just like mongo-shell if you have a dynamic senario it's very usfull for you. I'm noob in go and programing sorry for bad explantion so I tryed to make anything simple for noobs like me.


## Contents
- [Create Connetion](#)
- [CRUD]()
- [Options](#)

## Install
exute this command on your terminal
```
    go get github.com/aamirmousavi/MongoLeaf
```
then in your go code import leaf package
```go
import (
   "github.com/aamirmousavi/MongoLeaf"
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

## Database

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

## Colletion

as you now in mongo each database can have many collations so when you define a database you can call or create a collation in that database just like this example
```go
myCollation := branch.DataBase("DbName").Colletion("myCollation")
```
if database or colletion you calling is not exsists it will create them

## CRUD
### Read Many
for reading many recoard in a mongo colletion 
