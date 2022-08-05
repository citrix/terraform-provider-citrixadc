resource "citrixadc_opoption" "tf_opoption" {
  snmpset              = "ENABLED"
  snmptraplogging      = "ENABLED"
  partitionnameintrap  = "ENABLED"
  snmptraplogginglevel = "WARNING"
}
