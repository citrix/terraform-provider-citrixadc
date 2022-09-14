resource "citrixadc_lsnclient_nsacl_binding" "tf_lsnclient_nsacl_binding" {
  clientname = "my_lsnclient"
  aclname    = "my_acl"
}
