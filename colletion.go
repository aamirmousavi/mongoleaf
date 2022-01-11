package mongoleaf

import "context"

func (hand colletion) Aggregate(pipline, optionQuery string) ([]map[string]interface{}, error) {
	piplineMap, err := jsonToArrayMap(pipline)
	if err != nil {
		return nil, err
	}

	op, err := option.aggregate(optionQuery)
	if err != nil {
		return nil, err
	}

	cursor, err := hand.Colletion.Aggregate(context.TODO(), piplineMap, &op)
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (hand colletion) Find(filter, optionQuery string) ([]map[string]interface{}, error) {
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

func (hand colletion) FindOne(filter, optionQuery string) (map[string]interface{}, error) {
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

func (hand colletion) UpdateMany(filter, update, optionQuery string) (map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	updateMap, err := jsonToMap(update)
	if err != nil {
		return nil, err
	}
	op, err := option.update(optionQuery)
	if err != nil {
		return nil, err
	}
	res, err := hand.Colletion.UpdateMany(context.TODO(), filterMap, updateMap, &op)
	if err != nil {
		return nil, err
	}
	return structToMap(res)
}

func (hand colletion) UpdateOne(filter, update, optionQuery string) (map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	updateMap, err := jsonToMap(update)
	if err != nil {
		return nil, err
	}
	op, err := option.update(optionQuery)
	if err != nil {
		return nil, err
	}
	res, err := hand.Colletion.UpdateOne(context.TODO(), filterMap, updateMap, &op)
	if err != nil {
		return nil, err
	}
	return structToMap(res)
}

func (hand colletion) DeleteMany(filter, optionQuery string) (map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	op, err := option.delete(optionQuery)
	if err != nil {
		return nil, err
	}
	res, err := hand.Colletion.DeleteMany(context.TODO(), filterMap, &op)
	if err != nil {
		return nil, err
	}
	return structToMap(res)
}

func (hand colletion) DeleteOne(filter, optionQuery string) (map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	op, err := option.delete(optionQuery)
	if err != nil {
		return nil, err
	}
	res, err := hand.Colletion.DeleteOne(context.TODO(), filterMap, &op)
	if err != nil {
		return nil, err
	}
	return structToMap(res)
}

func (hand colletion) InsertMany(optionQuery string, documents ...string) (map[string]interface{}, error) {
	records := make([]interface{}, 0)
	for i := range documents {
		d, err := jsonToMap(documents[i])
		if err != nil {
			return nil, err
		}
		records = append(records, d)
	}
	op, err := option.insertMany(optionQuery)
	if err != nil {
		return nil, err
	}
	res, err := hand.Colletion.InsertMany(context.TODO(), records, &op)
	if err != nil {
		return nil, err
	}
	return structToMap(res)
}

func (hand colletion) InsertOne(optionQuery, document string) (map[string]interface{}, error) {
	documentMap, err := jsonToMap(document)
	if err != nil {
		return nil, err
	}
	op, err := option.insertOne(optionQuery)
	if err != nil {
		return nil, err
	}

	res, err := hand.Colletion.InsertOne(context.TODO(), documentMap, &op)
	if err != nil {
		return nil, err
	}
	return structToMap(res)
}

func (hand colletion) CountDocuments(filter, optionQuery string) (int64, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return 0, err
	}
	op, err := option.countDocuments(optionQuery)
	if err != nil {
		return 0, err
	}
	return hand.Colletion.CountDocuments(context.TODO(), filterMap, &op)
}

func (hand colletion) Distinct(fieldName, optionQuery string) ([]interface{}, error) {
	op, err := option.distinct(optionQuery)
	if err != nil {
		return nil, err
	}
	return hand.Colletion.Distinct(context.TODO(), fieldName, &op)
}

func (hand colletion) FindOneAndUpdateDistinct(filter, update, optionQuery string) (map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	updateMap, err := jsonToMap(update)
	if err != nil {
		return nil, err
	}
	res := hand.Colletion.FindOneAndUpdate(context.TODO(), filterMap, updateMap)

	return structToMap(res)
}

func (hand colletion) FindOneAndReplace(filter, replacement, optionQuery string) (map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	replaceObj, err := jsonToMap(replacement)
	if err != nil {
		return nil, err
	}
	op, err := option.findOneAndReplace(optionQuery)
	if err != nil {
		return nil, err
	}
	res := hand.Colletion.FindOneAndReplace(context.TODO(), filterMap, replaceObj, &op)

	return structToMap(res)
}

func (hand colletion) FindOneAndDelete(filter, optionQuery string) (map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	op, err := option.findOneAndDelete(optionQuery)
	if err != nil {
		return nil, err
	}
	res := hand.Colletion.FindOneAndDelete(context.TODO(), filterMap, &op)
	return structToMap(res)
}

func (hand colletion) Drop() error {
	return hand.Colletion.Drop(context.TODO())
}

func (hand colletion) EstimatedDocumentCount() (int64, error) {
	return hand.Colletion.EstimatedDocumentCount(context.TODO())
}
