package config

import (
	"fmt"
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	rdb := NewRedisClient()
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis successfully")
}
