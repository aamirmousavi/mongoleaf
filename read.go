package mongoleaf

import (
	"context"
)

func (hand colletion) Aggregate(filter string, optionQuery string) ([]map[string]interface{}, error) {
	filterMap, err := jsonToArrayMap(filter)
	if err != nil {
		return nil, err
	}

	op, err := option.aggregate(optionQuery)
	if err != nil {
		return nil, err
	}

	cursor, err := hand.Colletion.Aggregate(context.TODO(), filterMap, &op)
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (hand colletion) ReadMany(filter string, optionQuery string) ([]map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}

	op, err := option.find(optionQuery)
	if err != nil {
		return nil, err
	}

	cursor, err := hand.Colletion.Find(context.TODO(), filterMap, &op)
	if err != nil {
		return nil, err
	}
	var results []map[string]interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (hand colletion) ReadOne(filter string, optionQuery string) (map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}

	op, err := option.findOne(optionQuery)
	if err != nil {
		return nil, err
	}

	var results map[string]interface{}
	err = hand.Colletion.FindOne(context.TODO(), filterMap, &op).Decode(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
