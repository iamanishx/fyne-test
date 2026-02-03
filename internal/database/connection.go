package database

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sql.DB
}

func NewDB() (*DB, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	appDir := filepath.Join(home, ".passvault")
	if err := os.MkdirAll(appDir, 0700); err != nil {
		return nil, err
	}

	dbPath := filepath.Join(appDir, "vault.db")
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

func (db *DB) InitSchema() error {
	query := `
	PRAGMA foreign_keys = ON;
	CREATE TABLE IF NOT EXISTS secrets (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		format TEXT,
		content TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	CREATE TABLE IF NOT EXISTS fields (
		id TEXT,
		secret_id TEXT,
		key TEXT,
		value BLOB,
		is_sensitive BOOLEAN,
		FOREIGN KEY(secret_id) REFERENCES secrets(id) ON DELETE CASCADE
	);
	`
	_, err := db.conn.Exec(query)
	if err != nil {
		return err
	}

	return db.ensureSecretColumns()
}

func (db *DB) ensureSecretColumns() error {
	rows, err := db.conn.Query("PRAGMA table_info(secrets)")
	if err != nil {
		return err
	}
	defer rows.Close()

	cols := map[string]bool{}
	for rows.Next() {
		var cid int
		var name, ctype string
		var notnull int
		var dflt interface{}
		var pk int
		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dflt, &pk); err != nil {
			return err
		}
		cols[name] = true
	}

	if !cols["format"] {
		if _, err := db.conn.Exec("ALTER TABLE secrets ADD COLUMN format TEXT"); err != nil {
			return err
		}
	}
	if !cols["content"] {
		if _, err := db.conn.Exec("ALTER TABLE secrets ADD COLUMN content TEXT"); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}
