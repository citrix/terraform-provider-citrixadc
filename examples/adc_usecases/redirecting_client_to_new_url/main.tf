resource "citrixadc_responderaction" "responder_action" {
  name    = "act_redirect"
  type    = "redirect"
  target  = "\"http://www.example.com/404.html\""
  comment = "Other-Comment"
}
resource "citrixadc_responderpolicy" "responder_policy" {
  name   = "pol_redirect"
  action = citrixadc_responderaction.responder_action.name
  rule   = "CLIENT.IP.SRC.IN_SUBNET(222.222.0.0/16)"
}
resource "citrixadc_responderglobal_responderpolicy_binding" "binding" {
  policyname = citrixadc_responderpolicy.responder_policy.name
  priority   = 500
}