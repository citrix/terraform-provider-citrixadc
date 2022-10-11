resource "citrixadc_lbvserver" "lb_1" {
  name        = var.lbvserver1_name
  ipv46       = var.lbvserver1_ip
  port        = var.lbvserver1_port
  servicetype = var.lbvserver1_servicetype
  state       = "ENABLED"
}
resource "citrixadc_lbvserver" "lb_2" {
  name        = var.lbvserver2_name
  ipv46       = var.lbvserver2_ip
  port        = var.lbvserver2_port
  servicetype = var.lbvserver2_servicetype
  state       = "ENABLED"
}
resource "citrixadc_service" "tf_service1" {
  name        = var.service1_name
  servicetype = var.service1_servicetype
  ipaddress   = var.service1_ip
  ip          = var.service1_ip
  port        = var.service1_port
  state       = "ENABLED"
}
resource "citrixadc_service" "tf_service2" {
  name        = var.service2_name
  servicetype = var.service2_servicetype
  ipaddress   = var.service2_ip
  ip          = var.service2_ip
  port        = var.service2_port
  state       = "ENABLED"
}
resource "citrixadc_service" "tf_service3" {
  name        = var.service3_name
  servicetype = var.service3_servicetype
  ipaddress   = var.service3_ip
  ip          = var.service3_ip
  port        = var.service3_port
  state       = "ENABLED"
}
resource "citrixadc_service" "tf_service4" {
  name        = var.service4_name
  servicetype = var.service4_servicetype
  ipaddress   = var.service4_ip
  ip          = var.service4_ip
  port        = var.service4_port
  state       = "ENABLED"
}
resource "citrixadc_lbvserver_service_binding" "tf_binding1" {
  name        = citrixadc_lbvserver.lb_1.name
  servicename = citrixadc_service.tf_service1.name
}
resource "citrixadc_lbvserver_service_binding" "tf_binding2" {
  name        = citrixadc_lbvserver.lb_1.name
  servicename = citrixadc_service.tf_service2.name
}
resource "citrixadc_lbvserver_service_binding" "tf_binding3" {
  name        = citrixadc_lbvserver.lb_2.name
  servicename = citrixadc_service.tf_service3.name
}
resource "citrixadc_lbvserver_service_binding" "tf_binding4" {
  name        = citrixadc_lbvserver.lb_2.name
  servicename = citrixadc_service.tf_service4.name
}
resource "citrixadc_csvserver" "tf_csvserver" {
  name        = var.csvserver_name
  ipv46       = var.csvserver_ipv46
  port        = var.csvserver_port
  servicetype = var.csvserver_servicetype
  state       = "ENABLED"
  stateupdate = "ENABLED"
}

resource "citrixadc_csaction" "csaction_html" {
  name            = "csaction_HTML"
  targetlbvserver = citrixadc_lbvserver.lb_1.name
}
resource "citrixadc_csaction" "csaction_image" {
  name            = "csaction_Image"
  targetlbvserver = citrixadc_lbvserver.lb_2.name
}
resource "citrixadc_cspolicy" "tf_cspolicy1" {
  policyname = var.cspolicy1_name
  rule       = var.cspolicy1_rule
  action     = citrixadc_csaction.csaction_html.name
}
resource "citrixadc_cspolicy" "tf_cspolicy2" {
  policyname = var.cspolicy2_name
  rule       = var.cspolicy2_rule
  action     = citrixadc_csaction.csaction_html.name
}
resource "citrixadc_cspolicy" "tf_cspolicy3" {
  policyname = var.cspolicy3_name
  rule       = var.cspolicy3_rule
  action     = citrixadc_csaction.csaction_image.name
}
resource "citrixadc_cspolicy" "tf_cspolicy4" {
  policyname = var.cspolicy4_name
  rule       = var.cspolicy4_rule
  action     = citrixadc_csaction.csaction_image.name
}

resource "citrixadc_csvserver_cspolicy_binding" "tf_bind1" {
  name       = citrixadc_csvserver.tf_csvserver.name
  policyname = citrixadc_cspolicy.tf_cspolicy1.policyname
  priority   = 100
}
resource "citrixadc_csvserver_cspolicy_binding" "tf_bind2" {
  name       = citrixadc_csvserver.tf_csvserver.name
  policyname = citrixadc_cspolicy.tf_cspolicy2.policyname
  priority   = 200
}
resource "citrixadc_csvserver_cspolicy_binding" "tf_bind3" {
  name       = citrixadc_csvserver.tf_csvserver.name
  policyname = citrixadc_cspolicy.tf_cspolicy3.policyname
  priority   = 300
}
resource "citrixadc_csvserver_cspolicy_binding" "tf_bind4" {
  name       = citrixadc_csvserver.tf_csvserver.name
  policyname = citrixadc_cspolicy.tf_cspolicy4.policyname
  priority   = 400
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = var.sslcertkey_name
  cert    = var.sslcertkey_cert
  key     = var.sslcertkey_key
}
resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
  vservername = citrixadc_csvserver.tf_csvserver.name
  certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
}
