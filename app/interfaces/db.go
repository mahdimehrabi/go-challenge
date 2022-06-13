package interfaces

import "context"

type Model interface {
	SliceToModel(data []interface{}) error
}

type DB interface {
	//for execution of query
	Exec(
		ctx context.Context,
		query string,
		parameters []interface{}) error

	//get signle row
	QueryRow(
		ctx context.Context,
		query string,
		parameters []interface{},
		scans []interface{}) error

	//get multiple rows
	Query(ctx context.Context,
		query string,
		parameters []interface{},
		model *Model,
	) error
}
