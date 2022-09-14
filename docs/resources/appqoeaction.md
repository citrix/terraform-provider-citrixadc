---
subcategory: "Appqoe"
---

# Resource: appqoeaction

The appqoeaction resource is used to create appqoeaction.


## Example usage

```hcl
resource "citrixadc_appqoeaction" "tf_appqoeaction" {
  name        = "my_appqoeaction"
  priority    = "LOW"
  respondwith = "NS"
  delay       = 40
}
```


## Argument Reference

* `name` - (Required) Name for the AppQoE action. Must begin with a letter, number, or the underscore symbol (_). Other characters allowed, after the first character, are the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), and colon (:) characters. This is a mandatory argument
* `altcontentpath` - (Optional) Path to the alternative content service to be used in the ACS
* `altcontentsvcname` - (Optional) Name of the alternative content service to be used in the ACS
* `customfile` - (Optional) name of the HTML page object to use as the response
* `delay` - (Optional) Delay threshold, in microseconds, for requests that match the policy's rule. If the delay statistics gathered for the matching request exceed the specified delay, configured action triggered for that request, if there is no action then requests are dropped to the lowest priority level
* `dosaction` - (Optional) DoS Action to take when vserver will be considered under DoS attack and corresponding rule matches. Mandatory if AppQoE actions are to be used for DoS attack prevention.
* `dostrigexpression` - (Optional) Optional expression to add second level check to trigger DoS actions. Specifically used for Analytics based DoS response generation
* `maxconn` - (Optional) Maximum number of concurrent connections that can be open for requests that matches with rule.
* `numretries` - (Optional) Retry count
* `polqdepth` - (Optional) Policy queue depth threshold value. When the policy queue size (number of requests queued for the policy binding this action is attached to) increases to the specified polqDepth value, subsequent requests are dropped to the lowest priority level.
* `priority` - (Optional) Priority for queuing the request. If server resources are not available for a request that matches the configured rule, this option specifies a priority for queuing the request until the server resources are available again. If priority is not configured then Lowest priority will be used to queue the request.
* `priqdepth` - (Optional) Queue depth threshold value per priorirty level. If the queue size (number of requests in the queue of that particular priorirty) on the virtual server to which this policy is bound, increases to the specified qDepth value, subsequent requests are dropped to the lowest priority level.
* `respondwith` - (Optional) Responder action to be taken when the threshold is reached. Available settings function as follows:             ACS - Serve content from an alternative content service                   Threshold : maxConn or delay             NS - Serve from the Citrix ADC (built-in response)                  Threshold : maxConn or delay
* `retryonreset` - (Optional) Retry on TCP Reset
* `retryontimeout` - (Optional) Retry on request Timeout(in millisec) upon sending request to backend servers
* `tcpprofile` - (Optional) Bind TCP Profile based on L2/L3/L7 parameters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appqoeaction. It has the same value as the `name` attribute.


## Import

An appqoeaction can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
