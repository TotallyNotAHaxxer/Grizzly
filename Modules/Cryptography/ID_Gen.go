package Grizzly_Encryption

import "crypto/rand"

func Gen_Code(l int) string {
	ll := len(Chars)
	b := make([]byte, l)
	rand.Read(b)
	for i := 0; i < l; i++ {
		b[i] = Chars[int(b[i])%ll]
	}
	return string([]byte(b))
}
