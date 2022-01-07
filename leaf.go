package mongoleaf

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Branch struct {
	Client *mongo.Client
}

type dataBase struct {
	DB *mongo.Database
}

type colletion struct {
	Colletion *mongo.Collection
}

func New(URI string) (Branch, error) {

	c, err := connect(URI)
	if err != nil {
		return Branch{}, err
	}

	return Branch{Client: c}, nil
}
func connect(URI string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// log.Printf("mongodb Faild to send Ping to MongoDB with '%v' URI. \t Error Value: %v. \n", URI, err)
		return nil, err
	}
	return client, nil
}

func (branch Branch) DataBase(dataBaseName string) dataBase {
	return dataBase{DB: branch.Client.Database(dataBaseName)}
}

func (db dataBase) Colletion(colletionName string) colletion {
	return colletion{Colletion: db.DB.Collection(colletionName)}
}
