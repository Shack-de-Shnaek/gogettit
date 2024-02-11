package desktop

import (
	"io/fs"
	"os"
)

func GetDesktopFiles(filter *string) ([]fs.DirEntry, error) {
	return os.ReadDir("/usr/share/applications")
}
