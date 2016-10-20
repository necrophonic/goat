// Package initialise prepares workspaces to be used as test packs
//
// The chosen target to initialise must be a valid folder that exists and is empty
//
//	// Initialise the current working directory
//	if err := initialise.CurrentFolder(); err != nil {
//		log.Errorf("Failed to initialise: %v", err)
//	}
//
//	// Initialise a specified folder
//	if err := initialise.Folder("/path/to/folder"); err != nil {
//		log.Errorf("Failed to initialise: %v", err)
//	}
//
// In both instances, the following structure will be created:
//
//	+ ./
//	  + .gitignore
//	  + config/		# Holds configuration for the test pack
//	    + default.json
//	  + pages/		# Object page definitions
//	  + tests/		# Test specifications
//	  + results/		# Test output
//
package initialise

import (
	"errors"
	"io"
	"os"

	"github.com/necrophonic/log"
)

// Folder initialises a given folder as an empty GoAT test pack. Returns an error
// if the folder could not be initialised.
func Folder(target string) error {
	log.Debugf("initialiseing folder [%s]", target)

	if err := testFolderIsReadyToinitialise(target); err != nil {
		return errors.New("folder not ready to initialise: " + err.Error())
	}

	// Create the main folders
	os.Mkdir(target+"/tests", 0755)
	os.Mkdir(target+"/pages", 0755)
	os.Mkdir(target+"/results", 0755)
	os.Mkdir(target+"/config", 0755)

	// Create a gitignore as "target" contents shouldn't be commited in a test pack
	if err := writeGitIgnore(target); err != nil {
		return err
	}

	// Create a default configuration file
	if err := writeDefaultConfig(target); err != nil {
		return err
	}

	log.Infof("folder [%s] initialised", target)
	return nil
}

// CurrentFolder initialises the current working directory as a test pack. Returns an error
// if the current folder could not be initialised.
func CurrentFolder() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	return Folder(wd)
}

func writeGitIgnore(target string) error {
	log.Trace("writing .gitignore")
	gitignore, err := os.Create(target + "/.gitignore")
	if err != nil {
		return errors.New("failed to create .gitignore in target [" + target + "/.gitignore]: " + err.Error())
	}
	defer gitignore.Close()
	gitignore.WriteString("results/*")
	return nil
}

func writeDefaultConfig(target string) error {
	log.Trace("writing config/default.json")
	config, err := os.Create(target + "/config/default.json")
	if err != nil {
		return errors.New("failed to create config/default.json in target [" + target + "/config/default.json]: " + err.Error())
	}
	defer config.Close()
	config.WriteString("{" + "}") // TODO
	return nil
}

func testFolderIsReadyToinitialise(target string) error {
	log.Trace("testing folder is ready to initialise")
	if _, err := os.Stat(target); err != nil {
		if os.IsNotExist(err) {
			// Make the "not exists" error a bit more friendly
			return errors.New("target [" + target + "] does not exist")
		}
		return err
	}

	dir, err := os.Open(target)
	if err != nil {
		return errors.New("target [" + target + "] could not be opened: " + err.Error())
	}
	defer dir.Close()

	fi, err := dir.Stat()
	if err != nil {
		return err
	}

	if !fi.Mode().IsDir() {
		return errors.New("target [" + target + "] is not a folder")
	}

	if _, err = dir.Readdirnames(1); err != io.EOF {
		return errors.New("target [" + target + "] is not empty")
	}

	return nil
}
