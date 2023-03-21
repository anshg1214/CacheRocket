package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const POST_URL = "https://jsonplaceholder.typicode.com/posts"

const TODO_URL = "https://jsonplaceholder.typicode.com/todos"

const PORT = "8080"

var REDIS_URL = os.Getenv("REDIS_URL")

var CACHE_POST bool

var CACHE_TODO bool

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	if REDIS_URL == "" {
		REDIS_URL = "localhost:6379"
	}

	cachePostEnv := os.Getenv("CACHE_POST")

	if cachePostEnv == "false" {
		CACHE_POST = false
	} else {
		CACHE_POST = true
	}

	cacheTodoEnv := os.Getenv("CACHE_TODO")

	if cacheTodoEnv == "false" {
		CACHE_TODO = false
	} else {
		CACHE_TODO = true
	}
}
