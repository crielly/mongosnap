package s3upload

import (
	"compress/gzip"
	"io"
	"log"
	"os"

	"github.com/rlmcpherson/s3gof3r"
)

// log.Fatal() implies os.Exit(1)
func logerror(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

// S3upload streams compressed output to S3
func S3upload(toarchive, s3bucket, object string) {
	keys, err := s3gof3r.EnvKeys()

	// Open bucket we want to write a file to
	s3 := s3gof3r.New("", keys)
	bucket := s3.Bucket(s3bucket)

	// open a PutWriter for S3 upload
	s3writer, err := bucket.PutWriter(object, nil, nil)
	logerror(err)
	defer s3writer.Close()

	// Open a compressed writer to handle gzip and pass it to S3 writer
	zipwriter := gzip.NewWriter(s3writer)
	defer zipwriter.Close()

	// Open files we want archived
	file, err := os.Open(toarchive)
	logerror(err)
	defer file.Close()

	// Pass opened file to compression writer
	_, err = io.Copy(zipwriter, file)
	logerror(err)

}
