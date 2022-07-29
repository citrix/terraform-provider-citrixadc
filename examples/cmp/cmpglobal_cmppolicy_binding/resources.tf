resource "citrixadc_cmpglobal_cmppolicy_binding" "tf_cmpglobal_cmppolicy_binding" {
  globalbindtype = "SYSTEM_GLOBAL"
  priority   = 50
  policyname =citrixadc_cmppolicy.tf_cmppolicy.name
}

resource "citrixadc_cmppolicy" "tf_cmppolicy" {
    name = "tf_cmppolicy"
    rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
    resaction = "COMPRESS"
}