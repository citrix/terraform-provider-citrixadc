resource "citrixadc_mapbmr" "tf_mapbmr" {
  name           = "tf_mapbmr"
  ruleipv6prefix = "2001:db8:abcd:12::/64"
  psidoffset     = 6
  eabitlength    = 16
  psidlength     = 8
}
resource "citrixadc_mapbmr_bmrv4network_binding" "tf_binding" {
  name    = citrixadc_mapbmr.tf_mapbmr.name
  network = "1.2.3.0"
  netmask = "255.255.255.0"
}