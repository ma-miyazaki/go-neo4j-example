package internal

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Neo4jClient struct {
	driver  neo4j.Driver
	session neo4j.Session
}

var Neo4j = newNeo4jClient()

func newNeo4jClient() (client Neo4jClient) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)

	if driver, err = neo4j.NewDriver("bolt://neo4j:7687", neo4j.NoAuth(), func(config *neo4j.Config) {
		config.Encrypted = false
	}); err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	client = Neo4jClient{
		driver:  driver,
		session: session,
	}
	return
}

func (client Neo4jClient) Close() {
	if client.session != nil {
		client.session.Close()
	}
	if client.driver != nil {
		client.driver.Close()
	}
}

func (client Neo4jClient) Run(query string, params map[string]interface{}) (result neo4j.Result, err error) {
	result, err = client.session.Run(query, params)
	if err != nil {
		return
	}
	err = result.Err()
	return
}

func (client Neo4jClient) DeleteAll() (neo4j.Result, error) {
	return client.Run("MATCH (n) DETACH DELETE n", nil)
}
