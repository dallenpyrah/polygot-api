package repositories

import (
	"context"
	"polygot-api/providers"
)

type FileRetrievalRepository struct {
	dbConnectionProvider providers.DbConnectionProvider
}

func NewFileRetrievalRepository(dbConnectionProvider providers.DbConnectionProvider) FileRetrievalRepository {
	return FileRetrievalRepository{
		dbConnectionProvider: dbConnectionProvider,
	}
}

func (f FileRetrievalRepository) GetFileUploadStatus(id int64) (string, error) {
	f.dbConnectionProvider.OpenConnection()

	connection, _ := f.dbConnectionProvider.GetConnection()

	sql := "SELECT status FROM file_upload_status WHERE file_upload_id = $1 ORDER BY id DESC LIMIT 1"

	var status string
	err := connection.QueryRow(context.Background(), sql, id).Scan(&status)
	if err != nil {
		return "", err
	}

	f.dbConnectionProvider.CloseConnection()

	return status, nil
}
