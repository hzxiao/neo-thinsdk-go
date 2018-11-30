package neo

import "testing"

func TestSign(t *testing.T) {
	privateKey, err := NewSigningKey()
	if err != nil {
		t.Fatal(err)
	}
	msg := "hello world"
	signature, err := Sign([]byte(msg), privateKey)
	if err != nil {
		t.Fatal(err)
	}

	if !Verify([]byte(msg), signature, &privateKey.PublicKey) {
		t.Fatal("ecosa sign verify error")
	}
}
