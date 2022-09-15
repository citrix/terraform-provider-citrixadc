resource "citrixadc_hanode_routemonitor_binding" "tf_hanode_routemonitor_binding" {
  hanode_id    = 0
  routemonitor = "10.222.74.128"
  netmask      = "255.255.255.192"
}
