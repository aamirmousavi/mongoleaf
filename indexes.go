package mongoleaf

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
)

func (hand colletion) CreateIndex(indexModel, indexOption string) (string, error) {
	var model mongo.IndexModel
	err := json.Unmarshal([]byte(indexModel), &model)
	if err != nil {
		return "", err
	}
	op, err := option.Database.Colletion.Indexes.create(indexOption)
	if err != nil {
		return "", err
	}
	return hand.Colletion.Indexes().CreateOne(context.TODO(), model, &op)
}
