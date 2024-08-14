/*
Copyright © 2024  github.com/Meiste-r
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/robtec/newsapi/api"
	"github.com/spf13/cobra"
)

var Country string
var Category string
var Query string
var HowOften int

var checknewsCmd = &cobra.Command{
	Use:   "checknews",
	Short: "Checks for news.",
	Long: `Checks for news in a specific country of a specific topic.
	use news checknews -c (country(short)) -t (business, entertainment, general, health, science, sports ortechnology) -n (number inbetween 0 and 127 to set how many results shall been shown) -q (query).`,
	Run: func(cmd *cobra.Command, args []string) {
		//http client, api key, url
		httpClient := http.Client{}
		key := "API-KEY"
		url := "https://newsapi.org"

		//newsapi client
		client, err := api.New(&httpClient, key, url)
		if err != nil {
			fmt.Println(err)
			return
		}

		//options for search
		opts := api.Options{Country: Country, Category: Category, Q: Query, SortBy: "popularity"}

		//searching
		thl, err := client.TopHeadlines(opts)
		if err != nil {
			fmt.Println(err)
			return
		}

		if HowOften > len(thl.News.Articles) {
			HowOften = len(thl.News.Articles)

		}
		fmt.Println(HowOften)
		fmt.Println(len(thl.News.Articles))

		//printing results
		for i := 0; i < HowOften-1; i++ {
			fmt.Println("Title: " + thl.Articles[i].Title)
			fmt.Println("=================================")
			fmt.Println("Published at: " + thl.Articles[i].PublishedAt)
			fmt.Println("Source: " + thl.Articles[i].Source.Name)
			fmt.Println("=================================")
			fmt.Println("URL: " + thl.Articles[i].URL)
			fmt.Println("")
			fmt.Println("")

		}
	},
}

func init() {
	//adding flags
	checknewsCmd.PersistentFlags().StringVarP(&Country, "country", "c", "", "Passes a country to search for")
	checknewsCmd.PersistentFlags().StringVarP(&Category, "category", "t", "", "Passes a category to search for")
	checknewsCmd.PersistentFlags().StringVarP(&Query, "query", "q", "", "Passes a query to search for")
	checknewsCmd.PersistentFlags().IntVarP(&HowOften, "number", "n", 10, "Passes in how many results shall been shown")

	rootCmd.AddCommand(checknewsCmd)

}
