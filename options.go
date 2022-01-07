package mongoleaf

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var option optionType

type optionType struct{}

func (_ optionType) update(option string) (options.UpdateOptions, error) {
	var op options.UpdateOptions
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.UpdateOptions{}, err
	}
	return op, nil
}
func (_ optionType) delete(option string) (options.DeleteOptions, error) {
	var op options.DeleteOptions
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.DeleteOptions{}, err
	}
	return op, nil
}
func (_ optionType) insertMany(option string) (options.InsertManyOptions, error) {
	var op options.InsertManyOptions
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.InsertManyOptions{}, err
	}
	return op, nil
}
func (_ optionType) insertOne(option string) (options.InsertOneOptions, error) {
	var op options.InsertOneOptions
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.InsertOneOptions{}, err
	}
	return op, nil
}
func (_ optionType) aggregate(option string) (options.AggregateOptions, error) {
	var op options.AggregateOptions
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.AggregateOptions{}, err
	}
	return op, nil
}
func (_ optionType) find(option string) (options.FindOptions, error) {
	var op options.FindOptions
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.FindOptions{}, err
	}
	return op, nil
}

func (_ optionType) findOne(option string) (options.FindOneOptions, error) {
	var op options.FindOneOptions
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.FindOneOptions{}, err
	}
	return op, nil
}
