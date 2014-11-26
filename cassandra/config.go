package cassandra

import (
	"encoding/json"
	"log"
	"os"
)

type CassandraConfig struct {
	Hosts        []string
	Keyspace     string
	ColumnFamily string
}

func LoadConfig(filename string) CassandraConfig {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	config := CassandraConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
