resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
  name       = "new_syslogaction"
  serverip   = "20.3.3.3"
  serverport = 54
  loglevel = [
    "ERROR",
    "NOTICE",
  ]
}

resource "citrixadc_auditsyslogpolicy" "tf_policy" {
  name   = "new_auditsyslogpolicy"
  rule   = "ns_true"
  action = citrixadc_auditsyslogaction.tf_syslogaction.name
}

resource "citrixadc_vpnglobal_auditsyslogpolicy_binding" "tf_bind" {
  policyname             = citrixadc_auditsyslogpolicy.tf_policy.name
  priority               = 300
  gotopriorityexpression = "NEXT"
}