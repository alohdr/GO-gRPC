package postgres

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"gitlab.com/learnProto/pkg/v1/utils/constants"
	"gitlab.com/learnProto/pkg/v1/utils/converter"
)

type CustomMain struct {
	UserId      string         `db:"user_id,omitempty"`
	Pass        string         `db:"pass,omitempty"`
	DelFlag     bool           `db:"del_flag,omitempty"`
	Description sql.NullString `db:"description,omitempty"`
	CreId       string         `db:"cre_id,omitempty"`
	CreTime     time.Time      `db:"cre_time,omitempty"`
	ModId       string         `db:"mod_id,omitempty"`
	ModTime     time.Time      `db:"mod_time,omitempty"`
	ModTs       int            `db:"mod_ts,omitempty"`
}

type Registrasi struct {
	UserId string `db:"user_id"`
	Pass   string `db:"pass"`
}

type Id struct {
	UserId string `db: "pass"`
}

func (c *Conn) CustomMainSelect(filter *CustomMain) ([]*CustomMain, error) {
	// create query syntax
	qsyntax := fmt.Sprintf(`SELECT * FROM %s WHERE`, constants.Table_Custom_Main)

	// get filter struct metadata
	fil := reflect.ValueOf(*filter)

	for i := 0; i < fil.NumField(); i++ {
		if !fil.Field(i).IsZero() {
			qsyntax = fmt.Sprintf(`%s %s = '%v' AND`, qsyntax, converter.CamelToSnakeCase(fil.Type().Field(i).Name), fil.Field(i))
		}
	}

	qsyntax = strings.TrimRight(qsyntax, "AND")
	qsyntax = fmt.Sprintf(`%s;`, qsyntax)

	// create query syntax
	db, err := sql.Open(POSTGRES, c.Conn)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(qsyntax)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []*CustomMain

	//Fetch data to struct
	for rows.Next() {
		var rowsScan CustomMain
		err := rows.Scan(&rowsScan.UserId, &rowsScan.Pass, &rowsScan.DelFlag, &rowsScan.Description,
			&rowsScan.CreId, &rowsScan.CreTime, &rowsScan.ModId, &rowsScan.ModTime, &rowsScan.ModTs)

		if err != nil {
			return nil, err
		}

		// Append for every next row
		rowsScanArr = append(rowsScanArr, &rowsScan)
	}

	return rowsScanArr, nil
}

func (c *Conn) CustomMainRegister(input *Registrasi) (*CustomMain, error) {
	// create query syntax
	qsyntax := fmt.Sprintf(`INSERT  INTO %s (user_id, pass) VALUES($1, $2)`, constants.Table_Custom_Main)

	qsyntax = fmt.Sprintf(`%s;`, qsyntax)

	// create query syntax
	db, err := sql.Open(POSTGRES, c.Conn)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(qsyntax, input.UserId, input.Pass)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return &CustomMain{
		UserId: input.UserId,
		Pass:   input.Pass,
	}, nil
}

func (c *Conn) CustomMainDeleteRegister(id *Id) (*Id, error) {
	qsyntax := fmt.Sprintf(`DELETE  FROM %s WHERE %s = $1`, constants.Table_Custom_Main, constants.Field_User_Id)

	// create query syntax
	db, err := sql.Open(POSTGRES, c.Conn)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(qsyntax, id.UserId)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return &Id{
		UserId: id.UserId,
	}, nil
}
