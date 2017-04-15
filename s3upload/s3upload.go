package s3upload

import (
	"github.com/crielly/mongosnap/logger"
	zipTool "github.com/pierrre/archivefile/zip"
	"github.com/rlmcpherson/s3gof3r"
)

// Zip archives a file or directory
func Zip(dir, s3bucket, object string) error {

	keys, err := s3gof3r.EnvKeys()
	if err != nil {
		logger.Error.Println(err)
	}

	// Open bucket we want to write a file to
	s3 := s3gof3r.New("", keys)
	bucket := s3.Bucket(s3bucket)

	// open a PutWriter for S3 upload
	s3writer, err := bucket.PutWriter(object, nil, nil)
	if err != nil {
		logger.Error.Println(err)
	}
	defer s3writer.Close()

	progress := func(archivePath string) {
		logger.Info.Println(archivePath)
	}

	return zipTool.Archive(dir, s3writer, progress)
}
