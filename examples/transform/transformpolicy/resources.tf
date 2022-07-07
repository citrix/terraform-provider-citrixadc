resource "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformaction" "tf_trans_action1" {
  name = "tf_trans_action1"
  profilename = citrixadc_transformprofile.tf_trans_profile.name
  priority = 100
  requrlfrom = "http://m3.mydomain.com/(.*)"
  requrlinto = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto = "https://m3.mydomain.com/$1"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}
