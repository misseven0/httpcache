package diskcache

import (
	"os"
	"testing"

	"github.com/misseven0/httpcache/test"
)

func TestDiskCache(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "httpcache")
	if err != nil {
		t.Fatalf("TempDir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	test.Cache(t, New(tempDir))
}
