// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type SetupCommand struct {
	Meta
}

// Help satisfies the cli.Command Help function.
func (a *SetupCommand) Help() string {
	helpText := `
Usage: nomad setup <subcommand> [options] [args]

  This command groups helper subcommands used for setting up Consul and Vault.

  Setup Consul for Nomad:

      $ nomad setup consul -method-name-services="nomad-workloads"

  Please see the individual subcommand help for detailed usage information.
`
	return strings.TrimSpace(helpText)
}

// Synopsis satisfies the cli.Command Synopsis function.
func (a *SetupCommand) Synopsis() string { return "Interact with setup helpers" }

// Name returns the name of this command.
func (a *SetupCommand) Name() string { return "setup" }

// Run satisfies the cli.Command Run function.
func (a *SetupCommand) Run(_ []string) int { return cli.RunResultHelp }
