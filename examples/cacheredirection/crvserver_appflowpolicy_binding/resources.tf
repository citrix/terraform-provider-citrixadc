resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_crvserver_appflowpolicy_binding" "crvserver_appflowpolicy_binding" {
	name = citrixadc_crvserver.crvserver.name
	policyname = "tf_appflowpolicy"
	gotopriorityexpression = "END"
    labelname = citrixadc_crvserver.crvserver.name
	invoke = true
	labeltype = "reqvserver"
	priority = 1
}