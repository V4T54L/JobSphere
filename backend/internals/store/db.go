package store

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", os.Getenv("DB_PATH"))

	if err != nil {
		return nil, err
	}

	err = createTable(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   username TEXT NOT NULL UNIQUE,
	   password TEXT NOT NULL,
	   email TEXT NOT NULL UNIQUE,
	   created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	   updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	   is_admin BOOLEAN DEFAULT 0,
	   profile_picture TEXT
   );
    
   CREATE TABLE IF NOT EXISTS jobs (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   title TEXT NOT NULL,
	   description TEXT NOT NULL,
	   company TEXT NOT NULL,
	   location TEXT NOT NULL,
	   salary TEXT NOT NULL,
	   user_id INTEGER NOT NULL,
	   created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	   FOREIGN KEY (user_id) REFERENCES users(id)
	);

	INSERT OR IGNORE INTO users (username, password, email, is_admin) VALUES('admin', '$2a$10$yCCIwf.RIsZJtAfgmyji5u0nUK0lfm8oHXqZuaG9XrYTUVnWrcmyi', 'admin@test.com', true)
   `)
	// Admin password : Test@123

	return err
}