resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}

resource "citrixadc_lbvserver_transformpolicy_binding" "tf_binding" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_transformpolicy.tf_trans_policy.name
    priority = 100
    bindpoint = "REQUEST"
    gotopriorityexpression = "END"
}
