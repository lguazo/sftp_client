package sftp

import (
	"fmt"
	"os"

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
		// modTime = f.ModTime().Format("2006-01-02 15:04:05")
		modTime = f.ModTime().Format("01-02-2006 15:04:05")
		size = fmt.Sprintf("%12d", f.Size())

		if f.Name() == "file.txt" {
			fmt.Printf("File found itTTTT %s", f.Name())
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

	for _, f := range files {
		if f.Name() == "lic_status.dat" {
			fmt.Printf("File %s found", f.Name())
		}
	}

	return
}
