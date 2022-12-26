package downloader

import (
	"github.com/MagonxESP/dropper/internal/domain"
	"testing"
)

func AssertDownloadedFile(t *testing.T, actual *domain.DownloadedFile, expectedFilename string) {
	if actual.Name != expectedFilename {
		t.Errorf("expected %s got %s", expectedFilename, actual.Name)
	}

	if actual.Content == nil {
		t.Error("file content is empty")
	}
}
