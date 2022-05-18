resource "citrixadc_nsdiameter" "tf_nsdiameter" {
  identity               = "citrixadc.com"
  realm                  = "com"
  serverclosepropagation = "OFF"
}