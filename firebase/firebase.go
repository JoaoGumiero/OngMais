package firebase

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/JoaoGumiero/OngMais/entities"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	Client *firestore.Client
	once   sync.Once
)

// InitFirebase inicializa e retorna o cliente Firestore.
func InitFirebase() *firestore.Client {
	once.Do(func() {
		ctx := context.Background()

		// Load env variables wiith the godo lib
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		// Check if credential it's not empty
		credentialsPath := os.Getenv("FIREBASE_CREDENTIALS")
		if credentialsPath == "" {
			log.Fatalf("FIREBASE_CREDENTIALS environment variable is not set")
		}

		// Caminho (tenho que arrumar) para o arquivo JSON que contém sua chave de serviço
		sa := option.WithCredentialsFile(credentialsPath)

		// Project ID from the environment variable
		projectID := os.Getenv("FIREBASE_PROJECT_ID")
		if projectID == "" {
			log.Fatalf("error getting Firestore client: project id is required to access Firestore")
		}

		// Inicializa o aplicativo com uma conta de serviço, concedendo acesso total ao Firebase
		app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectID}, sa)
		if err != nil {
			log.Fatalf("error initializing app: %v", err)
		}

		// Obtem um cliente Firestore a partir do aplicativo
		Client, err = app.Firestore(context.TODO())
		if err != nil {
			log.Fatalf("error getting Firestore client: %v", err)
		}

		log.Println("Init Firebase completed", Client)
	})
	return Client
}

func StoreStates(states []entities.SimplifiedState, c *firestore.Client) {
	if c == nil {
		log.Fatalf("Store States: Firestore client is not initialized")
	}
	ctx := context.Background()
	for _, state := range states {
		_, _, err := c.Collection("br-states").Add(ctx, state)
		if err != nil {
			log.Fatalf("Failed adding state: %v", err)
		}
	}
}

func StoreCities(cities []entities.SimplifiedCity, c *firestore.Client) {
	if c == nil {
		log.Fatalf("Store Cities: Firestore client is not initialized")
	}
	ctx := context.Background()
	for _, city := range cities {
		_, _, err := c.Collection("br-cities").Add(ctx, city)
		if err != nil {
			log.Fatalf("Failed adding city: %v", err)
		}
	}
}
