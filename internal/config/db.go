package config

import "fmt"

type DBConfig struct {
	PostgresDatabase string `envconfig:"POSTGRES_DATABASE" required:"true"`
	PostgresUser     string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresPort     uint16 `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresHost     string `envconfig:"POSTGRES_HOST" required:"true"`
}

func (c *DBConfig) ConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDatabase,
	)
}
