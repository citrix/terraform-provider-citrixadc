resource "citrixadc_ipset" "tf_ipset" {
  name = "tf_test_ipset"
  #td   = "100"
  nsipbinding = [
    citrixadc_nsip.nsip1.ipaddress,
    citrixadc_nsip.nsip2.ipaddress,
  ]

  nsip6binding = [
    citrixadc_nsip6.nsip6_1.ipv6address,
  ]
}

# ipv4
resource "citrixadc_nsip" "nsip1" {
  ipaddress = "10.1.1.1"
  type      = "VIP"
  netmask   = "255.255.255.0"
}
resource "citrixadc_nsip" "nsip2" {
  ipaddress = "10.2.2.2"
  type      = "SNIP"
  netmask   = "255.255.255.0"
}

# ipv6
resource "citrixadc_nsip6" "nsip6_1" {
  ipv6address = "2009::2/64"
  type        = "VIP"
  icmp        = "DISABLED"
}