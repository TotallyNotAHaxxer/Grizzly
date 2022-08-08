package main

import (
	"fmt"
	"log"
	"net/http"

	Grizzly_Encryption "main/Modules/Cryptography"
	Grizzly_Assassins "main/Modules/Server_Helpers"
	Grizzly_Hardware "main/Modules/User-Information"
	Grizzly_YAML "main/Modules/YML"
	Grizzly_Dash "main/server_html/Dashboard"
)

var (
	Passwords_Typed int
	Websites_Typed  int
)

var Dashbd Grizzly_Dash.Grizzly_Dash_Data

func Processor(writer http.ResponseWriter, Request *http.Request) {
	if Request.URL.Path == "/server_html/Dashboard" {
		Grizzly_Assassins.Check_Kill("server_html/Passwords/Index.html")
		Grizzly_Assassins.Check_Kill("server_html/User/User.html")
		http.ServeFile(writer, Request, "server_html/Dashboard/Dash.html")
	}
	if Request.URL.Path == "/server_html/Post_Req" {
		Grizzly_Assassins.Check_Kill("server_html/Passwords/Index.html")
		Grizzly_Assassins.Check_Kill("server_html/User/User.html")
		http.ServeFile(writer, Request, "server_html/Post_Req/index.html")
	}
	if Request.URL.Path == "/server_html/User" {
		Grizzly_Assassins.Check_Kill("server_html/Passwords/Index.html")
		Grizzly_Assassins.Check_Kill("server_html/User/User.html")
		Grizzly_Hardware.Call_Gen()
		http.ServeFile(writer, Request, "server_html/User/User.html")
	}
	if Request.URL.Path == "/server_html/Passwords" {
		Grizzly_Assassins.Check_Kill("server_html/Passwords/Index.html")
		Grizzly_Assassins.Check_Kill("server_html/User/User.html")
		Grizzly_YAML.Reader("Modules/Storage/FN/DT.yaml")
		http.ServeFile(writer, Request, "server_html/Passwords/Index.html")
	}

	switch Request.Method {
	case "GET":
		Grizzly_Dash.Run_Format_Tmpl("Modules/Storage/FN/DT.yaml")
		Grizzly_Assassins.Check_Kill("server_html/Passwords/Index.html")
		Grizzly_Assassins.Check_Kill("server_html/User/User.html")
		http.ServeFile(writer, Request, "server_html/Dashboard/Dash.html")
	case "POST":
		if err := Request.ParseForm(); err != nil {
			log.Fatal(err)
			return
		} else {
			fmt.Println("no error")
		}
		Website_name := Request.FormValue("website")
		Password := Request.FormValue("password")
		if Website_name != "" && Password != "" {
			Passwords_Typed++
			Websites_Typed++
			Dashbd.Passwords_Saved_This_Session = Grizzly_Dash.Convert(Passwords_Typed)
			Dashbd.Websites_Saved_This_Session = Grizzly_Dash.Convert(Websites_Typed)
			Grizzly_Encryption.Run_Encrypt(Password, Website_name)
			http.ServeFile(writer, Request, "server_html/Post_Req/index.html")
		} else {
			fmt.Fprintf(writer, "You may be seeing this page due to a value being NILL/NULL/VOID or in other simpler terms empty, did you miss a field in the input? if so that may be why you are seeing this page, make sure all fields are correctly formatted and have data inside them before submitting it")
		}
	}
}

func main() {
	Grizzly_Hardware.Generate_User_Information()
	CSS := Grizzly_Hardware.Return_CSS()
	for _, k := range CSS {
		Grizzly_Hardware.HTML_Template_User_Information += k
	}
	http.HandleFunc("/", Processor)
	fmt.Println("[+]> http://localhost:5501")
	http.ListenAndServe(":5501", nil)
}
