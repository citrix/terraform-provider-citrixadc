resource "citrixadc_rnat6" "tf_rnat6" {
  name             = "my_rnat6"
  network          = "2003::/64"
  srcippersistency = "ENABLED"
}
