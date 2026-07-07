package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/rest-sh/restish/cli"
	"github.com/rest-sh/restish/openapi"
	"github.com/spf13/cobra"
)

func BenchmarkFormats(b *testing.B) {
	inputs := []struct {
		Name string
		URL  string
	}{
		{
			Name: "small",
			URL:  "https://github.com/OAI/OpenAPI-Specification/raw/d1cc440056f1c7bb913bcd643b15c14ee1c409f4/examples/v3.0/uspto.json",
		},
		{
			Name: "large",
			URL:  "https://github.com/github/rest-api-description/blob/83cdec7384b62ef6f54bad60270544d6fc6f22cd/descriptions/api.github.com/api.github.com.json?raw=true",
		},
	}

	cli.Init("benchmark", "1.0.0")
	cli.Defaults()
	cli.AddLoader(openapi.New())

	for _, t := range inputs {
		resp, err := http.Get(t.URL)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode >= 300 {
			panic("non-success status from server, check URLs are still working")
		}

		dummy := &cobra.Command{}
		doc, err := cli.Load(t.URL, dummy)
		if err != nil {
			panic(err)
		}

		dataJSON, err := json.Marshal(doc)
		if err != nil {
			panic(err)
		}

		fmt.Printf("json: %d\n", len(dataJSON))

		b.Run(t.Name+"-json-marshal", func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				json.Marshal(doc)
			}
		})

		b.Run(t.Name+"-json-unmarshal", func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				var tmp cli.API
				json.Unmarshal(dataJSON, &tmp)
			}
		})
	}
}
