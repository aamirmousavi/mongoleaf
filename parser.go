package mongoleaf

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func JSON(v interface{}) string {
	json, _ := interToJson(v)

	return json
}
func JSONPretty(v interface{}) string {
	_json, _ := json.MarshalIndent(v, " ", "	")
	return string(_json)
}
func jsonToMap(str string) (map[string]interface{}, error) {
	var _map map[string]interface{}
	if str == "" || str == "{}" {
		return _map, nil
	}
	err := json.Unmarshal([]byte(str), &_map)
	if err != nil {
		return nil, err
	}
	theMap, err := convertId(_map)
	if err != nil {
		return nil, err
	}
	return theMap, nil
}
func structToMap(theStruct interface{}) (map[string]interface{}, error) {
	_json, err := json.Marshal(theStruct)
	if err != nil {
		return nil, err
	}
	var theMap map[string]interface{}
	err = json.Unmarshal(_json, &theMap)
	if err != nil {
		return nil, err
	}
	return theMap, nil
}
func interToJson(_map interface{}) (string, error) {
	json, err := json.Marshal(_map)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func jsonToArrayMap(str string) ([]map[string]interface{}, error) {
	var _arrMap []map[string]interface{}
	err := json.Unmarshal([]byte(str), &_arrMap)
	if err != nil {
		return nil, err
	}
	for i := range _arrMap {
		_arrMap[i], err = convertId(_arrMap[i])
		if err != nil {
			return nil, err
		}
	}
	return _arrMap, nil
}
func convertIdToObject(value interface{}) (primitive.ObjectID, error) {
	if strHex, ok := value.(string); ok {
		obj, err := primitive.ObjectIDFromHex(strHex)
		if err != nil {
			return primitive.ObjectID{}, err
		}
		return obj, nil
	} else {
		return primitive.ObjectID{}, fmt.Errorf("your objectId is invaild")
	}
}

func convertId(theMap map[string]interface{}) (map[string]interface{}, error) {
	var err error
	for key, _val := range theMap {
		if _, valIsStr := _val.(string); valIsStr {
			if key == "_id" {
				theMap[key], err = convertIdToObject(_val)
				if err != nil {
					return nil, err
				}
			}
		} else {

			if v, ok := _val.(map[string]interface{}); ok {
				theMap[key], err = convertId(v)
				if err != nil {
					return nil, err
				}
			} else if v_a, ok := _val.([]interface{}); ok {
				for index, item := range v_a {
					test, ok := item.(map[string]interface{})
					if ok {
						v_a[index], err = convertId(test)
						if err != nil {
							return nil, err
						}
						theMap[key] = v_a
					}
				}
			}
		}
	}
	return theMap, nil
}
