package domain

const (
	BucketFileSystemKind = "filesystem"
)

type Bucket struct {
	// The Bucket Name
	Name string `yaml:"name" json:"name"`
	// Kind is the Bucket type
	Kind string `yaml:"kind" json:"kind"`
	// Spec is the configuration related to the Bucket Kind
	Spec interface{} `yaml:"spec" json:"spec"`
}

type BucketRepository interface {
	All() ([]Bucket, error)
	FindByName(name string) (*Bucket, error)
}
