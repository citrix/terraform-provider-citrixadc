resource "citrixadc_tmglobal_auditsyslogpolicy_binding" "tf_tmglobal_auditsyslogpolicy_binding" {
  policyname = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
  priority   = 100
}

resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
    name = "tf_auditsyslogpolicy"
    rule = "ns_true"
    action = "tf_syslogaction"
}