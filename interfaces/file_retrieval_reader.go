package interfaces

type FileRetrievalReader interface {
	GetFileUploadStatus(id int64) (string, error)
}
