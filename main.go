package main

import "os"
import "os/exec"

// import "os/signal"
// import "syscall"
import "time"

func main() {
	cmd := os.Args[1]
	args := os.Args[2:]
	// sigs := make(chan os.Signal, 1)

	// signal.Notify(sigs)

	for {
		last_started_at := time.Now()
		subprocess := exec.Command(cmd, args...)
		subprocess.Stdin = os.Stdin
		subprocess.Stdout = os.Stdout
		subprocess.Stderr = os.Stderr
		err := subprocess.Run()

		if err == nil {
			os.Exit(0)
		} else if time.Now().Sub(last_started_at) < 10*time.Second {
			os.Exit(1)
		}
	}
}
