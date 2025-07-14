package printdir

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// Tree prints a directory tree using the standard library
func Tree(w io.Writer, path string) error {
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if p == path {
			return nil
		}

		depth := strings.Count(p, string(os.PathSeparator))
		indent := strings.Repeat("│   ", depth-1)

		name := info.Name()
		if info.IsDir() {
			fmt.Fprintf(w, "%s├── 📁 %s\n", indent, name)
		} else {
			fmt.Fprintf(w, "%s├── 📄 %s\n", indent, name)
		}

		return nil
	})
	if err != nil {
		log.Printf("Failed to inspect the folder: %v", err)
		return err
	}

	return nil
}

// TreeAfero prints a directory tree using the afero package
func TreeAfero(w io.Writer, fs afero.Fs, path string) error {
	err := afero.Walk(fs, path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if p == path {
			return nil
		}

		depth := strings.Count(p, string(os.PathSeparator))
		indent := strings.Repeat("│   ", depth-1)

		name := info.Name()
		if info.IsDir() {
			fmt.Fprintf(w, "%s├── 📁 %s\n", indent, name)
		} else {
			fmt.Fprintf(w, "%s├── 📄 %s\n", indent, name)
		}

		return nil
	})
	if err != nil {
		log.Printf("Failed to inspect the folder: %v", err)
		return err
	}

	return nil
}

// TreeFS prints a directory tree using the standard library fs.FS
func TreeFS(w io.Writer, fsys fs.FS, path string) error {
	if path == "" {
		path = "."
	}

	// Calculate the base depth to handle relative paths correctly
	baseDepth := strings.Count(path, "/")
	if path != "." {
		baseDepth++
	}

	err := fs.WalkDir(fsys, path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if p == path {
			return nil
		}

		// Calculate the depth relative to the base path
		depth := strings.Count(p, "/") - baseDepth
		if depth < 0 {
			depth = 0
		}

		indent := strings.Repeat("│   ", depth)

		name := filepath.Base(p)
		if d.IsDir() {
			fmt.Fprintf(w, "%s├── 📁 %s\n", indent, name)
		} else {
			fmt.Fprintf(w, "%s├── 📄 %s\n", indent, name)
		}

		return nil
	})
	if err != nil {
		log.Printf("Failed to inspect the folder: %v", err)
		return err
	}

	return nil
}