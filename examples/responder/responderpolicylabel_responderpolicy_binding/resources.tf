resource "citrixadc_responderpolicylabel_responderpolicy_binding" "tf_responderpolicylabel_responderpolicy_binding" {
	labelname = citrixadc_responderpolicylabel.tf_responderpolicylabel.labelname
	policyname = citrixadc_responderpolicy.tf_responderpolicy.name
	priority = 5  
	gotopriorityexpression = "END"
	invoke = "false"
}

resource "citrixadc_responderpolicylabel" "tf_responderpolicylabel" {
	labelname = "tf_responderpolicylabel"
	policylabeltype = "HTTP"
}

resource "citrixadc_responderpolicy" "tf_responderpolicy" {
	name    = "tf_responderpolicy"
	action = "NOOP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}