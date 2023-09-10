package siopao

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// Copy copies the contents of the given source (file) into the destination.
func (file *File) Copy(dest string) error {
	destination := Open(dest)
	_, err := write(destination, true, func() (*any, error) {
		err := file.openRead()
		if err != nil {
			return nil, err
		}
		defer file.close()
		_, err = io.Copy(destination.file, file.file)
		return nil, err
	})
	return err
}

// CopyWithHash works similar to Copy but also creates a hash of the contents.
func (file *File) CopyWithHash(dest string) (*string, error) {
	destination := Open(dest)
	return write(destination, true, func() (*string, error) {
		err := file.openRead()
		if err != nil {
			return nil, err
		}
		defer file.close()
		hash := sha256.New()
		teeReader := io.TeeReader(file.file, hash)
		if _, err = io.Copy(destination.file, teeReader); err != nil {
			return nil, err
		}
		sum := hex.EncodeToString(hash.Sum(nil))
		return &sum, nil
	})
}
