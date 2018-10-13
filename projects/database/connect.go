package database

import (
	"context"
	"fmt"
	"os"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	// databaseName is the name of the application database
	databaseName = "machinable"
	// resourceDefinitions is the collection for storing resource definitions
	resourceDefinitions = "%s.definitions"
	// resourceFormat is the string format for a resource, this should include an account specifier as well
	resourceFormat = "%s.resource.%s"
	// collections is the collection of user collections
	collections = "%s.collections"
	// collectionFormat is the string format for a user's created collection
	collectionFormat = "%s.collections.%s"
	// userCollection is the collection of users for a project
	userCollection = "%s.users"
	// tokenCollection is the collection of api tokens for a project
	tokenCollection = "%s.tokens"
)

// TokenDocs returns the formatted string of the collection name of the collection that stores project api tokens
func TokenDocs(projectSlug string) string {
	return fmt.Sprintf(tokenCollection, projectSlug)
}

// UserDocs returns the formatted string of the collection name of the collection that stores project users
func UserDocs(projectSlug string) string {
	return fmt.Sprintf(userCollection, projectSlug)
}

// ResourceDefinitions returns the formatted string of the collection name of the collection that stores resource definitons for a project
func ResourceDefinitions(projectSlug string) string {
	return fmt.Sprintf(resourceDefinitions, projectSlug)
}

// ResourceDocs returns the formatted string of the collection name of the collection that stores resources (documents) for a project for a resource (path name)
func ResourceDocs(projectSlug, resourcePath string) string {
	return fmt.Sprintf(resourceFormat, projectSlug, resourcePath)
}

// CollectionNames returns the formatted string of the collection name of the collection that stores the list of project collections
func CollectionNames(projectSlug string) string {
	return fmt.Sprintf(collections, projectSlug)
}

// CollectionDocs returns the formatted string of the collection name of the collection that stores the list of documents for a project collection
func CollectionDocs(projectSlug, collectionName string) string {
	return fmt.Sprintf(collectionFormat, projectSlug, collectionName)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Connect returns a *mongo.Database connection
func Connect() *mongo.Database {
	host := getEnv("MONGO_HOST", "localhost")
	port := getEnv("MONGO_PORT", "27017")
	client, err := mongo.Connect(context.Background(), fmt.Sprintf("mongodb://%s:%s", host, port), nil)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return client.Database(databaseName)
}

// Collection returns a *mongo.Collection connection
func Collection(col string) *mongo.Collection {
	return Connect().Collection(col)
}
