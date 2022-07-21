resource "citrixadc_transformpolicy" "tf_trans_policy" {
  name        = "tf_trans_policy"
  profilename = "pro_1"
  rule        = "http.REQ.URL.CONTAINS(\"test_url\")"
}
resource "citrixadc_transformpolicylabel" "transformpolicylabel" {
  labelname       = "label_1"
  policylabeltype = "httpquic_req"
}
resource "citrixadc_transformpolicylabel_transformpolicy_binding" "transformpolicylabel_transformpolicy_binding" {
  policyname = citrixadc_transformpolicy.tf_trans_policy.name
  labelname  = citrixadc_transformpolicylabel.transformpolicylabel.labelname
  priority   = 2
}