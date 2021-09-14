# Since the appqoepolicy resource is not yet available on Terraform,
# the tf_appqoepolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add appqoe action tf_appqoeaction -priority MEDIUM
# add appqoe policy tf_appqoepolicy -rule true -action tf_appqoeaction

resource "citrixadc_csvserver_appqoepolicy_binding" "tf_csvserver_appqoepolicy_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	policyname = "tf_appqoepolicy"
	bindpoint = "REQUEST"
	priority = 5
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name        = "tf_csvserver"
	ipv46       = "10.10.10.33"
	port        = 80
	servicetype = "HTTP"
}