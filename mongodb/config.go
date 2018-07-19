package mongodb

import (
	"crypto/tls"
	"net"
	"net/url"

	"github.com/globalsign/mgo"
)

type Config struct {
	URL string
}

func (c *Config) loadAndValidate() (*mgo.Session, error) {
	mURI, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	qs := mURI.Query()
	ssl := qs.Get("ssl") == "true"
	qs.Del("ssl") // won't parse otherwise
	mURI.RawQuery = qs.Encode()

	dialInfo, err := mgo.ParseURL(mURI.String())
	if err != nil {
		return nil, err
	}

	if ssl {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			tlsConfig := &tls.Config{}
			return tls.Dial("tcp", addr.String(), tlsConfig)
		}
	}

	return mgo.DialWithInfo(dialInfo)
}
