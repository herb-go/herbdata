package xmlencoding

import (
	"encoding/xml"

	"github.com/herb-go/herbdata/dataencoding"
)

var Encoding = &dataencoding.Encoding{
	Marshal:   xml.Marshal,
	Unmarshal: xml.Unmarshal,
}
