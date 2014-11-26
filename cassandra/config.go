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

func LoadConfig(filename string) (CassandraConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return CassandraConfig{}, err
	}

	decoder := json.NewDecoder(file)
	config := CassandraConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
		return CassandraConfig{}, err
	}
	return config, nil
}
