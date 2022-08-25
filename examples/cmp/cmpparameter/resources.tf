resource "citrixadc_cmpparameter" "tf_cmpparameter" {
  cmplevel    = "optimal"
  quantumsize = 20
  servercmp   = "OFF"
}
