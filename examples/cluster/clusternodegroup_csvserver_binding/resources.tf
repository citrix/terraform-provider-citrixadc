resource "citrixadc_clusternodegroup_csvserver_binding" "tf_clusternodegroup_csvserver_binding" {
  name = "my_cs_group"
  vserver = "my_csvserver"
}