package repository

import (
	"github.com/MagonxESP/dropper/internal/domain"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type YamlBucketRepository struct{}

func NewYamlBucketRepository() *YamlBucketRepository {
	return &YamlBucketRepository{}
}

func readBucketYaml(path string) (*domain.Bucket, error) {
	var bucket domain.Bucket
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(content, &bucket); err != nil {
		return nil, err
	}

	return &bucket, nil
}

func getBucketsDirPath() string {
	path := os.Getenv("DROPPER_BUCKETS_DIR")

	if path == "" {
		path = "buckets"
	}

	return path
}

// All get all buckets
func (y *YamlBucketRepository) All() ([]domain.Bucket, error) {
	var buckets []domain.Bucket

	err := filepath.WalkDir(getBucketsDirPath(), func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && (strings.HasSuffix(d.Name(), ".yml") || strings.HasSuffix(d.Name(), ".yaml")) {
			bucket, err := readBucketYaml(path)

			if err != nil {
				return err
			}

			buckets = append(buckets, *bucket)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return buckets, nil
}

// FindByName find a bucket by name, if it not found returns nil
func (y *YamlBucketRepository) FindByName(name string) (*domain.Bucket, error) {
	buckets, err := y.All()

	if err != nil {
		return nil, err
	}

	for _, bucket := range buckets {
		if bucket.Name == name {
			return &bucket, nil
		}
	}

	return nil, nil
}
