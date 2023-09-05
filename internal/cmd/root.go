package cmd

import (
	"fmt"
	"github.com/j3yzz/sheriff/internal/cmd/migration"
	"github.com/j3yzz/sheriff/internal/cmd/server"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const asciiArt = `
███████ ██   ██ ███████ ██████  ██ ███████ ███████ 
██      ██   ██ ██      ██   ██ ██ ██      ██      
███████ ███████ █████   ██████  ██ █████   █████   
     ██ ██   ██ ██      ██   ██ ██ ██      ██      
███████ ██   ██ ███████ ██   ██ ██ ██      ██
`

func Execute() {
	fmt.Println(asciiArt)
	cmd := &cobra.Command{
		Use:   "sheriff",
		Short: "A simple and secure authentication and user management system for your web applications.",
	}

	migration.Register(cmd)
	server.Register(cmd)

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
