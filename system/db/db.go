package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	if _, err := os.Stat("./falconism.db"); os.IsNotExist(err) {
		os.Create("./falconism.db")
	}

	db, err := sql.Open("sqlite3", "./falconism.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
	   "post_datetime"  REAL,
	   "post_title"  TEXT,
	   "post_slug"  TEXT,
	   "post_content"  TEXT,
	   "ping_status"  INTEGER,
	   "to_ping"  INTEGER DEFAULT 1,
	   "sitemap_status"  INTEGER,
	   "to_sitemap"  INTEGER DEFAULT 1,
	   "meta_description"  TEXT
	);`)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS config (
		"parameter"  TEXT,
		"value"  TEXT
		);
	`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.Close()
}
