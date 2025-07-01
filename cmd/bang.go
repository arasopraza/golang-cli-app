package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var bangCmd = &cobra.Command{
	Use:   "meme",
	Short: "Responds with a meme message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bang! 💥 Making HTTP request...")

		resp, err := http.Get("https://api.imgflip.com/get_memes")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			return
		}

		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Printf("Response: %s\n", string(body))
	},
}

func init() {
	rootCmd.AddCommand(bangCmd)
}
