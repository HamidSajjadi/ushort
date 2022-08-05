package repositories

import "github.com/HamidSajjadi/ushort/internal"

type URLModel struct {
	Source    string
	Shortened string
	Views     int32
}

type URLRepository interface {
	GetOne(shortenedURL string) (url *URLModel, err error)
	Save(sourceURL string, shortURL string) (url *URLModel, err error)
	IncViews(shortenedURL string) (err error)
}

type InMemoryURLRepos struct {
	sourceToURL    map[string]*URLModel
	shortenedToURL map[string]*URLModel
	maxID          int32
}

func NewInMemoryRepo() URLRepository {
	return &InMemoryURLRepos{
		sourceToURL:    make(map[string]*URLModel),
		shortenedToURL: make(map[string]*URLModel),
		maxID:          0,
	}
}
func (i InMemoryURLRepos) GetOne(shortenedURL string) (url *URLModel, err error) {
	url, ok := i.shortenedToURL[shortenedURL]
	if !ok {
		err = internal.NotFoundErr
	}
	return
}

func (i InMemoryURLRepos) Save(sourceURL string, shortURL string) (url *URLModel, err error) {

	if _, ok := i.sourceToURL[sourceURL]; ok {
		return nil, internal.ConflictErr
	}
	if _, ok := i.shortenedToURL[shortURL]; ok {
		return nil, internal.ConflictErr
	}

	url = &URLModel{
		Source:    sourceURL,
		Shortened: shortURL,
		Views:     0,
	}

	i.sourceToURL[sourceURL] = url
	i.shortenedToURL[shortURL] = url
	return url, nil
}

func (i InMemoryURLRepos) IncViews(shortenedURL string) (err error) {
	if val, ok := i.shortenedToURL[shortenedURL]; ok {
		val.Views++
		return nil
	}
	return internal.NotFoundErr
}
