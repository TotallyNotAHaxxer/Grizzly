package Hardware

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	Grizzly_Yaml "main/Modules/YML"
	"os"
)

var Hi, Fi, Oi = Runner()

var HTML_Template_User_Information = `
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
    <!--
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
      -->
      <div class="sales-boxes"> 
        <div class="recent-sales box">
          <div class="title">User Information</div>
          <div class="sales-details">
            <ul class="details">
              <li class="topic">Title</li>
              <li><a href="#">Current Working Directory</a></li>
              <li><a href="#">Current Home Directory</a></li>
              <li><a href="#">Current File Directory</a></li>
              <li><a href="#">Operating System Name</a></li>
              <li><a href="#">Operating System Architecture</a></li>
              <li><a href="#">Operating System Hostname</a></li>
              <li><a href="#">Operating System IP</a></li>
              <li><a href="#">Operating System Uptime</a></li>
              <li><a href="#">Operating System User</a></li>
              <li><a href="#">Operating System User ID</a></li>
              <li><a href="#">Current Kernal Architecture</a></li>
              <li><a href="#">Current Kernal Version</a></li>
              <li><a href="#">Current Platform Version</a></li>
              <li><a href="#">Current Platform Name</a></li>
              <li><a href="#">Current Platform Family</a></li>
              <li><a href="#">Current Number of proc's</a></li>
              <li><a href="#">CPU Vendor ID</a></li>
              <li><a href="#">CPU Family ID</a></li>
              <li><a href="#">CPU Model Name</a></li>
              <li><a href="#">CPU Number of Cores</a></li>
              <li><a href="#">CPU Index Number</a></li>
              <li><a href="#">CPU Speed in Mhz</a></li>
              <li><a href="#">Total memory free</a></li>
              <li><a href="#">Total memory </a></li>
              <li><a href="#">Total memory used</a></li>
            </ul>
          <ul class="details">
            <li class="topic">Value</li>
            <!--25 needed values-->
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%v</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%s</b></li>
            <li><b href="#">%v</b></li>
            <li><b href="#">%v</b></li>
            <li><b href="#">%v</b></li>
            <li><b href="#">%v</b></li>
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

// generate HTML
//Hi = Hardware Information
//Fi = File information
//Oi = Operating System information

func Generate_User_Information() {
	HTML_Template_User_Information = fmt.Sprintf(
		HTML_Template_User_Information,
		Grizzly_Yaml.Total_Passwords,
		Grizzly_Yaml.Total_Websites,
		Grizzly_Yaml.Total_Keys,
		Grizzly_Yaml.Res,
		Fi.User_Work_Dir,
		Fi.User_Home_Dir,
		Fi.User_Curr_Dir,
		Oi.User_OS_Name,
		Oi.User_OS_Arch,
		Oi.User_OS_Host,
		Oi.User_OS_IPA.IP,
		Oi.User_OS_IPA.Port,
		Oi.User_OS_IPA.Zone,
		Hi.User_OS_Hardware_Unit_Uptime_time,
		Hi.User_OS_Hardware_Unit_Hostname,
		Hi.User_OS_Hardware_Unit_Hostname_User_ID,
		Hi.User_OS_Hardware_Unit_Current_Known_Kernal_Arch_And_Information,
		Hi.User_OS_Hardware_Unit_Current_Known_Kernal_Version,
		Hi.User_OS_Hardware_Unit_Current_Known_Platform_Version,
		Hi.User_OS_Hardware_Unit_Current_Known_Platform_Name,
		Hi.User_OS_Hardware_Unit_Current_Known_Platform_Family,
		Hi.User_OS_Hardware_Unit_Number_Of_Processes_Currently_Running,
		Hi.User_OS_Hardware_Unit_CPU_Vendor_ID,
		Hi.User_OS_Hardware_Unit_CPU_Family,
		Hi.User_OS_Hardware_Unit_CPU_Model_Name,
		Hi.User_OS_Hardware_Unit_CPU_Number_Of_Cores,
		Hi.User_OS_Hardware_Unit_CPU_Index_number,
		Hi.User_OS_Hardware_Unit_CPU_Speed_By_Index,
		Hi.User_OS_Hardware_Unit_Memory_Total_Free,
		Hi.User_OS_Hardware_Unit_Memory_Total_Mem,
		Hi.User_OS_Hardware_Unit_Memory_Total_Usage_Precent,
	)
}

func Return_CSS() []string {
	var scanner []string
	f, x := os.Open("Modules/User-Information/style.txt")
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		scanner1 := bufio.NewScanner(f)
		for scanner1.Scan() {
			scanner = append(scanner, scanner1.Text())
		}
	}
	return scanner
}

func Call_Gen() {
	f := ioutil.WriteFile("server_html/User/User.html", []byte(HTML_Template_User_Information), 0600)
	if f != nil {
		log.Fatal(f)
	} else {
		fmt.Println("you are ready to go :)")
		fmt.Println(HTML_Template_User_Information)
	}
}
