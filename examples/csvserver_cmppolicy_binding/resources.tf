resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_cmppolicy" "tf_cmppolicy" {
    name = "tf_cmppolicy"
    rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
    resaction = "COMPRESS"
}

resource "citrixadc_csvserver_cmppolicy_binding" "tf_bind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_cmppolicy.tf_cmppolicy.name
    priority = 110
    bindpoint = "RESPONSE"
    gotopriorityexpression = "END"
}
