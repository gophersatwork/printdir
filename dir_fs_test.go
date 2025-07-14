package printdir

import (
	"bytes"
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"
)

// TestPrintDirTreeFSSuccess tests the successful case of TreeFS
func TestPrintDirTreeFSSuccess(t *testing.T) {
	// Create a memory filesystem using fstest.MapFS
	memFS := fstest.MapFS{
		"file1.txt":         &fstest.MapFile{Data: []byte("test content"), Mode: 0644},
		"subdir1/file2.txt": &fstest.MapFile{Data: []byte("test content"), Mode: 0644},
		"subdir2/file3.txt": &fstest.MapFile{Data: []byte("test content"), Mode: 0644},
	}

	// Create a buffer to capture output
	var buf bytes.Buffer

	// Call the function with the buffer
	err := TreeFS(&buf, memFS, ".")
	if err != nil {
		t.Fatalf("TreeFS returned an error: %v", err)
	}

	// Get the output
	output := buf.String()

	// Check if the output contains the expected directories and files
	expectedEntries := []string{
		"├── 📁 subdir1",
		"├── 📁 subdir2",
		"├── 📄 file1.txt",
		"├── 📄 file2.txt",
		"├── 📄 file3.txt",
	}

	for _, entry := range expectedEntries {
		if !strings.Contains(output, entry) {
			t.Errorf("Expected output to contain %q, but it didn't. Output: %s", entry, output)
		}
	}
}

// TestPrintDirTreeFSFailure tests the failure case of TreeFS
func TestPrintDirTreeFSFailure(t *testing.T) {
	// Create an empty filesystem
	memFS := fstest.MapFS{}

	// Create a buffer to capture output (though we don't expect any)
	var buf bytes.Buffer

	// Try to print a non-existent directory
	err := TreeFS(&buf, memFS, "nonexistent")
	if err == nil {
		t.Fatalf("Expected an error when printing a non-existent directory, but got none")
	}
}

// errorFS is a filesystem that always returns an error for any operation
type errorFS struct{}

func (e errorFS) Open(name string) (fs.File, error) {
	return nil, fs.ErrNotExist
}

// TestPrintDirTreeFSOpenError tests the case where fs.Open returns an error
func TestPrintDirTreeFSOpenError(t *testing.T) {
	// Create an error filesystem
	errFS := errorFS{}

	// Create a buffer to capture output (though we don't expect any)
	var buf bytes.Buffer

	// Try to print a directory
	err := TreeFS(&buf, errFS, ".")
	if err == nil {
		t.Fatalf("Expected an error when using errorFS, but got none")
	}
}