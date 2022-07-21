resource "citrixadc_filterpolicy" "tf_filterpolicy" {
    name = "tf_filterpolicy"
    reqaction = "DROP"
    rule = "REQ.HTTP.URL CONTAINS http://abcd.com"
}

resource "citrixadc_filterglobal_filterpolicy_binding" "tf_filterglobal" {
    policyname = citrixadc_filterpolicy.tf_filterpolicy.name
    priority = 200
    state = "ENABLED"
}
