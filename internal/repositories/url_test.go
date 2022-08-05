package repositories

import (
	"github.com/HamidSajjadi/ushort/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

var urls = []*URLModel{
	{
		Source:    "source-1",
		Shortened: "shortened-1",
		Views:     1,
	}, {
		Source:    "source-2",
		Shortened: "shortened-2",
		Views:     0,
	},
}

func initInMemoryURLRepos() *InMemoryURLRepos {

	shortenedToURL := make(map[string]*URLModel)
	sourceToURL := make(map[string]*URLModel)
	for _, url := range urls {
		shortenedToURL[url.Shortened] = url
		sourceToURL[url.Source] = url
	}
	return &InMemoryURLRepos{
		sourceToURL:    sourceToURL,
		shortenedToURL: shortenedToURL,
	}
}

type expected struct {
	input  string
	hasErr bool
	err    error
	res    interface{}
}

func TestInMemoryURLRepos_GetOne(t *testing.T) {

	testMap := []expected{
		{
			input:  urls[0].Shortened,
			hasErr: false,
			err:    nil,
			res:    urls[0],
		},
		{
			input:  "fake-url",
			hasErr: true,
			err:    internal.NotFoundErr,
			res:    nil,
		},
	}

	repos := initInMemoryURLRepos()
	for _, test := range testMap {
		resp, err := repos.GetOne(test.input)
		if test.hasErr {
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, test.err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, test.res, resp)
		}
	}

}

func TestInMemoryURLRepos_IncVisits(t *testing.T) {
	testMap := []expected{
		{
			input:  urls[0].Shortened,
			hasErr: false,
			err:    nil,
			res:    urls[0].Views + 1,
		},
		{
			input:  "fake-url",
			hasErr: true,
			err:    internal.NotFoundErr,
			res:    nil,
		},
	}

	repos := initInMemoryURLRepos()
	for _, test := range testMap {
		err := repos.IncViews(test.input)
		if test.hasErr {
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, test.err)
		} else {
			assert.Nil(t, err)
			url := repos.shortenedToURL[test.input]
			assert.Equal(t, test.res, repos.shortenedToURL[test.input].Views)
			assert.Equal(t, test.res, url.Views)
			assert.Equal(t, test.res, repos.sourceToURL[url.Source].Views)
		}
	}
}
