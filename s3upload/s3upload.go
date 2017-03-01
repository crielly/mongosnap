package s3upload

import (
	"compress/gzip"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	file, err := os.Open("/home/crielly/DEV/tc-vars.tfvars")
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	reader, writer := io.Pipe()

	go func() {
		gw := gzip.NewWriter(writer)
		io.Copy(gw, file)

		file.Close()
		gw.Close()
		writer.Close()
	}()

	uploader := s3manager.NewUploader(nil)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   reader,
		Bucket: aws.String("360pi-ops"),
		Key:    aws.String("mongosnap/somefile.gz"),
	})
	if err != nil {
		log.Fatalln("Failed to Upload", err)
	}

	log.Println("Successfully upload to", result.Location)
}
