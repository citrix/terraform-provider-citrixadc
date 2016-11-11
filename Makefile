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
	go get ./...
	godep update ./...

build:
	godep go build -o terraform-provider-netscaler .

test:
	TF_ACC=1 TF_LOG=INFO godep go test -v $(PROVIDER_ONLY_PKGS)

plan:
	@terraform plan

clean:
	rm terraform-provider-netscaler
