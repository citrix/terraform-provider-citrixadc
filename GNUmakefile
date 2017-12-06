# Copyright 2016 Citrix Systems, Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

PROVIDER_ONLY_PKGS=$(shell go list ./... | grep -v "/vendor/" | grep -v "tools")

default: build

update:
	dep ensure -update
	dep prune

build:
	go build -o terraform-provider-netscaler .

build-windows:
	GOOS=windows go build -a -installsuffix windows  -o terraform-provider-netscaler.exe

build-linux:
	GOOS=linux go build -a -installsuffix linux  -o terraform-provider-netscaler

test:
	TF_ACC=1 TF_LOG=INFO go test -v $(PROVIDER_ONLY_PKGS)

plan:
	@terraform plan

clean:
	rm terraform-provider-netscaler

clean-windows:
	rm terraform-provider-netscaler.exe

release: clean build
	tar cvzf terraform-provider-netscaler-darwin-amd64.tar.gz terraform-provider-netscaler

release-linux: clean build-linux
	tar cvzf terraform-provider-netscaler-linux-amd64.tar.gz terraform-provider-netscaler

release-windows: clean-windows build-windows
	zip terraform-provider-netscaler-windows-amd64.zip terraform-provider-netscaler.exe
