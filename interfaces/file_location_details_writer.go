package interfaces

type FileLocationDetailsWriter interface {
	InsertFileLocationDetails(fileName string) (int64, error)
	InsertFileUploadStatus(id int64, status string) error
}
