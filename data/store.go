package data

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
)

type UrlStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewUrlStore() *UrlStore {
	return &UrlStore{
		urls: make(map[string]string),
	}
}

func (s *UrlStore) Get(shortURL string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	longURL, exists := s.urls[shortURL]
	return longURL, exists
}

func (s *UrlStore) Set(longURL string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	shortURL := generateShortURL(longURL)
	s.urls[shortURL] = longURL
	return shortURL
}

func generateShortURL(longURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(longURL))
	return hex.EncodeToString(hasher.Sum(nil))[:6]
}
