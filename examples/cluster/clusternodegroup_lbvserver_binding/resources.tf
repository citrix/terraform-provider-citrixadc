resource "citrixadc_clusternodegroup_lbvserver_binding" "tf_clusternodegroup_lbvserver_binding" {
  name = "my_test_group"
  vserver = "my_lbvserver"
}