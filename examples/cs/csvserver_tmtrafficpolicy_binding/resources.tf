# Since the tmttrafficpolicy resource is not yet available on Terraform,
# the tf_tmttrafficpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add tmtrafficAction tf_tmtrafficaction -SSO ON -userExpression "AAA.USER.PASSWD"
# add tmtrafficPolicy tf_tmttrafficpolicy true tf_tmtrafficaction

resource "citrixadc_csvserver_tmtrafficpolicy_binding" "tf_csvserver_tmtrafficpolicy_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	policyname = "tf_tmttrafficpolicy"
	priority = 1
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "HTTP"
}