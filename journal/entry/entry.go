// Package entry provides functions for creating new entries for journals.
package entry 

import (
	"fmt"
	"strconv"

	"os"
	"os/exec"
	"io/ioutil"

	"time"
	"math/rand"
)

var tmpFilePath = "/tmp/gournal.txt"

// Entry represents a journal entry a user has made.
type Entry struct {
	Id string
	Content string
	CreatedAt time.Time
}

// PromptNewEntry opens up a text editor in the terminal for the user to type a new journal entry.
// A temporary file is created at tmpFilePath while the user types, and the file is truncated before the function returns.
// After the user saves and closes the text editor, an Entry type is initialized and returned along with any errors.
func PromptNewEntry() (*Entry, error) {
	newEntry := &Entry{
		Id: "",
		Content: "",
		CreatedAt: time.Time{},
	}

	editorCmd := exec.Command("vim", tmpFilePath)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout

	// opens vim and waits until the user exits
	if err := editorCmd.Run(); err != nil {
		return newEntry, err
	}

	content, err := ioutil.ReadFile(tmpFilePath)
	if err != nil {
		return newEntry, err
	}
	newEntry.Content = string(content)

	// clear the file at tmpFilePath
	tmpFile, err := os.OpenFile(tmpFilePath, os.O_TRUNC, 0660)
	defer tmpFile.Close()
	if err != nil {
		return newEntry, err
	}

	newEntry.Id = fmt.Sprintf("%s%05s", time.Now().Format("20060102150405"), strconv.Itoa(rand.Intn(100000)))
	newEntry.CreatedAt = time.Now()

	return newEntry, nil
}
