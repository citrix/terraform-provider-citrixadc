resource "citrixadc_vpnsamlssoprofile" "tf_vpnsamlssoprofile" {
  name                        = "tf_vpnsamlssoprofile"
  assertionconsumerserviceurl = "http://www.example.com"
  sendpassword                = "ON"
  relaystaterule              = "true"
  samlissuername              = "asdf"
  signaturealg                = "RSA-SHA1"
  digestmethod                = "SHA256"
  nameidformat                = "Unspecified"
}
