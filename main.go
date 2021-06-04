package main

import (
	"log"
	"os"

	"github.com/go-openapi/strfmt"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/greatestusername/synthetics-go-client/client"
	"github.com/greatestusername/synthetics-go-client/client/http_checks"
	pp "github.com/greatestusername/synthetics-go-client/models"
)

func main() {
	r := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	r.DefaultAuthentication = httptransport.APIKeyAuth("api-key", "header", os.Getenv("API_ACCESS_TOKEN"))
	r.SetDebug(false)

	client := http_checks.New(r, strfmt.Default)

	//Get Check with id example
	resp, err := client.CreateHTTPCheck(http_checks.NewCreateHTTPCheckParams().WithDefaults().WithCheckDetail(http_checks.CreateHTTPCheckBody{
		CreateHTTPCheckParamsBodyAllOf0: http_checks.CreateHTTPCheckParamsBodyAllOf0{},
		Connection:                      &http_checks.CreateHTTPCheckParamsBodyCreateHTTPCheckParamsBodyAO1Connection{},
		HTTPMethod:                      new(string),
		HTTPRequestBody:                 "",
		SuccessCriteria:                 []*http_checks.CreateHTTPCheckParamsBodySuccessCriteriaItems0{},
	}))
	//GetHTTPCheck(http_checks.NewGetHTTPCheckParams().WithDefaults().WithID(175389))
	if err != nil {
		log.Fatal(err)
	}
	pp.PrettyPrint(resp)

	//Get Checks example
	resp2, err2 := client.Checks.GetChecks(nil)
	if err2 != nil {
		log.Fatal(err2)
	}
	pp.PrettyPrint(resp2)

}
