package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {

	// Firebase Init //
	const serviceAccountKey = os.Getenv("YOUR_SERVICE_ACCOUNT_KEY_PATH")
	path_to_saKey, saErr := pathToSaKey(serviceAccountKey)
	if saErr != nil {
		panic(saErr)
	}

	// Read UID
	uid_file := "uid.txt"
	path_to_uid, uidErr := pathToUid(uid_file)
	if uidErr != nil {
		panic(uidErr)
	}
	uid := readUid(path_to_uid)

	// Initialize Cloud Firestore - my server
	ctx := context.Background()
	sa := option.WithCredentialsFile(path_to_saKey)
	app, errFB := firebase.NewApp(ctx, nil, sa)
	if errFB != nil {
		panic(errFB)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Update firebase that it is connected to wifi
	wifiConnected(ctx, client, uid)

	for {

		// Read Server IP from firebase
		serverMap, _, errIP := readMediaServerIP(ctx, client, uid)
		if errIP != nil {
			log.Printf("An error while reading server IP from custom func: %s", err)
		}
		serverIP := serverMap["ip"].(string)

		// Create password for stream
		password := generateRandomString(20)

		// Create url - assumption rtsp, mediamtx default rtsp port
		// url := fmt.Sprintf("rtsp://%s:%s@%s:8554/mystream", uid, password, serverIP)
		url := fmt.Sprintf("rtsp://%s:%s@%s:8554/%s", uid, password, serverIP, uid)

		// Notify Firebase of Stream
		newDoc := generateRandomString(28) // matches length of rand firebase doc title
		addStream(ctx, client, newDoc, url, uid, password, serverIP)

		fail, errLR := streamListenReceive(ctx, client, newDoc)
		if errLR != nil {
			log.Printf("An error occured while listening to stream status: %s", errLR)
		}
		if fail {
			continue
		}

		errUS := updateAsStreaming(ctx, client, newDoc)
		if errUS != nil {
			log.Printf("An error occured while updating to stream status to streaming: %s", errLR)
		}

		// Run ffmpeg rtsp stream
		// errFFMPEG := runFFMPEG_mp4(url)
		errFFMPEG := runFFMPEG(url)
		if errFFMPEG != nil {
			// Handle error
			log.Printf("An error occured while running ffmpeg to stream: %s", errFFMPEG)
		}

		// If at this point stream is broken - update firebase
		updateStreamAsBroken(ctx, client, newDoc)

	}
}
