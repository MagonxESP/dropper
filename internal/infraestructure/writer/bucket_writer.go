package writer

import (
	"errors"
	"fmt"
	"github.com/MagonxESP/dropper/internal/domain"
)

func GetDownloadedFileWriterFromBucket(bucket *domain.Bucket) (domain.DownloadedFileWriter, error) {
	switch bucket.Kind {
	case domain.BucketFileSystemKind:
		return NewFileSystemWriter(domain.NewFileSystemBucketSpecFromMap(bucket.Spec.(map[string]interface{}))), nil
	default:
		return nil, errors.New(fmt.Sprintf("missing writer for %s bucket kind", bucket.Kind))
	}
}
