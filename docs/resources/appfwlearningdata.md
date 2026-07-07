---
subcategory: "Application Firewall"
---

# Resource: appfwlearningdata

The appfwlearningdata resource manages the Citrix ADC Application-Firewall **learned-data** table.

~> **BEST-EFFORT MODEL** `appfwlearningdata` is **not** a normal CRUD object. NITRO exposes only `get(all)`, `count`, `delete`, and the `reset`/`export` actions for it — there is no `add`/`set` endpoint and no per-object identity to reconcile. This resource is therefore modeled as an **action resource**: applying it performs the NITRO **`reset`** action, which **clears** the App-Firewall learned data for the given profile / security check. `Read` and `Update` are no-ops, and `Delete` is a **state-only removal** (the NITRO delete endpoint carries no key selector in the metadata, so no delete request is issued on destroy). The reset/delete semantics here should be **verified on a live ADC** before relying on them.

~> **WARNING** Applying this resource **clears (resets) App-Firewall learned data**. This is a disruptive, non-reversible side effect. Use it deliberately, and note that all attributes are `RequiresReplace` — changing any input re-runs the reset action.


## Example usage

```hcl
resource "citrixadc_appfwlearningdata" "tf_appfwlearningdata" {
  profilename   = "my_appfwprofile"
  securitycheck = "startURL"
  starturl      = "^https?://[^/]+/$"
}
```


## Argument Reference

All arguments are optional (they are the inputs to the `reset`/`export` actions) and are `RequiresReplace`.

* `profilename` - Name of the profile.
* `securitycheck` - Name of the security check. Possible values: `startURL`, `cookieConsistency`, `fieldConsistency`, `crossSiteScripting`, `SQLInjection`, `fieldFormat`, `CSRFtag`, `XMLDoSCheck`, `XMLWSICheck`, `XMLAttachmentCheck`, `TotalXMLRequests`, `creditCardNumber`, `ContentType`.
* `starturl` - Start URL configuration.
* `cookieconsistency` - Cookie name.
* `fieldconsistency` - Form field name.
* `formactionurl_ffc` - Form action URL.
* `contenttype` - Content Type name.
* `crosssitescripting` - Cross-site scripting.
* `formactionurl_xss` - Form action URL.
* `as_scan_location_xss` - Location of cross-site scripting exception. Possible values: `FORMFIELD`, `HEADER`, `COOKIE`, `URL`.
* `as_value_type_xss` - XSS value type. Possible values: `Tag`, `Attribute`, `Pattern`.
* `as_value_expr_xss` - XSS value expressions for Tag, Attribute or Pattern.
* `sqlinjection` - Form field name.
* `formactionurl_sql` - Form action URL.
* `as_scan_location_sql` - Location of SQL injection exception. Possible values: `FORMFIELD`, `HEADER`, `COOKIE`.
* `as_value_type_sql` - SQL value type. Possible values: `Keyword`, `SpecialString`, `Wildchar`.
* `as_value_expr_sql` - SQL value expressions for Keyword, SpecialString or Wildchar.
* `fieldformat` - Field format name.
* `formactionurl_ff` - Form action URL.
* `csrftag` - CSRF Form Action URL.
* `csrfformoriginurl` - CSRF Form Origin URL.
* `creditcardnumber` - The object expression to be excluded from the safe commerce check.
* `creditcardnumberurl` - The URL for which the credit card numbers are bypassed from inspection.
* `xmldoscheck` - XML Denial of Service check.
* `xmlwsicheck` - Web Services Interoperability Rule ID.
* `xmlattachmentcheck` - XML Attachment Content-Type.
* `totalxmlrequests` - (Boolean) Total XML requests.
* `target` - Target filename for data to be exported.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwlearningdata resource. It is a synthetic value (`appfwlearningdata-config`), since the object exposes no per-object readable identity.
