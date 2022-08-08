package Grizzly_Encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func Run_Encrypt(password, website string) {
	var q Server_Data
	q.Password = password
	q.Website_Name = website
	q.AES_Key = Gen_Code(Key_len)
	q.AES_File = "Modules/Storage/TD/" + Gen_Code(10) + ".txt"
	Encrypted_Data := Encrypt(q.Password, q.AES_Key)
	Writer(Encrypted_Data, q.AES_File)
	Writer(fmt.Sprintf(YAML_Template_x1, q.AES_File, q.AES_File, q.AES_Key, q.Website_Name), YAML_File)

}

func Encrypt(plainstring, keystring string) string {
	plaintext := []byte(plainstring)
	key := []byte(keystring)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return string(ciphertext)
}
