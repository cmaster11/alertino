package alertino

import (
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/cmaster11/alertino/features/config"
	"github.com/cmaster11/alertino/platform"
	"github.com/cmaster11/alertino/platform/util"
)

type ArangoDBCollections struct {
	Sessions driver.Collection
}

func setupArangoDB(config *config.ArangoDBConfiguration) *ArangoDBCollections {
	var endpoints []string

	for _, endpointConfig := range config.Coordinators {
		endpoints = append(endpoints, fmt.Sprintf("http://%s:%d", endpointConfig.Host, endpointConfig.Port))
	}

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: endpoints,
	})
	util.PanicIfError(err)

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.Username, config.Password),
	})
	util.PanicIfError(err)

	db, err := client.Database(nil, config.Database)
	util.PanicIfError(err)

	// --- Initialize collections
	collectionSessions := arangoDBGetOrCreateCollection(db, "sessions", &driver.CreateCollectionOptions{})
	_, _, err = collectionSessions.EnsureTTLIndex(nil, platform.IndexFieldExpiresAt, 1, nil)
	util.PanicIfError(err)

	collections := &ArangoDBCollections{
		Sessions: collectionSessions,
	}

	return collections
}

func arangoDBGetOrCreateCollection(db driver.Database, collectionName string, options *driver.CreateCollectionOptions) driver.Collection {
	if collectionName == "" {
		panic("empty collection name")
	}

	exists, err := db.CollectionExists(nil, collectionName)
	util.PanicIfError(err)

	if !exists {
		collection, err := db.CreateCollection(nil, collectionName, options)
		util.PanicIfError(err)
		return collection
	}

	collection, err := db.Collection(nil, collectionName)
	util.PanicIfError(err)

	return collection
}
