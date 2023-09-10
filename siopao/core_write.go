package siopao

import "github.com/ShindouMihou/siopao/paopao"

func (file *File) wrt(trunc bool, bytes []byte) error {
	if _, err := write(file, trunc, func() (*any, error) {
		if _, err := file.file.Write(bytes); err != nil {
			return nil, err
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (file *File) wrtjson(trunc bool, t interface{}) error {
	return file.wrtmarshal(paopao.Marshal, trunc, t)
}

func (file *File) wrtmarshal(marshal paopao.Marshaller, trunc bool, t interface{}) error {
	bytes, err := marshal(t)
	if err != nil {
		return err
	}
	return file.wrt(trunc, bytes)
}

func (file *File) wrtany(trunc bool, t any) error {
	switch t.(type) {
	case string:
		return file.wrt(trunc, []byte(t.(string)))
	case []byte:
		return file.wrt(trunc, t.([]byte))
	default:
		return file.wrtjson(true, t)
	}
}
