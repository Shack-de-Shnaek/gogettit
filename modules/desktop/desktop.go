package desktop

import (
	"io/fs"
	"os"
)

func GetDesktopFiles(searchString *string) ([]fs.DirEntry, error) {
	var files []fs.DirEntry
	files, err := os.ReadDir("/usr/share/applications")
	if err != nil {
			
	}
	return files, err
}
