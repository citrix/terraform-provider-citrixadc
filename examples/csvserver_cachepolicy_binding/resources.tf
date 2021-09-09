# Since the cachepolicy resource is not yet available on Terraform,
# the tf_cachepolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add cache policy tf_cachepolicy -rule "http.req.url.query.contains(\"IssuePage\")" -action CACHE

resource "citrixadc_csvserver_cachepolicy_binding" "tf_csvserver_cachepolicy_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	policyname = "tf_cachepolicy"
	priority = 5       
	bindpoint = "REQUEST" 
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "HTTP"
}