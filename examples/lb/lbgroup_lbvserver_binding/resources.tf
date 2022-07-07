resource "citrixadc_lbgroup_lbvserver_binding" "tf_lbvserverbinding" {
	name = citrixadc_lbgroup.tf_lbgroup.name
	vservername = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name 		    = "tf_lbvserver"
	ipv46       = "1.1.1.8"
	port        = "80"
	servicetype = "HTTP"
}

resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup"
}