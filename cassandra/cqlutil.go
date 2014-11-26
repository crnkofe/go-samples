package cassandra

import (
	"errors"
	"fmt"

	"github.com/gocql/gocql"
)

type Dict map[string]interface{}
type Rows []Dict

var config struct {
	cassandraConfig CassandraConfig
	initialized     bool
	cluster         *gocql.ClusterConfig
}

func setupConnection(configFile string) bool {
	if !config.initialized {
		cassandraConfig, err := LoadConfig(configFile)

		// connect to the cluster
		cluster := gocql.NewCluster(cassandraConfig.Hosts...)
		cluster.Keyspace = cassandraConfig.Keyspace
		cluster.Consistency = gocql.Quorum
		config.cluster = cluster

		if err != nil {
			panic(err)
		} else {
			config.initialized = true
			config.cassandraConfig = cassandraConfig
		}
	}

	return true
}

func ReadRows(cf string, key string) (Rows, error) {
	if !setupConnection("cassandra.config") {
		return Rows{}, errors.New("cassandra: Invalid config or cluster unavailable.")
	}

	session, err := config.cluster.CreateSession()
	if err != nil {
		return Rows{}, err
	}
	defer session.Close()
	queryString := fmt.Sprintf(`SELECT * FROM "%s" WHERE key = '%s'`, cf, key)
	iter := session.Query(queryString).Iter()
	defer iter.Close()

	ret := make(Rows, 0)
	value := make(map[string]interface{})
	for iter.MapScan(value) {
		ret = append(ret, value)
	}

	return ret, nil
}
