resource "citrixadc_rewriteglobal_rewritepolicy_binding" "tf_rewriteglobal_rewritepolicy_binding" {
	policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
	priority = 5
	type = "REQ_DEFAULT"
	globalbindtype = "SYSTEM_GLOBAL"
	gotopriorityexpression = "END"
	invoke = "true"
	labelname = citrixadc_rewritepolicylabel.tf_rewritepolicylabel.labelname
	labeltype = "policylabel"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
	labelname = "tf_rewritepolicylabel"
	transform = "http_req"
}