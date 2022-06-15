resource "citrixadc_mapbmr" "tf_mapbmr" {
  name           = "tf_mapbmr"
  ruleipv6prefix = "2001:db8:abcd:12::/64"
  psidoffset     = 6
  eabitlength    = 16
  psidlength     = 8
}
resource "citrixadc_mapdmr" "tf_mapdmr" {
  name         = "tf_mapdmr"
  bripv6prefix = "2002:db8::/64"
}
resource "citrixadc_mapdomain" "tf_mapdomain" {
  name       = "tf_mapdomain"
  mapdmrname = citrixadc_mapdmr.tf_mapdmr.name
}
resource "citrixadc_mapdomain_mapbmr_binding" "tf_binding" {
  name       = citrixadc_mapdomain.tf_mapdomain.name
  mapbmrname = citrixadc_mapbmr.tf_mapbmr.name
}