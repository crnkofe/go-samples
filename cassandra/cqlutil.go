package cassandra

import (
	"errors"
	"github.com/gocql/gocql"
)

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

func ReadRow(key string) (string, error) {
	if !setupConnection("cassandra.config") {
		return "", errors.New("cassandra: Invalid config or cluster unavailable.")
	}

	session, _ := config.cluster.CreateSession()
	defer session.Close()

	/*
		// insert a tweet
		if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
			"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
			log.Fatal(err)
		}

		var id gocql.UUID
		var text string

		if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
			"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Tweet:", id, text)

		// list all tweets
		iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
		for iter.Scan(&id, &text) {
			fmt.Println("Tweet:", id, text)
		}
		if err := iter.Close(); err != nil {
			log.Fatal(err)
		}
	*/

	return "", nil
}
