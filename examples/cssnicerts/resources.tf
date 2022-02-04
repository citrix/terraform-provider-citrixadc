resource "citrixadc_sslcertkey" "cert2" {
  certkey       = "cert2"
  cert          = "/var/tmp/certificate2.crt"
  key           = "/var/tmp/key2.pem"
  expirymonitor = "DISABLED"
}

resource "citrixadc_sslcertkey" "cert3" {
  certkey       = "cert3"
  cert          = "/var/tmp/certificate3.crt"
  key           = "/var/tmp/key3.pem"
  expirymonitor = "DISABLED"
}

resource "citrixadc_csvserver" "cssni" {
  ipv46       = "10.202.11.11"
  name        = "terraform-cs"
  port        = 443
  servicetype = "SSL"
  ciphers     = ["DEFAULT"]


  snisslcertkeys = [
    "${citrixadc_sslcertkey.cert2.certkey}",
    "${citrixadc_sslcertkey.cert3.certkey}",
  ]

}

