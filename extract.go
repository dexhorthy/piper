package main


import (
    "log"
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)


func Extract(source DataSourceConfig, query string) (float64, error){

    result := 0.0

    db, err := sql.Open(source.Database, fmt.Sprintf("user=%v dbname=%v password=%v", source.User, source.Database, source.Password))

    if err != nil {
        log.Fatal(err)
        return result, err
    }

    rows, err := db.Query(query)

    if err != nil {
        log.Fatal(err)
        return result, err
    }


    if rows.Next() {
        rows.Scan(&result)
    }

    return result, nil
}
