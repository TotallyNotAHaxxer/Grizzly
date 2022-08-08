package Hardware

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type User_Information_OS struct {
	User_OS_Name string       // OS Name
	User_OS_Arch string       // OS Architecture
	User_OS_Host string       // OS Hostname
	User_OS_IPA  *net.UDPAddr // OS Internet Protocal Addr
}

type User_Information_Files struct {
	User_Home_Dir string // Home Directorty
	User_Work_Dir string // Working Directory
	User_Curr_Dir string // Current Directory
}

type User_Hardware_Information struct {
	User_OS_Hardware_Unit_CPU_Vendor_ID                             string
	User_OS_Hardware_Unit_CPU_Family                                string
	User_OS_Hardware_Unit_CPU_Model_Name                            string
	User_OS_Hardware_Unit_CPU_Number_Of_Cores                       string
	User_OS_Hardware_Unit_CPU_Index_number                          string
	User_OS_Hardware_Unit_CPU_Speed_By_Index                        string
	User_OS_Hardware_Unit_CPU_Utilization                           []string
	User_OS_Hardware_Unit_Uptime_time                               string
	User_OS_Hardware_Unit_Hostname                                  string
	User_OS_Hardware_Unit_Hostname_User_ID                          string
	User_OS_Hardware_Unit_Current_Known_Kernal_Arch_And_Information string
	User_OS_Hardware_Unit_Current_Known_Platform_Version            string
	User_OS_Hardware_Unit_Current_Known_Platform_Name               string
	User_OS_Hardware_Unit_Current_Known_Platform_Family             string
	User_OS_Hardware_Unit_Current_Known_Kernal_Version              string
	User_OS_Hardware_Unit_Number_Of_Processes_Currently_Running     string
	User_OS_Hardware_Unit_Memory_Total_Free                         string
	User_OS_Hardware_Unit_Memory_Total_Mem                          string
	User_OS_Hardware_Unit_Memory_Total_Usage_Precent                string
}

var XX error
var Y net.Conn
var cc User_Information_OS

func X(x error) error {
	if x != nil {
		return x
	}
	return nil
}

//structure
func (Hardware *User_Information_OS) Parse_All_OS_Information() *User_Information_OS {
	Y, XX = net.Dial("udp", "8.8.8.8:80")
	X(XX)
	defer Y.Close()
	Addr := Y.LocalAddr().(*net.UDPAddr)
	Hardware.User_OS_IPA = Addr
	Hardware.User_OS_Host, XX = os.Hostname()
	X(XX)
	Hardware.User_OS_Name = runtime.GOOS
	Hardware.User_OS_Arch = runtime.GOARCH
	return Hardware

}

// string
func Check_Nil(value string, data_to_append string) string {
	if data_to_append != "" {
		value = data_to_append
	} else {
		fmt.Println("\033[31m VALUE -> ", data_to_append, " RETURNED EMPTY \033[39m")
		value = "Data returned empty [ no data ]"
	}
	return value
}

//structure
func (Hardware *User_Information_Files) Parse_All_Files_Information() *User_Information_Files {
	Hardware.User_Curr_Dir, XX = os.Getwd()
	X(XX)
	Hardware.User_Home_Dir, XX = os.UserHomeDir()
	X(XX)
	Hardware.User_Work_Dir, XX = os.Getwd()
	X(XX)
	return Hardware

}

var f0,
	f1,
	f2,
	f3,
	f4,
	f5,
	f6,
	f7,
	f8,
	f9,
	f10,
	f11,
	f12,
	f13,
	f14,
	f15,
	f16,
	f17,
	f18 string

//structure
func (Hardware *User_Hardware_Information) Parse_All_Hardware_Information() *User_Hardware_Information {
	Mem, x := mem.VirtualMemory()
	X(x)
	pe, err := cpu.Percent(0, true)
	for ctx, p := range pe {
		Hardware.User_OS_Hardware_Unit_CPU_Utilization = append(Hardware.User_OS_Hardware_Unit_CPU_Utilization, "\n | Core #"+strconv.Itoa(ctx)+" -> "+strconv.FormatFloat(p, 'f', 2, 64)+"%  Used \t| ")
	}
	X(err)
	cpu, xx := cpu.Info()
	X(xx)
	host, xx1 := host.Info()
	X(xx1)
	f1 = Check_Nil(Hardware.User_OS_Hardware_Unit_Memory_Total_Mem, strconv.FormatUint(Mem.Total, 10))
	f2 = Check_Nil(Hardware.User_OS_Hardware_Unit_Memory_Total_Free, strconv.FormatUint(Mem.Free, 10))
	sped0 := Hardware.User_OS_Hardware_Unit_Memory_Total_Usage_Precent + " %"
	f3 = Check_Nil(sped0, strconv.FormatFloat(Mem.UsedPercent, 'f', 2, 64))
	f4 = Check_Nil(Hardware.User_OS_Hardware_Unit_CPU_Index_number, strconv.FormatInt(int64(cpu[0].CPU), 10))
	f5 = Check_Nil(Hardware.User_OS_Hardware_Unit_CPU_Vendor_ID, cpu[0].VendorID)
	f6 = Check_Nil(Hardware.User_OS_Hardware_Unit_CPU_Family, cpu[0].Family)
	f7 = Check_Nil(Hardware.User_OS_Hardware_Unit_CPU_Number_Of_Cores, strconv.FormatInt(int64(cpu[0].Cores), 10))
	f8 = Check_Nil(Hardware.User_OS_Hardware_Unit_CPU_Model_Name, cpu[0].ModelName)
	sped := Hardware.User_OS_Hardware_Unit_CPU_Speed_By_Index + " Mhz"
	f9 = Check_Nil(sped, strconv.FormatFloat(cpu[0].Mhz, 'f', 2, 64))
	f10 = Check_Nil(Hardware.User_OS_Hardware_Unit_Hostname, host.Hostname)
	f11 = Check_Nil(Hardware.User_OS_Hardware_Unit_Hostname_User_ID, host.HostID)
	f12 = Check_Nil(Hardware.User_OS_Hardware_Unit_Uptime_time, strconv.FormatUint(host.Uptime, 10))
	f13 = Check_Nil(Hardware.User_OS_Hardware_Unit_Number_Of_Processes_Currently_Running, strconv.FormatUint(host.Procs, 10))
	f14 = Check_Nil(Hardware.User_OS_Hardware_Unit_Current_Known_Kernal_Version, host.KernelVersion)
	f15 = Check_Nil(Hardware.User_OS_Hardware_Unit_Current_Known_Kernal_Arch_And_Information, host.KernelArch)
	f16 = Check_Nil(Hardware.User_OS_Hardware_Unit_Current_Known_Platform_Version, host.PlatformVersion)
	f17 = Check_Nil(Hardware.User_OS_Hardware_Unit_Current_Known_Platform_Family, host.PlatformFamily)
	f18 = Check_Nil(Hardware.User_OS_Hardware_Unit_Current_Known_Platform_Name, host.Platform)
	v_hw := User_Hardware_Information{
		User_OS_Hardware_Unit_Memory_Total_Mem:                          f1,
		User_OS_Hardware_Unit_Memory_Total_Free:                         f2,
		User_OS_Hardware_Unit_Memory_Total_Usage_Precent:                f3 + "%",
		User_OS_Hardware_Unit_CPU_Index_number:                          f4,
		User_OS_Hardware_Unit_CPU_Vendor_ID:                             f5,
		User_OS_Hardware_Unit_CPU_Family:                                f6,
		User_OS_Hardware_Unit_CPU_Number_Of_Cores:                       f7,
		User_OS_Hardware_Unit_CPU_Model_Name:                            f8,
		User_OS_Hardware_Unit_CPU_Speed_By_Index:                        f9 + " Mhz",
		User_OS_Hardware_Unit_Hostname:                                  f10,
		User_OS_Hardware_Unit_Hostname_User_ID:                          f11,
		User_OS_Hardware_Unit_Uptime_time:                               f12,
		User_OS_Hardware_Unit_Number_Of_Processes_Currently_Running:     f13,
		User_OS_Hardware_Unit_Current_Known_Kernal_Version:              f14,
		User_OS_Hardware_Unit_Current_Known_Kernal_Arch_And_Information: f15,
		User_OS_Hardware_Unit_Current_Known_Platform_Version:            f16,
		User_OS_Hardware_Unit_Current_Known_Platform_Family:             f17,
		User_OS_Hardware_Unit_Current_Known_Platform_Name:               f18,
	}
	return &v_hw
}

func Runner() (*User_Hardware_Information, *User_Information_Files, *User_Information_OS) {
	HW_H := User_Hardware_Information{}
	HW_F := User_Information_Files{}
	HW_O := User_Information_OS{}
	HW_0 := HW_H.Parse_All_Hardware_Information()
	HW_1 := HW_F.Parse_All_Files_Information()
	HW_2 := HW_O.Parse_All_OS_Information()
	return HW_0, HW_1, HW_2
}
