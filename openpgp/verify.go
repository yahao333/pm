package openpgp

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/openpgp"

	"github.com/yahao333/pm/config"
)

func Verify(packagePath string) bool {
	keyRingFile, err := os.Open(config.Conf.KeyPrintPath)
	if err != nil {
		log.Println(err)
		return false
	}
	defer keyRingFile.Close()
	keyRing, err := openpgp.ReadArmoredKeyRing(keyRingFile)
	if err != nil {
		log.Println(err)
		return false
	}
	// Check package signature.
	packageSignatureFileName := fmt.Sprintf("%s.asc", packagePath)
	isValid := checkArmoredDetachedSignature(keyRing, packagePath, packageSignatureFileName)
	if !isValid {
		return false
	}
	// Check sha256 signature.
	sha256FileName := fmt.Sprintf("%s.sha256", packagePath)
	sha256SignatureFileName := fmt.Sprintf("%s.asc", sha256FileName)
	isValid = checkArmoredDetachedSignature(keyRing, sha256FileName, sha256SignatureFileName)
	return isValid
}

func checkArmoredDetachedSignature(keyRing openpgp.EntityList, signed, signature string) bool {
	signedFile, err := os.Open(signed)
	if err != nil {
		log.Println(err)
		return false
	}
	defer signedFile.Close()
	signatureFile, err := os.Open(signature)
	if err != nil {
		log.Println(err)
		return false
	}
	defer signatureFile.Close()
	_, err = openpgp.CheckArmoredDetachedSignature(keyRing, signedFile, signatureFile)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
