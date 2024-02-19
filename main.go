package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func downloadFile(url, fileName string, useCurl bool) error {
	if useCurl {
		return downloadFileCurl(url, fileName)
	}

	u, err := url.Parse(url)
	if err != nil {
		return err
	}

	switch u.Scheme {
	case "http", "https":
		return downloadFileHTTP(url, fileName)
	case "ftp":
		return downloadFileFTP(url, fileName)
	default:
		return fmt.Errorf("unsupported protocol: %s", u.Scheme)
	}
}

func downloadFileHTTP(url, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func downloadFileFTP(url, fileName string) error {
	// Implement FTP download logic using Go's FTP library or your preferred FTP library.
	return fmt.Errorf("FTP download not implemented")
}

func downloadFileCurl(url, fileName string) error {
	cmd := exec.Command("curl", "-o", fileName, url)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing curl command: %v", err)
	}

	return nil
}

func makeExecutable(fileName string) error {
	err := os.Chmod(fileName, 0755)
	if err != nil {
		return err
	}

	return nil
}

func runFile(fileName string) error {
	cmd := exec.Command(fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	useCurl := flag.Bool("c", false, "Use curl for downloading files")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Please provide the URL and the local file name as arguments.")
		return
	}

	url := args[0]
	fileName := args[1]

	err := downloadFile(url, fileName, *useCurl)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		if *useCurl && strings.Contains(err.Error(), "executable file not found") {
			fmt.Println("Please make sure 'curl' is installed and available in your system's PATH.")
		}
		return
	}

	err = makeExecutable(fileName)
	if err != nil {
		fmt.Println("Error making file executable:", err)
		return
	}

	err = runFile(fileName)
	if err != nil {
		fmt.Println("Error running file:", err)
		return
	}
}
