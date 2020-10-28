package tomlencoding

import (
	"github.com/herb-go/herbdata/dataencoding"
	"github.com/pelletier/go-toml"
)

var Encoding = &dataencoding.Encoding{
	Marshal:   toml.Marshal,
	Unmarshal: toml.Unmarshal,
}
