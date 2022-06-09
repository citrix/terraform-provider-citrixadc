resource "citrixadc_nsip6" "test_nsip" {
  ipv6address = "23::30:20:23:34/64"
  type        = "VIP"
  icmp        = "DISABLED"
}
resource "citrixadc_ip6tunnel" "tf_ip6tunnel" {
  name   = "tf_ip6tunnel"
  remote = "2001:db8:0:b::/64"
  local  = trimsuffix(citrixadc_nsip6.test_nsip.ipv6address, "/64")
}