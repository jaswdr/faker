package faker

import (
	"testing"
)

func TestProjectRepositoryName(t *testing.T) {
	a := New().Project()
	NotExpect(t, "", a.RepositoryName())
}

func TestProjectCodeName(t *testing.T) {
	a := New().Project()
	NotExpect(t, "", a.CodeName())
}

func TestProjectCommitMessage(t *testing.T) {
	a := New().Project()
	NotExpect(t, "", a.CommitMessage())
}
