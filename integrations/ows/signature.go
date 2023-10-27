package ows

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"

	"github.com/DanielDDHM/world-coin-converter/config"
)

func SignRequest(path string, nonce string, body string) string {
	mac := hmac.New(sha512.New, []byte(config.GetConfig().OwsSecret))
	mac.Write([]byte(config.GetConfig().OwsKey))
	mac.Write([]byte(nonce))
	mac.Write([]byte(path))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}
