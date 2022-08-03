resource "citrixadc_responderglobal_responderpolicy_binding" "tf_responderglobal_responderpolicy_binding" {
  globalbindtype = "SYSTEM_GLOBAL"
  priority   = 50
  policyname =citrixadc_responderpolicy.tf_responderpolicy.name
}


resource "citrixadc_responderpolicy" "tf_responderpolicy" {
	name    = "tf_responderpolicy"
	action = "NOOP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}