package minioclient

import (
	"log"

	config "github.com/alexnorgaard/eventsapp"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func GetClient() (*minio.Client, error) {
	config := config.GetConf()
	endpoint := "localhost:9000"
	accessKeyID := config.S3.Access_key
	secretAccessKey := config.S3.Secret_key
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up
	return minioClient, err
}
