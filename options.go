package mongoleaf

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var option optionType

type optionType struct{}

func (_ optionType) update(option string) (options.UpdateOptions, error) {
	var op options.UpdateOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.UpdateOptions{}, err
	}
	return op, nil
}
func (_ optionType) delete(option string) (options.DeleteOptions, error) {
	var op options.DeleteOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.DeleteOptions{}, err
	}
	return op, nil
}
func (_ optionType) insertMany(option string) (options.InsertManyOptions, error) {
	var op options.InsertManyOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.InsertManyOptions{}, err
	}
	return op, nil
}
func (_ optionType) insertOne(option string) (options.InsertOneOptions, error) {
	var op options.InsertOneOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.InsertOneOptions{}, err
	}
	return op, nil
}
func (_ optionType) aggregate(option string) (options.AggregateOptions, error) {
	var op options.AggregateOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.AggregateOptions{}, err
	}
	return op, nil
}
func (_ optionType) find(option string) (options.FindOptions, error) {
	var op options.FindOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.FindOptions{}, err
	}
	return op, nil
}

func (_ optionType) findOne(option string) (options.FindOneOptions, error) {
	var op options.FindOneOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.FindOneOptions{}, err
	}
	return op, nil
}
func (_ optionType) countDocuments(option string) (options.CountOptions, error) {
	var op options.CountOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.CountOptions{}, err
	}
	return op, nil
}
func (_ optionType) distinct(option string) (options.DistinctOptions, error) {
	var op options.DistinctOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.DistinctOptions{}, err
	}
	return op, nil
}
func (_ optionType) findOneAndUpdate(option string) (options.FindOneAndUpdateOptions, error) {
	var op options.FindOneAndUpdateOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.FindOneAndUpdateOptions{}, err
	}
	return op, nil
}
func (_ optionType) findOneAndReplace(option string) (options.FindOneAndReplaceOptions, error) {
	var op options.FindOneAndReplaceOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.FindOneAndReplaceOptions{}, err
	}
	return op, nil
}
func (_ optionType) findOneAndDelete(option string) (options.FindOneAndDeleteOptions, error) {
	var op options.FindOneAndDeleteOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.FindOneAndDeleteOptions{}, err
	}
	return op, nil
}

//

func (_ optionType) createCollection(option string) (options.CreateCollectionOptions, error) {
	var op options.CreateCollectionOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.CreateCollectionOptions{}, err
	}
	return op, nil
}
func (_ optionType) createView(option string) (options.CreateViewOptions, error) {
	var op options.CreateViewOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.CreateViewOptions{}, err
	}
	return op, nil
}
func (_ optionType) listCollections(option string) (options.ListCollectionsOptions, error) {
	var op options.ListCollectionsOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.ListCollectionsOptions{}, err
	}
	return op, nil
}

func (_ optionType) listDatabaseNames(option string) (options.ListDatabasesOptions, error) {
	var op options.ListDatabasesOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.ListDatabasesOptions{}, err
	}
	return op, nil
}
func (_ optionType) session(option string) (options.SessionOptions, error) {
	var op options.SessionOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.SessionOptions{}, err
	}
	return op, nil
}
func (_ optionType) runCmd(option string) (options.RunCmdOptions, error) {
	var op options.RunCmdOptions
	if option == "" || option == "{}" {
		return op, nil
	}
	err := json.Unmarshal([]byte(option), &op)
	if err != nil {
		return options.RunCmdOptions{}, err
	}
	return op, nil
}
