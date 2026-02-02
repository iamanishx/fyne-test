package state

import (
	"sync"

	"passvault-fyne/internal/crypto"
)

type SecureState struct {
	mu         sync.RWMutex
	masterKey  []byte
	isUnlocked bool
}

func NewSecureState() *SecureState {
	return &SecureState{}
}

func (s *SecureState) SetMasterKey(key []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.masterKey = make([]byte, len(key))
	copy(s.masterKey, key)
	s.isUnlocked = true
}

func (s *SecureState) GetMasterKey() []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if !s.isUnlocked {
		return nil
	}
	key := make([]byte, len(s.masterKey))
	copy(key, s.masterKey)
	return key
}

func (s *SecureState) IsUnlocked() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isUnlocked
}

func (s *SecureState) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.masterKey != nil {
		for i := range s.masterKey {
			s.masterKey[i] = 0
		}
		s.masterKey = nil
	}
	s.isUnlocked = false
}

func (s *SecureState) Encrypt(plaintext []byte) ([]byte, []byte, error) {
	s.mu.RLock()
	key := s.masterKey
	s.mu.RUnlock()

	if key == nil {
		return nil, nil, nil
	}

	return crypto.Encrypt(plaintext, key)
}

func (s *SecureState) Decrypt(ciphertext, nonce []byte) ([]byte, error) {
	s.mu.RLock()
	key := s.masterKey
	s.mu.RUnlock()

	if key == nil {
		return nil, nil
	}

	return crypto.Decrypt(ciphertext, nonce, key)
}
