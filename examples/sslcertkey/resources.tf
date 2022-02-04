resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate.crt"
  key = "/var/tmp/key.pem"
  notificationperiod = 10
  expirymonitor = "ENABLED"
}
