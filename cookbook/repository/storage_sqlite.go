package repository

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

db, err := sql.Open()