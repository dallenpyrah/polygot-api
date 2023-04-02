package repositories

import (
	"context"
	"log"
	"polygot-api/providers"
)

type FileLocationDetailsRepository struct {
	dbConnectionProvider providers.DbConnectionProvider
	logger               *log.Logger
}

func NewFileLocationDetailsRepository(dbConnectionProvider providers.DbConnectionProvider, logger *log.Logger) FileLocationDetailsRepository {
	return FileLocationDetailsRepository{
		dbConnectionProvider: dbConnectionProvider,
		logger:               logger,
	}
}

func (f FileLocationDetailsRepository) InsertFileLocationDetails(fileName string) (int64, error) {
	f.dbConnectionProvider.OpenConnection()

	connection, _ := f.dbConnectionProvider.GetConnection()

	sql := "INSERT INTO file_upload (file_name) VALUES ($1) RETURNING id"

	var id int64
	err := connection.QueryRow(context.Background(), sql, fileName).Scan(&id)

	if err != nil {
		return 0, err
	}

	f.dbConnectionProvider.CloseConnection()

	return id, nil
}

func (f FileLocationDetailsRepository) InsertFileUploadStatus(fileUploadId int64, status string) error {
	f.dbConnectionProvider.OpenConnection()

	connection, _ := f.dbConnectionProvider.GetConnection()

	sql := "INSERT INTO file_upload_status (file_upload_id, status) VALUES ($1, $2)"

	_, err := connection.Exec(context.Background(), sql, fileUploadId, status)
	if err != nil {
		return err
	}

	f.dbConnectionProvider.CloseConnection()

	return nil
}
