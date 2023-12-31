package digital_signatures

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func SignDataWithRSA(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hashed := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func VerifySignatureWithRSA(data, signature []byte, publicKey *rsa.PublicKey) error {
	hashed := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return err
	}
	return nil
}

func ExportRSAPrivateKeyToPEM(privateKey *rsa.PrivateKey) (string, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	return string(pem.EncodeToMemory(pemBlock)), nil
}

func ExportRSAPublicKeyToPEM(publicKey *rsa.PublicKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	pemBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	return string(pem.EncodeToMemory(pemBlock)), nil
}

func ImportRSAPrivateKeyFromPEM(keyPEM string) (*rsa.PrivateKey, error) {
	pemBlock, _ := pem.Decode([]byte(keyPEM))
	if pemBlock == nil {
		return nil, errors.New("invalid PEM block")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func ImportRSAPublicKeyFromPEM(keyPEM string) (*rsa.PublicKey, error) {
	pemBlock, _ := pem.Decode([]byte(keyPEM))
	if pemBlock == nil {
		return nil, errors.New("invalid PEM block")
	}
	publicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}
	return rsaPublicKey, nil
}
