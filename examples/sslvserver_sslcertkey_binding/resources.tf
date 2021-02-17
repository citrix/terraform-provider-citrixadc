resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate2.crt"
  key = "/var/tmp/key2.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.44"
  name        = "tf_lbvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}


resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
    vservername = citrixadc_lbvserver.tf_lbvserver.name
    certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
    snicert = true
}
