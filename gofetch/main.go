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

	// hostname
	fmt.Println(getHostName())
	// os
	fmt.Println(getOS())
	// host
	fmt.Println(getHost())
	// kernel
	fmt.Println(getKernel())
	// uptime
	fmt.Println(getUptime())
	// shell
	fmt.Println(getShell())
	// resolution
	fmt.Println(getDisplayResolution())
	// DE
	fmt.Println(getDE())
	// WM
	fmt.Println(getWM())
	// Terminal
	fmt.Println(getTerm())
	// CPU
	fmt.Println(getCPU())
	// GPU
	fmt.Println(getGPU())
	// Memory
	fmt.Println(getMemoryStats())
}

func getWM() string {
	wm := "Quartz Compositor" // default for macOS
	candidates := []string{"Rectangle", "Amethyst", "yabai", "Magnet", "chunkwm"}

	out, _ := exec.Command("ps", "-A").Output()
	processList := string(out)

	for _, c := range candidates {
		if strings.Contains(processList, c) {
			wm = c
			break
		}
	}

	return fmt.Sprintf("%s", wm)

}

func getHostName() string {
	user := os.Getenv("USER")
	hostname, _ := os.Hostname()
	return fmt.Sprintf("%s@%s", user, hostname)
}

func getHost() string {
	host, _ := os.Hostname()
	return host
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
	arch := strings.Trim(string(archget), "\n")
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
	return fmt.Sprintf("%s (%d)", cpu, cpuCount)
}

func getMemoryStats() string {

	pageSizeBytes, _ := exec.Command("sysctl", "-n", "hw.pagesize").Output()
	memsizeBytes, _ := exec.Command("sysctl", "-n", "hw.memsize").Output()

	out, err := exec.Command("vm_stat").Output()
	if err != nil {
		panic(err)
	}

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

	memfloat, _ := strconv.ParseInt(memSizeStr, 10, 64)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 64)

	memory := float64(memfloat) / (1024 * 1024)
	freeMemory := float64(memory - (float64(activePages)*float64(pageSize))/(1024*1024))

	return fmt.Sprintf("%.fMiB / %.fMiB", freeMemory, memory)
}

func getShell() string {
	return os.Getenv("SHELL")
}

func getGPU() string {

	out, _ := exec.Command("sh", "-c", `ioreg -r -d1 -c "IOPCIDevice" | grep "model"`).Output()

	lines := strings.Split(string(out), "\n")
	var gpu []interface{}
	for _, line := range lines {
		if strings.Contains(line, "model") {
			name := strings.TrimSpace(strings.Split(line, "<\"")[1])
			name = strings.TrimSuffix(name, "\">")
			gpu = append(gpu, name)
		}
	}
	return fmt.Sprintf("%s %s", gpu...)
}

func getKernel() string {
	out, _ := exec.Command("uname", "-v").Output()
	outStr := strings.Trim(strings.Split(string(out), ":")[0], "\n")

	return fmt.Sprintf("%s", outStr)
}

func getDisplayResolution() string {
	out, _ := exec.Command("sh", "-c", `system_profiler SPDisplaysDataType | grep "Resolution:" | head -n 1`).Output()
	return fmt.Sprintf("%s", strings.Trim(strings.TrimSpace(string(out)), "\n"))
}

func getDE() string {
	return fmt.Sprintf("Aqua")
}
