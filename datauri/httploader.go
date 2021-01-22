package datauri

import (
	"net/url"

	"github.com/herb-go/fetcher"
)

var HTTPDataURILoader = DataURILoaderFunc(func(u *url.URL) ([]byte, error) {
	p := fetcher.NewPreset().With(fetcher.ParsedURL(u))
	data := []byte{}
	_, err := p.FetchAndParse(fetcher.Should200(fetcher.AsBytes(&data)))
	if err != nil {
		return nil, err
	}
	return data, nil
})

func init() {
	Register("http", HTTPDataURILoader)
	Register("https", HTTPDataURILoader)
}
