---
subcategory: "NS"
---

# Resource: nsfeature

The nsfeature resource is used to enable or disable ADC features.


## Example usage

```hcl
resource "citrixadc_nsfeature" "tf_nsfeature" {
    cs = true
    lb = true
    ssl = false
    appfw = false
}
```


## Argument Reference

The following arguments can set to `true` or `false` to enable or disable the corresponding feature.

* wl - (Optional)
* sp - (Optional)
* lb - (Optional)
* cs - (Optional)
* cr - (Optional)
* sc - (Optional)
* cmp - (Optional)
* pq - (Optional)
* ssl - (Optional)
* gslb - (Optional)
* hdosp - (Optional)
* cf - (Optional)
* ic - (Optional)
* sslvpn - (Optional)
* aaa - (Optional)
* ospf - (Optional)
* rip - (Optional)
* bgp - (Optional)
* rewrite - (Optional)
* ipv6pt - (Optional)
* appfw - (Optional)
* responder - (Optional)
* htmlinjection - (Optional)
* push - (Optional)
* appflow - (Optional)
* cloudbridge - (Optional)
* isis - (Optional)
* ch - (Optional)
* appqoe - (Optional)
* contentaccelerator - (Optional)
* rise - (Optional)
* feo - (Optional)
* lsn - (Optional)
* rdpproxy - (Optional)
* rep - (Optional)
* urlfiltering - (Optional)
* videooptimization - (Optional)
* forwardproxy - (Optional)
* sslinterception - (Optional)
* adaptivetcp - (Optional)
* cqa - (Optional)
* ci - (Optional)
* bot - (Optional)


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsfeature resource. It is a random string prefixed with "tf-nsfeature-"
