package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type RsaData struct {
	PrivateKey []byte
	PublicKey  []byte
}

func NewRsa() (RsaData, error) {
	data := RsaData{}
	generateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return data, err
	}
	private_key := pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(generateKey),
		Type:  "RSA PRIVATE KEY",
	})
	encode_public, err := x509.MarshalPKIXPublicKey(&generateKey.PublicKey)
	if err != nil {
		return data, err
	}
	public_key := pem.EncodeToMemory(&pem.Block{
		Bytes: encode_public,
		Type:  "PUBLIC KEY",
	})
	data.PrivateKey = private_key
	data.PublicKey = public_key
	return data, nil
}
func EncodeRSA(publicKey, origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}
func DecodeRSA(privateKey, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
