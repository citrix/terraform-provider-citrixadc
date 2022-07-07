resource "citrixadc_sslcertkey" "foo" {
  certkey            = "sample_ssl_cert"
  cert               = "/var/tmp/certificate1.crt"
  key                = "/var/tmp/key1.pem"
  notificationperiod = 10
  expirymonitor      = "ENABLED"
}
resource "citrixadc_vpnglobal_sslcertkey_binding" "tf_vpnglobal_slcertkey_binding" {
  certkeyname = citrixadc_sslcertkey.foo.certkey
}