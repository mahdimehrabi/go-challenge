package infrastracture

import (
	"challange/app/interfaces"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type PgxDB struct {
	logger interfaces.Logger
	Conn   *pgx.Conn
}

func NewPgxDB(logger interfaces.Logger) *PgxDB {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DBUsername"), os.Getenv("DBPassword"),
		os.Getenv("DBHost"), os.Getenv("DBPort"),
		os.Getenv("DBName"),
	)
	conn, err := pgx.Connect(context.TODO(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &PgxDB{
		logger: logger,
		Conn:   conn,
	}
}

//execute sql command and return rows affected count and err
func (db *PgxDB) Exec(
	ctx context.Context,
	query string,
	parameters []interface{}) (int64, error) {
	cmdTag, err := db.Conn.Exec(ctx, query, parameters...)
	return cmdTag.RowsAffected(), err
}
