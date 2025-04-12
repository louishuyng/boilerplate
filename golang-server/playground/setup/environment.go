package playground_setup

import (
	"fmt"
	"os"
	"rz-server/internal/common/interfaces"

	"github.com/joho/godotenv"
)

var _ interfaces.Environment = (*Environment)(nil)

type Environment struct {
}

func NewEnvironment() *Environment {
	err := godotenv.Load()

	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err))
	}

	return &Environment{}
}

func (e *Environment) GetEnv(space interfaces.SPACE, key string) string {
	prefix := ""

	switch space {
	case interfaces.SERVER:
		prefix = "SERVER_"
	case interfaces.DATABASE:
		prefix = "DB_"
	}

	environmentKey := fmt.Sprintf("%s%s", prefix, key)
	value := os.Getenv(environmentKey)

	if value == "" {
		panic(fmt.Sprintf("Environment variable %s not set", environmentKey))
	}

	return value
}
