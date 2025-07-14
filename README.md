# PrintDir

A Go library for printing directory trees with support for multiple filesystem implementations.

## Features

- Print directory trees using different filesystem implementations:
  - Standard library filepath.Walk
  - Standard library fs.FS
  - [Afero](github.com/spf13/afero) filesystem 
- Visual representation with emoji icons (📁 for directories, 📄 for files)
- Proper indentation based on directory depth

## Installation

```bash
go get github.com/gophersatwork/printdir
```

## Usage

### Using Standard Library filepath.Walk

```go
package main

import (
	"fmt"
	"os"
	"github.com/gophersatwork/printdir"
)

func main() {
	// Print the directory tree to stdout
	err := printdir.Tree(os.Stdout, "/path/to/directory")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// You can also print to a file
	file, err := os.Create("directory_tree.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	err = printdir.Tree(file, "/path/to/directory")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

### Using Standard Library fs.FS

```go
package main

import (
	"bytes"
	"fmt"
	"os"
	"github.com/gophersatwork/printdir"
)

func main() {
	// Create a directory filesystem
	fsys := os.DirFS("/path/to/directory")

	// Print the directory tree to stdout
	err := printdir.TreeFS(os.Stdout, fsys, ".")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// You can also print to a buffer
	var buf bytes.Buffer
	err = printdir.TreeFS(&buf, fsys, ".")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Directory tree:")
	fmt.Println(buf.String())
}
```

### Using Afero

```go
package main

import (
	"bytes"
	"fmt"
	"os"
	"github.com/spf13/afero"
	"github.com/gophersatwork/printdir"
)

func main() {
	// Create a real filesystem
	fs := afero.NewOsFs()

	// Print the directory tree to stdout
	err := printdir.TreeAfero(os.Stdout, fs, "/path/to/directory")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Or use an in-memory filesystem for testing
	memFs := afero.NewMemMapFs()
	memFs.MkdirAll("/path/to/directory", 0755)

	// Print to a buffer
	var buf bytes.Buffer
	err = printdir.TreeAfero(&buf, memFs, "/path/to/directory")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Directory tree:")
	fmt.Println(buf.String())
}
```

## API Reference

### Tree

```go
func Tree(w io.Writer, path string) error
```

Prints a directory tree using the standard library's filepath.Walk. Takes an io.Writer to write the output and a path to the directory.

### TreeFS

```go
func TreeFS(w io.Writer, fsys fs.FS, path string) error
```

Prints a directory tree using the standard library's fs.FS interface. Takes an io.Writer, an fs.FS implementation, and a path within the filesystem.

### TreeAfero

```go
func TreeAfero(w io.Writer, fs afero.Fs, path string) error
```

Prints a directory tree using the Afero filesystem abstraction. Takes an io.Writer, an Afero filesystem implementation, and a path within the filesystem.

## Testing

The library includes comprehensive tests for all implementations. The tests use both real temporary directories and in-memory filesystems to ensure proper functionality without affecting the user's filesystem.

To run the tests:

```bash
go test ./...
```