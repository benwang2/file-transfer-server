package env

import "os"

func LoadVars() {
	os.Setenv("DB_USERNAME", "username")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "mydb")
}
