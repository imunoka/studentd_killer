package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	fmt.Println("Starting studentd killer...")
	fmt.Println("Will check every 1 millisecond and terminate any process named 'studentd'")
	fmt.Printf("Running as PID: %d (you may need sudo for some processes)\n\n", os.Getpid())

	for {
		pids := findStudentdPIDs()

		for _, pid := range pids {
			killProcess(pid)
		}

//		if len(pids) == 0 {
//			fmt.Print(".")
//		}

		time.Sleep(1 * time.Millisecond)
	}
}

func findStudentdPIDs() []int {
	switch runtime.GOOS {
	case "darwin":
		return findPIDsMac()
	case "linux":
		return findPIDsLinux()
	default:
		log.Printf("Unsupported GOOS: %s — only darwin (macOS) and linux are supported\n", runtime.GOOS)
		return nil
	}
}

// macOS: use pgrep -x (exact process name match)
func findPIDsMac() []int {
	cmd := exec.Command("pgrep", "-x", "studentd")
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return nil // no processes found — normal
		}
		log.Printf("pgrep failed: %v", err)
		return nil
	}

	var pids []int
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var pid int
		_, err := fmt.Sscanf(line, "%d", &pid)
		if err == nil && pid > 0 {
			pids = append(pids, pid)
		}
	}
	return pids
}

// Linux fallback — /proc method
func findPIDsLinux() []int {
	var pids []int

	entries, err := os.ReadDir("/proc")
	if err != nil {
		return nil
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		var pid int
		_, err := fmt.Sscanf(entry.Name(), "%d", &pid)
		if err != nil {
			continue
		}

		cmdlinePath := fmt.Sprintf("/proc/%d/cmdline", pid)
		data, err := os.ReadFile(cmdlinePath)
		if err != nil {
			continue
		}

		// cmdline is null-separated
		cmd := strings.ReplaceAll(string(data), "\x00", " ")
		cmd = strings.TrimSpace(cmd)

		if cmd == "studentd" || strings.HasPrefix(cmd, "studentd ") ||
			strings.Contains(cmd, " studentd ") || strings.Contains(cmd, "/studentd ") {
			pids = append(pids, pid)
		}
	}

	return pids
}

func killProcess(pid int) {
	fmt.Printf("Found studentd → PID %d → sending SIGTERM... ", pid)

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("failed (cannot find process)")
		return
	}

	// Try polite termination first
	err = proc.Signal(os.Interrupt) // SIGINT
	if err == nil {
		fmt.Println("SIGINT sent")
		return
	}

	fmt.Printf("SIGINT failed: %v → trying SIGKILL... ", err)

	err = proc.Kill() // SIGKILL
	if err != nil {
		fmt.Printf("KILL failed: %v\n", err)
		return
	}
	fmt.Println("SIGKILL sent")
}
