// rfoutlet provides outlet control via cli and web interface for
// Raspberry PI 2/3.
//
// The transmitter and receiver logic has been ported from the great
// https://github.com/sui77/rc-switch C++ project to golang.
//
// rfoutlet comes with ready to use commands for transmitting and receiving
// remote control codes as well as a command for serving a web interface (see
// cmd/ directory). The pkg/ directory exposes the gpio package which contains
// the receiver and transmitter code.
package main

import (
	"fmt"
	"os"

	"github.com/martinohmann/rfoutlet/cmd"
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "rfoutlet",
		Short:         "A tool for interacting with remote controlled outlets",
		Long:          "rfoutlet is a tool for interacting with remote controlled outlets. It provides functionality to sniff and transmit the codes controlling the outlets.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	return cmd
}

func main() {
	rootCmd := newRootCommand()

	rootCmd.AddCommand(cmd.NewServeCommand())
	rootCmd.AddCommand(cmd.NewSniffCommand())
	rootCmd.AddCommand(cmd.NewTransmitCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}