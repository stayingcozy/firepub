package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func RunFFMPEG() {
	// Set up the command and arguments
	// cmd := exec.Command("ffmpeg", "-re", "-i", "/dev/video1", "-vcodec", "libvpx", "-cpu-used", "5", "-deadline", "1", "-g", "10", "-error-resilient", "1", "-auto-alt-ref", "1", "-f", "rtp", "rtp://127.0.0.1:5004?pkt_size=1200")
	cmd := exec.Command("ffmpeg", "-re", "-stream_loop", "-1", "-i", "video/test.mp4", "-s", "640x480", "-r", "30", "-vcodec", "libvpx", "-cpu-used", "5", "-deadline", "1", "-g", "10", "-error-resilient", "1", "-auto-alt-ref", "1", "-f", "rtp", "rtp://127.0.0.1:5004?pkt_size=1200")
	// 	"-input_format", "h264",  // Added line for input format
	// Set the output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ffmpeg command completed successfully")
}
