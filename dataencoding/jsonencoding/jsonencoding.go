package jsonencoding

import (
	"encoding/json"

	"github.com/herb-go/herbdata/dataencoding"
)

var Encoding = &dataencoding.Encoding{
	Marshal:   json.Marshal,
	Unmarshal: json.Unmarshal,
}
