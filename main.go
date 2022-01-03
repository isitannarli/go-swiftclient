package main

import (
	"flag"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/swauth"
	"os"
	"path/filepath"
)

func swiftClient(authUrl string, authUser string, authKey string) *gophercloud.ServiceClient {
	opts := swauth.AuthOpts{
		User: authUser,
		Key:  authKey,
	}

	providerClient, err := openstack.NewClient(authUrl)

	if err != nil {
		panic(err)
	}

	client, err := swauth.NewObjectStorageV1(providerClient, opts)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected")

	return client
}

func upload(filePath string, client *gophercloud.ServiceClient, containerName string) *objects.CreateHeader {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	createOpts := objects.CreateOpts{
		Content: file,
	}

	res := objects.Create(client, containerName, filepath.Base(file.Name()), createOpts)

	headers, err := res.Extract()

	if err != nil {
		panic(err)
	}

	return headers
}

func isDir(path string) bool {
	if pathAbs, err := filepath.Abs(path); err == nil {
		if fileInfo, err := os.Stat(pathAbs); !os.IsNotExist(err) && fileInfo.IsDir() {
			return true
		}
	}

	return false
}

func main() {
	var (
		authUrl       string
		authUser      string
		authKey       string
		containerName string
		path          string
	)

	flag.StringVar(&authUrl, "auth-url", "", "Auth url")
	flag.StringVar(&authUser, "auth-user", "", "Auth username")
	flag.StringVar(&authKey, "auth-key", "", "Auth key/password")

	flag.StringVar(&containerName, "container-name", "", "Container name (example: frontend/assets)")

	flag.StringVar(&path, "path", "", "File or directory path (example: ./file.txt OR ./directory/sub-directory)")

	flag.Parse()

	if len(authUrl) == 0 || len(authUser) == 0 || len(authKey) == 0 || len(containerName) == 0 || len(path) == 0 {
		fmt.Println("Args not valid!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	client := swiftClient(authUrl, authUser, authKey)

	var files []string

	if isDir(path) {
		err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if filePath == path {
				return nil
			}

			files = append(files, filePath)
			return nil
		})

		if err != nil {
			panic(err)
		}
	} else {
		files = append(files, path)
	}

	for _, file := range files {
		upload(file, client, containerName)
	}
}
