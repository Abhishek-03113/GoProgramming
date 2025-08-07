package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	getCPU()
}

func getHostName() string {
	user := os.Getenv("USER")
	hostname, _ := os.Hostname()
	return fmt.Sprintf("%s@%s", user, hostname)
}

func getTerm() string {
	return os.Getenv("TERM")
}

func getOS() string {
	name, _ := exec.Command("sw_vers", "--ProductName").Output()
	version, _ := exec.Command("sw_vers", "--ProductVersion").Output()
	build, _ := exec.Command("sw_vers", "--BuildVersion").Output()
	namestr := strings.Trim(string(name), "\n")
	versionstr := strings.Trim(string(version), "\n")
	buildstr := strings.Trim(string(build), "\n")
	archget, _ := exec.Command("uname", "-m").Output()
	arch := string(archget)
	return fmt.Sprintf("%s : %s (%s) %s", namestr, versionstr, buildstr, arch)
}

func getUptime() string {
	bootTime, _ := exec.Command("sysctl", "-n", "kern.boottime").Output()
	boot := string(bootTime)
	sec := strings.Split(boot, " ")[3]
	sec = strings.Trim(sec, ",")
	secs, _ := strconv.ParseInt(sec, 10, 64)
	uptime := time.Since(time.Unix(secs, 0))
	return uptime.String()
}

func getCPU() string {
	cpuCount := runtime.NumCPU()
	cpuver, _ := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
	cpu := string(cpuver)
	return fmt.Sprintf("%s (%d)", cpu, cpuCount)
}

func getMemoryStats() string {

	pageSizeBytes, _ := exec.Command("sysctl", "-n", "hw.pagesize").Output()
	memsizeBytes, _ := exec.Command("sysctl", "-n", "hw.memsize").Output()
	freePagesBytes, _ := exec.Command("sysctl", "-n", "vm.page_free_count").Output()

	pageSizeStr := strings.Trim(string(pageSizeBytes), "\n")
	memSizeStr := strings.Trim(string(memsizeBytes), "\n")
	freePagesStr := strings.Trim(string(freePagesBytes), "\n")

	fmt.Println(memSizeStr, pageSizeStr, freePagesStr)

	return ""
}
