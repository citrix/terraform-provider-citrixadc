resource "citrixadc_rewritepolicylabel_rewritepolicy_binding" "tf_rewritepolicylabel_rewritepolicy_binding" {
	labelname = citrixadc_rewritepolicylabel.tf_rewritepolicylabel.labelname
	policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
	gotopriorityexpression = "END"
	priority = 5   
}

resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
	labelname = "tf_rewritepolicylabel"
	transform = "http_req"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}