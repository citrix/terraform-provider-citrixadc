resource "citrixadc_ipsecprofile" "tf_ipsecprofile" {
  name                  = "my_ipsecprofile"
  ikeversion            = "V2"
  encalgo               = ["AES", "3DES"]
  hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
  livenesscheckinterval = 50
  psk                   = "GCC5VcY0TQ+0TfjGwCrR+cQthm5UnBPB"
}
