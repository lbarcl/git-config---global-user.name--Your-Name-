package config

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
)

type RSAKeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  []byte
}

var encryptionInstance *RSAKeyPair

func GetEncryption() RSAKeyPair {
	if encryptionInstance == nil {
		fmt.Println("Generating RSA key pair")

		var instance RSAKeyPair

		privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			log.Fatal(err)
		}
		instance.PrivateKey = privateKey

		publicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
		if err != nil {
			log.Fatal(err)
		}
		instance.PublicKey = publicKey

		encryptionInstance = &instance
		return instance

	} else {
		fmt.Println("Returning existing RSA key pair")
		return *encryptionInstance
	}

}

func GetRandomBytes(length int) []byte {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
