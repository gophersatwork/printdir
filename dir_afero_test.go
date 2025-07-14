package printdir

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/afero"
)

// TestPrintDirTreeAferoSuccess tests the successful case of TreeAfero
func TestPrintDirTreeAferoSuccess(t *testing.T) {
	// Create a memory filesystem
	fs := afero.NewMemMapFs()

	// Create a directory structure
	fs.MkdirAll("testdir/subdir1", 0755)
	fs.MkdirAll("testdir/subdir2", 0755)
	afero.WriteFile(fs, "testdir/file1.txt", []byte("test content"), 0644)
	afero.WriteFile(fs, "testdir/subdir1/file2.txt", []byte("test content"), 0644)

	// Create a buffer to capture output
	var buf bytes.Buffer

	// Call the function with the buffer
	err := TreeAfero(&buf, fs, "testdir")
	if err != nil {
		t.Fatalf("TreeAfero returned an error: %v", err)
	}

	// Get the output
	output := buf.String()

	// Check if the output contains the expected directories and files
	expectedEntries := []string{
		"├── 📁 subdir1",
		"├── 📁 subdir2",
		"├── 📄 file1.txt",
		"├── 📄 file2.txt",
	}

	for _, entry := range expectedEntries {
		if !strings.Contains(output, entry) {
			t.Errorf("Expected output to contain %q, but it didn't. Output: %s", entry, output)
		}
	}
}

// TestPrintDirTreeAferoFailure tests the failure case of TreeAfero
func TestPrintDirTreeAferoFailure(t *testing.T) {
	// Create a memory filesystem
	fs := afero.NewMemMapFs()

	// Create a buffer to capture output (though we don't expect any)
	var buf bytes.Buffer

	// Try to print a non-existent directory
	err := TreeAfero(&buf, fs, "nonexistent")
	if err == nil {
		t.Fatalf("Expected an error when printing a non-existent directory, but got none")
	}
}