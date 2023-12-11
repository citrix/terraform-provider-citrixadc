#create responder action for http to https redirection
resource "citrixadc_responderaction" "tf_responderaction" {
  name              = var.rs_action["name"]
  type              = var.rs_action["type"]
  bypasssafetycheck = "YES"
  target            = var.rs_action["target"]
}
#create HTTP to HTTPs responder redirection policy
resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name   = var.rs_policy["name"]
  action = citrixadc_responderaction.tf_responderaction.name
  rule   = var.rs_policy["rule"]
}
