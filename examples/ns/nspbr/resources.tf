resource "citrixadc_nspbr" "tf_nspbr" {
  name       = "my_nspbr"
  action     = "ALLOW"
  nexthop    = "true"
  nexthopval = "10.222.74.128"
}

