# Since the appflowpolicy resource is not yet available on Terraform,
# the tf_appflowpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add appflow collector col1 -IPaddress 10.10.5.5
# add appflow action test_action -collectors col1
# add appflow policy tf_appflowpolicy client.TCP.DSTPORT.EQ(22) test_action

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_appflowpolicy_binding" "tf_bind" {
  name                   = citrixadc_vpnvserver.tf_vpnvserver.name
  policy                 = "tf_appflowpolicy"
  bindpoint              = "ICA_REQUEST"
  priority               = 200
  gotopriorityexpression = "END"
}