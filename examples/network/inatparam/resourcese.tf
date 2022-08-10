resource "citrixadc_inatparam" "tf_inatparam" {
  nat46ignoretos    = "NO"
  nat46zerochecksum = "ENABLED"
  nat46v6mtu        = "1400"
}
