package main

import (
	"fmt"
	"os/exec"
	"time"
)

func runFFMPEG(url string) error {

	maxAttempts := 5
	attempt := 1

	for {

		// ffmpeg command, h264 video codec
		cmd := exec.Command(
			"ffmpeg", "-f", "v4l2 ", "-input_format", "h264","-i", "/dev/video1", "-c:v", "libx264",
			"-bf", "0", "-f", "rtsp", url,
		)

		err := cmd.Run()

		if err == nil {
			// Command succeeded, break out of the loop
			break
		}

		fmt.Printf("Attempt %d failed: %v\n", attempt, err)

		if attempt >= maxAttempts {
			return fmt.Errorf("max attempts reached, unable to start ffmpeg")
		}

		// Increment the attempt counter and wait for a moment before retrying
		attempt++
		time.Sleep(2 * time.Second)
	}

	return nil

}
