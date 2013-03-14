package btcreg

type Config struct {
    Port string
    PostgresServer string
}

var Conf Config

func LoadConfig() {
}

func (c Config) SqlDriver() string {
  return "sqlite3"
}

func (c Config) SqlOpen() string {
  return "btcreg.sqlite3"
}
