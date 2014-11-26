package cassandra

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

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

func WriteRow(cf string, columns []string, row Dict) error {
	if !setupConnection("cassandra.config") {
		return errors.New("cassandra: Invalid config or cluster unavailable.")
	}

	session, err := config.cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	values := make([]string, 0)
	for _, column := range columns {
		sval := fmt.Sprint(row[column])
		if _, err := strconv.Atoi(sval); err == nil {
			values = append(values, sval)
		} else {
			values = append(values, fmt.Sprintf(`'%s'`, sval))
		}
	}

	columnsQueryStringPart := strings.Join(columns, ", ")
	valuesQueryStringPart := strings.Join(values, ", ")

	queryString := fmt.Sprintf(
		`INSERT INTO "%s" (%s) VALUES (%s)`,
		cf,
		columnsQueryStringPart,
		valuesQueryStringPart)
	fmt.Println(queryString)
	err = session.Query(queryString).Exec()
	if err == nil {
		return nil
	} else {
		return errors.New("cassandra: Could not insert row.")
	}
}
