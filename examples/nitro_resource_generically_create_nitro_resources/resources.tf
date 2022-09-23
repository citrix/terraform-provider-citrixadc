resource "citrixadc_service" "tf_service" {

    name = "tf_service"
    ip = "192.168.43.33"
    servicetype  = "HTTP"
    port = 80
}

resource "citrixadc_nitro_resource" "tf_lbvserver" {
    workflows_file = "workflows.yaml"
    workflow = "lbvserver"

    # The following attributes changing will trigger the update function
    attributes = {
      ipv46       = "10.10.10.33"
    }

    # The following attributes changinge will trigger the delete and re create of the resource
    non_updateable_attributes = {
      name        = "tf_lbvserver"
      servicetype = "HTTP"
      port        = 80
    }
}

resource "citrixadc_nitro_resource" "tf_lbvserver_service_bind" {
    workflows_file = "workflows.yaml"
    workflow = "lbvserver_service_binding"

    # Bindings do not support update operation
    # Hence all attributes should be defined in the non_updateable_attributes map
    non_updateable_attributes = {
        name = citrixadc_nitro_resource.tf_lbvserver.non_updateable_attributes.name
        servicename = citrixadc_service.tf_service.name
        weight = 2
    }
}

# An unrelated resource to show case how non updateable object workflow
# should be used
resource "citrixadc_nitro_resource" "tf_patset" {
    workflows_file = "workflows.yaml"
    workflow = "policypatset"

    # Since update is not supported by the non_updateable_object workflow
    # all attributes should be defined in the non_updateable_attributes map

    non_updateable_attributes = {
      name = "tf_patset"
      comment = "Policy patset comment new"
    }
}
