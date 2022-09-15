resource "citrixadc_streamidentifier" "tf_streamidentifier" {
  name         = "my_streamidentifier"
  selectorname = "my_streamselector"
  samplecount  = 10
  sort         = "CONNECTIONS"
  snmptrap     = "ENABLED"
}
