package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type FilePathType int

const (
	NoneType FilePathType = iota
	FileType
	DirType
)

type Command struct {
	Dir    string   // directory where the command is executed (optional)
	Name   string   // command's name
	Args   []string // arguments to the command (optional)
	Output string   // output's Stdout bytes as string
	Lines  []string // split Output element into lines
	Error  error    // error occurs when the command is not able to be executed.
}

// CheckPathType checks if current fp (file's path) is a directory (DirType), a file (FileType) or none of these.
// This function wraps os.Stat() function that returns a FileInfo.
func CheckPathType(fp string) FilePathType {
	if info, err := os.Stat(fp); err == nil {
		if info.IsDir() {
			return DirType
		}
		return FileType
	}

	return NoneType
}

// Cmd is a function that executes a command in the specified directory.
// It takes a directory (dir) as optional, command, and optional arguments as parameters.
// It returns a Command struct containing the command's output and any errors encountered.
// If the directory is not empty, it checks if the path is a directory. If not, it returns an error.
// If the path is relative, it converts it to an absolute path.
//
// The command is executed using the exec.Command function with the specified directory if provided.
// The command's output is captured and stored in the Command struct's output and lines fields.
//
// If there is an error while executing the command, it is stored in the Command struct's error field.
// If the command produces any output, it is stored in the Command struct's output field as a string.
// The output is also split into lines and stored in the Command struct's lines field as a string slice.
//
// The Command struct is returned after the command execution.
//
// Example usage:
//
//	result := Cmd("/path/to/dir", "ls", "-l")
//	if result.Error != nil {
//	    fmt.Println("Error:", result.Error)
//	}
//	fmt.Println("Output:", result.Output)
//	fmt.Println("Lines:", result.Lines).
func Cmd(dir string, command string, args ...string) *Command {
	c := &Command{dir, command, args, "", nil, nil}

	if c.Dir != "" {
		if CheckPathType(dir) != DirType {
			c.Error = fmt.Errorf("specified dir is not a directory")
			return c
		}

		if !filepath.IsAbs(dir) {
			absPath, err := filepath.Abs(dir)

			if err != nil {
				c.Error = fmt.Errorf("it was unable to get absolute path from dir: " + err.Error())
				return c
			}

			c.Dir = absPath
		}
	}

	var out bytes.Buffer
	cmd := exec.Command(c.Name, c.Args...)
	if c.Dir != "" {
		cmd.Dir = c.Dir
	}
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		c.Error = err
		return c
	}

	if out.Len() > 0 {
		c.Output = out.String()
		c.Lines = strings.Split(strings.TrimSuffix(out.String(), "\n"), "\n")
	}

	return c
}

// GetFullCommand returns the full command string that will be executed.
// It concatenates the command name and its arguments with spaces in between.
// If there are no arguments, it returns only command's name
func (c *Command) GetFullCommand() string {
	if len(c.Args) > 0 {
		cmd := c.Name
		for _, e := range c.Args {
			cmd += " " + e
		}

		return cmd
	}

	return c.Name
}

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

// Raise throws a Panic when an error occurs and then whole application is interrupted.
func Raise(err error) {
	if err != nil {
		panic(err)
	}
}
