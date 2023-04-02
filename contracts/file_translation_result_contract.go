package contracts

type FileTranslationResultContract struct {
	Id       int64  `json:"id"`
	FileName string `json:"file_name"`
	Content  string `json:"content"`
}
