resource "citrixadc_spilloverpolicy" "policy1" {
  action  = "spillover"
  comment = "spillover"
  name    = "spilloverpolicy1"
  rule    = "true"
}
