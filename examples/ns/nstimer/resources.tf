resource "citrixadc_nstimer" "tf_nstimer" {
  name     = "tf_nstimer"
  interval = 10
  unit     = "SEC"
  comment  = "Testing"
}