package domain

type FileSystemBucketSpec struct {
	// DirPath is the absolute or relative path to the directory that contains the bucket files
	DirPath string `yaml:"dir_path" json:"dir_path"`
}

func NewFileSystemBucketSpecFromMap(values map[string]interface{}) *FileSystemBucketSpec {
	return &FileSystemBucketSpec{
		DirPath: values["dir_path"].(string),
	}
}
