package skynet

import (
	"bytes"
	"fmt"
	"github.com/NebulousLabs/go-skynet"
	"github.com/fatih/color"
	"github.com/jay-dee7/sia-box/config"
	"github.com/jay-dee7/sia-box/crypto"
	snappy "github.com/klauspost/compress/s2"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func Upload(path string, faster bool) error {
	t := time.Now()
	color.Yellow("upload started at: %s", t.Format(time.RFC1123))

	//checkForFileChanges()

	newFile, err := compress(path, faster)
	if err != nil {
		return err
	}
	defer cleanup(newFile)
	color.Red("config path inside upload: %s", newFile)
	hash, err := skynet.UploadFile(newFile, skynet.DefaultUploadOptions)
	if err != nil {
		return fmt.Errorf("error uploading file to skynet: %w", err)
	}

	color.Green("path has been synced to skynet successfully: \n \t %s", hash)
	color.Green("upload finished in: \t %d ms", time.Since(t).Milliseconds())
	return nil
}

func compress(filename string, faster bool) (string, error) {

	src, err := os.Open(filename)
	input, err := ioutil.ReadAll(src)
	if err != nil {
		return "", fmt.Errorf("error while reading data from source: %w", err)
	}

	tmpFile := fmt.Sprintf("%s-%d", filename, time.Now().UnixNano())
	file, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if faster {
		comp := snappy.Encode(nil, input)
		_, _ = io.Copy(file, bytes.NewBuffer(comp))
		return file.Name(), nil
	}
	cfg, err := config.Read()
	if err != nil {
		color.Red("error while reading config: %s", err)
		return "", err
	}

	encData, err := crypto.EncryptAES([]byte(cfg.Password), []byte("hello world"))
	if err != nil {
		color.Red("error in ecnrypt: %s", cfg.Password)
		return "", err
	}

	comp := snappy.EncodeBetter(nil, encData)

	_, _ = io.Copy(file, bytes.NewBuffer(comp))

	return file.Name(), nil
}

func cleanup(file string) {
	_ = os.Remove(file)
}
