resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.34.21"
  port        = "443"
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_csvserver" "tf_csvserver" {

  name        = "tf_csvserver"
  ipv46       = "10.202.11.11"
  port        = 9090
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"
}


resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslpolicy"
  rule   = "true"
  action = "NOOP"
}

# Import id is <vservername>,<policyname>
resource "citrixadc_sslvserver_sslpolicy_binding" "tf_binding_lb" {
    vservername = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_sslpolicy.tf_sslpolicy.name
    priority = 333
    type = "REQUEST"
}

resource "citrixadc_sslvserver_sslpolicy_binding" "tf_binding_cs" {
    vservername = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_sslpolicy.tf_sslpolicy.name
    priority = 333
    type = "REQUEST"
}
