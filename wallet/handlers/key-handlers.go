package handlers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"os"
)

func AddressGenerate(pubKey *ecdsa.PublicKey) (string, error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	pubHash := sha256.Sum256(pubBytes)
	return hex.EncodeToString(pubHash[:16]), nil
}

func KeysGenerate() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pubKey := &privKey.PublicKey
	return privKey, pubKey, nil
}

func KeysStore(keysPath string, privKey *ecdsa.PrivateKey, pubKey *ecdsa.PublicKey) error {
	privBytes, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		return err
	}
	pubBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(keysPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privBytes,
	})
	if err != nil {
		return err
	}
	err = pem.Encode(file, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})
	if err != nil {
		return err
	}
	return nil
}

func KeysLoad(keysPath string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	data, err := os.ReadFile(keysPath)
	if err != nil {
		return nil, nil, err
	}
	var privKey *ecdsa.PrivateKey
	var pubKey *ecdsa.PublicKey
	for len(data) > 0 {
		var block *pem.Block
		block, data = pem.Decode(data)
		if block == nil {
			break
		}
		switch block.Type {
		case "EC PRIVATE KEY":
			privKey, err = x509.ParseECPrivateKey(block.Bytes)
			if err != nil {
				return nil, nil, err
			}
		case "PUBLIC KEY":
			parsedPubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return nil, nil, err
			}
			var ok bool
			pubKey, ok = parsedPubKey.(*ecdsa.PublicKey)
			if !ok {
				return nil, nil, errors.New("public key is not ECDSA")
			}
		}
	}
	if privKey == nil || pubKey == nil {
		return nil, nil, errors.New("file missing required keys")
	}
	return privKey, pubKey, nil
}
