package nosql

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AdditionalProperty struct {
	NewProperty string `bson:"new_property,omitempty"`
}

type Person struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name,omitempty"`
	Additional AdditionalProperty `bson:"additional,omitempty"`
	// Properties need capitals to be properly marshalled to/from db
}

type DatabaseConnections struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (c *DatabaseConnections) Close() {
	c.client.Disconnect(context.TODO())
}

func Connect() *DatabaseConnections {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")
	return &DatabaseConnections{client: client}
}

func (c *DatabaseConnections) NewCollection() {
	c.collection = c.client.Database("test").Collection("persons")
}

func (c *DatabaseConnections) InsertDoc() {
	newPerson := Person{Name: "Shiba", Additional: AdditionalProperty{NewProperty: "Dog"}}
	//newPerson := Person{Name: "New Person"}

	_, err := c.collection.InsertOne(context.TODO(), newPerson)

	if err != nil {
		fmt.Println(err)
	}
}

func (c *DatabaseConnections) FindDoc() Person {
	var result Person
	filter := bson.M{"name": "testing"}

	err := c.collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.Name)
	return result
}

func (c *DatabaseConnections) FindAll() []Person {
	var results []Person

	docs, err := c.collection.Find(context.TODO(), bson.D{})
	defer docs.Close(context.TODO())

	if err != nil {
		fmt.Println(err)
	}

	err = docs.All(context.TODO(), &results)

	if err != nil {
		fmt.Println(err)
	}

	return results
}

func (c *DatabaseConnections) DeleteDocument() {
	result, err := c.collection.DeleteOne(context.TODO(), bson.D{{}})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.DeletedCount)
		fmt.Println("Successful delete")
	}
}
