package main


import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)


func Extract(source DataSourceConfig, query string) (float64, error){

    result := 0.0

    // TODO cache db connections
    db, err := sql.Open("postgres", fmt.Sprintf("user=%v dbname=%v password=%v host=%v port=%v", source.User, source.Database, source.Password, source.Host, source.Port))

    if err != nil {
        return result, err
    }

    rows, err := db.Query(query)

    if err != nil {
        return result, err
    }

    if rows.Next() {
        rows.Scan(&result)
    }

    return result, nil
}
