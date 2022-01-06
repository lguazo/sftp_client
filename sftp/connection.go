package sftp

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func Conn() *ssh.Client {
	// Get SFTP URL, Port, User & Password from environment
	host := os.Getenv("SFTP_URL")
	user := os.Getenv("SFTP_USER")
	pass := os.Getenv("SFTP_PASSWORD")
	port := 22

	fmt.Fprintf(os.Stdout, "Connecting to %s ...\n", host)

	var auths []ssh.AuthMethod

	// Use password authentication if provided
	if pass != "" {
		auths = append(auths, ssh.Password(pass))
	}

	var sshconfig ssh.Config
	sshconfig.SetDefaults()
	cipherOrder := sshconfig.Ciphers
	sshconfig.Ciphers = append(cipherOrder, "3des", "blowfish", "3des-cbc", "blowfish-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc")

	// Initialize client configuration
	config := ssh.ClientConfig{
		Config: sshconfig,
		User:   user,
		Auth:   auths,
		// Uncomment to ignore host key check
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	addr := fmt.Sprintf("%s:%d", host, port)

	// Connect to server
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connecto to [%s]: %v\n", addr, err)
		os.Exit(1)
	}

	// defer conn.Close()

	return conn

	// Create new SFTP client
	// sc, err := sftp.NewClient(conn)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to start SFTP subsystem: %v\n", err)
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println("Connection Succesfully..")
	// }
	// defer sc.Close()

	// return sc
	// checkSftpFile(*sc, "/Home/ce_broker")
	// SendEmail()

	// listFiles(*sc, "/")
}
