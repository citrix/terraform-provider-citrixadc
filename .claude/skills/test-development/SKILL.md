---
name: test-development
description: Write acceptance tests for terraform-provider-citrixadc resources in citrixadc_framework/acctests/
---

## Environment Setup

Acceptance tests require a live Citrix ADC instance. The following ADC testbeds are available:

| Mode | IP Address | `ADC_TESTBED` Value | Use For |
|---|---|---|---|
| **Standalone** (primary) | `10.101.132.121` | Not set / `UNSPECIFIED` | Most tests (default) |
| **Standalone** (secondary pool) | `10.101.132.151`–`10.101.132.155` | Not set / `UNSPECIFIED` | Fallback / additional standalone instances — retry here when a test fails on the primary due to an ADC-specific error, or use as extra lanes for parallel runs (see below) |
| **Cluster** | `10.101.132.133` | `CLUSTER` | Cluster-specific tests |
| **HA** | `10.101.132.141` | `HA_PAIR` or `HA` | HA-specific tests |

The primary (`10.101.132.121`) and the secondary standalone pool (`10.101.132.151`, `10.101.132.152`, `10.101.132.153`, `10.101.132.154`, `10.101.132.155`) are all interchangeable standalone instances that share the **same credentials** (`NS_LOGIN=nsroot`, `NS_PASSWORD='CADS123$%^'`). The secondary pool exists so a standalone test that fails for reasons specific to one box can be retried on a clean instance before concluding the provider code is at fault, and so independent test runs can be spread across separate boxes (one instance per lane) to avoid shared-ADC state contention.

### Running Tests Against Standalone ADC (default)

```bash
export NS_URL=http://10.101.132.121/
export NS_LOGIN=nsroot
export NS_PASSWORD='CADS123$%^'
TF_ACC=1 go test ./citrixadc_framework/acctests/ -v -run TestAccLbparameter_basic -timeout 120m
```

### Fallback: Retry on a Secondary Standalone ADC (`10.101.132.151`–`10.101.132.155`)

When a standalone test **fails on the primary (`10.101.132.121`) because of an ADC-specific error** — not a provider/code bug — re-run the exact same test against any secondary standalone instance (`10.101.132.151` through `10.101.132.155`), all of which have identical credentials. If it passes there, the failure was the box, not the code.

ADC-specific (box-specific) errors that warrant a fallback retry include:
- Connectivity / availability issues with the primary: connection refused/timeout, `5xx` from the NITRO endpoint, the instance being down or rebooting.
- Leftover/orphaned config on the primary from a previous broken run (e.g. a create that fails with an "already exists" conflict, or a destroy that can't find the entity) — a clean secondary side-steps it.
- Feature/licensing/state quirks specific to the primary box (e.g. a feature not enabled, a partition switched, disk full) that are unrelated to the resource under test.

Do **NOT** fall back for failures that indicate a real provider bug — schema/plan errors ("inconsistent result after apply", "unknown value"), ID-parse/import failures, wrong NITRO payloads, or assertion mismatches. Those reproduce on any ADC and must be fixed in code, not retried on another box.

```bash
# Same test, same credentials, a secondary standalone instance (.151–.155)
export NS_URL=http://10.101.132.151/   # or .152 / .153 / .154 / .155
export NS_LOGIN=nsroot
export NS_PASSWORD='CADS123$%^'
TF_ACC=1 go test ./citrixadc_framework/acctests/ -v -run TestAccLbparameter_basic -timeout 120m
```

Leave `ADC_TESTBED` unset (`UNSPECIFIED`) for any secondary — they are standalone boxes, so standalone skip conditions still apply. When reporting results, note which standalone instance produced the passing run if a fallback was used (e.g. "passed on secondary `.152` after a connectivity failure on `.121`").

### Running Tests Against Cluster ADC

```bash
export NS_URL=http://10.101.132.133/
export NS_LOGIN=nsroot
export NS_PASSWORD='CADS123$%^'
export ADC_TESTBED=CLUSTER
TF_ACC=1 go test ./citrixadc_framework/acctests/ -v -run TestAccCluster -timeout 120m
```

### Running Tests Against HA ADC

```bash
export NS_URL=http://10.101.132.141/
export NS_LOGIN=nsroot
export NS_PASSWORD='CADS123$%^'
export ADC_TESTBED=HA_PAIR
TF_ACC=1 go test ./citrixadc_framework/acctests/ -v -run TestAccHanode -timeout 120m
```

### ADC Testbed Skip Conditions

Tests use skip conditions to run only on the appropriate testbed. Two global variables are initialized in `provider_test.go`:

```go
var isCpxRun bool      // true if NS_URL contains "localhost" (CPX container)
var adcTestbed string   // from ADC_TESTBED env var, defaults to "UNSPECIFIED"
```

**Skip pattern for Cluster-only tests:**
```go
if adcTestbed != "CLUSTER" {
    t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
}
```

**Skip pattern for HA-only tests:**
```go
if adcTestbed != "HA_PAIR" {
    t.Skipf("ADC testbed is %s. Expected HA_PAIR.", adcTestbed)
}
```

**Skip pattern for CPX-incompatible tests:**
```go
if isCpxRun {
    t.Skip("Feature not available in CPX")
}
```

**Common `ADC_TESTBED` values:** `CLUSTER`, `HA_PAIR`, `HA`, `STANDALONE_HSM`, `STANDALONE_NON_DEFAULT_SSL_PROFILE`, `STANDALONE_12CORES`, `INSTALLER`, `UNSPECIFIED`

When writing tests, add the appropriate skip condition if the test only applies to a specific ADC mode. Most tests (CRUD, ephemeral, datasource) run against Standalone and need no skip condition.

## Test File Structure

Each resource has a test file at `citrixadc_framework/acctests/<resource_name>_test.go`.

All test files use:
- Package: `package citrixadc`
- Provider factories: `testAccProtoV6ProviderFactories`
- Pre-check: `testAccPreCheck(t)`
- Client helper: `testAccGetFrameworkClient()` (from `helpers_test.go`)

## Resource Categories

Resources fall into 4 categories. Know which you are testing:

| Category | Traits | Create API | Delete API |
|---|---|---|---|
| **singleton** | No unique attrs, no parent | `UpdateUnnamedResource` | No delete (state-only removal) |
| **named_resource** | Has unique attrs, no parent | `AddResource` | `DeleteResource` |
| **global_binding** | Has unique attrs, no parent attrs, has `delete_arg_attrs` | `AddResource` | `DeleteResourceWithArgsMap` (empty name) |
| **binding_with_parent** | Has parent attrs + delete_arg_attrs | `AddResource` | `DeleteResourceWithArgsMap` (parent value) |

### Singleton Resources (e.g., lbparameter, aaacertparams)

- No destroy check needed (resource always exists on ADC)
- `testAccCheckDestroy` should verify the resource still exists (it's never truly deleted)
- ID is static (e.g., `"lbparameter-config"`)

### Named Resources (e.g., lbvserver, sslcertkey)

- Destroy check verifies the resource no longer exists on ADC
- ID is the resource name or `key:UrlEncode(value)` pairs

### Binding Resources (global_binding, binding_with_parent)

- ID uses `key:UrlEncode(value)` format for multiple unique attrs
- Destroy check uses `FindResourceArrayWithParams` + array filtering
- Use `utils.ParseIdString()` to parse IDs in check functions

#### Sourcing config for participating entities

A binding resource joins two (or more) entities, so its test config must first **create the participating entities** before declaring the binding. Don't invent that config from scratch — **reuse the working HCL from the existing acceptance test of each participating entity**.

How to identify the participants: split the binding name on `_..._binding`. The leading token is the parent/primary entity, and the embedded token(s) are the bound entity. For example:

| Binding resource | Participating entities | Acceptance tests to reference |
|---|---|---|
| `metricsprofile_servicegroup_binding` | `metricsprofile`, `servicegroup` | `metricsprofile_test.go`, `servicegroup_test.go` |
| `clusternodegroup_lbvserver_binding` | `clusternodegroup`, `lbvserver` | `clusternodegroup_test.go`, `lbvserver_test.go` |
| `sslvserver_sslcertkey_binding` | `sslvserver`, `sslcertkey` | `sslvserver_test.go`, `sslcertkey_test.go` |

Procedure:
1. For each participating entity, look for `citrixadc_framework/acctests/<entity>_test.go`. If present, lift the `resource "citrixadc_<entity>" ...` block from its `_basic_step1` config constant (drop attributes irrelevant to the binding; keep the required ones) and adapt the instance label.
2. Wire the binding to those resources by **reference**, not hardcoded literals — e.g. `servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname` — and add `depends_on` where ordering matters.
3. If a participating entity has **no** acceptance test yet, fall back to the resource's own schema (`citrixadc_framework/<entity>/resource_schema.go`) for required attributes, then to web search (NetScaler docs) per the agent's research step.
4. Participating-entity acceptance tests live almost entirely in `citrixadc_framework/acctests/<entity>_test.go` (the directory has 700+ test files, including common entities like `servicegroup`, `lbvserver`, `sslvserver`, `sslcertkey`, `clusternodegroup`). The legacy SDK v2 `citrixadc/` package has only a handful of test files, so it is rarely useful — check the Framework acctests dir first, and only glance at `citrixadc/` as a last resort.

## Writing Basic CRUD Tests

### Standard Pattern

```go
func TestAcc<Resource>_basic(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        CheckDestroy:             testAccCheck<Resource>Destroy,
        Steps: []resource.TestStep{
            {
                Config: testAcc<Resource>_basic_step1,
                Check: resource.ComposeTestCheckFunc(
                    testAccCheck<Resource>Exist("citrixadc_<resource>.tf_<resource>", nil),
                    resource.TestCheckResourceAttr("citrixadc_<resource>.tf_<resource>", "<attr>", "<value>"),
                ),
            },
            {
                Config: testAcc<Resource>_basic_step2,  // Updated values
                Check: resource.ComposeTestCheckFunc(
                    // Verify updated values
                ),
            },
        },
    })
}
```

### Exist Check Function

```go
func testAccCheck<Resource>Exist(n string, id *string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[n]
        if !ok {
            return fmt.Errorf("Not found: %s", n)
        }
        if rs.Primary.ID == "" {
            return fmt.Errorf("No <resource> ID is set")
        }
        if id != nil {
            if *id != "" && *id != rs.Primary.ID {
                return fmt.Errorf("Resource ID has changed!")
            }
            *id = rs.Primary.ID
        }

        client, err := testAccGetFrameworkClient()
        if err != nil {
            return fmt.Errorf("Failed to get test client: %v", err)
        }
        data, err := client.FindResource(service.<NitroType>.Type(), "<resource_name_or_empty>")
        if err != nil {
            return err
        }
        if data == nil {
            return fmt.Errorf("<resource> %s not found", n)
        }
        return nil
    }
}
```

### Destroy Check Function

For **singleton** resources (always exist on ADC):
```go
func testAccCheck<Resource>Destroy(s *terraform.State) error {
    client, err := testAccGetFrameworkClient()
    if err != nil {
        return fmt.Errorf("Failed to get test client: %v", err)
    }
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "citrixadc_<resource>" {
            continue
        }
        _, err := client.FindResource(service.<NitroType>.Type(), rs.Primary.ID)
        if err == nil {
            // Singleton resources always exist - this is expected
            return fmt.Errorf("<resource> %s still exists", rs.Primary.ID)
        }
    }
    return nil
}
```

For **binding** resources (parse ID, filter array):
```go
func testAccCheck<Resource>Destroy(s *terraform.State) error {
    client, err := testAccGetFrameworkClient()
    if err != nil {
        return fmt.Errorf("Failed to get test client: %v", err)
    }
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "citrixadc_<resource>" {
            continue
        }
        idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"attr1", "attr2"}, nil)
        if err != nil {
            return err
        }
        // Use FindResourceArrayWithParams + filter to check binding no longer exists
    }
    return nil
}
```

## Ephemeral / Write-Only Attribute Testing

When a resource has secret attributes (marked `x-secret-attr` in metadata), the generated code expands them into 3 schema attributes:

| Attribute | Schema Flags | Purpose |
|---|---|---|
| `<name>` | Optional, Sensitive | Backward-compatible legacy path |
| `<name>_wo` | Optional, Sensitive, WriteOnly | New ephemeral path (value not persisted in state) |
| `<name>_wo_version` | Optional, Computed, Default(1) | Version tracker to trigger updates |

### Important Behaviors

1. **Sensitive attribute (`<name>`)**: Value IS persisted in state (encrypted). Terraform detects changes by comparing old vs new values.
2. **WriteOnly attribute (`<name>_wo`)**: Value is NOT persisted in state. To trigger an update when the value changes, the user must bump `<name>_wo_version`.
3. **Both can coexist**: If both are set, `_wo` takes precedence (applied last in `GetPayloadFromConfig`).
4. **Neither is returned by NITRO API**: Secret values are never returned by the ADC — `SetAttrFromGet` skips them and retains the value from config/state.

### Test Pattern: Backward-Compatible Path (Sensitive Attribute)

Pass sensitive values via `TF_VAR_*` environment variables and Terraform `variable` blocks. Use different variable names per step to force Terraform to detect a change.

```go
const testAcc<Resource>_secret_step1 = `
    variable "<resource>_<secret_attr>" {
      type      = string
      sensitive = true
    }

    resource "citrixadc_<resource>" "test" {
        <secret_attr> = var.<resource>_<secret_attr>
        // ... other attrs
    }
`

const testAcc<Resource>_secret_step2 = `
    variable "<resource>_<secret_attr>_2" {
      type      = string
      sensitive = true
    }

    resource "citrixadc_<resource>" "test" {
        <secret_attr> = var.<resource>_<secret_attr>_2
        // ... other attrs
    }
`

func TestAcc<Resource>_secret_backward_compat(t *testing.T) {
    t.Setenv("TF_VAR_<resource>_<secret_attr>", "value1")
    t.Setenv("TF_VAR_<resource>_<secret_attr>_2", "value2")
    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        CheckDestroy:             testAccCheck<Resource>Destroy,
        Steps: []resource.TestStep{
            {
                Config: testAcc<Resource>_secret_step1,
                Check: resource.ComposeTestCheckFunc(
                    testAccCheck<Resource>Exist("citrixadc_<resource>.test", nil),
                    // Verify non-secret attributes that confirm the secret was applied
                ),
            },
            {
                Config: testAcc<Resource>_secret_step2,
                Check: resource.ComposeTestCheckFunc(
                    testAccCheck<Resource>Exist("citrixadc_<resource>.test", nil),
                    // Verify update was applied
                ),
            },
        },
    })
}
```

### Test Pattern: Ephemeral Path (WriteOnly Attribute)

Use `<name>_wo` with `<name>_wo_version`. Bump the version between steps to trigger an update.

```go
const testAcc<Resource>_wo_step1 = `
    variable "<resource>_<secret_attr>_wo" {
      type      = string
      sensitive = true
    }

    resource "citrixadc_<resource>" "test" {
        <secret_attr>_wo         = var.<resource>_<secret_attr>_wo
        <secret_attr>_wo_version = 1
        // ... other attrs
    }
`

const testAcc<Resource>_wo_step2 = `
    variable "<resource>_<secret_attr>_wo_2" {
      type      = string
      sensitive = true
    }

    resource "citrixadc_<resource>" "test" {
        <secret_attr>_wo         = var.<resource>_<secret_attr>_wo_2
        <secret_attr>_wo_version = 2
        // ... other attrs
    }
`

func TestAcc<Resource>_secret_wo_ephemeral(t *testing.T) {
    t.Setenv("TF_VAR_<resource>_<secret_attr>_wo", "ephemeral_value1")
    t.Setenv("TF_VAR_<resource>_<secret_attr>_wo_2", "ephemeral_value2")
    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        CheckDestroy:             testAccCheck<Resource>Destroy,
        Steps: []resource.TestStep{
            {
                Config: testAcc<Resource>_wo_step1,
                Check: resource.ComposeTestCheckFunc(
                    testAccCheck<Resource>Exist("citrixadc_<resource>.test", nil),
                    // Verify wo_version is tracked in state
                    resource.TestCheckResourceAttr("citrixadc_<resource>.test", "<secret_attr>_wo_version", "1"),
                ),
            },
            {
                Config: testAcc<Resource>_wo_step2,
                Check: resource.ComposeTestCheckFunc(
                    testAccCheck<Resource>Exist("citrixadc_<resource>.test", nil),
                    // Verify version bumped — confirms update was triggered
                    resource.TestCheckResourceAttr("citrixadc_<resource>.test", "<secret_attr>_wo_version", "2"),
                ),
            },
        },
    })
}
```

### Key Points for Ephemeral Tests

- **DO NOT** use `resource.TestCheckResourceAttr` on `<name>_wo` — write-only values are not stored in state and cannot be checked.
- **DO** check `<name>_wo_version` — this IS stored in state and confirms the update path was triggered.
- **DO** check observable side-effects (e.g., if setting `cookiepassphrase` also requires `useencryptedpersistencecookie = "ENABLED"`, verify that attr).
- **Use `t.Setenv()`** instead of `os.Setenv()` — it auto-cleans up after the test.
- **Use different variable names** per step (e.g., `_wo` and `_wo_2`) to ensure Terraform sees a config change.

## Datasource Tests

```go
const testAcc<Resource>DataSource_basic = `
    resource "citrixadc_<resource>" "test" {
        // ... create resource
    }

    data "citrixadc_<resource>" "test" {
        <unique_attr> = citrixadc_<resource>.test.<unique_attr>
        depends_on = [citrixadc_<resource>.test]
    }
`

func TestAcc<Resource>DataSource_basic(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: testAcc<Resource>DataSource_basic,
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttr("data.citrixadc_<resource>.test", "<attr>", "<value>"),
                ),
            },
        },
    })
}
```

For **singleton** datasources (e.g., lbparameter), no unique attr is needed — just use `depends_on`.

## ID Format and ParseIdString

Resources use `key:UrlEncode(value)` format for IDs with multiple unique attributes. When parsing IDs in test check functions, use `utils.ParseIdString`:

```go
import "github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

idMap, _, err := utils.ParseIdString(bindingId, []string{"attr1", "attr2"}, []string{"optional_attr"})
if err != nil {
    return fmt.Errorf("Error parsing ID: %v", err)
}
parentValue := idMap["attr1"]
```

Arguments:
- `legacyAttrOrder []string` — ordered attribute names matching SDK v2 comma-separated ID format
- `legacyOptionalAttrs []string` — attributes that may be absent in legacy IDs (suffix `?` in `resource_id_mapping.json`)

## Reference Examples

- **Singleton with ephemeral**: `lbparameter_test.go` — `TestAccLbparameter_cookiepassphrase_backward_compat`, `TestAccLbparameter_cookiepassphrase_wo_ephemeral`
- **Named resource with ephemeral**: `sslcertkey_test.go` — `TestAccSslcertkey_basic` (passplain_wo path), `TestAccSslcertkey_passplain` (backward-compat path)
- **Named resource basic CRUD**: `aaagroup_test.go`
- **Binding with parent**: `aaagroup_aaauser_binding_test.go`
- **Datasource**: `lbparameter_test.go` — `TestAccLbparameterDataSource_basic`
