package mycrypto

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCrypto(t *testing.T) {
	key := []byte("1234567890123456")
	msg := []byte("1234567890")
	cipher, err := Encrypt(msg, key)
	if err != nil {
		t.Errorf("encrypto failed:%v", err)
	}
	plain, err := Decrypt(cipher, key)
	if err != nil {
		t.Errorf("decrypto failed:%v", err)
	}
	fmt.Println(msg)
	fmt.Println(plain)
	if !reflect.DeepEqual(msg, plain) {
		t.Errorf("fail")
	}

}
