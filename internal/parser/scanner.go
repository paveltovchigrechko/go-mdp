package parser

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Scanner struct {
	root      string
	extension string
	filenames []string
}

func NewScanner(root, ext string) *Scanner {
	s := &Scanner{
		root:      root,
		extension: ext,
	}
	return s
}

func (s *Scanner) ScanFilesWithExt() {
	dir := os.DirFS(s.root)

	fs.WalkDir(dir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println(err)
		}
		if strings.HasSuffix(d.Name(), s.extension) {
			absPath, err := filepath.Abs(path)
			if err != nil {
				log.Println(err)
			}
			s.filenames = append(s.filenames, absPath)
		}

		return nil
	})
}
