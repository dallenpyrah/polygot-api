package interfaces

type FileLocationDetailsWriter interface {
	InsertFileLocationDetails(fileName string) (int64, error)
	InsertFileUploadStatus(id int64, status string) error
	GetFileUploadStatus(id int64) (string, error)
}
