resource "citrixadc_rewriteaction" "request_server_replace" {
  name              = "Action-Rewrite-Request_Server_Replace"
  type              = "replace"
  target            = "HTTP.REQ.HOSTNAME.SERVER"
  stringbuilderexpr = "\"Web.hq.example.net\""
}
resource "citrixadc_rewriteaction" "response_server_replace" {
  name              = "Action-Rewrite-Response_Server_Replace"
  type              = "replace"
  target            = "HTTP.RES.HEADER(\"Server\")"
  stringbuilderexpr = "\"www.example.com\""
}

resource "citrixadc_rewritepolicy" "request_server_replace" {
  name        = "Rewrite-Request_Server_Replace"
  rule        = "HTTP.REQ.HOSTNAME.SERVER.EQ(\"www.example.com\")"
  action      = citrixadc_rewriteaction.request_server_replace.name
  undefaction = "NOREWRITE"
}
resource "citrixadc_rewritepolicy" "response_server_replace" {
  name   = "Rewrite-Response_Server_Replace"
  rule   = "HTTP.REQ.HEADER(\"Server\").EQ(\"Web.hq.example.net\")"
  action = citrixadc_rewriteaction.response_server_replace.name
}


resource "citrixadc_rewriteglobal_rewritepolicy_binding" "binding_request_server_replace" {
  policyname             = citrixadc_rewritepolicy.request_server_replace.name
  priority               = 500
  type                   = "REQ_DEFAULT"
  gotopriorityexpression = "END"
}
resource "citrixadc_rewriteglobal_rewritepolicy_binding" "binding_response_server_replace" {
  policyname             = citrixadc_rewritepolicy.response_server_replace.name
  priority               = 500
  type                   = "RES_DEFAULT"
  gotopriorityexpression = "END"
}
