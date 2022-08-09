resource "citrixadc_clusternodegroup_authenticationvserver_binding" "tf_clusternodegroup_authenticationvserver_binding" {
  name     = "my_authentication_group"
  vserver = "my_authentication_server"
}
