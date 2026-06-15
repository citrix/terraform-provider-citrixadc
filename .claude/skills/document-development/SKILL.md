---
name: document-development
description: Write and update Terraform registry-style documentation for citrixadc resources under docs/
---

## Documentation Structure

```
docs/
  index.md                              — Provider overview and configuration
  resources/<resource_name>.md          — Resource documentation
  data-sources/<resource_name>.md       — Datasource documentation
```

## Doc File Layout

Every resource doc follows this structure:

```markdown
---
subcategory: "<Category>"
---

# Resource: <resource_name>

<One-line description of the resource.>


## Example usage

<HCL code blocks showing how to use the resource.>


## Argument Reference

<Bulleted list of all configurable attributes.>


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - <Description of ID format.>


## Import

<Import command example (omit for singletons).>
```

Datasource docs follow the same structure but use `# Data Source:` heading and `data "<type>" "<name>"` blocks.

## Subcategory Conventions

Determine the subcategory by looking up the resource's module in `resource_module_mapping.py` (`RESOURCE_MODULE_MAPPING` dict), then mapping the module name to the subcategory using this table:

| Module (from `resource_module_mapping.py`) | Subcategory |
|---|---|
| `aaa` | `"AAA"` |
| `authentication` | `"Authentication"` |
| `authorization` | `"Authorization"` |
| `lb` | `"Load Balancing"` |
| `ssl` | `"SSL"` |
| `cs` | `"Content Switching"` |
| `gslb` | `"GSLB"` |
| `network` | `"Network"` |
| `ns` | `"NS"` |
| `snmp` | `"SNMP"` |
| `ha` | `"HA"` |
| `cluster` | `"Cluster"` |
| `system` | `"System"` |
| `basic` | `"Basic"` |
| `dns` | `"DNS"` |
| `cache` | `"Cache"` |
| `cmp` | `"Compression"` |
| `cr` | `"Cache Redirection"` |
| `responder` | `"Responder"` |
| `rewrite` | `"Rewrite"` |
| `transform` | `"Transform"` |
| `appfw` | `"Application Firewall"` |
| `bot` | `"Bot"` |
| `vpn` | `"VPN"` |
| `tm` | `"Traffic Management"` |
| `audit` | `"Audit"` |
| `policy` | `"Policy"` |
| `spillover` | `"Spillover"` |
| `appflow` | `"AppFlow"` |
| `stream` | `"Stream"` |
| `feo` | `"Front End Optimization"` |
| `ica` | `"ICA"` |
| `rdp` | `"RDP"` |
| `db` | `"Database"` |
| `subscriber` | `"Subscriber"` |
| `contentinspection` | `"Content Inspection"` |
| `videooptimization` | `"Video Optimization"` |
| `appqoe` | `"AppQoE"` |
| `utility` | `"Utility"` |
| `ntp` | `"NTP"` |
| `quic` | `"QUIC"` |

**How to look up**: Find the resource name in `resource_module_mapping.py`, get its module value, then use the table above to get the subcategory string.

If the module is not in this table, fall back to checking existing docs: `grep -h "^subcategory:" docs/resources/<prefix>*.md`.

## Resource Categories and Their Doc Patterns

### Singleton Resources (e.g., lbparameter, aaacertparams)

- No `## Import` section (singletons always exist on ADC)
- ID description: `It is a unique string prefixed with "<resource_name>-config"`
- All attributes are Optional (no Required)

### Named Resources (e.g., aaagroup, sslcertkey)

- Include `## Import` section with example
- ID references the name attribute: `It has the same value as the \`<name_attr>\` attribute.`
- Import example: `terraform import citrixadc_<resource>.tf_<resource> <name_value>`

### Binding Resources (e.g., aaagroup_aaauser_binding)

- Include `## Import` section
- ID is concatenation of unique attrs: `It is the concatenation of \`attr1\` and \`attr2\` attributes separated by a comma.`
- Import example: `terraform import citrixadc_<resource>.tf_<resource> val1,val2`

## Argument Reference Formatting

Each argument is a bullet point with this format:

```markdown
* `<attr_name>` - (<Required|Optional>[, Sensitive][, WriteOnly]) <Description>. [Possible values: [ VAL1, VAL2 ]]
```

### Flag Annotations

| Flag | When to use |
|---|---|
| `Required` | Attribute has `is_required: true` in metadata |
| `Optional` | Attribute has `is_required: false` |
| `Sensitive` | Attribute has `is_sensitive: true` (secret value persisted in state) |
| `WriteOnly` | Attribute has `is_write_only: true` (value NOT persisted in state) |

### Type-Specific Formatting

- **Enums**: Append `Possible values: [ VAL1, VAL2 ]`
- **Integers with range**: Append `Minimum value = X Maximum value = Y`
- **Strings with length**: Append `Minimum length = X`
- **Defaults**: Mention default value when it exists (e.g., `Defaults to \`"ENABLED"\`.`)

## Documenting Secret/Ephemeral Attributes

When a resource has attributes marked `x-secret-attr` in metadata, the code generation expands them into three schema attributes. All three must be documented:

### 1. Original Sensitive Attribute (`<name>`)

```markdown
* `<name>` - (Optional, Sensitive) <Description from metadata>. The value is persisted in Terraform state (encrypted). See also `<name>_wo` for an ephemeral alternative.
```

### 2. Write-Only Attribute (`<name>_wo`)

```markdown
* `<name>_wo` - (Optional, Sensitive, WriteOnly) Same as `<name>`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `<name>_wo_version`. If both `<name>` and `<name>_wo` are set, `<name>_wo` takes precedence.
```

### 3. Version Tracker (`<name>_wo_version`)

```markdown
* `<name>_wo_version` - (Optional) An integer version tracker for `<name>_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
```

### Example Blocks for Secret Attributes

Always provide two HCL examples when a resource has secret attributes:

**Sensitive path** (backward-compatible, value persisted in state):

```markdown
### Using <name> (sensitive attribute - persisted in state)

\```hcl
variable "<resource>_<name>" {
  type      = string
  sensitive = true
}

resource "citrixadc_<resource>" "example" {
  <name> = var.<resource>_<name>
  // ... other required attrs
}
\```
```

**Write-only/ephemeral path** (value NOT persisted in state):

```markdown
### Using <name>_wo (write-only/ephemeral - NOT persisted in state)

The `<name>_wo` attribute provides an ephemeral path for <description>. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `<name>_wo_version`.

\```hcl
variable "<resource>_<name>" {
  type      = string
  sensitive = true
}

resource "citrixadc_<resource>" "example" {
  <name>_wo         = var.<resource>_<name>
  <name>_wo_version = 1
  // ... other required attrs
}
\```

To rotate the secret, update the variable value and bump the version:

\```hcl
resource "citrixadc_<resource>" "example" {
  <name>_wo         = var.<resource>_<name>
  <name>_wo_version = 2  # Bumped to trigger update
  // ... other required attrs
}
\```
```

## Datasource Documentation

Datasource docs follow the same structure but simpler:

```markdown
---
subcategory: "<Category>"
---

# Data Source: <resource_name>

The <resource_name> data source allows you to retrieve information about <description>.


## Example usage

\```terraform
data "citrixadc_<resource>" "example" {
  <lookup_attr> = "<value>"
}

output "<attr>" {
  value = data.citrixadc_<resource>.example.<attr>
}
\```


## Argument Reference

* `<lookup_attr>` - (Required) <Description>.

## Attribute Reference

In addition to the arguments, the following attributes are available:

<List of all read-only attributes with descriptions.>
```

For **singleton datasources** (e.g., lbparameter), no lookup attribute is needed — all attributes are computed.

## Finding Information for Documentation

### From Metadata

The primary source of truth for attribute names, types, descriptions, and flags is the metadata file:

```
tfdata/<resource_name>.json
```

Key fields to extract:
- `option_name` — attribute name
- `description` — description text (join array elements)
- `type` — data type (string, integer, boolean, number)
- `is_required` — Required vs Optional
- `x-secret-attr` — needs ephemeral documentation
- `x-unique-attr` — part of resource identity (for Import section)
- `is_get_id` + `is_delete_id` — parent/name attribute
- `default` — default value to mention

### From Resource Schema

If metadata is unavailable, read the generated schema file:

```
citrixadc_framework/<resource>/resource_schema.go
```

Look for `schema.Schema` block — each attribute has `Required`, `Optional`, `Computed`, `Sensitive`, `WriteOnly`, `Default`, and `Description` fields.

### From SDK v2 Resources

For legacy resources not yet in `citrixadc_framework/`, read:

```
citrixadc/resource_citrixadc_<resource>.go
```

The schema is defined in the `resourceCitrixAdc<Name>()` function.

### Subcategory Lookup

Look up the resource in `resource_module_mapping.py` to get its module, then use the module-to-subcategory mapping table in the "Subcategory Conventions" section above.

Fallback — check existing docs for the same resource family:

```bash
grep -h "^subcategory:" docs/resources/<prefix>*.md | sort -u
```

## Writing Guidelines

1. **Keep descriptions concise** — Use the NITRO API description from metadata but trim excessive detail. One or two sentences is usually enough for the Argument Reference.
2. **Use consistent terminology** — "Citrix ADC" (not "NetScaler" or "ADC" alone), "NITRO API", "Terraform state".
3. **Show realistic examples** — Use meaningful resource names (e.g., `my_group`, `servercert1`) not generic placeholders. Show dependent resources when needed (e.g., a binding example should reference its parent).
4. **Sensitive values in variables** — Always use `variable` blocks with `sensitive = true` for secret attributes in examples. Never hardcode secrets.
5. **Match existing style** — Read 2-3 existing docs in the same subcategory before writing a new one to match conventions.
6. **No emojis** — Keep documentation professional and plain-text friendly.
7. **HCL formatting** — Use 2-space indentation in HCL blocks. Align `=` signs within a resource block for readability.

## Reference Examples

- **Singleton**: `docs/resources/lbparameter.md` — no Import section, static ID, ephemeral attribute documentation
- **Singleton (simple)**: `docs/resources/aaacertparams.md` — minimal singleton doc
- **Named resource**: `docs/resources/aaagroup.md` — Import section, name-based ID
- **Binding**: `docs/resources/aaagroup_aaauser_binding.md` — comma-separated Import, concatenated ID
- **Datasource**: `docs/data-sources/sslcertkey.md` — lookup by name, output examples
- **Resource with secret (legacy style)**: `docs/resources/aaakcdaccount.md` — secret attr without ephemeral docs (needs updating)
