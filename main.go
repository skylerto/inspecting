package main

import (
  "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
  "os"
)

type Results struct {
	Version string `json:"version"`
	Controls []struct {
		Status string `json:"status"`
		CodeDesc string `json:"code_desc"`
		RunTime float64 `json:"run_time"`
		StartTime string `json:"start_time"`
		Message string `json:"message,omitempty"`
		SkipMessage string `json:"skip_message,omitempty"`
		Resource string `json:"resource,omitempty"`
	} `json:"controls"`
	OtherChecks []interface{} `json:"other_checks"`
	Profiles []struct {
		Name string `json:"name"`
		Title string `json:"title"`
		Maintainer string `json:"maintainer"`
		Copyright string `json:"copyright"`
		CopyrightEmail string `json:"copyright_email"`
		License string `json:"license"`
		Summary string `json:"summary"`
		Version string `json:"version"`
		Supports []struct {
			OsFamily string `json:"os-family"`
		} `json:"supports"`
		Controls []struct {
			Title string `json:"title"`
			Desc string `json:"desc"`
			Impact float64 `json:"impact"`
			Refs []interface{} `json:"refs"`
			Tags struct {
			} `json:"tags"`
			Code string `json:"code"`
			SourceLocation struct {
				Ref string `json:"ref"`
				Line int `json:"line"`
			} `json:"source_location"`
			ID string `json:"id"`
			Results []struct {
				Status string `json:"status"`
				CodeDesc string `json:"code_desc"`
				RunTime float64 `json:"run_time"`
				StartTime string `json:"start_time"`
			} `json:"results"`
		} `json:"controls"`
		Groups []struct {
			Title interface{} `json:"title"`
			Controls []string `json:"controls"`
			ID string `json:"id"`
		} `json:"groups"`
		Attributes []struct {
			Name string `json:"name"`
			Options struct {
				Default string `json:"default"`
				Description string `json:"description"`
			} `json:"options"`
		} `json:"attributes"`
		Sha256 string `json:"sha256"`
	} `json:"profiles"`
	Platform struct {
		Name string `json:"name"`
		Release string `json:"release"`
	} `json:"platform"`
	Statistics struct {
		Duration float64 `json:"duration"`
	} `json:"statistics"`
}

func main() {
  file := os.Args[1]

	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s", b)

  res := Results{}
  json.Unmarshal([]byte(b), &res)
  fmt.Println(res)

}
