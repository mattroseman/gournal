package main

import (
	"os"
	"math/rand"
	"time"
	"log"

	"github.com/urfave/cli/v2"

	"github.com/mattroseman/gournal/journal"
	"github.com/mattroseman/gournal/journal/entry"
)

var defaultJournalName = "default"

func main() {
	// seed random for later usage
	rand.Seed(time.Now().Unix())

	app := cli.App{
		Name: "gournal",
		Usage: "Command line app to create journal entries",
		Action: newEntry,
		Commands: []*cli.Command{
			{
				Name: "new",
				Usage: "create a new journal entry",
				Action: newEntry,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newEntry(c *cli.Context) error {
	entry, err := entry.PromptNewEntry()
	if err != nil {
		return err
	}

	j, err := journal.Get(defaultJournalName)
	if err != nil {
		return err
	}

	j.AddEntry(*entry)

	if err = j.Save(); err != nil {
		return err
	}

	return nil
}
