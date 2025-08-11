package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	getMemoryStats()
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
	cpu := strings.Trim(string(cpuver), "\n")
	return fmt.Sprintf("%s (%d) \n", cpu, cpuCount)
}

func getMemoryStats() string {

	pageSizeBytes, _ := exec.Command("sysctl", "-n", "hw.pagesize").Output()
	memsizeBytes, _ := exec.Command("sysctl", "-n", "hw.memsize").Output()
	inactivePagesBytes, _ := exec.Command("sysctl", "-n", "vm.page_reusable_count").Output()

	out, err := exec.Command("vm_stat").Output()
	if err != nil {
		panic(err)
	}

	// Helper to extract page counts
	parse := func(key string) uint64 {
		re := regexp.MustCompile(key + ":[ \t]+([0-9]+)\\.")
		match := re.FindStringSubmatch(string(out))
		if len(match) < 2 {
			return 0
		}
		val, _ := strconv.ParseUint(match[1], 10, 64)
		return val
	}

	activePages := parse("Pages active")

	pageSizeStr := strings.Trim(string(pageSizeBytes), "\n")
	memSizeStr := strings.Trim(string(memsizeBytes), "\n")
	inactivePagesStr := strings.Trim(string(inactivePagesBytes), "\n")
	fmt.Println(inactivePagesStr)

	memfloat, _ := strconv.ParseInt(memSizeStr, 10, 64)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 64)

	memory := float64(memfloat) / (1024 * 1024 * 1024)
	freeMemory := float64(memory-(float64(activePages)*float64(pageSize))) / (1024 * 1024 * 1024)

	fmt.Println(memory, freeMemory)
	return ""
}
