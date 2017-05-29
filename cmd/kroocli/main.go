package main

import (
	"fmt"
	"os"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/cli"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr = ":8082"
	)

	shell := ishell.New()
	shell.SetHomeHistoryPath(".kroocli_history")
	shell.Println("Kontainer.ooo interactive shell")

	logger := log.NewLogfmtLogger(os.Stdout)

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	cli.InitShell(shell, conn, logger)

	shell.Start()
}
