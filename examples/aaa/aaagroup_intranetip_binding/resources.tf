resource "citrixadc_aaagroup_intranetip_binding" "tf_aaagroup_intranetip_binding" {
  groupname  = "my_group"
  intranetip = "10.222.73.160"
  netmask    = "255.255.255.192"
}
