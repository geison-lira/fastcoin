package handlers

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
)

func AddressGenerate(pubKey *ecdsa.PublicKey) (string, error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	pubHash := sha256.Sum256(pubBytes)
	return hex.EncodeToString(pubHash[:16]), nil
}
