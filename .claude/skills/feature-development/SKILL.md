---
name: feature-development
description: Add new resources to the Plugin Framework provider and debug existing resource implementations in citrixadc_framework/
---

## Architecture Overview

The provider uses **terraform-plugin-mux** to serve two implementations:
1. **`citrixadc/`** — Legacy SDK v2 provider (~500+ resources)
2. **`citrixadc_framework/`** — New Plugin Framework provider (actively growing)

Each Framework resource lives in its own subdirectory with 4 files:
- `resource_<name>.go` — CRUD logic
- `resource_schema.go` — Schema definition, model struct, payload/state helpers
- `datasource_<name>.go` — Datasource read logic
- `datasource_schema.go` — Datasource schema definition

## Code Generation System

Resources are generated using Jinja2 templates + Python. **Do not hand-write resource files** — use the generation pipeline.

### Key Files

| File | Purpose |
|---|---|
| `generate_all.py` | Main generator script — processes metadata, renders templates, updates provider.go |
| `jinja_templates/resource.go.j2` | Template for resource CRUD implementation |
| `jinja_templates/resource_schema.go.j2` | Template for resource schema, model struct, payload/state functions |
| `jinja_templates/datasource.go.j2` | Template for datasource implementation |
| `jinja_templates/datasource_schema.go.j2` | Template for datasource schema |
| `resources.txt` | List of resource names to generate (one per line) |
| `tfdata/<resource_name>.json` | Metadata for each resource (attributes, types, flags) |
| `resource_id_mapping.json` | Legacy SDK v2 ID attribute order for backward-compatible parsing |
| `resource_module_mapping.py` | Maps resource names to NITRO Go module package names |

### Running the Generator

```bash
# Generate all resources listed in resources.txt
python3 generate_all.py
```

The generator:
1. Reads each resource name from `resources.txt`
2. Loads metadata from `tfdata/<resource_name>.json`
3. Renders 4 Go files per resource into `citrixadc_framework/<resource_name>/`
4. Cleans up unused `utils` imports
5. Updates `vendor/.../service/resources.go` with new resource enum entries
6. Updates `citrixadc_framework/provider/provider.go` with import/registration entries

### Post-Generation Steps

After running the generator:
```bash
# Format generated code
make fmt

# Build and verify compilation
make build

# Install provider locally for testing
make install
```

## Metadata Format (`tfdata/<resource_name>.json`)

Each metadata file is a JSON array of attribute objects:

```json
[
    {
        "option_name": "attributename",
        "type": "string",
        "description": ["Human-readable description of the attribute."],
        "is_get_id": false,
        "is_delete_id": false,
        "is_updateable": true,
        "is_required": false,
        "default": "ENABLED",
        "x-unique-attr": false,
        "x-secret-attr": false
    }
]
```

### Attribute Fields

| Field | Type | Purpose |
|---|---|---|
| `option_name` | string | NITRO API attribute name (snake_case) |
| `type` | string | `"string"`, `"integer"`, `"boolean"`, `"number"`, `"string[]"` |
| `description` | string[] | Description lines (joined for Go string literal) |
| `is_get_id` | bool | Used as the resource identifier for GET/READ API calls |
| `is_delete_id` | bool | Used as the parent identifier for DELETE API calls |
| `is_updateable` | bool | `false` = RequiresReplace plan modifier (forces recreation) |
| `is_required` | bool | `true` = Required in schema; `false` = Optional+Computed |
| `default` | any | Default value for the attribute (generates Default: in schema) |
| `x-unique-attr` | bool | Attribute is part of the resource's unique identity |
| `x-secret-attr` | bool | Attribute is a secret — triggers ephemeral triple expansion |

### How Attributes Drive Resource Classification

The generator classifies resources into 4 categories based on attribute flags:

```
parent_attrs     = attrs where is_get_id=true AND is_delete_id=true
delete_arg_attrs = attrs where x-unique-attr=true AND is_delete_id=false
```

| Category | parent_attrs | delete_arg_attrs | Create API | Delete API |
|---|---|---|---|---|
| **singleton** | none | none | `UpdateUnnamedResource` | No-op (state removal only) |
| **named_resource** | present | none | `AddResource` | `DeleteResource` |
| **global_binding** | none | present | `UpdateUnnamedResource` | `DeleteResourceWithArgs/Map` |
| **binding_with_parent** | present | present | `UpdateUnnamedResource` | `DeleteResourceWithArgsMap` |

### How Attributes Drive ID Generation

| Pattern | Condition | ID Format |
|---|---|---|
| `static` | No `x-unique-attr` attrs | `"<resource_name>-config"` |
| `single_unique` | One `x-unique-attr` attr | Plain value (e.g., `"my-vserver"`) |
| `multiple_unique` | Multiple `x-unique-attr` attrs | `"attr1:urlEncode(val1),attr2:urlEncode(val2)"` |

### How Attributes Drive Read Pattern

| Pattern | Condition | API Call |
|---|---|---|
| `simple_find` | No `is_get_id`, no `x-unique-attr` | `FindResource(type, "")` |
| `find_with_id` | One attr with both `is_get_id` and `x-unique-attr` | `FindResource(type, id)` |
| `array_filter_no_id` | No `is_get_id`, has `x-unique-attr` | `FindResourceArrayWithParams` + filter |
| `array_filter_with_id` | Has `is_get_id` and multiple `x-unique-attr` | `FindResourceArrayWithParams(name=id)` + filter |
| `no_get` | NITRO doc exposes no `get`/`get (all)` verb | No API call — Read preserves existing state |

When a resource has no GET endpoint, Read must be a no-op that re-stores the prior state unchanged; drift detection is impossible by definition. Datasource generation should be suppressed entirely for these resources (a datasource with no query is meaningless). Detect this from the NITRO doc Operations section, not from metadata flags.

## Secret Attribute Expansion (`x-secret-attr`)

When an attribute has `"x-secret-attr": true`, the generator expands it into 3 schema attributes:

| Generated Attribute | Schema Flags | Purpose |
|---|---|---|
| `<name>` | Optional, Sensitive | Backward-compatible legacy path (persisted in state) |
| `<name>_wo` | Optional, Sensitive, WriteOnly | Ephemeral path (NOT persisted in state) |
| `<name>_wo_version` | Optional, Computed, Default(1) | Version tracker to trigger updates |

Key behaviors:
- In `GetPayloadFromthePlan`: `_wo` and `_wo_version` are **skipped**
- In `GetPayloadFromtheConfig`: `_wo` value is read from config and mapped to the original NITRO field
- In `SetAttrFromGet`: All 3 secret attrs are **skipped** (NITRO never returns secrets)
- In `Update`: Change detection checks `<name>` OR `<name>_wo_version`
- If both `<name>` and `<name>_wo` are set, `_wo` takes precedence (applied last)

## Legacy ID Backward Compatibility

The `resource_id_mapping.json` file maps resource names to their SDK v2 comma-separated ID attribute order:

```json
{
    "aaagroup_aaauser_binding": "groupname,username",
    "lbvserver_service_binding": "name,servicename,weight?"
}
```

- `?` suffix marks optional trailing attributes
- The generator passes this to `ParseIdString()` so imported SDK v2 state IDs can be parsed correctly
- New IDs use `key:urlEncode(value)` format; `ParseIdString` handles both formats transparently

## Adding a New Resource — Step by Step

### 1. Create the Metadata File

Create `tfdata/<resource_name>.json` with the attribute array. Source attribute information from the NITRO API documentation or existing SDK v2 resource file (`citrixadc/resource_citrixadc_<name>.go`).

Pay attention to:
- Set `is_get_id: true` + `is_delete_id: true` on the parent/name attribute
- Set `x-unique-attr: true` on all identity attributes
- Set `x-secret-attr: true` on password/passphrase attributes
- Set `is_updateable: false` on attributes that require resource recreation
- Set `is_required: true` on mandatory attributes

### 2. Add to resources.txt

Add the resource name (one per line) to `resources.txt`.

### 3. Add Legacy ID Mapping (if migrating from SDK v2)

If this resource existed in the SDK v2 provider, add its ID format to `resource_id_mapping.json`:
```json
{
    "<resource_name>": "attr1,attr2,optional_attr?"
}
```

### 4. Add Module Mapping

Ensure the resource has an entry in `resource_module_mapping.py`:
```python
RESOURCE_MODULE_MAPPING = {
    "<resource_name>": "<nitro_module>",  # e.g., "lb", "aaa", "ssl"
}
```

### 5. Run the Generator

```bash
python3 generate_all.py
```

### 6. Verify and Build

```bash
make fmt
make build
```

### 7. Register in Provider (if not auto-registered)

The generator auto-updates `citrixadc_framework/provider/provider.go`, but verify:
- Import line exists for the new package
- Resource is listed in `Resources()` function
- Datasource is listed in `DataSources()` function

### 8. Remove from SDK v2 Provider (if migrating)

If migrating from SDK v2, remove the old registration from `citrixadc/provider.go` `providerResources()` to avoid duplicate type name conflicts from the muxer.

## Debugging Existing Provider Implementations

### Debugging Workflow

When investigating issues in an existing Framework resource:

1. **Identify the resource category** — Read the resource's metadata (`tfdata/<name>.json`) to determine if it's a singleton, named_resource, global_binding, or binding_with_parent. This dictates which CRUD pattern and API calls are used.

2. **Read the generated source** — Check all 4 files in `citrixadc_framework/<name>/`:
   - `resource_schema.go` — Model struct, payload builders, state setters
   - `resource_<name>.go` — CRUD logic, read helper
   - `datasource_schema.go` — Datasource schema
   - `datasource_<name>.go` — Datasource read logic

3. **Compare with template output** — If the resource was generated, regenerate it and diff against the current file to see if manual edits have diverged from the template.

4. **Check the NITRO struct** — Verify the Go struct exists and has the expected fields in `vendor/github.com/citrix/adc-nitro-go/resource/config/<module>/`.

### Common Bug Categories

#### Schema Mismatches

**Symptom**: Terraform plan shows unexpected diffs, attributes reset to defaults, or "inconsistent result" errors after apply.

**Diagnosis**:
- Check if `is_computed` is set correctly — attributes without defaults that are Optional-only (not Computed) will show diffs if the API returns a value not in config.
- Check if `default` value in metadata matches what the ADC actually returns. Mismatched defaults cause perpetual diffs.
- For secret attrs: verify `SetAttrFromGet` skips them (NITRO never returns secrets). If not skipped, Terraform sees a change from the config value to empty string.

**Fix**: Update `tfdata/<name>.json` metadata and regenerate, or fix the template if the pattern is wrong.

#### ID Parsing Failures

**Symptom**: Import fails, destroy check fails, or "unable to parse ID" errors.

**Diagnosis**:
- Check the ID format: singleton uses static `"<name>-config"`, single_unique uses plain value, multiple_unique uses `key:urlEncode(value)` pairs.
- For resources migrated from SDK v2: verify `resource_id_mapping.json` has the correct legacy attribute order. `ParseIdString()` needs this to handle old `"val1,val2"` format IDs.
- Check if optional attrs (marked with `?` in ID mapping) are handled — `ParseIdString` returns `optionalAbsent` map for these.

**Fix**: Update `resource_id_mapping.json` and regenerate, or fix `ParseIdString` call arguments in the template.

#### NITRO API Call Failures

**Symptom**: "Unable to create/update/delete <resource>" errors during apply.

**Diagnosis**:
- **Wrong API method**: Check if the resource category matches the API call (e.g., singleton should use `UpdateUnnamedResource`, not `AddResource`).
- **Wrong service type**: Check `resource_module_mapping.py` and `generate_all.py` `get_nitro_service_type()` — the enum in `vendor/.../service/resources.go` must match.
- **Missing NITRO struct fields**: The Go struct in `vendor/.../resource/config/<module>/<name>.go` must have all fields referenced in the payload builder.
- **Delete args wrong**: For bindings, verify `delete_arg_attrs` in the template correctly identifies which attributes are passed to `DeleteResourceWithArgsMap`.

**Fix**: Depends on root cause — may need metadata fix, template fix, or vendor struct update.

#### Change Detection Issues

**Symptom**: Updates not triggered when attributes change, or unnecessary updates on every plan.

**Diagnosis**:
- In `Update()`, each attribute is checked with `!data.<Attr>.Equal(state.<Attr>)`. If an attribute is missing from this check, changes are silently ignored.
- For secret attrs: change detection checks the `<name>` value OR `<name>_wo_version`. If the user sets `_wo` without bumping `_wo_version`, no update is triggered (by design).
- Attributes with `is_updateable: false` get `RequiresReplace` plan modifiers — they are not checked in `Update()` because Terraform forces recreation instead.

**Fix**: Verify the attribute is not excluded from change detection in the template's Update block. Check `is_secret`, `is_write_only`, `is_version_tracker`, and `needs_plan_modifier` flags.

#### Datasource Shared Model Issues

**Symptom**: Datasource compilation errors or unexpected schema attributes.

**Diagnosis**:
- Datasource shares the `ResourceModel` struct with the resource. This means `_wo` and `_wo_version` fields must exist in the datasource schema even though they're semantically meaningless for reads.
- The datasource schema template (`datasource_schema.go.j2`) does NOT set `Sensitive`, `WriteOnly`, or `Default` — these are resource-only concerns.

**Fix**: Ensure the datasource schema includes all fields from the model struct, with appropriate Optional/Computed flags.

### Debugging Tools and Techniques

#### Enable Terraform Debug Logging

```bash
export TF_LOG=DEBUG
export TF_LOG_PROVIDER=DEBUG
terraform apply
```

This shows:
- `tflog.Debug/Trace` messages from resource CRUD operations
- NITRO client HTTP request/response details
- Attribute change detection decisions

#### Enable NITRO Client Logging

The NITRO client uses `hclog` with configurable log levels. Set via environment:
```bash
export TF_LOG=TRACE
```

This reveals:
- HTTP method, URL, and request body for every NITRO API call
- Response status codes and error payloads
- Resource find/filter operations

#### Inspect Generated vs. Expected Code

```bash
# Regenerate a single resource and compare
python3 generate_all.py  # generates all resources in resources.txt
diff citrixadc_framework/<resource>/resource_<resource>.go <backup>
```

#### Test with Acceptance Tests

```bash
# Run a specific resource's test
TF_ACC=1 go test ./citrixadc_framework/acctests/ -v -run TestAcc<Resource>_basic -timeout 120m
```

Requires environment variables:
```bash
export NS_URL=http://<ADC_IP>/
export NS_LOGIN=nsroot
export NS_PASSWORD='CADS123$%^'
```

#### Check NITRO API Directly

```bash
# Verify resource exists on ADC
curl -s -u nsroot:'CADS123$%^' http://<ADC_IP>/nitro/v1/config/<resource_type>/<resource_name>

# List all resources of a type
curl -s -u nsroot:'CADS123$%^' http://<ADC_IP>/nitro/v1/config/<resource_type>
```

### Error Flow in Framework Resources

All errors flow through Terraform's diagnostics system:

```
NITRO Client Error
  → resource CRUD function catches error
  → resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to <op> <resource>, got error: %s", err))
  → Terraform reports to user
```

Key error categories from `resp.Diagnostics.AddError`:
- `"Client Error"` — NITRO API call failed (network, auth, resource not found)
- `"Parse Error"` — ID string parsing failed (malformed ID, missing attributes)
- `"Configuration Error"` — Provider configuration invalid (missing credentials)

### Provider Configuration Debugging

The provider validates configuration in `Configure()` (`citrixadc_framework/provider/provider.go`):

| Parameter | Environment Variable | Required |
|---|---|---|
| `endpoint` | `NS_URL` | Yes |
| `username` | `NS_LOGIN` | Yes (default: `nsroot`) |
| `password` | `NS_PASSWORD` | Yes (default: `nsroot`) |
| `proxied_ns` | `_MPS_API_PROXY_MANAGED_INSTANCE_IP` | No |
| `partition` | `NS_PARTITION` | No |

If provider data is nil in `Configure()` of a resource/datasource, it means the provider itself failed to configure — check provider block configuration and environment variables.

## Template Variables Reference

### resource.go.j2 Variables

| Variable | Source | Example |
|---|---|---|
| `package_name` | resource name | `lbparameter` |
| `resource_struct_name` | PascalCase + "Resource" | `LbparameterResource` |
| `model_name` | PascalCase + "ResourceModel" | `LbparameterResourceModel` |
| `pascal_name` | PascalCase | `Lbparameter` |
| `nitro_service_type` | Service type mapping | `Lbparameter` |
| `resource_var_name` | lowercase resource name | `lbparameter` |
| `get_payload_func` | Payload function name | `lbparameterGetThePayloadFromthePlan` |
| `set_attr_func` | State setter function name | `lbparameterSetAttrFromGet` |
| `attributes` | Processed attribute list | (see metadata) |
| `read_pattern` | Read pattern analysis | `{pattern: 'simple_find', ...}` |
| `id_generation` | ID generation pattern | `{pattern: 'static', ...}` |
| `legacy_id_attr_order` | Legacy ID attr names | `["groupname", "username"]` |
| `legacy_id_optional_attrs` | Optional legacy attrs | `["type"]` |

### resource_schema.go.j2 Variables

| Variable | Source | Example |
|---|---|---|
| `model_name` | PascalCase + "ResourceModel" | `LbparameterResourceModel` |
| `resource_type_name` | PascalCase + "Resource" | `LbparameterResource` |
| `nitro_info.import_path` | Module mapping | `"github.com/citrix/adc-nitro-go/resource/config/lb"` |
| `nitro_info.type_name` | Module + SentenceCase | `lb.Lbparameter` |
| `function_names` | Function name dict | `{get_payload_func: ..., set_attr_func: ...}` |
| `required_imports` | Import analysis | `{types: true, int64default: true, ...}` |
| `id_generation` | ID pattern analysis | `{pattern: 'static', unique_attrs: [...]}` |

## NITRO API Client Methods

| Method | Usage |
|---|---|
| `AddResource(type, name, &payload)` | Create named resources |
| `UpdateResource(type, name, &payload)` | Update named resources |
| `UpdateUnnamedResource(type, &payload)` | Create/update singletons and bindings |
| `DeleteResource(type, name)` | Delete named resources |
| `DeleteResourceWithArgsMap(type, name, argsMap)` | Delete bindings (conditional args) |
| `FindResource(type, name)` | Read single resource |
| `FindResourceArrayWithParams(params)` | Read array of resources (for bindings) |

## Utility Functions (`citrixadc_framework/utils/`)

| Function | Purpose |
|---|---|
| `ParseIdString(id, legacyOrder, legacyOptional)` | Parse both new and legacy ID formats |
| `UrlEncode(value)` / `UrlDecode(encoded)` | Encode/decode values for composite IDs |
| `ConvertToInt64(value)` | Convert interface{} (int, float64, string) to int64 |
| `IntPtr(i)` / `BoolPtr(b)` | Create pointers for optional NITRO struct fields |

## Common Issues and Solutions

### "Duplicate type name" from muxer
Remove the old SDK v2 resource registration from `citrixadc/provider.go` when migrating a resource to the Framework.

### Generated code doesn't compile
1. Check that the NITRO Go struct exists in `vendor/github.com/citrix/adc-nitro-go/resource/config/<module>/`
2. Verify `resource_module_mapping.py` has the correct module name
3. Run `make fmt` to fix formatting issues

### Resource enum missing in resources.go
The generator auto-updates `vendor/.../service/resources.go`. If it fails, manually add the enum entry matching the `SentenceCase_with_underscores` pattern (e.g., `Lbparameter`, `Aaagroup_aaauser_binding`).

### Secret attribute not expanding
Ensure the metadata has `"x-secret-attr": true` on the attribute. The expansion happens in `expand_secret_attributes()` in `generate_all.py`.

## Modifying Templates

When modifying Jinja2 templates, keep in mind:
- Templates use `{%- ... -%}` for whitespace control
- The generator passes processed attributes (after `expand_secret_attributes`), not raw metadata
- Test template changes by regenerating a known resource and diffing the output
- All 4 templates share the same `ResourceModel` struct — datasource schema must include `_wo` and `_wo_version` fields even though they are semantically meaningless for reads

## Rename support (`newname` attribute via NITRO `rename` action)

Some NITRO resources expose a `rename` action (`?action=rename`) plus a `newname`
attribute alongside the primary name/key attribute. The codegen treats `newname`
like any other attribute — typically Optional+Computed with a `RequiresReplace`
plan modifier — and emits a no-op Update. That is wrong: a `newname` change should
drive an **in-place rename**, not a destroy/recreate. This mirrors the SDK v2
convention (canonical SDK v2 example: `citrixadc/resource_citrixadc_appfwpolicy.go`,
which detects the change and calls `client.ActOnResource(type, &payload, "rename")`).

**Canonical Framework example:** `citrixadc_framework/lbpolicylabel`.

When to apply: the NITRO doc Operations section lists a `rename` verb AND the model
has a `newname` (or equivalent) attribute distinct from the primary key. Confirm the
rename payload shape from the doc's `rename` anchor — for lbpolicylabel it is POST
`?action=rename` with body `{"lbpolicylabel":{"labelname":<old>,"newname":<new>}}`
(both mandatory).

Recipe:

1. **Schema flags for `newname`:** make it `Optional` only. Remove `Computed`
   (it is a pure user input, never echoed by GET — Computed causes known-after-apply
   churn) and remove the `RequiresReplace` plan modifier (RequiresReplace would force
   recreation instead of letting the change reach Update). The primary key attribute
   (e.g. `labelname`) stays `Required` + `RequiresReplace` — changing the key itself
   is still a recreate, exactly like SDK v2 `ForceNew` on `name`.

2. **Exclude `newname` from the add/create payload.** It is rename-only; the add
   POST must not carry it. In `xxxGetThePayloadFromthePlan`, skip the `newname` field
   (leave a comment noting it is rename-only).

3. **Wire the rename branch in Update.** Read both prior `state` and `plan`. If
   `newname` changed and is non-null/non-empty, build a payload of
   `{<keyattr>: <current live name>, newname: plan.newname}` and call
   `r.client.ActOnResource(service.X.Type(), &payload, "rename")`. **The rename source
   (the payload key value) must be the CURRENT LIVE name, which is held in `state.Id`,
   NOT `state.<keyattr>`.** The key attribute stays pinned to the originally-configured
   value (see step 4), so it is stale on a second rename — only `state.Id` tracks the
   live name. (On the first rename `state.Id == state.<keyattr>`, which is why using the
   key attribute appears to work until you rename twice or destroy.) After a successful
   rename, set `data.Id = types.StringValue(newName)` so the ID tracks the live object
   for the read-back and all future reads. Because the resource has no `set` endpoint,
   no other branch is needed — every other attribute is RequiresReplace and never
   reaches Update.

4. **The ID is the live name; Read AND Delete must key off `data.Id`, never the key
   attribute.** This is the load-bearing rule (a missed Delete fix here leaves the
   renamed object dangling on a destroy). After a rename the live object is named
   `newName` (so `id` = newName) while the configured key attribute still holds the
   OLD value:
   - **Read:** `FindResource(service.X.Type(), data.Id.ValueString())` — by ID, not key.
   - **Delete:** `DeleteResource(service.X.Type(), data.Id.ValueString())` — by ID, not
     `data.<keyattr>` (which is stale/old after a rename and would delete the wrong or a
     non-existent object, leaving the renamed object behind).
   - **`xxxSetAttrFromGet`** must NOT blindly overwrite the key attribute from the GET
     response or it will clobber the user's configured value and trigger a spurious
     RequiresReplace diff. Guard it: only adopt the GET key value when the model's key
     is null/empty (covers import, where state carries only the ID — this is also what
     makes an imported resource usable); otherwise preserve the existing value. In
     Update, additionally capture the plan's key + newname before the read-back and
     restore them afterward (belt-and-suspenders).
   - **Lifecycle that proves it:** create(A)→id=A,key=A; rename(newname=B)→live=B, id=B,
     key stays A (config key=A,newname=B → clean plan); read by id=B works; second
     rename(newname=C)→renames B→C via state.Id; delete by id=C removes the live object
     (no dangling). This is exactly the bug class to verify with a rename acceptance
     test (a recreate would leave the object under the old name since `newname` is
     excluded from the add payload — assert the new name exists and the old name is gone).

5. **Datasource is unaffected.** `newname` is rename-only and never a read-back field;
   the datasource keeps using its own `xxxSetAttrFromGetForDatasource` setter (which
   sets the field to null) and sets its own ID. No datasource code changes are needed.

How this differs from the plain RequiresReplace approach: RequiresReplace destroys
and recreates the resource on a name change (losing any server-side state / bindings
attached to it and incurring downtime); the rename action mutates the name in place
on the appliance, preserving everything. Use rename whenever NITRO offers the verb.

## Reference Examples

- **Singleton**: `lbparameter` — static ID, no delete, `UpdateUnnamedResource`
- **Named resource**: `aaagroup` — single unique attr as ID, `AddResource`/`DeleteResource`
- **Named resource with secret**: `sslcertkey` — `passplain` expanded to `passplain`, `passplain_wo`, `passplain_wo_version`
- **Binding with parent**: `aaagroup_aaauser_binding` — parent `groupname`, delete arg `username`
- **Singleton with secret**: `lbparameter` — `cookiepassphrase` expanded to 3 attrs
- **Rename support (`newname`)**: `lbpolicylabel` — in-place rename via NITRO
  `?action=rename` in Update; SDK v2 origin `citrixadc/resource_citrixadc_appfwpolicy.go`
