package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

func main() {
	// Run the vm_stat command
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

	// Extract stats
	freePages := parse("Pages free")
	activePages := parse("Pages active")
	inactivePages := parse("Pages inactive")
	speculativePages := parse("Pages speculative")
	wiredPages := parse("Pages wired down")
	purgeablePages := parse("Pages purgeable")
	reusablePages := parse("Pages reusable")

	pageSize := uint64(4096) // bytes

	// Compute available memory (like Neofetch)
	availablePages := freePages + speculativePages + reusablePages
	availableBytes := availablePages * pageSize

	// Compute total memory from sysctl
	totalBytesRaw, err := exec.Command("sysctl", "-n", "hw.memsize").Output()
	if err != nil {
		panic(err)
	}
	totalBytes, _ := strconv.ParseUint(string(totalBytesRaw[:len(totalBytesRaw)-1]), 10, 64)

	// Used = total - available
	usedBytes := totalBytes - availableBytes

	// Print results
	fmt.Printf("Total Memory: %.2f GiB\n", float64(totalBytes)/(1024*1024*1024))
	fmt.Printf("Used Memory:  %.2f GiB\n", float64(usedBytes)/(1024*1024*1024))
	fmt.Printf("Avail Memory: %.2f GiB\n", float64(availableBytes)/(1024*1024*1024))

	// Debug: print all stats in pages
	fmt.Println("\n--- Raw Page Stats ---")
	fmt.Printf("Free: %d\n", freePages)
	fmt.Printf("Active: %d\n", activePages)
	fmt.Printf("Inactive: %d\n", inactivePages)
	fmt.Printf("Speculative: %d\n", speculativePages)
	fmt.Printf("Wired: %d\n", wiredPages)
	fmt.Printf("Purgeable: %d\n", purgeablePages)
	fmt.Printf("Reusable: %d\n", reusablePages)
}
