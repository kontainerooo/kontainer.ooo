package template_test

import (
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing/template"
	"github.com/lib/pq"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check", func() {
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

	Describe("Config", func() {
		It("Should validate a config", func() {
			c := template.NewCheck(template.Nginx)
			err := c.Config(&routing.RouterConfig{
				RefID: 1,
				Name:  "name",
				ListenStatement: &routing.ListenStatement{
					IPAddress: abstraction.Inet("127.0.0.1"),
					Port:      1337,
					Keyword:   "ssl",
				},
				ServerName: pq.StringArray{"domain.com"},
			}, false)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should return an error if the RefID isn't set", func() {
			c := template.NewCheck(template.Nginx)
			err := c.Config(&routing.RouterConfig{}, false)
			Ω(err).Should(BeEquivalentTo(template.ErrNoRefID))
		})

		It("Should return an error if the Name isn't set", func() {
			c := template.NewCheck(template.Nginx)
			err := c.Config(&routing.RouterConfig{
				RefID: 1,
			}, false)
			Ω(err).Should(BeEquivalentTo(template.ErrNoName))
		})

		It("Should return an error if the ListenStatement isn't set", func() {
			c := template.NewCheck(template.Nginx)
			err := c.Config(&routing.RouterConfig{
				RefID: 1,
				Name:  "name",
			}, false)
			Ω(err).Should(BeEquivalentTo(template.ErrNoListenStatement))
		})

		It("Should return an error if the ServerName isn't set", func() {
			c := template.NewCheck(template.Nginx)
			err := c.Config(&routing.RouterConfig{
				RefID: 1,
				Name:  "name",
				ListenStatement: &routing.ListenStatement{
					IPAddress: abstraction.Inet("127.0.0.1"),
					Port:      1337,
					Keyword:   "ssl",
				},
			}, false)
			Ω(err).Should(BeEquivalentTo(template.ErrEmptyServerName))
		})

		XIt("Should return an error if the AccesLog isn't set", func() {
			c := template.NewCheck(template.Nginx)
			err := c.Config(&routing.RouterConfig{
				RefID: 1,
				Name:  "name",
				ListenStatement: &routing.ListenStatement{
					IPAddress: abstraction.Inet("127.0.0.1"),
					Port:      1337,
					Keyword:   "ssl",
				},
			}, false)
			Ω(err).Should(BeEquivalentTo(template.ErrNoName))
		})

		XIt("Should return an error if the ErrorLog isn't set", func() {
			c := template.NewCheck(template.Nginx)
			err := c.Config(&routing.RouterConfig{
				RefID: 1,
				Name:  "name",
				ListenStatement: &routing.ListenStatement{
					IPAddress: abstraction.Inet("127.0.0.1"),
					Port:      1337,
					Keyword:   "ssl",
				},
			}, false)
			Ω(err).Should(BeEquivalentTo(template.ErrNoName))
		})

		XIt("Should return an error if the RootPath isn't set", func() {
			c := template.NewCheck(template.Nginx)
			err := c.Config(&routing.RouterConfig{
				RefID: 1,
				Name:  "name",
				ListenStatement: &routing.ListenStatement{
					IPAddress: abstraction.Inet("127.0.0.1"),
					Port:      1337,
					Keyword:   "ssl",
				},
			}, false)
			Ω(err).Should(BeEquivalentTo(template.ErrNoName))
		})
	})

	XIt("Should return an error if the SSLSettings aren't properly set", func() {
		c := template.NewCheck(template.Nginx)
		err := c.Config(&routing.RouterConfig{
			RefID: 1,
			Name:  "name",
			ListenStatement: &routing.ListenStatement{
				IPAddress: abstraction.Inet("127.0.0.1"),
				Port:      1337,
				Keyword:   "ssl",
			},
		}, false)
		Ω(err).Should(BeEquivalentTo(template.ErrNoName))
	})

	XIt("Should return an error if the LocationRules aren't properly set", func() {
		c := template.NewCheck(template.Nginx)
		err := c.Config(&routing.RouterConfig{
			RefID: 1,
			Name:  "name",
			ListenStatement: &routing.ListenStatement{
				IPAddress: abstraction.Inet("127.0.0.1"),
				Port:      1337,
				Keyword:   "ssl",
			},
		}, false)
		Ω(err).Should(BeEquivalentTo(template.ErrNoName))
	})
})
