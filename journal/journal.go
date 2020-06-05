// Package journal provides functions for creating and editing journals
package journal

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/json"

	"github.com/mattroseman/gournal/journal/entry"
)

var saveDirName = ".gournal"

// Journal represents a journal which contains an array of entries
type Journal struct {
	Name string
	Entries []entry.Entry
	NumEntries int
}

// getJournalPath finds the user's home directory and joins it with the save dir and the given journal name to get the file path to the journal file.
// It returns a string that's the file path to the journal with the given name, and an error if something goes wrong getting the user's home directory.
func getJournalPath(name string) (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(userHomeDir, saveDirName, name + ".json"), nil
}

// Get reads and unmarshals a journal JSON file.
// It returns a Journal instance and an error if somethin wen't wrong with reading the save file.
// If no journal file was found, a zero instance of Journal is returned.
func Get(name string) (*Journal, error) {
	journal := &Journal{
		Name: name,
	}

	journalFilePath, err := getJournalPath(name)
	if err != nil {
		return journal, err
	}

	// check to see if the journal file already exists
	if _, err := os.Stat(journalFilePath); os.IsNotExist(err) {
		return journal, nil
	} else if err != nil {
		return journal, err
	}

	// read the journal file
	data, err := ioutil.ReadFile(journalFilePath)
	if err != nil {
		return journal, err
	}

	// unmarshal the journal file data into the journal instance
	if err := json.Unmarshal(data, journal); err != nil {
		return journal, err
	}

	return journal, nil
}

// Save updates the save file for this journal with it's current values.
// It returns an error if something goes wrong saving the journal to a file.
func (j Journal) Save() error {
	// marshal journal to JSON
	data, err := json.Marshal(j)
	if err != nil {
		return err
	}

	// truncate and rewrite to the journal file
	journalFilePath, err := getJournalPath(j.Name)
	if err != nil {
		return err
	}

	journalFile, err := os.OpenFile(journalFilePath, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0660)
	defer journalFile.Close()
	if err != nil {
		return err
	}

	if _, err := journalFile.Write(data); err != nil {
		return err
	}

	return nil
}

// AddEntry adds the given entry to the Journal instance j.
func (j *Journal) AddEntry(entry entry.Entry) {
	j.Entries = append(j.Entries, entry)
	j.NumEntries++

	return
}
