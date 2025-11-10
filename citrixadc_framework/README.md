# CitrixADC Framework Provider

This directory contains a sample provider implementation based on Terraform Plugin Framework that works alongside the existing SDK v2 provider.

## Structure

```
citrixadc_framework/
├── provider.go                   # Main framework provider implementation
├── resource_lbvserver.go         # Sample framework resource
├── datasource_nsversion.go       # Sample framework data source
└── provider_test.go              # Basic provider tests
```

## Features

### Provider Configuration
The framework provider mirrors the existing SDK v2 provider configuration:
- Same provider schema (username, password, endpoint, etc.)
- Environment variable support (NS_LOGIN, NS_PASSWORD, NS_URL)
- Uses the same Nitro client for consistent API behavior

### Sample Resources and Data Sources

#### Resource: `citrixadc_lbvserver_fw`
- Framework-based load balancer virtual server resource
- Demonstrates CRUD operations using the Nitro client
- Type-safe configuration with proper validation
- Import state support

#### Data Source: `citrixadc_nsversion_fw`
- Framework-based NetScaler version data source
- Shows how to read data from the API using framework patterns

## Usage Examples

### Framework Resource
```hcl
resource "citrixadc_lbvserver_fw" "example" {
  name        = "example-lb-fw"
  servicetype = "HTTP"
  ipv46       = "192.168.1.100"
  port        = 80
}
```

### Framework Data Source
```hcl
data "citrixadc_nsversion_fw" "version" {}

output "ns_version" {
  value = data.citrixadc_nsversion_fw.version.version
}
```

## Building

### Standalone Framework Provider
```bash
go build -o terraform-provider-citrixadc-framework main_framework_standalone.go
```

### With SDK v2 Provider (Current)
```bash
go build -o terraform-provider-citrixadc .
```

## Benefits

1. **Type Safety**: Compile-time type checking with framework types
2. **Modern API**: Uses latest Terraform provider patterns
3. **Better Validation**: Built-in framework validation capabilities
4. **Consistent Client**: Reuses same Nitro client as SDK v2 provider
5. **Gradual Migration**: Can coexist with existing SDK v2 resources

## Testing

```bash
go test ./citrixadc_framework/
```

## Implementation Notes

- Framework resources use `_fw` suffix to avoid conflicts with SDK v2 resources
- Provider uses same configuration schema for consistency
- All CRUD operations delegate to the same Nitro client
- Type-safe configuration with proper null/unknown handling
- Proper diagnostic error handling

This implementation provides a foundation for gradually migrating from SDK v2 to the more modern Plugin Framework while maintaining backward compatibility.