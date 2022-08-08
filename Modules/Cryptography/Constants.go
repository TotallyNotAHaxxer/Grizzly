package Grizzly_Encryption

type Server_Data struct {
	Password     string
	Website_Name string
	AES_Key      string
	AES_File     string
}

var YAML_Template_x1 = `File %s:
     FileName: "%s"
     NN_P: "%s"
     Website: "%s"
`

var YAML_File = "Modules/Storage/FN/DT.yaml"
var X error
var Key_len = len("****************")
var Chars = "abcdefghijklmnopqrstuvwxyz1234567890"
