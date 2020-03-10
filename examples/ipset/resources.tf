resource "citrixadc_ipset" "tf_ipset" {
  name = "tf_test_ipset"
  #td   = "100"
  nsipbinding = [
    citrixadc_nsip.nsip1.ipaddress,
    citrixadc_nsip.nsip2.ipaddress,
  ]
}

resource "citrixadc_nsip" "nsip1" {
  ipaddress = "1.1.1.1"
  type      = "VIP"
  netmask   = "255.255.255.0"
}
resource "citrixadc_nsip" "nsip2" {
  ipaddress = "2.2.2.2"
  type      = "SNIP"
  netmask   = "255.255.255.0"
}
