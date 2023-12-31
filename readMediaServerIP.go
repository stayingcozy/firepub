package main

import (
	"context"

	"cloud.google.com/go/firestore"
)

func readMediaServerIP(ctx context.Context, client *firestore.Client, uid string) (map[string]interface{}, string, error) {

	var err error
	var doc *firestore.DocumentSnapshot

	iter := client.Collection("mediaServers").Where("uid", "==", uid).Documents(ctx)
	for {
		doc, err = iter.Next()
		// if err == iterator.Done {
		// 	break
		// }
		if err != nil {
			return nil, "", err
		}
		// fmt.Println(doc.Data())

		streamData := doc.Data()
		streamDocName := doc.Ref.ID

		if streamData != nil {
			return streamData, streamDocName, nil
		}
	}

}
