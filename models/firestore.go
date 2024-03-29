package models

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
)

var fc *firestore.Client
var ctx context.Context
var FSEnabled bool

type DataCollected struct {
	Email string
}

func ConnectFirestore() {
	FSEnabled = os.Getenv("FS_ENABLED") == "true"
	if !FSEnabled {
		log.Println("Firestore integration disabled")
		return
	}
	// Set up Firestore client options
	ctx = context.Background()
	opt := option.WithCredentialsFile("./service-account.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to configure Firebase: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	// Assign client to global FirestoreClient variable
	fc = client
	log.Println("Firestore client initialized!")
}

func ReadFromFirestore(collection string, id string) (*firestore.DocumentSnapshot, error) {
	if !FSEnabled {
		log.Println("FS not enabled. No attempt to read data will be performed.")
		return nil, nil
	}
	docRef := fc.Collection(collection).Doc(id)
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}
	if doc.Exists() {
		//fmt.Printf("Document data: %#v\n", doc.Data())
		return doc, nil
	}
	fmt.Println("No such document!")
	return nil, nil
}

func WriteToFirestore(collectionName string, doc interface{}) (string, error) {
	if !FSEnabled {
		log.Println("FS not enabled. No attempt to write data will be performed.")
		return "MOCK_ID", nil
	}
	data, _, err := fc.Collection(collectionName).Add(ctx, doc)
	id := data.ID
	if err != nil {
		return "", err
	}
	return id, nil
}

func UpdateFirestoreDocument(collectionName string, docId string, docData interface{}) (string, error) {
	// Get a reference to the document
	docRef := fc.Collection(collectionName).Doc(docId)

	// Set the document data
	_, err := docRef.Set(ctx, docData)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Document successfully written with ID: %s!\n", docId), nil
}

func CheckIfEmailAlreadyExists(collection string, email string) (bool, error) {
	if !FSEnabled {
		log.Println("FS not enabled. No attempt to read email data will be performed.")
		return false, nil
	}
	iter := fc.Collection(collection).Where("email", "==", email).Limit(1).Documents(ctx)
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
