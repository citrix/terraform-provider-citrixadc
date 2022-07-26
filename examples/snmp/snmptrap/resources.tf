resource "citrixadc_snmptrap" "tf_snmptrap" {
  severity        = "Major"
  trapclass       = "specific"
  trapdestination = "192.168.2.2"
}