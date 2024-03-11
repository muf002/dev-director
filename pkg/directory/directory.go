package directory

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type searchResult struct {
	dirName  string
	found    bool
	duration time.Duration
}

func searchDirectory(dirName string, wg *sync.WaitGroup, results chan<- searchResult, editor string) {
	start := time.Now()

	var found bool
	filepath.WalkDir("/Users/usmanfarooq/Documents/", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			return nil // Skip files, only process directories
		}
		if d.Name() == dirName {
			openInEditor(path, editor)
			found = true
			return filepath.SkipDir
		}
		return nil
	})
	end := time.Now()
	duration := end.Sub(start)
	results <- searchResult{dirName, found, duration}
	wg.Done()
}

func OpenDirectories(dirNames []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose the editor to open directories:")
	fmt.Println("1. VSCode")
	fmt.Println("2. Neovim")
	fmt.Print("Enter choice (1 or 2): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var editor string
	switch choice {
	case "1":
		editor = "vscode"
	case "2":
		editor = "nvim"
	default:
		fmt.Println("Invalid choice")
		return
	}
	fmt.Println(dirNames)
	start := time.Now()
	var wg sync.WaitGroup
	results := make(chan searchResult, len(dirNames))
	for _, dirName := range dirNames {
		wg.Add(1)
		go searchDirectory(dirName, &wg, results, editor)
	}
	wg.Wait()
	close(results)

	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Time taken to open all files: %s\n", duration)

}

func openInEditor(dir, editor string) {
	var editorCmd *exec.Cmd
	switch editor {
	case "vscode":
		editorCmd = exec.Command("code", "--new-window", dir)
	case "nvim":
		script := fmt.Sprintf(`
            tell application "iTerm"
                create window with default profile
                tell current session of current window
                    write text "cd %s && nvim ."
                end tell
            end tell`, dir)
		editorCmd = exec.Command("osascript", "-e", script)
	default:
		fmt.Println("No valid editor selected")
		return
	}

	if err := editorCmd.Start(); err != nil {
		fmt.Printf("Failed to open directory %s in editor: %s\n", dir, err)
	}
	// Implement full-screen logic if required and possible
}
