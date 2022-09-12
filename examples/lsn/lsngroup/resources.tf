resource "citrixadc_lsngroup" "tf_lsngroup" {
  groupname     = "my_lsngroup"
  clientname    = "my_lsnclient"
  logging       = "DISABLED"
  nattype       = "DYNAMIC"
}
