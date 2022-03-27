package interop

import (
	"crypto/rand"

	tj "github.com/Hyperledger-TWGC/tjfoc-gm/sm2"
	tjx509 "github.com/Hyperledger-TWGC/tjfoc-gm/x509"
)

type TJSM2 struct {
	PrivateKey *tj.PrivateKey
	PublicKey  *tj.PublicKey
}

func NewTJSM2() (*TJSM2, error) {
	PrivateKey, err := tj.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &TJSM2{PrivateKey: PrivateKey, PublicKey: &PrivateKey.PublicKey}, nil
}

func TJImportKey(privPEM []byte, pubPEM []byte) (*TJSM2, error) {
	PrivateKey, err := tjx509.ReadPrivateKeyFromPem(privPEM, nil)
	if err != nil {
		return nil, err
	}
	PublicKey, err := tjx509.ReadPublicKeyFromPem(pubPEM)
	if err != nil {
		return nil, err
	}
	return &TJSM2{PrivateKey: PrivateKey, PublicKey: PublicKey}, nil
}

func (instance *TJSM2) ExportKey() (privPEM []byte, pubPEM []byte, err error) {
	privPEM, err = tjx509.WritePrivateKeyToPem(instance.PrivateKey, nil)
	if err != nil {
		return
	}
	pubPEM, err = tjx509.WritePublicKeyToPem(instance.PublicKey)
	return
}

func (instance *TJSM2) Encrypt(msg []byte) ([]byte, error) {
	encrypted, err := instance.PublicKey.EncryptAsn1(msg, rand.Reader)
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

func (instance *TJSM2) Decrypt(encrypted []byte) ([]byte, error) {
	decrypted, err := instance.PrivateKey.DecryptAsn1(encrypted)