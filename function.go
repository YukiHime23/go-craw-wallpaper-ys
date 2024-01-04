package downloadAL

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
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
