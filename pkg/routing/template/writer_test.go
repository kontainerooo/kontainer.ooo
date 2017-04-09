package template_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing/template"
	"github.com/lib/pq"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Writer", func() {
	Describe("New Writer", func() {
		It("Should return new writer", func() {
			w, err := template.NewWriter(template.Nginx, "/tmp")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(w).ShouldNot(BeNil())
		})

		It("Should return an error if router does not exist", func() {
			_, err := template.NewWriter(1000, "/tmp")
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if path is no directory", func() {
			_, err := template.NewWriter(template.Nginx, "/etc/hosts")
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if path does not exist", func() {
			_, err := template.NewWriter(template.Nginx, "-")
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Create File", func() {
		var testPath = "/tmp/test-template-kroo/"
		BeforeEach(func() {
			err := os.Mkdir(testPath, os.ModeDir|os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.RemoveAll(testPath)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should write Conf file", func() {
			w, _ := template.NewWriter(template.Nginx, testPath)

			c := &routing.RouterConfig{}

			err := w.CreateFile(c)
			Ω(err).ShouldNot(HaveOccurred())

			b, err := ioutil.ReadFile(fmt.Sprintf("%s/%d_%s.conf", testPath, c.RefID, c.Name))
			Ω(err).ShouldNot(HaveOccurred())
			Ω(b).ShouldNot(BeEmpty())
		})
	})

	Context("Cache", func() {
		Describe("New Cache", func() {
			It("Should return a new cache", func() {
				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				Ω(cache).ShouldNot(BeNil())
			})
		})

		Describe("SetConf", func() {
			It("Should put a config into the cache", func() {
				id, name := uint(1), "name"
				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(&routing.RouterConfig{
					RefID: id,
					Name:  name,
				})

				conf, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(conf.RefID).Should(BeEquivalentTo(id))
				Ω(conf.Name).Should(BeEquivalentTo(name))
			})
		})

		Describe("EditConf", func() {

			It("Should edit the listenstatement", func() {
				id, name := uint(1), "name"
				conf := &routing.RouterConfig{
					RefID: id,
					Name:  name,
				}

				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(conf)

				conf.ServerName = pq.StringArray{"test.com"}
				cache.EditConf(conf)

				cc, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cc).Should(BeEquivalentTo(conf))
			})

			It("Should edit the server name", func() {
				id, name := uint(1), "name"
				conf := &routing.RouterConfig{
					RefID: id,
					Name:  name,
				}

				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(conf)

				inet, _ := abstraction.NewInet("::1")
				conf.ListenStatement = &routing.ListenStatement{
					IPAddress: inet,
				}
				cache.EditConf(conf)

				cc, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cc).Should(BeEquivalentTo(conf))
			})

			It("Should edit the access log", func() {
				id, name := uint(1), "name"
				conf := &routing.RouterConfig{
					RefID: id,
					Name:  name,
				}

				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(conf)

				conf.AccessLog = routing.Log{
					Path:    "/var/kroo/1/name/alog",
					Keyword: "access",
				}
				cache.EditConf(conf)

				cc, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cc).Should(BeEquivalentTo(conf))
			})

			It("Should edit the error log", func() {
				id, name := uint(1), "name"
				conf := &routing.RouterConfig{
					RefID: id,
					Name:  name,
				}

				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(conf)

				conf.ErrorLog = routing.Log{
					Path:    "/var/kroo/1/name/elog",
					Keyword: "error",
				}
				cache.EditConf(conf)

				cc, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cc).Should(BeEquivalentTo(conf))
			})

			It("Should edit the ssl settings", func() {
				id, name := uint(1), "name"
				conf := &routing.RouterConfig{
					RefID: id,
					Name:  name,
				}

				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(conf)

				conf.ErrorLog = routing.Log{
					Path:    "/var/kroo/1/name/elog",
					Keyword: "error",
				}
				cache.EditConf(conf)

				cc, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cc).Should(BeEquivalentTo(conf))
			})

			It("Should edit the error log", func() {
				id, name := uint(1), "name"
				conf := &routing.RouterConfig{
					RefID: id,
					Name:  name,
				}

				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(conf)

				conf.SSLSettings = routing.SSLSettings{
					Certificate:         "cert",
					CertificateKey:      "cert.key",
					Ciphers:             []string{"DES"},
					Protocols:           []string{"TLSv1"},
					Curve:               "idonthaveaclue",
					PreferServerCiphers: true,
				}
				cache.EditConf(conf)

				cc, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cc).Should(BeEquivalentTo(conf))
			})

			It("Should edit the location rules", func() {
				id, name := uint(1), "name"
				conf := &routing.RouterConfig{
					RefID: id,
					Name:  name,
				}

				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(conf)

				conf.LocationRules = routing.LocationRules{&routing.LocationRule{}}
				cache.EditConf(conf)

				cc, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cc).Should(BeEquivalentTo(conf))
			})

			It("Should edit the root path", func() {
				id, name := uint(1), "name"
				conf := &routing.RouterConfig{
					RefID: id,
					Name:  name,
				}

				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(conf)

				conf.RootPath = "/var/root"
				cache.EditConf(conf)

				cc, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cc).Should(BeEquivalentTo(conf))
			})
		})

		Describe("RemoveConf", func() {
			It("Should remove a config from the cache", func() {
				id, name := uint(1), "name"
				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return nil })
				cache.SetConf(&routing.RouterConfig{
					RefID: id,
					Name:  name,
				})

				cache.RemoveConf(id, name)

				conf, err := cache.GetConf(id, name)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(conf.RefID).Should(BeEquivalentTo(0))
			})
		})

		Describe("UpdateConf", func() {
			It("Should update a conf using the get function", func() {
				id, name := uint(1), "name"
				conf := routing.RouterConfig{
					RefID: id,
					Name:  name,
				}

				cache := template.NewCache(func(id uint, s string, r *routing.RouterConfig) error {
					r.RefID = id
					r.Name = s
					return nil
				})
				cache.SetConf(&conf)

				conf.RootPath = "change"

				cc := cache.UpdateConf(id, name)

				Ω(cc.RefID).Should(BeEquivalentTo(id))
				Ω(cc.Name).Should(BeEquivalentTo(name))
			})

			It("Should return nil if an error occurs", func() {
				cache := template.NewCache(func(uint, string, *routing.RouterConfig) error { return errors.New("") })

				cc := cache.UpdateConf(0, "")
				Ω(cc).Should(BeNil())
			})
		})

		Context("Check", func() {
			Describe("Listen Statement", func() {
				It("Should validate a Listen Statement", func() {
					c := template.NewCheck(template.Nginx)
					err := c.ListenStatement(&routing.ListenStatement{
						IPAddress: abstraction.Inet("127.0.0.1"),
						Port:      1337,
						Keyword:   "ssl",
					})
					Ω(err).ShouldNot(HaveOccurred())
				})

				XIt("Should return an error if the ip is not in the available pool", func() {
				})

				It("Should return an error if the port is out of range", func() {
					c := template.NewCheck(template.Nginx)
					err := c.ListenStatement(&routing.ListenStatement{
						Port: 80,
					})
					Ω(err).Should(BeEquivalentTo(template.ErrPortRange))
				})

				It("Should return an error if the keyword isnt available in the router", func() {
					c := template.NewCheck(template.Nginx)
					err := c.ListenStatement(&routing.ListenStatement{
						Port:    1337,
						Keyword: "asdf",
					})
					Ω(err).Should(BeEquivalentTo(template.ErrKeyword))
				})

				It("Should return an error if the keyword isnt available", func() {
					c := template.NewCheck(2)
					err := c.ListenStatement(&routing.ListenStatement{
						Port:    1337,
						Keyword: "asdf",
					})
					Ω(err).Should(BeEquivalentTo(template.ErrKeyword))
				})
			})

			Describe("ServerName", func() {
				It("Should validate a ServerName", func() {
					c := template.NewCheck(template.Nginx)
					err := c.ServerName(pq.StringArray{"domain.com"})
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("Should return an error if the servername holds an invalid name", func() {
					c := template.NewCheck(template.Nginx)
					err := c.ServerName(pq.StringArray{"quatsch"})
					Ω(err).Should(BeEquivalentTo(template.ErrInvalidName))
				})
			})

			XDescribe("Path", func() {
				It("Should validate a path", func() {
					_ = template.NewCheck(template.Nginx)
				})

				It("Should return an error if ", func() {
					_ = template.NewCheck(template.Nginx)
				})
			})

			XDescribe("Log", func() {
				It("Should validate a log", func() {
					_ = template.NewCheck(template.Nginx)
				})

				It("Should return an error if ", func() {
					_ = template.NewCheck(template.Nginx)
				})
			})

			XDescribe("SSLSettings", func() {
				It("Should validate ssl settings", func() {
					_ = template.NewCheck(template.Nginx)
				})

				It("Should return an error if ", func() {
					_ = template.NewCheck(template.Nginx)
				})
			})

			XDescribe("LocationRule", func() {
				It("Should validate a location rule", func() {
					_ = template.NewCheck(template.Nginx)
				})

				It("Should return an error if ", func() {
					_ = template.NewCheck(template.Nginx)
				})
			})

			XDescribe("LocationRules", func() {
				It("Should validate LocationRules", func() {
					_ = template.NewCheck(template.Nginx)
				})

				It("Should return an error if ", func() {
					_ = template.NewCheck(template.Nginx)
				})
			})
		})

	})
})
