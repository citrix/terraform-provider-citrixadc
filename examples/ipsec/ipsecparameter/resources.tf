resource "citrixadc_ipsecparameter" "tf_ipsecparameter" {
  ikeversion            = "V2"
  encalgo               = ["AES", "3DES"]
  hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
  livenesscheckinterval = 50
}
