package cli

import (
	auth "OpenList/Go/service/auth"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var CLI_COLOR = struct {
	Reset     string
	Bold      string
	Cyan      string
	Green     string
	Yellow    string
	BlueUnder string
	Gray      string
	Red       string
}{
	Reset:     "\033[0m",
	Bold:      "\033[1m",
	Cyan:      "\033[36m",
	Green:     "\033[32m",
	Yellow:    "\033[33m",
	BlueUnder: "\033[4;34m",
	Gray:      "\033[90m",
	Red:       "\033[31m",
}

func HandleCLI() bool {
	cmd := os.Args[1]
	reader := bufio.NewReader(os.Stdin)

	switch cmd {
	case "setPassword":
		newPassword := ""
		if len(os.Args) >= 3 {
			newPassword = strings.TrimSpace(os.Args[2])
			if err := auth.SetSingleUserPassword(newPassword); err != nil {
				fmt.Printf("%s✗ Error: %v%s\n", CLI_COLOR.Red, err, CLI_COLOR.Reset)
				return false
			}
			fmt.Printf("%s%s✓ Password updated successfully.%s\n", CLI_COLOR.Bold, CLI_COLOR.Green, CLI_COLOR.Reset)
			return false
		}

		fmt.Printf("\n%s%s── Change Password ──%s\n\n", CLI_COLOR.Bold, CLI_COLOR.Cyan, CLI_COLOR.Reset)

		fmt.Printf("%s  New password%s (min 8 chars): \033[8m", CLI_COLOR.Cyan, CLI_COLOR.Reset)
		line1, _ := reader.ReadString('\n')
		fmt.Print("\033[0m\n")
		newPassword = strings.TrimSpace(line1)

		fmt.Printf("%s  Confirm password%s: \033[8m", CLI_COLOR.Cyan, CLI_COLOR.Reset)
		line2, _ := reader.ReadString('\n')
		fmt.Print("\033[0m\n")
		confirm := strings.TrimSpace(line2)

		if newPassword != confirm {
			fmt.Printf("\n%s%s✗ Passwords do not match.%s\n\n", CLI_COLOR.Bold, CLI_COLOR.Red, CLI_COLOR.Reset)
			return false
		}

		if err := auth.SetSingleUserPassword(newPassword); err != nil {
			fmt.Printf("\n%s%s✗ Error: %v%s\n\n", CLI_COLOR.Bold, CLI_COLOR.Red, err, CLI_COLOR.Reset)
			return false
		}

		fmt.Printf("\n%s%s✓ Password updated. You can now login.%s\n\n", CLI_COLOR.Bold, CLI_COLOR.Green, CLI_COLOR.Reset)
		return false

	case "help":
		fmt.Printf("\n%s%sOpenList CLI%s\n", CLI_COLOR.Bold, CLI_COLOR.Cyan, CLI_COLOR.Reset)
		fmt.Printf("%s──────────────────────────────%s\n", CLI_COLOR.Gray, CLI_COLOR.Reset)
		fmt.Printf("  %shelp%s          Show this help\n", CLI_COLOR.Yellow, CLI_COLOR.Reset)
		fmt.Printf("  %ssetPassword%s   Change the admin password\n", CLI_COLOR.Yellow, CLI_COLOR.Reset)
		fmt.Printf("%s──────────────────────────────%s\n\n", CLI_COLOR.Gray, CLI_COLOR.Reset)
		return false

	default:
		fmt.Printf("%s✗ Unknown command: %s%s\n", CLI_COLOR.Red, cmd, CLI_COLOR.Reset)
		fmt.Printf("  Run %shelp%s for available commands.\n", CLI_COLOR.Yellow, CLI_COLOR.Reset)
		return false
	}
}
