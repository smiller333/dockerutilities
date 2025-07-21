// Package cmd provides command-line interface functionality for dockerutils.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate autocompletion scripts for various shells",
	Long: `Generate autocompletion scripts for dockerutils for the specified shell.

The completion script can be sourced to enable autocompletion for dockerutils commands,
flags, and arguments in your shell.

Available shells:
  bash       Generate autocompletion script for bash
  zsh        Generate autocompletion script for zsh  
  fish       Generate autocompletion script for fish
  powershell Generate autocompletion script for PowerShell

Examples:
  # Generate bash completion script
  dockerutils completion bash > /etc/bash_completion.d/dockerutils
  
  # Generate zsh completion script
  dockerutils completion zsh > "${fpath[1]}/_dockerutils"
  
  # Generate fish completion script
  dockerutils completion fish > ~/.config/fish/completions/dockerutils.fish
  
  # Generate PowerShell completion script
  dockerutils completion powershell > dockerutils.ps1

Installation Instructions:

  Bash:
    # Linux:
    dockerutils completion bash | sudo tee /etc/bash_completion.d/dockerutils
    # macOS (using Homebrew):
    dockerutils completion bash | sudo tee /usr/local/etc/bash_completion.d/dockerutils
    
  Zsh:
    # Add to your ~/.zshrc:
    autoload -U compinit; compinit
    # Then generate the completion file:
    dockerutils completion zsh > "${fpath[1]}/_dockerutils"
    
  Fish:
    dockerutils completion fish > ~/.config/fish/completions/dockerutils.fish
    
  PowerShell:
    # Add to your PowerShell profile:
    dockerutils completion powershell | Out-String | Invoke-Expression`,
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
