# Since the contentinspectionpolicy resource is not yet available on Terraform,
# the tf_contentinspectionpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add contentinspection policy tf_contentinspectionpolicy -rule true -action NOINSPECTION

resource "citrixadc_lbvserver_contentinspectionpolicy_binding" "tf_lbvserver_contentinspectionpolicy_binding" {
	bindpoint = "REQUEST"
	gotopriorityexpression = "END"
	name = citrixadc_lbvserver.tf_lbvserver.name
	policyname = "tf_contentinspectionpolicy"
	priority = 1    
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver"
	ipv46       = "10.10.10.33"
	port        = 80
	servicetype = "HTTP"
}
