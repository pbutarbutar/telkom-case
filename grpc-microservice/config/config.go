package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (c configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(paths ...Path) Config {
	pathsString := make([]string, len(paths))
	for i, path := range paths {
		pathsString[i] = string(path)
	}
	err := godotenv.Load(pathsString...)
	if err != nil {
		log.Println(err)
	}
	return &configImpl{}
}

type Path string

func NewPaths() []Path {
	return []Path{".env"}
}
