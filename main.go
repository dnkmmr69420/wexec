package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}

	fileURL := os.Args[1]
	err := downloadFile(fileURL)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}

	filePath, err := filepath.Abs(filepath.Base(fileURL))
	if err != nil {
		fmt.Printf("Error getting absolute file path: %v\n", err)
		return
	}

	err = makeExecutable(filePath)
	if err != nil {
		fmt.Printf("Error making file executable: %v\n", err)
		return
	}

	err = runFile(filePath)
	if err != nil {
		fmt.Printf("Error running file: %v\n", err)
		return
	}
}

func downloadFile(fileURL string) error {
	u, err := url.Parse(fileURL)
	if err != nil {
		return err
	}

	filename := filepath.Base(u.Path)
	if filename == "" {
		filename = "downloaded_file"
	}

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	var resp *http.Response
	if u.Scheme == "ftp" {
		resp, err = ftpDownload(fileURL)
	} else {
		resp, err = http.Get(fileURL)
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func ftpDownload(fileURL string) (*http.Response, error) {
	// Implement FTP download logic here
	// This is just a placeholder function
	return nil, fmt.Errorf("FTP download not implemented")
}

func makeExecutable(filePath string) error {
	if runtime.GOOS != "windows" {
		return os.Chmod(filePath, 0755)
	}
	return nil
}

func runFile(filePath string) error {
	cmd := exec.Command(filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
