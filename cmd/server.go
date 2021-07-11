// Copyright (c) Inlets Author(s) 2021. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/inlets/inlets/pkg/server"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// serverCmd represents the server sub command.
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: `Start the tunnel server.`,
	Long: `Start the tunnel server on a machine with a publicly-accessible IPv4 IP address such as a VPS.

Note: You can pass the --token argument followed by a token value to both the server and client to prevent unauthorized connections to the tunnel.`,
	RunE: runServer,
	Example: `  # Bind the data and control plane to 80 and 8080
  inlets server --port 80 \
    --control-port 8080
  
  # Bind the control-plane to 127.0.0.1:
  inlets server --port 80 \
    --control-port 8001 \
    --control-addr 127.0.0.1`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {

	serverCmd.Flags().IntP("port", "p", 8000, "port for server and for tunnel")
	serverCmd.Flags().IntP("control-port", "c", 8001, "control port for tunnel")

	serverCmd.Flags().String("data-addr", "0.0.0.0", "address the server should serve tunneled services on")
	serverCmd.Flags().String("control-addr", "0.0.0.0", "address tunnel clients should connect to")

	serverCmd.Flags().StringP("token", "t", "", "token for authentication")
	serverCmd.Flags().StringP("token-from", "f", "", "read the authentication token from a file")

	serverCmd.Flags().Bool("print-token", false, "prints the token in server mode")

	serverCmd.Flags().Bool("disable-transport-wrapping", false, "disable wrapping the transport that removes CORS headers for example")

	inletsCmd.AddCommand(serverCmd)
}

// runServer does the actual work of reading the arguments passed to the server sub command.
func runServer(cmd *cobra.Command, _ []string) error {

	fmt.Printf("%s", WelcomeMessage)
	log.Printf("Starting server - version %s", getVersion())

	tokenFile, err := cmd.Flags().GetString("token-from")
	if err != nil {
		return errors.Wrap(err, "failed to get 'token-from' value.")
	}

	var token string
	if len(tokenFile) > 0 {
		fileData, err := ioutil.ReadFile(tokenFile)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("unable to load file: %s", tokenFile))
		}

		// new-lines will be stripped, this is not configurable and is to
		// make the code foolproof for beginners
		token = strings.TrimRight(string(fileData), "\n")
	} else {
		tokenVal, err := cmd.Flags().GetString("token")
		if err != nil {
			return errors.Wrap(err, "failed to get 'token' value.")
		}
		token = tokenVal
	}

	if tokenEnv, ok := os.LookupEnv("TOKEN"); ok && len(tokenEnv) > 0 {
		fmt.Printf("Token read from environment variable.\n")
		token = tokenEnv
	}

	printToken, err := cmd.Flags().GetBool("print-token")
	if err != nil {
		return errors.Wrap(err, "failed to get 'print-token' value.")
	}

	if len(token) > 0 && printToken {
		log.Printf("Server token: %q", token)
	}

	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return errors.Wrap(err, "failed to get the 'port' value.")
	}

	controlPort, err := cmd.Flags().GetInt("control-port")
	if err != nil {
		return errors.Wrap(err, "failed to get the 'control-port' value.")
	}

	disableWrapTransport, err := cmd.Flags().GetBool("disable-transport-wrapping")
	if err != nil {
		return errors.Wrap(err, "failed to get the 'disable-transport-wrapping' value.")
	}

	dataAddr, err := cmd.Flags().GetString("data-addr")
	if err != nil {
		return errors.Wrap(err, "failed to get the 'data-addr' value.")
	}

	controlAddr, err := cmd.Flags().GetString("control-addr")
	if err != nil {
		return errors.Wrap(err, "failed to get the 'control-addr' value.")
	}

	if port == controlPort {
		return fmt.Errorf("the --port and --control-port cannot be the same value")
	}

	inletsServer := server.Server{
		Port:        port,
		ControlPort: controlPort,
		DataAddr:    dataAddr,
		ControlAddr: controlAddr,
		Token:       token,

		DisableWrapTransport: disableWrapTransport,
	}

	inletsServer.Serve()
	return nil
}
