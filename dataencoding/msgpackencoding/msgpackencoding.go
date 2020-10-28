package msgpackencoding

import (
	"github.com/herb-go/herbdata/dataencoding"
	"github.com/vmihailenco/msgpack"
)

var Encoding = &dataencoding.Encoding{
	Marshal: func(v interface{}) ([]byte, error) {
		return msgpack.Marshal(v)
	},
	Unmarshal: func(data []byte, v interface{}) error {
		return msgpack.Unmarshal(data, v)

	},
}
