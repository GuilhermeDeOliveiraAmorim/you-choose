package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
)

func HashEmailWithHMAC(email string) (string, []exceptions.ProblemDetails) {
	key := []byte(config.SECRETS_VAR.JWT_SECRET)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(email))

	return hex.EncodeToString(h.Sum(nil)), nil
}
