package siopao

import "github.com/ShindouMihou/siopao/paopao"

// Write writes, or appends if the file exists, the content to the file.
// Anything other than string and []byte is marshaled into Json with the paopao.Marshal.
func (file *File) Write(t any) error {
	return file.wrtany(false, t)
}

// Overwrite overwrites the file and writes the content to the file.
// Anything other than string and []byte is marshaled into Json with the paopao.Marshal.
func (file *File) Overwrite(t any) error {
	return file.wrtany(true, t)
}

// WriteMarshal works like Write, but marshals anything other than string and []byte with the provided marshal.
func (file *File) WriteMarshal(marshal paopao.Marshaller, t any) error {
	return file.wrtmarshal(marshal, false, t)
}

// OverwriteMarshal works like Overwrite, but marshals anything other than string and []byte with the provided marshal.
func (file *File) OverwriteMarshal(marshal paopao.Marshaller, t any) error {
	return file.wrtmarshal(marshal, true, t)
}
