resource "citrixadc_authorizationpolicy" "authorize" {
	name   = "tp-authorize-1"
	rule   = "true"
	action = "DENY"
  }
  resource "citrixadc_authorizationpolicylabel" "authorizationpolicylabel" {
	labelname = "trans_http_url"
  }
  resource "citrixadc_authorizationpolicylabel_authorizationpolicy_binding" "authorizationpolicylabel_authorizationpolicy_binding" {
	policyname = citrixadc_authorizationpolicy.authorize.name
	labelname = citrixadc_authorizationpolicylabel.authorizationpolicylabel.labelname
	priority = 2
  }