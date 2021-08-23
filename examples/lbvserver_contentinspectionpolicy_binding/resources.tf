resource "citrixadc_lbvserver_contentinspectionpolicy_binding" "tf_lbvserver_contentinspectionpolicy_binding" {
	bindpoint = "REQUEST"
	gotopriorityexpression = "END"
	name = "tf_lbvserver"
	policyname = "tf_contentinspectionpolicy"
	priority = 1    
}