package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/koltyakov/gosip"
)

// Webs ...
type Webs struct {
	client    *gosip.SPClient
	config    *RequestConfig
	endpoint  string
	modifiers map[string]string
}

// Conf ...
func (webs *Webs) Conf(config *RequestConfig) *Webs {
	webs.config = config
	return webs
}

// Select ...
func (webs *Webs) Select(oDataSelect string) *Webs {
	if webs.modifiers == nil {
		webs.modifiers = make(map[string]string)
	}
	webs.modifiers["$select"] = oDataSelect
	return webs
}

// Expand ...
func (webs *Webs) Expand(oDataExpand string) *Webs {
	if webs.modifiers == nil {
		webs.modifiers = make(map[string]string)
	}
	webs.modifiers["$expand"] = oDataExpand
	return webs
}

// Filter ...
func (webs *Webs) Filter(oDataFilter string) *Webs {
	if webs.modifiers == nil {
		webs.modifiers = make(map[string]string)
	}
	webs.modifiers["$filter"] = oDataFilter
	return webs
}

// Top ...
func (webs *Webs) Top(oDataTop int) *Webs {
	if webs.modifiers == nil {
		webs.modifiers = make(map[string]string)
	}
	webs.modifiers["$top"] = string(oDataTop)
	return webs
}

// OrderBy ...
func (webs *Webs) OrderBy(oDataOrderBy string, ascending bool) *Webs {
	direction := "asc"
	if !ascending {
		direction = "desc"
	}
	if webs.modifiers == nil {
		webs.modifiers = make(map[string]string)
	}
	if webs.modifiers["$orderby"] != "" {
		webs.modifiers["$orderby"] += ","
	}
	webs.modifiers["$orderby"] += fmt.Sprintf("%s %s", oDataOrderBy, direction)
	return webs
}

// Get ...
func (webs *Webs) Get() ([]byte, error) {
	apiURL, _ := url.Parse(webs.endpoint)
	query := url.Values{}
	for k, v := range webs.modifiers {
		query.Add(k, trimMultiline(v))
	}
	apiURL.RawQuery = query.Encode()
	sp := NewHTTPClient(webs.client)
	headers := map[string]string{}
	if webs.config != nil {
		headers = webs.config.Headers
	}
	return sp.Get(apiURL.String(), headers)
}

// Add ...
func (webs *Webs) Add(title string, url string, metadata map[string]interface{}) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/add", webs.endpoint)

	if metadata == nil {
		metadata = make(map[string]interface{})
	}

	metadata["__metadata"] = map[string]string{
		"type": "SP.WebCreationInformation",
	}

	metadata["Title"] = title
	metadata["Url"] = url

	// metadata["Description"]

	// Default values
	if metadata["Language"] == nil {
		metadata["Language"] = 1033
	}
	if metadata["UseSamePermissionsAsParentSite"] == nil {
		metadata["UseSamePermissionsAsParentSite"] = true
	}
	if metadata["WebTemplate"] == nil {
		metadata["WebTemplate"] = "STS"
	}

	parameters, _ := json.Marshal(metadata)

	body := trimMultiline(`{
		"parameters": ` + fmt.Sprintf("%s", parameters) + `
	}`)

	sp := NewHTTPClient(webs.client)
	headers := getConfHeaders(webs.config)

	headers["Accept"] = "application/json;odata=verbose"
	headers["Content-Type"] = "application/json;odata=verbose;charset=utf-8"

	return sp.Post(endpoint, []byte(body), headers)
}