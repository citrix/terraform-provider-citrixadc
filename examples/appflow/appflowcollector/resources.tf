resource "citrixadc_appflowcollector" "tf_appflowcollector" {
  name      = "tf_collector"
  ipaddress = "192.168.2.2"
  transport = "logstream"
  port      =  80
}
