package application

import "github.com/MagonxESP/dropper/internal/domain"

type BucketFinder struct {
	Repository domain.BucketRepository
}

func NewBucketFinder(repository domain.BucketRepository) *BucketFinder {
	return &BucketFinder{Repository: repository}
}

func (f *BucketFinder) All() ([]domain.Bucket, error) {
	return f.Repository.All()
}

func (f *BucketFinder) FindByName(name string) (*domain.Bucket, error) {
	return f.Repository.FindByName(name)
}
