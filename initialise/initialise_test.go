package initialise

import (
	"testing"

	"github.com/necrophonic/assert"
)

const (
	errStub = "folder not ready to initialise: "
)

func TestInitialiseFolderDoesNotExist(t *testing.T) {
	assert.ErrorEquals(Folder("./invalid"), errStub+"target [./invalid] does not exist", t)
}

func TestInitialiseFolderIsNotEmpty(t *testing.T) {
	assert.ErrorEquals(Folder("."), errStub+"target [.] is not empty", t)
}

func TestInitialiseFolderIsNotAFolder(t *testing.T) {
	assert.ErrorEquals(Folder("./initialise.go"), errStub+"target [./initialise.go] is not a folder", t)
}
