---
subcategory: "Generic"
---

# Resource: nitro\_resource

The nitro\_resource resource is used to generically create NITRO resources.

It will handle the creation, update and deletion of the resource in the manner
described by the workflow map entry referenced in the configuration block.

The execution is governed by the data found in the corresponding workflow entry.

Browse through the workflows.yaml file to see what endpoints are available.

The workflows.yaml file can be found [here](https://github.com/citrix/terraform-provider-citrixadc/blob/master/citrixadc/testdata/workflows.yaml).

## Limitations

nitro\_resource does come with some limitations.

The comparison of each attribute value is dependent on the string representation of the specific attribute.

As such complex attribute values such as lists or maps cannot be reliably compared and hence are not supported.

Also if the value provided in the configuration has any difference with the string representation returned by the
NITRO API a false update will be issued on subsequent resource executions. To avoid this make sure the attribute
is written just as the string representation of the NITRO API retrieved value.


## Example usage

```hcl
resource "citrixadc_nitro_resource" "tf_lbvserver" {
    workflows_file = "workflows.yaml"
    workflow = "lbvserver"

    # The following attributes changing will trigger the update function
    attributes = {
      ipv46       = "10.10.10.33"
    }

    # The following attributes changing will trigger the delete and re create of the resource
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
        servicename = "service1"
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

```


## Argument Reference

* `workflows_file` - (Required) The path to the workflows yaml file that contains the workflow definitions.
* `workflow` - (Required) The map key which points to the specific entry to use from the workflows map.
* `attributes` - (Optional) The map containing the attributes of the resource.

    Any change to these attribute values will trigger an update operation.

* `non_updateable_attributes` - (Optional) The map containing the non updateable attributes of the resource.

    Any change to these attribute values will trigger delete and re create of the resource.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nitro\_resource. Depending on the workflow type this can be a single identifying attributes or a pair of
    attributes separated by a comma `,`.


## Import

A nitro\_resource can be imported using its id

```shell
terraform import citrixadc_nitro_resource.tf_object object_id

terraform import citrixadc_nitro_resource.tf_biding primary_id,secondary_id
```
