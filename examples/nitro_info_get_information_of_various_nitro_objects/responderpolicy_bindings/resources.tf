data "citrixadc_nitro_info" "csvs_bindings" {
    workflow = yamldecode(file("../workflows/responderpolicy_csvserver_binding.yaml"))
    primary_id = "tf_responder_policy"
}

output "csvs_object_output" {
    value = [ for item in data.citrixadc_nitro_info.csvs_bindings.nitro_list: item.object ]
}

data "citrixadc_nitro_info" "lbvs_bindings" {
    workflow = yamldecode(file("../workflows/responderpolicy_lbvserver_binding.yaml"))
    primary_id = "tf_responder_policy"
}

output "lbvs_object_output" {
    value = [ for item in data.citrixadc_nitro_info.lbvs_bindings.nitro_list: item.object ]
}

data "citrixadc_nitro_info" "global_bindings" {
    workflow = yamldecode(file("../workflows/responderpolicy_responderglobal_binding.yaml"))
    primary_id = "tf_responder_policy"
}

output "global_object_output" {
    value = [ for item in data.citrixadc_nitro_info.global_bindings.nitro_list: item.object ]
}

locals {
  csbinds = [ for item in data.citrixadc_nitro_info.csvs_bindings.nitro_list: item.object ]
  lbbinds = [ for item in data.citrixadc_nitro_info.lbvs_bindings.nitro_list: item.object ]
  globalbinds = [ for item in data.citrixadc_nitro_info.global_bindings.nitro_list: item.object ]
  allbinds = setunion(local.csbinds, local.lbbinds, local.globalbinds)
}

output "concat_object_output" {
  value = local.allbinds
}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name   = "tf_responder_policy"
  action = "NOOP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"

  globalbinding {
    invoke                 = true
    labeltype              = "vserver"
    labelname              = citrixadc_lbvserver.test_lbvserver.name
    type                   = "REQ_OVERRIDE"
    gotopriorityexpression = "END"
    priority               = 600
  }
}

resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}


resource "citrixadc_csvserver_responderpolicy_binding" "tf_bind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_responderpolicy.tf_responder_policy.name
    priority = 100
    bindpoint = "REQUEST"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.44"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver_responderpolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_responderpolicy.tf_responder_policy.name
    priority = 120
    bindpoint = "REQUEST"
}

resource "citrixadc_lbvserver" "tf_lbvserver2" {
  ipv46       = "10.10.10.55"
  name        = "tf_lbvserver2"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver_responderpolicy_binding" "tf_bind2" {
    name = citrixadc_lbvserver.tf_lbvserver2.name
    policyname = citrixadc_responderpolicy.tf_responder_policy.name
    priority = 110
    bindpoint = "REQUEST"
}

resource "citrixadc_lbvserver" "test_lbvserver" {

  ipv46       = "10.10.10.66"
  name        = "test_lbvserver"
  port        = 80
  servicetype = "HTTP"

}
