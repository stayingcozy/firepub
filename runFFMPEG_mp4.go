package main

import (
	"fmt"
	"os/exec"
)

func runFFMPEG_mp4(user string, password string, server_ip string) {

	// ffmpeg command, h264 video codec
	cmdString := fmt.Sprintf(
		"ffmpeg -re -stream_loop -1 -i test.mp4 -c:v libx264 -bf 0 -f rtsp rtsp://%s:%s@%s:8554/mystream",
		user, password, server_ip,
	)

	cmd := exec.Command("bash", "-c", cmdString)

	err := cmd.Run()
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("ffmpeg command completed successfully")
}
