resource "citrixadc_auditnslogpolicy" "policy1" {
    name = "policy1"
    rule = "true"
    action = "SETASLEARNNSLOG_ACT"
}