package main

import (
	"io"
	"net/http"
	"os"
	"time"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

func download(remoteURL, localPath string) (err error) {
	var (
		outFile  *os.File
		response *http.Response
	)

	// First, try and create localPath for writing as outFile
	if outFile, err = os.Create(localPath); err == nil {
		defer outFile.Close()

		// Second, send a GET request to use remoteURL as response
		if response, err = httpClient.Get(remoteURL); err == nil {
			defer response.Body.Close()

			// Finally, write the resulting response.Body to outFile
			_, err = io.Copy(outFile, response.Body)
		}
	}

	// Return whichever error occurred first
	return err
}
