---
subcategory: "Policy"
---

# Data Source: policytracing

The policytracing data source retrieves captured policy-tracing records from the Citrix ADC, equivalent to the `show policy tracing` operation. Policy tracing records the policy-evaluation history for transactions; this data source lets you read those records and optionally narrow them with filter inputs such as transaction ID, detail level, and cluster node.


## Example usage

```terraform
data "citrixadc_policytracing" "tf_policytracing" {
  detail = "all"
}

output "policytracing_protocoltype" {
  value = data.citrixadc_policytracing.tf_policytracing.protocoltype
}

output "policytracing_filterexpr" {
  value = data.citrixadc_policytracing.tf_policytracing.filterexpr
}
```


## Argument Reference

The following arguments are optional filters used to select the policy-tracing record to return. The first record matching every supplied filter is returned.

* `transactionid` - (Optional) Unique ID to identify the current transaction.
* `detail` - (Optional) Show detailed information of the captured records. Defaults to `"all"`. Possible values: [ brief, all ]
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `capturesslhandshakepolicies` - (Optional) Set it to yes if you need to capture the SSL handshake policies.
* `filterexpr` - (Optional) Policy tracing filter expression. For example: `http.req.url.startswith("/this")`.
* `protocoltype` - (Optional) Protocol type for which policy records need to be collected. Defaults to `"HTTP"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policytracing data source. It is a synthetic constant string `"policytracing-query"`.
* `transactionid` - Unique ID to identify the current transaction.
* `detail` - Detail level of the captured records.
* `nodeid` - Unique number that identifies the cluster node.
* `capturesslhandshakepolicies` - Indicates whether the SSL handshake policies are captured.
* `filterexpr` - Policy tracing filter expression.
* `protocoltype` - Protocol type for which policy records are collected.
