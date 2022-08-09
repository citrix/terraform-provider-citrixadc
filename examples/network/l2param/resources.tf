resource "citrixadc_l2param" "tf_l2param" {
  mbfpeermacupdate   = 20
  maxbridgecollision = 30
  bdggrpproxyarp     = "DISABLED"
}
