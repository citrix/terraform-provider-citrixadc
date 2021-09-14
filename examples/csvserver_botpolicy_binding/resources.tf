resource "citrixadc_csvserver_botpolicy_binding" "tf_csvserver_botpolicy_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	policyname = citrixadc_botpolicy.tf_botpolicy.name
	priority = 5
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "HTTP"
}

resource citrixadc_botpolicy tf_botpolicy {
	name = "tf_botpolicy"  
	profilename = "BOT_BYPASS"
	rule  = "true"
}