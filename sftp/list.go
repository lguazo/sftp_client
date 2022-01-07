package sftp

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lguazo/sftp_client/email"
	"github.com/pkg/sftp"
)

func ListFiles(sc sftp.Client, remoteDir string) (err error) {
	fmt.Fprintf(os.Stdout, "Listing [%s] ...\n\n", remoteDir)

	files, err := sc.ReadDir(remoteDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list remote dir: %v\n", err)
		return
	}

	for _, f := range files {
		var name, modTime, size string

		name = f.Name()
		modTime = f.ModTime().Format("01-02-2006 15:04:05")
		size = fmt.Sprintf("%12d", f.Size())

		if f.Name() == "file.txt" {
			fmt.Printf("File %s found", f.Name())
		}

		if f.IsDir() {
			name = name + "/"
			modTime = ""
			size = "DIR"
		}

		// Output each file name and size in bytes
		fmt.Fprintf(os.Stdout, "%19s %12s %s\n", modTime, size, name)
	}

	return
}

func CheckSftpFile(sc sftp.Client, remoteDir string) (err error) {
	files, err := sc.ReadDir(remoteDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list remote dir: %v\n", err)
		return
	}

	var fileName, fileDate string
	var fileState bool

	fileName = os.Getenv("FILE_NAME")
	fileDate = os.Getenv("FILE_DATE")
	fileWhen := strings.ToLower(os.Getenv("FILE_CONDITION"))

	currentDate := time.Now().Format("01/02/2006")
	c := time.Now().Format("01022006")

	for _, f := range files {
		var modTime string

		modTime = f.ModTime().Format("01/02/2006")

		if fileWhen == "now" {

			if f.Name() == fileName && modTime == currentDate {
				fmt.Printf("File %s found", f.Name())
				fileState = true
				break
			}

		} else if fileWhen == "customdate" {

			if f.Name() == fileName && modTime == fileDate {
				fmt.Printf("File %s found", f.Name())
				fileState = true
				break
			} else {
				fileState = false
			}

		} else if fileWhen == "customfilename" {

			customFile := fileName + "." + c

			if f.Name() == customFile && modTime == currentDate {
				fmt.Printf("File %s found", f.Name())
				fileState = true
				break
			} else {
				fileState = false
			}

		}
	}

	if fileState == false {
		email.SendEmail()
	}

	return
}
