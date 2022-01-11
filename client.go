package mongoleaf

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (branch Branch) Connect() error {
	return branch.Client.Connect(context.TODO())
}

func (branch Branch) Disconnect() error {
	return branch.Client.Disconnect(context.TODO())
}
func (branch Branch) ListDatabaseNames(filter, optionQuery string) ([]string, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	op, err := option.listDatabaseNames(optionQuery)
	if err != nil {
		return nil, err
	}
	return branch.Client.ListDatabaseNames(context.TODO(), filterMap, &op)
}
func (branch Branch) ListDatabases(filter, optionQuery string) (map[string]interface{}, error) {
	filterMap, err := jsonToMap(filter)
	if err != nil {
		return nil, err
	}
	op, err := option.listDatabaseNames(optionQuery)
	if err != nil {
		return nil, err
	}
	res, err := branch.Client.ListDatabases(context.TODO(), filterMap, &op)
	if err != nil {
		return nil, err
	}
	return structToMap(res)
}

func (branch Branch) Ping() error {
	return branch.Client.Ping(context.TODO(), readpref.Primary())
}
