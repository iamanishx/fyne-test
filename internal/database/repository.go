package database

import (
	"database/sql"
)

func (db *DB) SaveSecret(secret *SecretEntry) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		"INSERT OR REPLACE INTO secrets (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)",
		secret.ID, secret.Name, secret.CreatedAt, secret.UpdatedAt,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM fields WHERE secret_id = ?", secret.ID)
	if err != nil {
		return err
	}

	for _, field := range secret.Fields {
		_, err = tx.Exec(
			"INSERT INTO fields (id, secret_id, key, value, is_sensitive) VALUES (?, ?, ?, ?, ?)",
			field.ID, secret.ID, field.Key, field.Value, field.IsSensitive,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (db *DB) GetSecrets() ([]SecretEntry, error) {
	rows, err := db.conn.Query("SELECT id, name, created_at, updated_at FROM secrets ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var secrets []SecretEntry
	for rows.Next() {
		var s SecretEntry
		if err := rows.Scan(&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		secrets = append(secrets, s)
	}
	return secrets, nil
}

func (db *DB) GetSecret(id string) (*SecretEntry, error) {
	var s SecretEntry
	err := db.conn.QueryRow(
		"SELECT id, name, created_at, updated_at FROM secrets WHERE id = ?", id,
	).Scan(&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	rows, err := db.conn.Query("SELECT id, key, value, is_sensitive FROM fields WHERE secret_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f Field
		if err := rows.Scan(&f.ID, &f.Key, &f.Value, &f.IsSensitive); err != nil {
			return nil, err
		}
		s.Fields = append(s.Fields, f)
	}

	return &s, nil
}

func (db *DB) DeleteSecret(id string) error {
	_, err := db.conn.Exec("DELETE FROM secrets WHERE id = ?", id)
	return err
}
