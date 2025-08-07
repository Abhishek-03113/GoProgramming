package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
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
	return fmt.Sprintf("%s : %s (%s)", namestr, versionstr, buildstr)
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
