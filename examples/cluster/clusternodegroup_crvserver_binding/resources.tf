resource "citrixadc_clusternodegroup_crvserver_binding" "tf_clusternodegroup_crvserver_binding" {
  name = "my_cr_group"
  vserver = "my_crvserver"
}