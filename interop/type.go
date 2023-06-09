
package interop

type SM2 interface {
	Encrypt(msg []byte) ([]byte, error)
	Decrypt(encrypted []byte) ([]byte, error)
	Sign(msg []byte) ([]byte, error)
	Verify(msg []byte, sign []byte) bool
	ExportKey() (privPEM []byte, pubPEM []byte, err error)
}