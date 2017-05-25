// Package cli is
package cli

import (
	"github.com/go-kit/kit/log"

	"github.com/abiosoft/ishell"
	"google.golang.org/grpc"
)

// InitShell adds all available kontaineroo commands to an ishell instance
func InitShell(sh *ishell.Shell, conn *grpc.ClientConn, logger log.Logger) {
	sh.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "authenticate yourself",
		Func: func(c *ishell.Context) {
			c.Println("not yet implemented")
		},
	})
}
