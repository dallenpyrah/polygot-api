package interfaces

type FileLocationDetailsWriter interface {
	InsertFileLocationDetails(fileName string) int64
}
