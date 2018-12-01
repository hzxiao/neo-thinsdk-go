package neo

import (
	"testing"
)

func TestCompressPublicKey(t *testing.T) {
	key, err := NewSigningKey()
	if err != nil {
		t.Fatal(err)
	}
	compressed := CompressPublicKey(&key.PublicKey)
	uncompressed, err := DecompressPublicKey(compressed)
	if err != nil {
		t.Fatal(err)
	}
	result := comparePublicKey(&key.PublicKey, uncompressed)
	if result != true {
		t.Fatal("result does not match!")
	}
}
