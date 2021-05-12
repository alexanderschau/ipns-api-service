package api

import (
	"crypto/rand"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
)

func GenKey(contentHash string) (*shell.Key, error) {
	token := randToken()
	key, err := s.KeyGen(ctx, token)
	if err != nil {
		return nil, err
	}

	err = UpdateKey(token, contentHash)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func randToken() string {
	b := make([]byte, 10)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
