package db

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
    "github.com/sirupsen/logrus"
)

// Connect initializes a connection to the PostgreSQL database.
func Connect() (*sql.DB, error) {
//    host := "127.0.0.1";
//    port := "5432";
//    user := "newuser";
//    password := "newpassword";
//    dbname := "go-dev";
    
//    dsn := fmt.Sprintf(
//        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
//        host,
//        port,
//        user,
//        password,
//        dbname,
//    )
    dsn := "user=postgres.maqaoisdvzuvcgxaxkpi password=Winder#0925@ host=aws-0-us-west-1.pooler.supabase.com port=6543 dbname=postgres";
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to open database connection: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    logrus.Info("Connected to PostgreSQL database successfully")
    return db, nil
}

// Migrate runs the database migrations to ensure the database schema is up-to-date.
func Migrate(db *sql.DB) error {
    // Define the schema (you may use a more sophisticated migration tool)
    schema := `
        CREATE TABLE IF NOT EXISTS queries (
            id SERIAL PRIMARY KEY,
            domain VARCHAR(255) NOT NULL,
            result TEXT NOT NULL,
            queried_at TIMESTAMP NOT NULL
        );
    `

    _, err := db.Exec(schema)
    if err != nil {
        return fmt.Errorf("failed to run database migrations: %w", err)
    }

    logrus.Info("Database migrations completed successfully")
    return nil
}
