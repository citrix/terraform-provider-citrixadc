resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
  limitidentifier  = "tf_nslimitidentifier"
  threshold        = 1
  timeslice        = 1000
  limittype        = "BURSTY"
  mode             = "REQUEST_RATE"
  maxbandwidth     = 0
  trapsintimeslice = 1
}