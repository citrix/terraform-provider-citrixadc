resource "citrixadc_nd6" "tf_nd6" {
  neighbor = "2001::3"
  mac      = "e6:ec:41:50:b1:d1"
  ifnum    = "LO/1"
}
