package Grizzly_YAML

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

var Total_Websites, Total_Passwords, Total_Keys, Grizz_Dat, Res string
var Website, Passwrd, Total_Ke int

func SetW(filepath string) error {
	err := os.Chmod(filepath, 0222)
	return err
}

func Decrypt_Files(filename, p string) (string, error) {
	var ct []byte
	ct, x := ioutil.ReadFile(filename)
	if x != nil {
		return "", x
	}
	b, x2 := aes.NewCipher([]byte(p))
	if x2 != nil {
		return "", x2
	}
	iv := ct[:aes.BlockSize]
	ct = ct[aes.BlockSize:]
	Crypt := cipher.NewCFBDecrypter(b, iv)
	Crypt.XORKeyStream(ct, ct)
	return string(ct), nil
}

func Reader(filename string) {
	f, x := ioutil.ReadFile(filename)
	phrase := "Grizzly Password Managment Solutions   :: Got error when trying to read file -> " + filename + " -> "
	E(x, phrase)
	data := make(map[string]Data)
	err2 := yaml.Unmarshal(f, &data)
	phrase = "Grizzly Password Managment Solutions   :: Got error when trying to unmarshal the yaml structure -> "
	E(err2, phrase)
	fi, x := os.Open("Modules/YML/HTML_Top.txt")
	if x != nil {
		log.Fatal(x)
	}
	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	var Template_HTM string
	for scanner.Scan() {
		Template_HTM += scanner.Text()
	}
	Template_HTM += "<tbody>\n"
	Website = 0
	Passwrd = 0
	Total_Ke = 0
	for _, v := range data {
		Template_HTM += "<tr>\n"
		Template_HTM += "\t<td>" + v.Website + "</td>\n"
		Pass, x := Decrypt_Files(v.Filename, v.Key)
		if x != nil {
			fmt.Println("[!] Error when decrypting file -> ")
		}
		Template_HTM += "\t<td>" + Pass + "</td>\n"
		Template_HTM += "\t<td>" + v.Filename + "</td>\n"
		Template_HTM += "\t<td>" + v.Key + "</td>\n"
		Template_HTM += "</tr>\n"
		Website++
		Passwrd++
		Total_Ke++
	}
	a := strconv.Itoa(Website)
	b := strconv.Itoa(Passwrd)
	c := strconv.Itoa(Total_Ke)
	d := Website * Passwrd * Total_Ke
	Total_Websites = a
	Total_Passwords = b
	Total_Keys = c
	Res = strconv.Itoa(d)
	Template_HTM += "</tbody>\n"
	Dt := Template_HTM + Template_CSS
	time.Sleep(1 * time.Second)
	File_Writer("server_html/Passwords/Index.html", Dt)
}
