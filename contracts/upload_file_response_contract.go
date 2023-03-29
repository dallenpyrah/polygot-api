package contracts

type UploadFileResponseContract struct {
	FileName  string `json:"file_name"`
	RequestId int64  `json:"request_id"`
}
