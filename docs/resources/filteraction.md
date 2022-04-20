---
subcategory: "Filter"
---

# Resource: filteraction

The filteraction resource is used to create Filter Action Resource.


## Example usage

```hcl
resource "citrixadc_filteraction" "tf_filteraction" {
  name  = "tf_filteraction"
  qual  = "corrupt"
  value = "X-Forwarded-For"
}
```


## Argument Reference

* `name` - (Required) Name for the filtering action. Must begin with a letter, number, or the underscore character (_). Other characters allowed, after the first character, are the hyphen (-), period (.) hash (#), space ( ), at sign (@), equals (=), and colon (:) characters. Choose a name that helps identify the type of action. The name of a filter action cannot be changed after it is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `qual` - (Required) Qualifier, which is the action to be performed. The qualifier cannot be changed after it is set. The available options function as follows: ADD - Adds the specified HTTP header. RESET - Terminates the connection, sending the appropriate termination notice to the user's browser. FORWARD - Redirects the request to the designated service. You must specify either a service name or a page, but not both. DROP - Silently deletes the request, without sending a response to the user's browser.  CORRUPT - Modifies the designated HTTP header to prevent it from performing the function it was intended to perform, then sends the request/response to the server/browser. ERRORCODE. Returns the designated HTTP error code to the user's browser (for example, 404, the standard HTTP code for a non-existent Web page).
* `page` - (Optional) HTML page to return for HTTP requests (For use with the ERRORCODE qualifier).
* `respcode` - (Optional) Response code to be returned for HTTP requests (for use with the ERRORCODE qualifier).
* `servicename` - (Optional) Service to which to forward HTTP requests. Required if the qualifier is FORWARD.
* `value` - (Optional) String containing the header_name and header_value. If the qualifier is ADD, specify <header_name>:<header_value>. If the qualifier is CORRUPT, specify only the header_name


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the filteraction. It has the same value as the `name` attribute.


## Import

A filteraction can be imported using its name, e.g.

```shell
terraform import citrixadc_filteraction.tf_filteraction tf_filteraction
```
