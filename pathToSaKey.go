package main

import (
	"os/user"
	"path/filepath"
)

func pathToSaKey(sa_file string) (string, error) {
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	// Read Firebase UID produced from QR Setup
	path_to_sakey := filepath.Join(currentUser.HomeDir, "firepub", sa_file)

	return path_to_sakey, nil
}