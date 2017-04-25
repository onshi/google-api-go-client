// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	webmasters "google.golang.org/api/webmasters/v3"
)

func init() {
	scopes := []string{
		webmasters.WebmastersScope,
		webmasters.WebmastersReadonlyScope,
	}
	registerDemo("searchconsole", strings.Join(scopes, " "), searchconsoleMain)
}

// searchconsoleMain is an example that demonstrates calling the Search Console API.
//
// Example usage:
//   go build -o go-api-demo *.go
//   go-api-demo -clientid="my-clientid" -secret="my-secret" searchconsole
func searchconsoleMain(client *http.Client, argv []string) {
	if len(argv) != 0 {
		fmt.Fprintln(os.Stderr, "Usage: searchconsole")
		return
	}

	svc, err := webmasters.New(client)
	if err != nil {
		log.Fatalf("Unable to create SearchConsole service: %v", err)
	}

	showSearchConsoleServiceInfo(svc)
}

func showSearchConsoleServiceInfo(svc *webmasters.Service) {
	showSitesInfo(svc)
	showSearchAnalytics(svc, "https://unamo.com/")
}

func showSitesInfo(svc *webmasters.Service) {
	siteListResponse, err := svc.Sites.List().Do()

	if err != nil {
		log.Fatalf("Unable to get list of sites")
	}

	fmt.Fprintln(os.Stderr, "Site List:")
	for _, site := range siteListResponse.SiteEntry {
		fmt.Printf("Permission level: %v\n", site.PermissionLevel)
		fmt.Printf("Site Url: %v\n", site.SiteUrl)
		fmt.Printf("Force Send Fields: %v\n", site.ForceSendFields)
		fmt.Printf("Null Fields: %v\n", site.NullFields)
	}
}

func showSearchAnalytics(svc *webmasters.Service, siteUrl string) {
	queryRequest := &webmasters.SearchAnalyticsQueryRequest{
		EndDate:   "2017-04-25",
		StartDate: "2017-04-23",
	}

	searchAnalytics, err := svc.Searchanalytics.Query(siteUrl, queryRequest).Do()

	if err != nil {
		log.Fatalf("Unable to get list of searchAnalytics")
	}

	fmt.Fprintln(os.Stderr, "searchAnalytics:")
	for _, apiDataRow := range searchAnalytics.Rows {
		fmt.Printf("Clicks: %v\n", apiDataRow.Clicks)
		fmt.Printf("Ctr: %v\n", apiDataRow.Ctr)
		fmt.Printf("Impressions: %v\n", apiDataRow.Impressions)
		fmt.Printf("Keys: %v\n", apiDataRow.Keys)
		fmt.Printf("Positiones: %v\n", apiDataRow.Position)
		fmt.Printf("Null Fields: %v\n", apiDataRow.ForceSendFields)
		fmt.Printf("Null Fields: %v\n", apiDataRow.NullFields)
	}

	fmt.Printf("ResponseAggregationType: %v\n", searchAnalytics.ResponseAggregationType)
}
