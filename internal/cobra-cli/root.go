package cobra_cli

import (
	"fmt"
	"os"

	"github.com/heisenberg8055/graper/internal/crawler"
	"github.com/spf13/cobra"
)

var recrawler bool
var rootCmd = &cobra.Command{
	Use:        "scraper",
	Short:      "A cli application for checking dead links in html websites",
	Long:       "A cli application for checking dead links in html websites recursively in the same domain and list out all the dead links",
	SuggestFor: []string{"crawler", "scrape"},
	Example:    "scraper <url>, scraper --url <url>, scraper -r <url>",
	Args:       cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		crawler.Crawler(recrawler, args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occured: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&recrawler, "url", "r", false, "Crawler recurs to find all the dead links inside the given url of same domain")

}
