resource "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "tf_auditnslogglobal_auditnslogpolicy_binding" {
  policyname = "SETASLEARNNSLOG_ADV_POL"
  priority   = 100
  globalbindtype = "SYSTEM_GLOBAL"
}