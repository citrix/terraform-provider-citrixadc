resource "citrixadc_snmpalarm" "tf_snmpalarm" {
  trapname = "PORT-ALLOC-EXCEED"
  thresholdvalue = 10
  normalvalue    = 5
  time           = 60
  state          = "DISABLED"
  severity       = "Minor"
}
