resource "citrixadc_service" "tf_service" {
  servicetype         = "HTTP"
  name                = "tf_service"
  ipaddress           = "10.77.33.22"
  ip                  = "10.77.33.22"
  port                = "80"
  state               = "ENABLED"
  wait_until_disabled = true
}
resource "citrixadc_lbmonitor" "tf_monitor" {
  monitorname = "tf_monitor"
  type        = "HTTP"
}
resource "citrixadc_service_lbmonitor_binding" "tf_binding" {
  name         = citrixadc_service.tf_service.name
  monitor_name = citrixadc_lbmonitor.tf_monitor.monitorname
  monstate     = "ENABLED"
  weight       = 2
}