package config

import (
	"fmt"
	"strconv"
)

// DBENGINE get DB engine
func DBENGINE() string {
	return GetEnv("DB_ENGINE", "postgres")
}

// DBCONFIG get DB Config
func DBCONFIG() string {
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "5432")
	user := GetEnv("DB_USER", "postgres")
	pass := GetEnv("DB_PASS", "password")
	dbname := GetEnv("DB_NAME", "postgres")
	ssl := GetEnv("DB_SSL", "disable")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		pass,
		dbname,
		ssl,
	)
}

// JWTSECRETKEY get secret key for jwt
func JWTSECRETKEY() string {
	return GetEnv("JWT_SECRET_KEY", "secret")
}

// JWTEXPIRATIONHOUR get expire time for jwt
func JWTEXPIRATIONHOUR() int {
	val, _ := strconv.Atoi(GetEnv("JWT_EXPIRATION_HOUR", "24"))
	return val
}
