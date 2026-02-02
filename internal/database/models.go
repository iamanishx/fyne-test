package database

import "time"

type SecretEntry struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Fields    []Field   `json:"fields"`
}

type Field struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Value       []byte `json:"value"`
	IsSensitive bool   `json:"is_sensitive"`
}

type EncryptedData struct {
	Ciphertext []byte `json:"ciphertext"`
	Nonce      []byte `json:"nonce"`
}
