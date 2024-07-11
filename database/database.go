package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Ed1123/todo/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func readEnvVar(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		log.Fatalf("%s not found in .env file", key)
	}
	return value
}

func Connect() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := readEnvVar("DATABASE_HOST")
	dbName := readEnvVar("DATABASE_NAME")
	user := readEnvVar("DATABASE_USER")
	password := readEnvVar("DATABASE_PASSWORD")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		host, user, password, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func CreateTask(db *sql.DB, task models.Task, todo models.ToDo) error {
	sqlStatement := `
		INSERT INTO tasks
		name, description, date_created, date_updated, todo_id
		VALUES ($1, $2, NOW(), NOW(), $3)`
	_, err := db.Exec(sqlStatement, task.Name, task.Description, todo.Id)
	return err
}
