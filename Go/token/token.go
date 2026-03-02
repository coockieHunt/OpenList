package token

import (
	"crypto/rand"
	"encoding/hex"
)

//gerate rand token
func genToken() string {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
		return ""
	}
    
    return hex.EncodeToString(bytes)
}