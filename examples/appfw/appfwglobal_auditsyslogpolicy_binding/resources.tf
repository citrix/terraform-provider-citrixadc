resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
  name       = "tf_syslogaction"
  serverip   = "10.78.60.33"
  serverport = 514
  loglevel = [
    "ERROR",
    "NOTICE",
  ]
}
resource "citrixadc_auditsyslogpolicy" "tf_policy" {
  name   = "tf_auditsyslogpolicy"
  rule   = "ns_true"
  action = citrixadc_auditsyslogaction.tf_syslogaction.name
}
resource "citrixadc_appfwglobal_auditsyslogpolicy_binding" "tf_binding" {
  policyname = citrixadc_auditsyslogpolicy.tf_policy.name
  priority   = 90
  state      = "DISABLED"
  type       = "NONE"
}