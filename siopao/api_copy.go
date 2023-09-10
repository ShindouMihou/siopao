package siopao

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
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
func (file *File) CopyWithHash(kind ChecksumKind, dest string) (*string, error) {
	destination := Open(dest)
	return write(destination, true, func() (*string, error) {
		err := file.openRead()
		if err != nil {
			return nil, err
		}
		defer file.close()
		hsh := sha256.New()
		switch kind {
		case Sha256Checksum:
			hsh = sha256.New()
		case Md5Checksum:
			hsh = md5.New()
		case Sha512Checksum:
			hsh = sha512.New()
		default:
			return nil, errors.New("unsupported checksum kind")
		}
		teeReader := io.TeeReader(file.file, hsh)
		if _, err = io.Copy(destination.file, teeReader); err != nil {
			return nil, err
		}
		sum := hex.EncodeToString(hsh.Sum(nil))
		return &sum, nil
	})
}
