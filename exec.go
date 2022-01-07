package mongoleaf

import "context"

func (hand colletion) UpdateMany(filter string, update string, optionQuery string) (map[string]interface{}, error) {
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
	results, err := structToMap(res)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (hand colletion) UpdateOne(filter string, update string, optionQuery string) (map[string]interface{}, error) {
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
	results, err := structToMap(res)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (hand colletion) DeleteMany(filter string, optionQuery string) (map[string]interface{}, error) {
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
	results, err := structToMap(res)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (hand colletion) DeleteOne(filter string, optionQuery string) (map[string]interface{}, error) {
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
	results, err := structToMap(res)
	if err != nil {
		return nil, err
	}
	return results, nil
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
	results, err := structToMap(res)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (hand colletion) InsertOne(document string, optionQuery string) (map[string]interface{}, error) {
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
	results, err := structToMap(res)
	if err != nil {
		return nil, err
	}
	return results, nil
}
