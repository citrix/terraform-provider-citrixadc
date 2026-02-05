---
subcategory: "Rewrite"
---

# Data Source: citrixadc_rewriteaction

The `citrixadc_rewriteaction` data source is used to retrieve information about an existing rewrite action configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_rewriteaction" "example" {
  name = "my_rewriteaction"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the rewrite action to retrieve.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the rewrite action (same as name).
* `comment` - Comment. Can be used to preserve information about this rewrite action.
* `newname` - New name for the rewrite action (if renamed).
* `refinesearch` - Specify additional criteria to refine the results of the search.
* `search` - Search facility that is used to match multiple strings in the request or response. Used in the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types.
* `stringbuilderexpr` - Expression that specifies the content to insert into the request or response at the specified location, or that replaces the specified string.
* `target` - Default syntax expression that specifies which part of the request or response to rewrite.
* `type` - Type of user-defined rewrite action. Possible values include:
  * `noop` - No operation.
  * `delete` - Delete the specified target.
  * `insert_http_header` - Insert an HTTP header.
  * `delete_http_header` - Delete an HTTP header.
  * `corrupt_http_header` - Corrupt an HTTP header.
  * `insert_before` - Insert before the specified target.
  * `insert_after` - Insert after the specified target.
  * `replace` - Replace the specified target.
  * `replace_http_res` - Replace HTTP response.
  * `delete_all` - Delete all occurrences.
  * `replace_all` - Replace all occurrences.
  * `insert_before_all` - Insert before all occurrences.
  * `insert_after_all` - Insert after all occurrences.
  And many more action types for various protocols and use cases.
