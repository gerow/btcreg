package btcreg

type Config struct {
    Port string
    PostgresServer string
}

var Conf Config
var confLoaded bool = false

func LoadConfig() {
  if confLoaded {
    return
  }
}

func (c Config) SqlDriver() string {
  return "sqlite3"
}

func (c Config) SqlOpen() string {
  return "btcreg.sqlite3"
}
