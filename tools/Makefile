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

GENERATED_PKGS=$(shell go list ./... )

JSON_DIR = ../../../chiradeep/go-nitro/jsonconfig
SCHEMA_EXEC = generate_schema

default: build 

build:
	go build -v -o $(SCHEMA_EXEC)  .

generate-test-build:
	go build ${GENERATED_PKGS}

clean:
	rm generate_schema

fmt:
	go fmt ${GENERATED_PKGS}

schema_only:
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/cs/cspolicy.json -b $(JSON_DIR)/cs/csvserver_cspolicy_binding.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/cs/csvserver.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/cs/csaction.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/lb/lbmonitor.json -d monitorname
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/lb/lbvserver.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/ns/nsacl.json -d aclname
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/ns/nsacls.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/gslb/gslbservice.json -d servicename
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/gslb/gslbsite.json -d sitename
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/gslb/gslbvserver.json -b $(JSON_DIR)/gslb/gslbvserver_domain_binding.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/network/inat.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/network/ipset.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/basic/service.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/basic/servicegroup.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/ssl/sslcertkey.json -d certkey
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/ssl/sslprofile.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/basic/server.json
	@# Responder stuff
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/responder/responderpolicy.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/responder/responderaction.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/responder/responderglobal_responderpolicy_binding.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/responder/responderpolicylabel.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/cs/csvserver_responderpolicy_binding.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/lb/lbvserver_responderpolicy_binding.json
	@# Rewrite stuff
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/rewrite/rewriteaction.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/rewrite/rewritepolicy.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/rewrite/rewriteglobal_rewritepolicy_binding.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/rewrite/rewritepolicylabel.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/cs/csvserver_rewritepolicy_binding.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/lb/lbvserver_rewritepolicy_binding.json
	@# Cluster
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/cluster/clusterinstance.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/cluster/clusternode.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/cluster/cluster.json
	./$(SCHEMA_EXEC) -i $(JSON_DIR)/cluster/clusterfiles.json
	


schema: schema_only fmt
