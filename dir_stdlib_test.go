package printdir

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestPrintDirTreeStdlibSuccess tests the successful case of Tree
func TestPrintDirTreeStdlibSuccess(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "printdir_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a directory structure
	testDir := filepath.Join(tempDir, "testdir")
	subDir1 := filepath.Join(testDir, "subdir1")
	subDir2 := filepath.Join(testDir, "subdir2")

	os.MkdirAll(subDir1, 0755)
	os.MkdirAll(subDir2, 0755)

	os.WriteFile(filepath.Join(testDir, "file1.txt"), []byte("test content"), 0644)
	os.WriteFile(filepath.Join(subDir1, "file2.txt"), []byte("test content"), 0644)

	// Create a buffer to capture output
	var buf bytes.Buffer

	// Call the function with the buffer
	err = Tree(&buf, testDir)
	if err != nil {
		t.Fatalf("Tree returned an error: %v", err)
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

// TestPrintDirTreeStdlibFailure tests the failure case of Tree
func TestPrintDirTreeStdlibFailure(t *testing.T) {
	// Create a buffer to capture output (though we don't expect any)
	var buf bytes.Buffer

	// Try to print a non-existent directory
	err := Tree(&buf, "nonexistent_directory_that_should_not_exist")
	if err == nil {
		t.Fatalf("Expected an error when printing a non-existent directory, but got none")
	}
}