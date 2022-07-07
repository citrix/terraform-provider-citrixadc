resource "citrixadc_csvserver" "test_csvserver" {

  ipv46       = "10.10.10.22"
  name        = "test_csvserver"
  port        = 80
  servicetype = "HTTP"

}
resource "citrixadc_lbvserver" "test_lbvserver" {

  ipv46       = "10.10.10.33"
  name        = "test_lbvserver"
  port        = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_test_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

  csvserverbinding {
    name                   = citrixadc_csvserver.test_csvserver.name
    bindpoint              = "REQUEST"
    priority               = 114
    gotopriorityexpression = "END"
  }

  globalbinding {
    gotopriorityexpression = "END"
    labelname              = citrixadc_lbvserver.test_lbvserver.name
    labeltype              = "reqvserver"
    priority               = 205
    invoke                 = true
    type                   = "REQ_DEFAULT"
  }

}

resource "citrixadc_rewritepolicylabel" "rewrite_policylabel" {
  labelname = "tf_test_label"
  transform = "http_req"
  comment   = "Some comment other"
}

resource "citrixadc_rewriteaction" "rewrite_action" {
  bypasssafetycheck = "NO"
  name              = "tf_test_rewrite_action"
  target            = "HTTP.REQ.HOSTNAME"
  type              = "delete"
}
