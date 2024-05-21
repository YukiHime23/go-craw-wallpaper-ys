package gocrawal

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/YukiHime23/go-craw-al/models"
)

func DownloadFile(URL, fileName string, pathTo string) error {
	fmt.Println("-> Start download <-")

	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	//Create an empty file
	if fileName == "" {
		fileName = path.Base(URL)
	}

	file, err := os.Create(pathTo + "/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the field
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("-> download done \"" + fileName + "\" <-")
	return nil
}

func IntInArray(arr []int, str int) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func CreateFolder(path string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error getting home directory: %v", err)
	}

	// Change to the home directory
	err = os.Chdir(homeDir)
	if err != nil {
		return "", fmt.Errorf("Error changing to home directory: %v", err)
	}

	// Create the new folder
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return "", fmt.Errorf("Error creating new folder: %v", err)
	}

	// Print the full path of the new folder
	newFolderPath := filepath.Join(homeDir, path)
	fmt.Println("New folder created at:", newFolderPath)
	return newFolderPath, nil
}

func CrawURL(queue <- chan models.AzurLane, name string)  {
  for _, v := range queue {
    fmt.Println("test")
    time.Sleep(time.Second)
  }
  fmt.Printf("Worker %s done and exit\n", name)
}

func StartCraw(list []models.AzurLane) <-chan models.AzurLane {
	queue := make(chan models.AzurLane, 100)

	go func() {
    for _, v := range list {
      queue <- v
			fmt.Printf("File %s has been enqueued\n", v.FileName)
    }

		close(queue)
	}()

	return queue
}


