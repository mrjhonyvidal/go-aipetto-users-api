package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
)

func GetTokenMd5(text string) string {
	hasher := md5.New()
	defer hasher.Reset()
	hasher.Write([]byte(text))
	hash := hex.EncodeToString(hasher.Sum(nil))

	u := uuid.New()
	token := hash + u.String()

	return token
}