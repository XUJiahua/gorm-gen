package template

import (
	"io/fs"
	"testing"
)

func TestEmbed(t *testing.T) {
	fs.WalkDir(BaseTemplates, ".", func(path string, d fs.DirEntry, err error) error {

		return nil
	})
}
