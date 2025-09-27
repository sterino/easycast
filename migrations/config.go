package migrations

type Config struct {
	Host                     string `env:"DB_HOST,notEmpty"`
	Port                     string `env:"DB_PORT,notEmpty"`
	User                     string `env:"DB_USER,notEmpty"`
	Password                 string `env:"DB_PASSWORD,notEmpty"`
	DBName                   string `env:"DB_NAME,notEmpty"`
	SSLMode                  string `env:"DB_SSL_MODE,notEmpty"             envDefault:"disable"`
	MaxConns                 int    `env:"DB_MAX_CONNS"                     envDefault:"15"`
	MaxIdleConnections       int    `env:"DB_MAX_IDLE_CONNS"                envDefault:"10"`
	MaxConnIdleTimeInSeconds int    `env:"DB_MAX_CONN_IDLE_TIME_IN_SECONDS" envDefault:"300"`  // 5 minutes
	MaxConnLifeTimeInSeconds int    `env:"DB_MAX_CONN_LIFETIME_IN_SECONDS"  envDefault:"1500"` // 25 minutes

	// needed for tests, do nothing in production
	RunningInCI bool `env:"DB_RUNNING_IN_CI" envDefault:"false"`
}
