resource "citrixadc_rnat6_nsip6_binding" "tf_rnat6_nsip6_binding" {
  name  = "my_rnat6"
  natip6 = "2001:db8:85a3::8a2e:370:7334"
}