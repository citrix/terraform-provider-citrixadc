resource "citrixadc_vrid6" "tf_vrid6" {
  vrid6_id             = 3
  priority             = 30
  preemption           = "DISABLED"
  sharing              = "DISABLED"
  tracking             = "NONE"
  trackifnumpriority   = 0
  preemptiondelaytimer = 0
}