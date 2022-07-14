package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gitlab.com/learnProto/pkg/v1/config"
	"gitlab.com/learnProto/pkg/v1/utils/constants"
)

const (
	POSTGRES string = "postgres"
)

type Conn struct {
	Conn string
}

// New initializes the postgres writers
func New(config *config.Config) (*Conn, error) {
	// Set postgres och config
	conn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Pass,
		config.Postgres.Dbname)

	db, err := sql.Open(POSTGRES, conn)
	if err != nil {
		return nil, err
	}

	// Test postgres OCH connection
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %v limit 1", constants.Table_Custom_Main))
	defer db.Close()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return &Conn{
		Conn: conn,
	}, nil
}
