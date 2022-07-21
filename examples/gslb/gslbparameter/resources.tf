resource "citrixadc_gslbparameter" "tf_gslbparameter" {
  ldnsentrytimeout = 50
  rtttolerance     = 10
  ldnsmask         = "255.255.255.255"
}


