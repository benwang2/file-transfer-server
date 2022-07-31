package env

import "os"

func LoadVars() {
	os.Setenv("DB_USERNAME", "username")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_HOST", "host")
	os.Setenv("DB_PORT", "port")
	os.Setenv("DB_NAME", "dbname")
}
