package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/lab5e/go-userapi"
	"github.com/lab5e/go-userapi/apitools"
)

// Simple example that queries the profile. You can get an API token from the
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

	fmt.Println("Requesting profile...")

	start := time.Now()

	// Finally -- the request. There are three return values from the call:
	// * the profile
	// * response object
	// * error
	// The response object can usually be ignored unless you want to inspect the
	// internal state in the client.
	profile, response, err := userapiClient.ProfileApi.UserGetUserProfile(ctx).Execute()
	if err != nil {
		fmt.Println("Got error requesting profile: ", err)
		fmt.Println("Response is ", response)
		return
	}

	fmt.Println("Done, request took ", time.Since(start))

	// ...and print the result to screen. The tabwriter package prints out a
	// table on screen.
	fmt.Println("Profile:")
	fmt.Println("========")

	table := tabwriter.NewWriter(os.Stdout, 2, 4, 3, ' ', 0)

	table.Write([]byte(fmt.Sprintf("User ID:\t%s\n", *profile.UserId)))
	table.Write([]byte(fmt.Sprintf("Name:\t%s\n", *profile.Name)))
	table.Write([]byte(fmt.Sprintf("Profile URL:\t%s\n", *profile.ProfileUrl)))
	table.Write([]byte(fmt.Sprintf("Avatar URL:\t%s\n", *profile.AvatarUrl)))

	table.Flush()

}
