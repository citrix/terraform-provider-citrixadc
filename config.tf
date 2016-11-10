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

resource "netscaler_lbvserver" "my-lb-vserver" {
  name = "sample_lb"
  ipv46 = "10.71.136.150"
  port = 443
  servicetype = "SSL"
  lbmethod = "ROUNDROBIN"
  persistencetype = "COOKIEINSERT"
}

resource "netscaler_lbvserver" "my-lb-vserver2" {
  name = "sample_lb2"
  ipv46 = "10.71.136.151"
  servicetype = "SSL"
  port = 443
}

resource "netscaler_service" "backend_1" {
  lbvserver = "${netscaler_lbvserver.my-lb-vserver2.name}"
  ip = "10.123.43.55"
  servicetype = "HTTP"
  port = 80
}

resource "netscaler_service" "backend_2" {
  lbvserver = "${netscaler_lbvserver.my-lb-vserver2.name}"
  ip = "10.33.44.54"
  servicetype = "HTTP"
  port = 80
  clttimeout = 360
}

resource "netscaler_csvserver" "foo-cs" {
  name = "sample_cs"
  ipv46 = "10.71.138.151"
  servicetype = "SSL"
  port = 443

}
