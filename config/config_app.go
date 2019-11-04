package config

// Represents the configuration for the app
type AppConfig struct {
	ListenAddr *string `short:"a" long:"listenAddr" description:"Address (HOST:PORT) to serve API endpoints on"`

	// MongoDB connection string
	MongoDBConnString string `short:"m" long:"mongoDBConnString" required:"true" description:"MongoDB connection string"`
}
