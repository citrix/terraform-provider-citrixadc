# Since the appflowpolicy resource is not yet available on Terraform,
# the tf_appflowpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add appflow collector col1 -IPaddress 10.10.5.5
# add appflow action test_action -collectors col1
# add appflow policy tf_appflowpolicy client.TCP.DSTPORT.EQ(22) test_action

resource "citrixadc_lbvserver_appflowpolicy_binding" "tf_lbvserver_appflowpolicy_binding" {
	name = citrixadc_lbvserver.tf_lbvserver.name
	policyname = "tf_appflowpolicy"
	labelname = citrixadc_lbvserver.tf_lbvserver.name
	gotopriorityexpression = "END"
	invoke = true
	labeltype = "reqvserver"
	priority = 1
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver"
	ipv46       = "10.10.10.33"
	port        = 80
	servicetype = "HTTP"
}