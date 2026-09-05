---
subcategory: "DNS"
---

# Data Source: dnspolicylabel

This data source is used to retrieve information about an existing DNS policy label.

## Example Usage

```hcl
data "citrixadc_dnspolicylabel" "example" {
  labelname = "blue_label"
}
```

## Argument Reference

* `labelname` - (Required) Name of the DNS policy label.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the DNS policy label (same as `labelname`).
* `newname` - The new name of the DNS policy label.
* `transform` - The type of transformations allowed by the policies bound to the label.

### Read-only dnspolicylabel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_dnspolicylabel` resource). They are GET-only/Computed. Any attribute the appliance does not return is `null`.

* `numpol` - Number of polices bound to label.
* `hits` - Number of times policy label was invoked.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label invocation.
* `invoke_labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
* `flowtype` - Flowtype of the bound dns policy.
* `description` - Description of the policylabel.
* `isdefault` - A value of true is returned if it is a default dns policylabel.
