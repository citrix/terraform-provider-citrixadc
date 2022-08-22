resource "citrixadc_cmppolicylabel_cmppolicy_binding" "tf_cmppolicylabel_cmppolicy_binding" {
  policyname = citrixadc_cmppolicy.tf_cmppolicy.name
  labelname  = "my_cmppolicy_label"
  priority   = 100
}

resource "citrixadc_cmppolicy" "tf_cmppolicy" {
  name      = "tf_cmppolicy"
  rule      = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
  resaction = "COMPRESS"
}
