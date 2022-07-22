# Since the tmttrafficpolicy resource is not yet available on Terraform,
# the tf_tmttrafficpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add tmtrafficAction tf_tmtrafficaction -SSO ON -userExpression "AAA.USER.PASSWD"
# add tmtrafficPolicy tf_tmttrafficpolicy true tf_tmtrafficaction

resource "citrixadc_lbvserver_tmtrafficpolicy_binding" "tf_lbvserver_tmtrafficpolicy_binding" {
	name = citrixadc_lbvserver.tf_lbvserver.name
	policyname = "tf_tmttrafficpolicy"
	priority = 1
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver"
	ipv46       = "10.10.10.33"
	port        = 80
	servicetype = "HTTP"
}