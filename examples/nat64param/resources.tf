resource "citrixadc_nat64param" "tf_nat64param" {
  nat64ignoretos    = "NO"
  nat64zerochecksum = "ENABLED"
  nat64v6mtu        = 1280
  nat64fragheader   = "ENABLED"
}