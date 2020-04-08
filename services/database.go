package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Exec(command string) error {
	yugabyteHost := viper.GetString(`YUGABYTEDB.HOST`)
	yugabytePort := viper.GetInt(`YUGABYTEDB.PORT`)
	yugabyteUser := viper.GetString(`YUGABYTEDB.USER`)
	yugabytePassword := viper.GetString(`YUGABYTEDB.PASSWORD`)
	yugabyteDBName := viper.GetString(`YUGABYTEDB.DBNAME`)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", yugabyteHost, yugabytePort, yugabyteUser, yugabytePassword, yugabyteDBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(command); err != nil {
		log.Fatal(err)

		return fmt.Errorf(err.Error())
	}
	return nil
}

func Query(command string) (interface{}, error) {
	yugabyteHost := viper.GetString(`YUGABYTEDB.HOST`)
	yugabytePort := viper.GetInt(`YUGABYTEDB.PORT`)
	yugabyteUser := viper.GetString(`YUGABYTEDB.USER`)
	yugabytePassword := viper.GetString(`YUGABYTEDB.PASSWORD`)
	yugabyteDBName := viper.GetString(`YUGABYTEDB.DBNAME`)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", yugabyteHost, yugabytePort, yugabyteUser, yugabytePassword, yugabyteDBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		print(err)
	}

	rows, err := db.Query(command)
	if err != nil {
		print(err)
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	json.Marshal(tableData)

	return tableData, nil
}
