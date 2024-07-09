package firebase

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// InitFirebase inicializa e retorna o cliente Firestore.
func InitFirebase() *firestore.Client {
	ctx := context.Background()

	// Caminho (tenho que arrumar) para o arquivo JSON que contém sua chave de serviço
	sa := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_PATH"))

	// Inicializa o aplicativo com uma conta de serviço, concedendo acesso total ao Firebase
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	// Obtenha um cliente Firestore a partir do aplicativo
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Firestore client: %v", err)
	}

	return client
}
