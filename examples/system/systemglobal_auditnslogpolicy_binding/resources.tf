resource "citrixadc_systemglobal_auditnslogpolicy_binding" "tf_systemglobal_auditnslogpolicy_binding" {
  policyname = "tf_auditnslogpolicy"
  priority   = 50
}