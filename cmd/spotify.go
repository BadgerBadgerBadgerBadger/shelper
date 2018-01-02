package cmd

// import "fmt"
import "github.com/spf13/cobra"

import "scionofbytes.me/shelper/spotify"

var cmd = &cobra.Command{
	Use:   "spotify",
	Short: "You da man, and Spotify yo bitch!",
}

var searchSelf bool

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search through the spotify music catalogue or your own saved songs.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate your spotify account to get an app token.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		spotify.Authenticate
	},
}

func init() {
	searchCmd.Flags().BoolVarP(&searchSelf, "searchSelf", "p", false, "If set, searches through your playlists.")
	cmd.AddCommand(searchCmd)
	rootCmd.AddCommand(cmd)
}
