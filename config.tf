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

provider "netscaler" {
  username     = "nsroot"
  password    = "nsroot"
  endpoint = "http://10.71.136.250/"
}

resource "netscaler_lb" "my-lb-vserver" {
  name = "sample_lb"
  vip = "10.71.136.150"
  port = 443
  service_type = "SSL"
  lb_method = "ROUNDROBIN"
  persistence_type = "COOKIEINSERT"
}

resource "netscaler_lb" "my-lb-vserver2" {
  name = "sample_lb2"
  vip = "10.71.136.151"
  port = 443
}

resource "netscaler_svc" "backend_1" {
  lb = "${netscaler_lb.my-lb-vserver2.name}"
  ip = "10.33.44.55"
  port = 80
}

resource "netscaler_svc" "backend_2" {
  lb = "${netscaler_lb.my-lb-vserver2.name}"
  ip = "10.33.44.54"
  port = 80
}
