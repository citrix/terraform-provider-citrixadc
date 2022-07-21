resource "citrixadc_mapdmr" "tf_mapdmr" {
  name         = "tf_mapdmr"
  bripv6prefix = "2002:db8::/64"
}
resource "citrixadc_mapdomain" "tf_mapdomain" {
  name       = "tf_mapdomain"
  mapdmrname = citrixadc_mapdmr.tf_mapdmr.name
}