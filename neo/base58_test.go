package neo

import (
	"bytes"
	"testing"
)

func TestB58encoding(t *testing.T)  {
	checkB58Encoding(t, []byte("1"))

	checkB58Encoding(t, []byte("hello world"))
}

func checkB58Encoding(t *testing.T, data []byte)  {
	e := b58encode(data)
	d, err := b58decode(e)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, d) {
		t.Fatalf("invalid base-58 encoding error")
	}
}

func TestBase58CheckEncoding(t *testing.T) {
	checkB58CheckEncoding(t, []byte("1"))

	checkB58CheckEncoding(t, []byte("hello world"))
}

func checkB58CheckEncoding(t *testing.T, data []byte)  {
	e := Base58CheckEncode(0, data)
	v, d, err := Base58CheckDecode(e)
	if err != nil {
		t.Fatal(err)
	}

	if v != 0 {
		t.Fatal("base-58 check encoding invalid ver error")
	}
	if !bytes.Equal(data, d) {
		t.Fatalf("invalid base-58 check encoding error")
	}
}
