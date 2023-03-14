/*
Copyright © 2023 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(jocker completion bash)

# To load completions for each session, execute once:
Linux:
  $ jocker completion bash > /etc/bash_completion.d/jocker
MacOS:
  $ jocker completion bash > /usr/local/etc/bash_completion.d/jocker

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ jocker completion zsh > "${fpath[1]}/_jocker"

# You will need to start a new shell for this setup to take effect.

Fish:

$ jocker completion fish | source

# To load completions for each session, execute once:
$ jocker completion fish > ~/.config/fish/completions/jocker.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			_ = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			_ = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			_ = cmd.Root().GenFishCompletion(os.Stdout, true)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
