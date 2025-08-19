// Package cmd provides command-line interface functionality for dockerutils.
// Copyright (c) 2025 Docker Utils Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate autocompletion scripts for various shells",
	Long: `Generate autocompletion scripts for dockerutils to enable command and flag completion in your shell.

This command generates shell-specific completion scripts that provide:
- Command name completion (server, completion, etc.)
- Flag and option completion (--port, --host, etc.)
- Argument completion for supported commands
- Context-aware suggestions based on current command

Available shells:
  bash       Generate autocompletion script for bash
  zsh        Generate autocompletion script for zsh  
  fish       Generate autocompletion script for fish
  powershell Generate autocompletion script for PowerShell

Examples:
  # Generate and install bash completion (Linux)
  dockerutils completion bash | sudo tee /etc/bash_completion.d/dockerutils
  
  # Generate and install bash completion (macOS)
  dockerutils completion bash | sudo tee /usr/local/etc/bash_completion.d/dockerutils
  
  # Generate zsh completion script
  dockerutils completion zsh > "${fpath[1]}/_dockerutils"
  
  # Generate fish completion script
  dockerutils completion fish > ~/.config/fish/completions/dockerutils.fish
  
  # Generate PowerShell completion script
  dockerutils completion powershell > dockerutils.ps1

Quick Installation:

  Bash (Linux/macOS):
    dockerutils completion bash | sudo tee /etc/bash_completion.d/dockerutils
    
  Zsh:
    # Add to ~/.zshrc: autoload -U compinit; compinit
    dockerutils completion zsh > "${fpath[1]}/_dockerutils"
    
  Fish:
    dockerutils completion fish > ~/.config/fish/completions/dockerutils.fish
    
  PowerShell:
    dockerutils completion powershell | Out-String | Invoke-Expression

After installation, restart your shell or run 'source ~/.bashrc' (bash) to enable completion.`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			return cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			return cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			return cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			return cmd.Help()
		}
	},
}

func init() {
	// Add the completion command to the root command
	rootCmd.AddCommand(completionCmd)
}
