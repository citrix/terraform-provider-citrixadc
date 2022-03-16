resource "citrixadc_authenticationepaaction" "tf_epaaction" {
  name            = "tf_epaaction"
  csecexpr        = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
  defaultepagroup = "new_group"
  deletefiles     = "old_files"
  killprocess     = "old_process"
}