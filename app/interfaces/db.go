package interfaces

import "context"

type DB interface {
	Query(ctx context.Context, query string, scans []interface{})
}
