resource "citrixadc_vrid" "tf_vrid" {
  vrid_id              = 3
  priority             = 30
  preemption           = "DISABLED"
  sharing              = "ENABLED"
  tracking             = "NONE"
  trackifnumpriority   = 0
  preemptiondelaytimer = 0
}