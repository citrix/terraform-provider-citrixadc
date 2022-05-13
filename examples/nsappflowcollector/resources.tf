resource "citrixadc_nsappflowcollector" "tf_appflowcollector" {
  name      = "tf_appflowcollector"
  ipaddress = "1.2.4.1"
  port      = 30
}