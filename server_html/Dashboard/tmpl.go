package dashboard

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	Grizzly_YAML "main/Modules/YML"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

var HTML = `
<!DOCTYPE html>
<html lang="en" dir="ltr">
  <head>
    <meta charset="UTF-8">
    <title> Grizzly </title>
    <link href='https://unpkg.com/boxicons@2.0.7/css/boxicons.min.css' rel='stylesheet'>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
   </head>
<body>
  <div class="sidebar">
    <div class="logo-details">
      <i class='''></i>
      <span class="logo_name">Grizzly</span>
    </div>
    <ul class="nav-links">
      <li><a href="/server_html/Dashboard" class="active"><i class='bx bx-grid-alt' ></i><span class="links_name">Dashboard</span></a></li>
      <li><a href="/server_html/Passwords"><i class='bx bx-box' ></i><span class="links_name">Passwords</span></a></li>
      <li><a href="/server_html/User"><i class='bx bx-pie-chart-alt-2' ></i><span class="links_name">User Information</span></a></li>
      <li><a href="#"><i class='bx bx-coin-stack' ></i><span class="links_name">Documentation</span></a></li>
      <li><a href="/server_html/Post_Req"><i class='bx bx-book-alt' ></i><span class="links_name">Save a New Password</span></a></li>
   </ul>
  </div>
  <section class="home-section">
    <nav>
      <div class="sidebar-button">
        <i class='bx bx-menu sidebarBtn'></i>
        <span class="dashboard">Dashboard</span>
      </div>
    </nav>

    <div class="home-content">
      <div class="overview-boxes">
        <div class="box">
          <div class="right-side">
            <div class="box-topic">Total Passwords</div>
            <div class="number">%s</div>
          </div>
        </div>
        <div class="box">
          <div class="right-side">
            <div class="box-topic">Total Websites</div>
            <div class="number">%s</div>
          </div>
        </div>
        <div class="box">
          <div class="right-side">
            <div class="box-topic">Total AES Keys</div>
            <div class="number">%s</div>
          </div>
        </div>
        <div class="box">
          <div class="right-side">
            <div class="box-topic">Grizzly Data</div>
            <div class="number">%s</div>
          </div>
        </div>
      </div>
      <div class="sales-boxes"> 
        <div class="recent-sales box">
          <div class="title">Grizzly Log</div>
          <div class="sales-details">
            <ul class="details">
              <li class="topic">Log Date</li>
              <li><a href="#">%s</a></li>
            </ul>
            <ul class="details">
            <li class="topic">Passwords saved</li>
            <li><a href="#">%s</a></li>
          </ul>
          <ul class="details">
            <li class="topic">Websites saved</li>
            <li><a href="#">%s</a></li>
          </ul>
          </div>
        </div>
      </div>
    </div>
  </section>

  <script>
   let sidebar = document.querySelector(".sidebar");
let sidebarBtn = document.querySelector(".sidebarBtn");
sidebarBtn.onclick = function() {
  sidebar.classList.toggle("active");
  if(sidebar.classList.contains("active")){
  sidebarBtn.classList.replace("bx-menu" ,"bx-menu-alt-right");
}else
  sidebarBtn.classList.replace("bx-menu-alt-right", "bx-menu");
}
 </script>

</body>
</html>
`

type Grizzly_Dash_Data struct {
	Total_Pass                   string
	Total_Websites               string
	Total_Keys                   string
	Total_Amount_Of_Logged_Data  string
	Date_Of_Log                  string
	Passwords_Saved_This_Session string
	Websites_Saved_This_Session  string
}

type Data struct {
	Filename string `yaml:"FileName"`
	Key      string `yaml:"NN_P"`
	Website  string `yaml:"Website"`
}

var Dash Grizzly_Dash_Data

func Convert(f int) string {
	a := strconv.Itoa(f)
	return a
}

func Run_Format_Tmpl(filename string) {
	var (
		template, template_Date string
		Passes                  int
		Websites                int
		Keys                    int
		Total                   int
	)

	f, x := ioutil.ReadFile(filename)
	phrase := "Grizzly Password Managment Solutions   :: Got error when trying to read file -> " + filename + " -> "
	if x != nil {
		fmt.Println(phrase, x)
	}
	data := make(map[string]Data)
	err2 := yaml.Unmarshal(f, &data)
	phrase = "Grizzly Password Managment Solutions   :: Got error when trying to unmarshal the yaml structure -> "
	if err2 != nil {
		fmt.Println(phrase, err2)
	}
	y, m, d := time.Now().Local().Date()
	template_Date = "%v  %v  %v"
	val_date := fmt.Sprintf(template_Date, y, m, d)
	Dash.Date_Of_Log = val_date
	for _, v := range data {
		if v.Key != "" {
			Keys++
		}
		if v.Website != "" {
			Websites++
		}
		pass, x := Grizzly_YAML.Decrypt_Files(v.Filename, v.Key)
		if x != nil {
			fmt.Println("[!] Could not decrypt file -> ", x)
		} else {
			if pass != "" {
				Passes++
			}
		}
	}
	Total = Keys + Websites
	Dash.Total_Amount_Of_Logged_Data = Convert(Total)
	Dash.Total_Keys = Convert(Keys)
	Dash.Total_Pass = Convert(Passes)
	Dash.Total_Websites = Convert(Websites)
	if Dash.Passwords_Saved_This_Session == "" {
		Dash.Passwords_Saved_This_Session = "None during this session"
	}
	if Dash.Websites_Saved_This_Session == "" {
		Dash.Websites_Saved_This_Session = "None during this session"
	}
	template = fmt.Sprintf(HTML,
		Dash.Total_Pass,
		Dash.Total_Websites,
		Dash.Total_Keys,
		Dash.Total_Amount_Of_Logged_Data,
		Dash.Date_Of_Log,
		Dash.Passwords_Saved_This_Session,
		Dash.Websites_Saved_This_Session)
	f2, x := os.Open("server_html/Dashboard/CSS.txt")
	if x != nil {
		log.Fatal(x)
	} else {
		defer f2.Close()
		scanner := bufio.NewScanner(f2)
		for scanner.Scan() {
			template += scanner.Text()
		}
	}
	time.Sleep(1 * time.Second)
	Grizzly_YAML.File_Writer("server_html/Dashboard/Dash.html", template)
}
