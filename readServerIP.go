package main

import (
	"context"

	"cloud.google.com/go/firestore"
)

func readServerIP(ctx context.Context, client *firestore.Client, desiredServer string) (map[string]interface{}, error) {

	dsnap, err := client.Collection("serverIP").Doc(desiredServer).Get(ctx)
	if err != nil {
		return nil, err
	}

	serverMap := dsnap.Data()

	return serverMap, nil
}
