/*
Copyright 2016 Citrix Systems, Inc, All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package netscaler

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//NitroParams encapsulates options to create a NitroClient
type NitroParams struct {
	Url       string
	Username  string
	Password  string
	ProxiedNs string
	SslVerify bool
}

//NitroClient has methods to configure the NetScaler
//It abstracts the REST operations of the NITRO API
type NitroClient struct {
	url       string
	statsURL  string
	username  string
	password  string
	proxiedNs string
	client    *http.Client
}

//NewNitroClient returns a usable NitroClient. Does not check validity of supplied parameters
//This is for backwards compatibility.
//Please use NewNitroClientFromParams
func NewNitroClient(url string, username string, password string) *NitroClient {
	c := new(NitroClient)
	c.url = strings.Trim(url, " /") + "/nitro/v1/config/"
	c.statsURL = strings.Trim(url, " /") + "/nitro/v1/stat/"
	c.username = username
	c.password = password
	c.client = &http.Client{}
	return c
}

//NewNitroClientFromParams returns a usable NitroClient. Does not check validity of supplied parameters
func NewNitroClientFromParams(params NitroParams) (*NitroClient, error) {
	u, err := url.Parse(params.Url)
	if err != nil {
		return nil, fmt.Errorf("Supplied URL %s is not a URL", params.Url)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, fmt.Errorf("Supplied URL %s does not have a HTTP/HTTPS scheme", params.Url)
	}
	c := new(NitroClient)
	c.url = strings.Trim(params.Url, " /") + "/nitro/v1/config/"
	c.statsURL = strings.Trim(params.Url, " /") + "/nitro/v1/stat/"
	c.username = params.Username
	c.password = params.Password
	c.proxiedNs = params.ProxiedNs
	if params.SslVerify {
		c.client = &http.Client{}
	} else {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.client = &http.Client{Transport: tr}
	}
	return c, nil
}

//NewNitroClientFromEnv returns a usable NitroClient. Parameters url, username and password can be passed in
//as the first three positional parameters. Otherwise, it tries to read these values from
//environment variable NS_URL, NS_LOGIN and NS_PASSWORD
func NewNitroClientFromEnv() (*NitroClient, error) {
	url := os.Getenv("NS_URL")
	username := os.Getenv("NS_LOGIN")
	password := os.Getenv("NS_PASSWORD")
	if url == "" || username == "" || password == "" {
		return nil, fmt.Errorf("Could not completely determine login parameters from env: NS_URL, NS_LOGIN, NS_PASSWORD")
	}
	proxiedNs := os.Getenv("_MPS_API_PROXY_MANAGED_INSTANCE_IP")
	sslverifyStr := os.Getenv("NS_SSLVERIFY")
	sslVerify := true
	if sslverifyStr != "" {
		var err error
		sslVerify, err = strconv.ParseBool(sslverifyStr)
		if err != nil {
			return nil, fmt.Errorf("Could not parse env variable NS_SSLVERIFY: valid values are true and false")
		}
	}
	nitroParams := NitroParams{
		Url:       url,
		Username:  username,
		Password:  password,
		ProxiedNs: proxiedNs,
		SslVerify: sslVerify,
	}
	return NewNitroClientFromParams(nitroParams)
}
