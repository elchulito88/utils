// This shows an example of how to generate a SSH RSA Private/Public key pair and save it locally

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

// Paths for the CreateSSHKey function
type Paths struct {
	Path string
}

// CreateSSHKey is used to create SSH Keys
func (p Paths) CreateSSHKey(src string) {

	savePrivateFileTo := "./" + p.Path
	savePublicFileTo := "./" + p.Path + ".pub"

	bitSize := 4096

	privateKey, err := GeneratePrivateKey(bitSize)
	if err != nil {
		log.Fatal(err.Error())
	}

	publicKeyBytes, err := GeneratePublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	privateKeyBytes := EncodePrivateKeyToPEM(privateKey)

	err = WriteKeyToFile(privateKeyBytes, savePrivateFileTo)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = WriteKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GeneratePrivateKey creates a RSA Private Key of specified byte size
func GeneratePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	log.Println("Private Key generated")
	return privateKey, nil
}

// EncodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func EncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

// GeneratePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func GeneratePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	log.Println("Public key generated")
	return pubKeyBytes, nil
}

// WriteKeyToFile writes keys to a file
func WriteKeyToFile(keyBytes []byte, saveFileTo string) error {
	err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}

	log.Printf("Key saved to: %s", saveFileTo)
	return nil
}
