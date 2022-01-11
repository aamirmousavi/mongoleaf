package mongoleaf

import (
	"context"
)

func (db dataBase) Aggregate(pipline, optionQuery string) ([]map[string]interface{}, error) {
	piplineMap, err := jsonToArrayMap(pipline)
	if err != nil {
		return nil, err
	}
	op, err := option.aggregate(optionQuery)
	if err != nil {
		return nil, err
	}
	cursor, err := db.DB.Aggregate(context.TODO(), piplineMap, &op)
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (db dataBase) Drop() error {
	return db.DB.Drop(context.TODO())
}

func (db dataBase) CreateCollection(name, optionQuery string) error {
	op, err := option.createCollection(optionQuery)
	if err != nil {
		return err
	}
	return db.DB.CreateCollection(context.TODO(), name, &op)
}

func (db dataBase) CreateView(viewName, viewOn, pipline, optionQuery string) error {
	piplineMap, err := jsonToArrayMap(pipline)
	if err != nil {
		return err
	}
	op, err := option.createView(optionQuery)
	if err != nil {
		return err
	}
	return db.DB.CreateView(context.TODO(), viewName, viewOn, piplineMap, &op)
}

func (db dataBase) ListCollectionNames(filter, optionQuery string) ([]string, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	op, err := option.listCollections(optionQuery)
	if err != nil {
		return nil, err
	}
	return db.DB.ListCollectionNames(context.TODO(), filterMap, &op)
}

func (db dataBase) ListCollections(filter, optionQuery string) ([]map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	op, err := option.listCollections(optionQuery)
	if err != nil {
		return nil, err
	}
	cursor, err := db.DB.ListCollections(context.TODO(), filterMap, &op)
	if err != nil {
		return nil, err
	}
	var results []map[string]interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
func (db dataBase) Name() string {
	return db.DB.Name()
}

func (db dataBase) RunCommand(cmd, optionQuery string) (map[string]interface{}, error) {
	cmdMap, err := jsonToMap(cmd)
	if err != nil {
		return nil, err
	}
	op, err := option.runCmd(optionQuery)
	if err != nil {
		return nil, err
	}
	res := db.DB.RunCommand(context.TODO(), cmdMap, &op)
	return structToMap(res)
}

func (db dataBase) RunCommandCursor(cmd, optionQuery string) ([]map[string]interface{}, error) {
	cmdMap, err := jsonToMap(cmd)
	if err != nil {
		return nil, err
	}
	op, err := option.runCmd(optionQuery)
	if err != nil {
		return nil, err
	}
	cursor, err := db.DB.RunCommandCursor(context.TODO(), cmdMap, &op)
	if err != nil {
		return nil, err
	}
	var results []map[string]interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
