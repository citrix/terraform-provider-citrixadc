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

resource "netscaler_lbvserver" "sample_lb" {
  name = "sample_lb"
  ipv46 = "10.71.136.150"
  port = 443
  servicetype = "SSL"
  lbmethod = "ROUNDROBIN"
  persistencetype = "COOKIEINSERT"
  sslcertkey = "${netscaler_sslcertkey.foo-ssl-cert-2.certkey}"
}

resource "netscaler_lbvserver" "sample_lb2" {
  name = "sample_lb2"
  ipv46 = "10.71.136.151"
  servicetype = "SSL"
  port = 443
}

resource "netscaler_service" "backend_1" {
  lbvserver = "${netscaler_lbvserver.sample_lb2.name}"
  ip = "10.123.43.55"
  servicetype = "HTTP"
  port = 80
  lbmonitor = "${netscaler_lbmonitor.foo-monitor-2.monitorname}"
}

resource "netscaler_service" "backend_2" {
  lbvserver = "${netscaler_lbvserver.sample_lb2.name}"
  ip = "10.33.44.54"
  servicetype = "HTTP"
  port = 80
  clttimeout = 360
  lbmonitor = "${netscaler_lbmonitor.foo-monitor.monitorname}"
}

resource "netscaler_csvserver" "foo-cs" {
  name = "sample_cs"
  ipv46 = "10.71.139.151"
  servicetype = "SSL"
  port = 443
}

resource "netscaler_cspolicy" "foo-cspolicy" {
  policyname = "sample_cspolicy"
  rule = "CLIENT.IP.SRC.SUBNET(24).EQ(10.227.84.0)"
  csvserver = "${netscaler_csvserver.foo-cs.name}"
  targetlbvserver = "${netscaler_lbvserver.sample_lb2.name}"
  priority = 21
}

resource "netscaler_cspolicy" "foo-cspolicy2" {
  policyname = "sample_cspolicy2"
  rule = "CLIENT.IP.SRC.SUBNET(24).EQ(10.127.88.0)"
  csvserver = "${netscaler_csvserver.foo-cs.name}"
  targetlbvserver = "${netscaler_lbvserver.sample_lb2.name}"
  priority = 11
}

resource "netscaler_sslcertkey" "foo-ssl-cert" {
  certkey = "sample_ssl_cert"
  cert = "/var/certs/server.crt"
  key = "/var/certs/server.key"
  expirymonitor = "ENABLED"
  notificationperiod = 83
}

resource "netscaler_sslcertkey" "foo-ssl-cert-2" {
  certkey = "sample_ssl_cert2"
  cert = "/var/certs/server2.crt"
  key = "/var/certs/server2.key"
  expirymonitor = "ENABLED"
  notificationperiod = 33
}

resource "netscaler_lbmonitor" "foo-monitor" {
  monitorname = "sample_lb_monitor"
  type = "HTTP"
  interval = 350
  resptimeout = 250
}

resource "netscaler_lbmonitor" "foo-monitor-2" {
  monitorname = "sample_lb_monitor2"
  type = "HTTP"
  interval = 260
  units3 = "MSEC"
  resptimeout = 150
  units4 = "MSEC"
}

resource "netscaler_servicegroup" "backend_group" {
  servicegroupname = "backend_group_1"
  lbvserver = "${netscaler_lbvserver.sample_lb2.name}"
  servicetype = "HTTP"
  lbmonitor = "${netscaler_lbmonitor.foo-monitor.monitorname}"
  servicegroupmembers = ["172.20.0.20:200:50","172.20.0.101:80:10",  "172.20.0.10:80:40"]
}
