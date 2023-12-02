package main

import (
	"os/user"
	"path/filepath"
)

func pathToUid(uid_file string) (string, error) {
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	// Read Firebase UID produced from QR Setup
	path_to_uid := filepath.Join(currentUser.HomeDir, uid_file)

	return path_to_uid, nil

}