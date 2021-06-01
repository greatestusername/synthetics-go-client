package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/strfmt"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/greatestusername/synthetics-go-client/client"

	"github.com/greatestusername/synthetics-go-client/client/http_checks"
)

func main() {
	r := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	r.DefaultAuthentication = httptransport.APIKeyAuth("api-key", "header", os.Getenv("API_ACCESS_TOKEN"))
	r.SetDebug(true)

	client := apiclient.New(r, strfmt.Default)

	resp, err := client.HTTPChecks.GetHTTPCheck(http_checks.NewGetHTTPCheckParams().WithDefaults().WithID(175389))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)

}
