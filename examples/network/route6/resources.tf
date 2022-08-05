resource "citrixadc_route6" "tf_route6" {
  network  = "2001:db8:85a3::/64"
  vlan     = 2
  weight   = 5
  distance = 3
}
