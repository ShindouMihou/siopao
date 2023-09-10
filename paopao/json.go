package paopao

import "encoding/json"

type Marshaller func(v any) ([]byte, error)
type Unmarshaler func(data []byte, v any) error

var Marshal = json.Marshal
var Unmarshal = json.Unmarshal
