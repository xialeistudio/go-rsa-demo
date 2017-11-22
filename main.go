package main

import (
	"encoding/pem"
	"os"
	"io/ioutil"
	"errors"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"log"
)

func encrypt(data []byte) ([]byte, error) {
	fp, err := os.Open("public.key")
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	buf, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.New("load public key failed")
	}
	pubkey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pubkey.(*rsa.PublicKey), data)
}

func decrypt(data []byte) ([]byte, error) {
	fp, err := os.Open("privacy.key")
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	buf, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.New("load privacy key failed")
	}
	prikey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, prikey, data)
}

var logger = log.New(os.Stdout, "main ", log.LstdFlags|log.Lshortfile)

func main() {
	ret, err := encrypt([]byte("Hello World"))
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Println("Encrypt success")
	ret, err = decrypt(ret)
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Println("Decrypt success", string(ret))
}
