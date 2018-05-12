/*
Copyright 2016 Citrix Systems, Inc

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
	"os"
	"testing"
)

func TestClientCreate(t *testing.T) {
	t.Log("Create client from environment variables and supplied params")
	oldURL := os.Getenv("NS_URL")
	oldLogin := os.Getenv("NS_LOGIN")
	oldPwd := os.Getenv("NS_PASSWORD")
	os.Unsetenv("NS_URL")
	os.Unsetenv("NS_LOGIN")
	os.Unsetenv("NS_PASSWORD")

	_, err := NewNitroClientFromEnv()
	if err == nil {
		t.Error("Expected to fail in creating client")
	}

	os.Setenv("NS_URL", "http://127.0.0.1:32775")
	_, err = NewNitroClientFromEnv()
	if err == nil {
		t.Error("Expected to fail in creating client")
	}
	os.Setenv("NS_LOGIN", "nsroot")
	_, err = NewNitroClientFromEnv()
	if err == nil {
		t.Error("Expected to fail in creating client")
	}
	os.Setenv("NS_PASSWORD", "nsroot")
	_, err = NewNitroClientFromEnv()
	if err != nil {
		t.Error("Didnt expect to fail in creating client")
	}
	os.Setenv("_MPS_API_PROXY_MANAGED_INSTANCE_IP", "10.221.48.207")
	client, err := NewNitroClientFromEnv()
	if err != nil {
		t.Error("Didnt expect to fail in creating client")
	}
	if client.proxiedNs != "10.221.48.207" {
		t.Error("proxiedNS not set despite being set in the environment")
	}
	os.Setenv("NS_SSLVERIFY", "False")
	_, err = NewNitroClientFromEnv()
	if err != nil {
		t.Error("Didnt expect to fail in creating client with SSL verify option")
	}

	os.Setenv("NS_URL", oldURL)
	os.Setenv("NS_LOGIN", oldLogin)
	os.Setenv("NS_PASSWORD", oldPwd)
}

func TestClientCreateFromParams(t *testing.T) {
	t.Log("Create client from supplied params")

	params := NitroParams{
		Url:      os.Getenv("NS_URL"),
		Username: os.Getenv("NS_LOGIN"),
		Password: os.Getenv("NS_PASSWORD"),
	}
	client, err := NewNitroClientFromParams(params)
	if client == nil {
		t.Error("Expected to succeed in creating client")
	}

	params.SslVerify = false
	client, err = NewNitroClientFromParams(params)
	if client == nil {
		t.Error("Expected to succeed in creating client")
	}

	params.Url = "localhost:32770"
	client, err = NewNitroClientFromParams(params)
	if err == nil {
		t.Error("Expected to fail in creating client due to invalid URL")
	}

}
