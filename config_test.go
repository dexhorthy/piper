package main

import "fmt"

func ExampleLoadConfig() {
    config := LoadConfig("piper.yml")


    for _, pipeConfig := range config {
        fmt.Printf("Loaded Postgres Host: %v\n", pipeConfig.Source.Host)
        fmt.Printf("Loaded Postgres Port: %v\n", pipeConfig.Source.Port)
        fmt.Printf("Loaded Postgres Database: %v\n", pipeConfig.Source.Database)
        fmt.Printf("Loaded Postgres User: %v\n", pipeConfig.Source.User)
        fmt.Printf("Loaded Postgres Password: %v\n", pipeConfig.Source.Password)
        fmt.Printf("Loaded Query: %v\n", pipeConfig.Query)
        fmt.Printf("Loaded Dest: %v\n", pipeConfig.Dest)
        fmt.Printf("Loaded Graphite Host: %v\n", pipeConfig.Graphite.Host)
        fmt.Printf("Loaded Graphite Port: %v\n", pipeConfig.Graphite.Port)
        fmt.Printf("\n")
    }


    // Output:
    // Loaded Postgres Host: localhost
    // Loaded Postgres Port: 5432
    // Loaded Postgres Database: postgres
    // Loaded Postgres User: postgres
    // Loaded Postgres Password: postgres
    // Loaded Query: SELECT count(*) FROM bar
    // Loaded Dest: piper.bar.count
    // Loaded Graphite Host: localhost
    // Loaded Graphite Port: 2003
    //
    // Loaded Postgres Host: localhost
    // Loaded Postgres Port: 5432
    // Loaded Postgres Database: postgres
    // Loaded Postgres User: postgres
    // Loaded Postgres Password: postgres
    // Loaded Query: SELECT id FROM bar order by id desc limit 1
    // Loaded Dest: piper.bar.greatest_id
    // Loaded Graphite Host: localhost
    // Loaded Graphite Port: 2003

}
