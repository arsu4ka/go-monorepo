package config

import (
	"fmt"
	"os"
	"strconv"
)

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r Redis) GetAddress() string {
	return fmt.Sprintf("%s:%s", r.Host, fmt.Sprint(r.Port))
}

func (r Redis) GetConnectionString() string {
	authPart := ""
	if !(r.Username == "" && r.Password == "") {
		authPart = fmt.Sprintf("%s:%s@", r.Username, r.Password)
	}

	return fmt.Sprintf("redis://%s%s:%s", authPart, r.Host, fmt.Sprint(r.Port))
}

func GetRedisFromEnv() (Redis, error) {
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return Redis{}, err
	}

	host := os.Getenv("REDIS_HOST")
	username := os.Getenv("REDIS_USERNAME")
	password := os.Getenv("REDIS_PASSWORD")
	return Redis{Host: host, Port: port, Username: username, Password: password}, nil
}
