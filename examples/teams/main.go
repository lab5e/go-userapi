package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/lab5e/go-userapi"
	"github.com/lab5e/go-userapi/apitools"
)

// Simple example that queries the teams. You can get an API token from the
// "API Tokens" page at https://span.lab5e.com/
func main() {

	// Since nobody wants their credentials in checked-in code we'll use a
	// command line parameter for the token.
	token := ""
	flag.StringVar(&token, "token", "", "API Token for the user service")
	flag.Parse()

	if token == "" {
		fmt.Println("Missing the API token")
		flag.PrintDefaults()
		return
	}

	// Create the configuration object for the client
	config := userapi.NewConfiguration()

	// Set this to true to see the requests and responses
	config.Debug = false

	// Create the client
	userapiClient := userapi.NewAPIClient(config)

	// We'll need a context with the credentials included. This utility function
	// creaets a context for us. You probably want a new context for each request
	// you are doing if you have a long-running service.
	ctx, done := apitools.NewAuthenticatedContext(token, 10*time.Second)
	defer done()

	fmt.Println("Query teams...")

	start := time.Now()

	// Finally -- the request. There are three return values from the call:
	// * the list of teams
	// * response object
	// * error
	// The response object can usually be ignored unless you want to inspect the
	// internal state in the client.
	teamList, _, err := userapiClient.TeamsApi.ListTeams(ctx).Execute()
	if err != nil {
		fmt.Println("Got error requesting teams: ", err)
		return
	}
	fmt.Println("Done, request took ", time.Since(start))

	// ...and print the result to screen. The tabwriter package prints out a
	// table on screen.

	table := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)
	table.Write([]byte("Team ID\tTags\tMembers\n"))

	for _, team := range *teamList.Teams {
		table.Write([]byte(fmt.Sprintf("%s\t%s\t\n", *team.TeamId, tagList(*team.Tags))))

		for _, member := range *team.Members {
			table.Write([]byte(fmt.Sprintf("\t\t%s (%s)\n", *member.User.Email, *member.Role)))
		}
	}
	table.Flush()
}

func tagList(tags map[string]string) string {
	var t []string
	for k, v := range tags {
		t = append(t, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(t, ", ")
}
