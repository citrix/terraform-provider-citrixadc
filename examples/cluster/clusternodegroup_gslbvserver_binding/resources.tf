resource "citrixadc_clusternodegroup_gslbvserver_binding" "tf_clusternodegroup_gslbvserver_binding" {
  name = "my_gslb_group"
  vserver = "my_gslbvserver"
}