package repository

import (
	"database/sql"

	"github.com/kunalsambharwal36/infra-assignments/internal/domain"
)

type ConfigRepository struct {
	db *sql.DB
}

func NewConfigRepository() *ConfigRepository {
	return &ConfigRepository{
		db: ConnectDB(),
	}
}

// Create or Update Config (UPSERT)
func (r *ConfigRepository) CreateConfig(config domain.Config) error {

	query := `
	INSERT INTO configs (
		id,
		host,
		port,
		app_name,
		log_level
	)
	VALUES ($1, $2, $3, $4, $5)

	ON CONFLICT (id)

	DO UPDATE SET
		host = EXCLUDED.host,
		port = EXCLUDED.port,
		app_name = EXCLUDED.app_name,
		log_level = EXCLUDED.log_level;
	`

	_, err := r.db.Exec(
		query,
		config.ID,
		config.Host,
		config.Port,
		config.AppName,
		config.LogLevel,
	)

	return err
}

// Get Config by ID
func (r *ConfigRepository) GetConfig(id string) (domain.Config, error) {

	var config domain.Config

	query := `
	SELECT
		id,
		host,
		port,
		app_name,
		log_level
	FROM configs
	WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&config.ID,
		&config.Host,
		&config.Port,
		&config.AppName,
		&config.LogLevel,
	)

	return config, err
}