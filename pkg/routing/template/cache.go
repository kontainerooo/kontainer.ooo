package template

import "github.com/kontainerooo/kontainer.ooo/pkg/routing"

// Cache is a cache for routing.RouterConfig files
type Cache interface {
	// GetConf returns a routing.RouterConfig either from the cache or using the get function
	GetConf(refID uint, name string) (*routing.RouterConfig, error)

	// SetConf puts a conf into the cache
	SetConf(r *routing.RouterConfig) *routing.RouterConfig

	// EditConf edits an existing conf or gets it if its not exists and edits afterwards
	EditConf(r *routing.RouterConfig) *routing.RouterConfig

	// RemoveConf removes a routing.RouterConfig from the Cache
	RemoveConf(refID uint, name string)

	// UpdateConf removes a routing.RouterConfig and gets a new version from the Database
	UpdateConf(refID uint, name string) *routing.RouterConfig
}

type cache struct {
	m             map[uint]map[string]*routing.RouterConfig
	getRouterConf func(uint, string, *routing.RouterConfig) error
}

func (c *cache) GetConf(refID uint, name string) (*routing.RouterConfig, error) {
	ref, ok := c.m[refID]
	if ok {
		conf, ok := ref[name]
		if ok {
			return conf, nil
		}
	} else {
		c.m[refID] = make(map[string]*routing.RouterConfig)
	}

	conf := &routing.RouterConfig{}
	err := c.getRouterConf(refID, name, conf)
	if err != nil {
		return nil, err
	}

	c.m[refID][name] = conf
	return conf, nil
}

func (c *cache) changeConf(r *routing.RouterConfig, edit bool) *routing.RouterConfig {
	_, ok := c.m[r.RefID]
	if !ok {
		c.m[r.RefID] = make(map[string]*routing.RouterConfig)
	}

	conf, ok := c.m[r.RefID][r.Name]
	if edit && ok {
		if r.ListenStatement != nil {
			conf.ListenStatement = r.ListenStatement
		}

		if r.ServerName != nil && len(r.ServerName) != 0 {
			conf.ServerName = r.ServerName
		}

		if r.AccessLog != (routing.Log{}) {
			if r.AccessLog.Keyword != "" {
				conf.AccessLog.Keyword = r.AccessLog.Keyword
			}

			if r.AccessLog.Path != "" {
				conf.AccessLog.Path = r.AccessLog.Path
			}
		}

		if r.ErrorLog != (routing.Log{}) {
			if r.ErrorLog.Keyword != "" {
				conf.ErrorLog.Keyword = r.ErrorLog.Keyword
			}

			if r.ErrorLog.Path != "" {
				conf.ErrorLog.Path = r.ErrorLog.Path
			}
		}

		ssl := r.SSLSettings
		{
			if ssl.Certificate != "" {
				conf.SSLSettings.Certificate = ssl.Certificate
			}

			if ssl.CertificateKey != "" {
				conf.SSLSettings.CertificateKey = ssl.CertificateKey
			}

			if len(ssl.Ciphers) != 0 {
				conf.SSLSettings.Ciphers = ssl.Ciphers
			}

			if len(ssl.Protocols) != 0 {
				conf.SSLSettings.Protocols = ssl.Protocols
			}

			if len(ssl.Curve) != 0 {
				conf.SSLSettings.Curve = ssl.Curve
			}

			conf.SSLSettings.PreferServerCiphers = ssl.PreferServerCiphers
		}

		if len(r.LocationRules) != 0 {
			conf.LocationRules = r.LocationRules
		}

		if r.RootPath != "" {
			conf.RootPath = r.RootPath
		}

	} else if ok {
		c.m[r.RefID][r.Name] = r
	} else {
		conf, err := c.GetConf(r.RefID, r.Name)
		if err != nil {
			return nil
		}
		c.m[r.RefID][r.Name] = conf
		if edit {
			return c.changeConf(r, edit)
		}
	}

	return c.m[r.RefID][r.Name]
}

func (c *cache) SetConf(r *routing.RouterConfig) *routing.RouterConfig {
	return c.changeConf(r, false)
}

func (c *cache) EditConf(r *routing.RouterConfig) *routing.RouterConfig {
	return c.changeConf(r, true)
}

func (c *cache) RemoveConf(refID uint, name string) {
	_, ok := c.m[refID]
	if ok {
		c.m[refID][name] = nil
	}
}

func (c *cache) UpdateConf(refID uint, name string) *routing.RouterConfig {
	c.RemoveConf(refID, name)
	return c.EditConf(&routing.RouterConfig{
		RefID: refID,
		Name:  name,
	})
}

// NewCache returns a new Cache instance with the help of a get function
func NewCache(g func(uint, string, *routing.RouterConfig) error) Cache {
	return &cache{
		m:             make(map[uint]map[string]*routing.RouterConfig),
		getRouterConf: g,
	}
}
