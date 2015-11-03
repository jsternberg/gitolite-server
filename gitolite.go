package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/jsternberg/gitolite"
	"golang.org/x/crypto/ssh"
)

func loadPrivateKeyFromFile() (ssh.Signer, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	privateKeyPath := filepath.Join(user.HomeDir, ".ssh/id_rsa")

	pemBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load server private key: %s", err)
	}
	return ssh.ParsePrivateKey(pemBytes)
}

func realMain() int {
	privateKey, err := loadPrivateKeyFromFile()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	config := gitolite.DefaultConfig()
	config.AddHostKey(privateKey)

	server := gitolite.New(config)
	if err := server.ListenAndServe(":1997"); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(realMain())
}
