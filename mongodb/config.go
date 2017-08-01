package mongodb

import (
	"gopkg.in/mgo.v2"
)

type Config struct {
	URL string
}

func (c *Config) loadAndValidate() (*mgo.Session, error) {
	session, err := mgo.Dial(c.URL)
	if err != nil {
		return nil, err
	}

	return session, nil
}
