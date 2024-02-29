package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

// OpenBrowser is a function that opens a web browser with the given URL.
// It first checks the operating system and uses the appropriate command to open the browser.
//   - For Linux, it uses the "xdg-open" command. For Windows, it uses "rundll32" with "url.dll,FileProtocolHandler".
//   - For macOS, it uses the "open" command. If the platform is not supported, it returns an error.
//   - If there is an error while opening the browser, it prints the error message to the console.
func OpenBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
