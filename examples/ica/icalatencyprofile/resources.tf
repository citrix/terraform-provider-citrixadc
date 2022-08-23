resource "citrixadc_icalatencyprofile" "tf_icalatencyprofile" {
  name                     = "my_ica_latencyprofile"
  l7latencymonitoring      = "ENABLED"
  l7latencythresholdfactor = 120
  l7latencywaittime        = 100
}
