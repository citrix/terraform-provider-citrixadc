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
resource "citrixadc_responderpolicy" "responder_policy" {
  name   = "test_responder_policy"
  action = "NOOP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"

  csvserverbinding {
    priority               = 200
    name                   = citrixadc_csvserver.test_csvserver.name
    gotopriorityexpression = "END"
    invoke                 = false
    bindpoint              = "REQUEST"
  }

  globalbinding {
    invoke                 = true
    labeltype              = "vserver"
    labelname              = citrixadc_lbvserver.test_lbvserver.name
    type                   = "REQ_OVERRIDE"
    gotopriorityexpression = "END"
    priority               = 666
  }
}

resource "citrixadc_responderaction" "responder_action" {
  name    = "test_action_2"
  type    = "respondwith"
  target  = "HTTP.REQ.HEADER(\"jallopy\")"
  comment = "other comment"
}

resource "citrixadc_rewritepolicylabel" "rewrite_policylabel" {
  labelname = "tf_test_label"
  transform = "http_req"
  comment   = "Some comment other"
}
