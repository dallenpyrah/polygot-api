package providers

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
)

type DbConnectionProvider struct {
	connection *pgx.Conn
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")        // Look for the config in the current directory
	viper.AddConfigPath("./config") // Optionally, look for the config in the ./config folder

	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read the configuration file")
	}
}

func (p *DbConnectionProvider) OpenConnection() error {
	dsn := viper.GetString("database.dsn")
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		return err
	}

	p.connection = conn
	return nil
}

func (p *DbConnectionProvider) GetConnection() (*pgx.Conn, error) {
	if p.connection != nil {
		return p.connection, nil
	}

	err := p.OpenConnection()

	if err != nil {
		return nil, err
	}

	return p.connection, nil
}

func (p *DbConnectionProvider) CloseConnection() error {
	if p.connection == nil {
		return nil
	}

	err := p.connection.Close(context.Background())
	if err != nil {
		return err
	}

	p.connection = nil
	return nil
}
