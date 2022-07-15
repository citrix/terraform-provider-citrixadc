resource "citrixadc_transformpolicy" "tf_trans_policy" {
  name        = "tf_trans_policy"
  profilename = "tf_trans_profile"
  rule        = "http.REQ.URL.CONTAINS(\"test_url\")"
}
resource "citrixadc_transformglobal_transformpolicy_binding" "transformglobal_transformpolicy_binding" {
  policyname = citrixadc_transformpolicy.tf_trans_policy.name
  priority   = 2
  type       = "REQ_DEFAULT"
}