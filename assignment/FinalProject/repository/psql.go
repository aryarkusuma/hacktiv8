package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DbInit(ctx context.Context) *pgx.Conn {
	// urlExample := "postgres://arya:arya@localhost:5432/shuttle"
	conn, err := pgx.Connect(ctx, "postgres://arya:arya@localhost:5432/shuttling")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	var name string
	err = conn.QueryRow(ctx, "select shuttle_type from shuttles where id=$1", "bed57f59-5022-40a0-8b1c-afdd6352e390").Scan(&name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name)

	return conn
}
