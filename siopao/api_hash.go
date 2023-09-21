package siopao

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io"
)

type ChecksumKind string

const (
	Sha512Checksum ChecksumKind = "sha512"
	Sha256Checksum ChecksumKind = "sha256"
	Md5Checksum    ChecksumKind = "md5"
)

// Checksum gets the checksum hash of the file's contents.
func (file *File) Checksum(kind ChecksumKind) (string, error) {
	f, err := file.openRead()
	if err != nil {
		return "", err
	}
	defer file.close(f)

	var hsh hash.Hash
	switch kind {
	case Sha256Checksum:
		hsh = sha256.New()
	case Md5Checksum:
		hsh = md5.New()
	case Sha512Checksum:
		hsh = sha512.New()
	default:
		return "", errors.New("unsupported checksum kind")
	}
	if _, err := io.Copy(hsh, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hsh.Sum(nil)), nil
}
