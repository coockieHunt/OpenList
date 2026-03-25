package cli

import (
	"OpenList/Go/auth"
	"bufio"
	"fmt"
	"os"
)

var CLI_COLOR = struct {
	Reset     string
	Bold      string
	Cyan      string
	Green     string
	Yellow    string
	BlueUnder string
	Gray      string
}{
	Reset:     "\033[0m",
	Bold:      "\033[1m",
	Cyan:      "\033[36m",
	Green:     "\033[32m",
	Yellow:    "\033[33m",
	BlueUnder: "\033[4;34m",
	Gray:      "\033[90m",
}

func HandleCLI() bool {
	cmd := os.Args[1]
	reader := bufio.NewReader(os.Stdin)

	switch cmd {
	case "genAdminToken":
		fmt.Print("\n")
		fmt.Printf("%s%s┌──────────────────────────────────────────────────────────┐%s\n", CLI_COLOR.Bold, CLI_COLOR.Cyan, CLI_COLOR.Reset)
		fmt.Printf("%s%s│                OPENLIST ADMIN GENERATOR                  │%s\n", CLI_COLOR.Bold, CLI_COLOR.Cyan, CLI_COLOR.Reset)
		fmt.Printf("%s%s└──────────────────────────────────────────────────────────┘%s\n", CLI_COLOR.Bold, CLI_COLOR.Cyan, CLI_COLOR.Reset)
		webURL := os.Getenv("WEB_URL")
		token, err := auth.GenerateAuthToken(0, true)
		if err != nil {
			fmt.Printf("Error generating auth token: %v\n", err)
			return false
		}

		var fullURL string
		if webURL != "" {
			fullURL = webURL + "/login?token=" + token
		} else {
			fullURL = "http://localhost:" + os.Getenv("WEB_PORT") + "/login?token=" + token
		}
		fmt.Printf("\n%s%sToken URL:%s %s\n", CLI_COLOR.Bold, CLI_COLOR.Cyan, CLI_COLOR.Reset, token)
		fmt.Printf("\n%s%sLogin URL:%s %s\n", CLI_COLOR.Bold, CLI_COLOR.Cyan, CLI_COLOR.Reset, fullURL)
		fmt.Printf("\n%s%sℹ️  INSTRUCTIONS%s\n", CLI_COLOR.Bold, CLI_COLOR.Yellow, CLI_COLOR.Reset)
		fmt.Printf("  1. Copy the URL above.\n")
		fmt.Printf("  2. Press %s[ENTER]%s to start the servers.\n", CLI_COLOR.Bold, CLI_COLOR.Reset)
		fmt.Printf("  3. Once started, paste the URL in your browser.\n")

		fmt.Printf("\n%s%s(OpenList will start and you can request a username)%s\n", CLI_COLOR.Gray, CLI_COLOR.Bold, CLI_COLOR.Reset)

		fmt.Printf("\n%s➜ Press Enter to continue...%s", CLI_COLOR.Green, CLI_COLOR.Reset)
		reader.ReadBytes('\n')

		return true
	case "help":
		fmt.Println("Usage: ./app [genAdminToken|help]")
		return false

	default:
		fmt.Printf("missing command: %s\n", cmd)
		return false
	}
}
