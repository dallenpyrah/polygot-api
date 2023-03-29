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

func (f *FileLocationDetailsRepository) InsertFileLocationDetails(fileName string) (int64, error) {
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
