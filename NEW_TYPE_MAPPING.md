# Terraform Plugin Framework Type Mapping Reference

This document provides a comprehensive mapping between types from metadata, Terraform Plugin Framework types, and Schema Attribute types.

## Basic Type Mappings

| Metadata Type | Framework Type | Schema Attribute | Import Package |
|---------|---------------|------------------|----------------|
| `string` | `types.String` | `schema.StringAttribute` | `github.com/hashicorp/terraform-plugin-framework/types` |
| `integer` | `types.Int64` | `schema.Int64Attribute` | `github.com/hashicorp/terraform-plugin-framework/types` |
| `boolean` | `types.Bool` | `schema.BoolAttribute` | `github.com/hashicorp/terraform-plugin-framework/types` |
| `number` | `types.Float64` | `schema.Float64Attribute` | `github.com/hashicorp/terraform-plugin-framework/types` |

## Complex Type Mappings

| Metadata Type | Framework Type | Schema Attribute | Import Package |
|---------|---------------|------------------|----------------|
| `string[]` | `types.List` | `schema.ListAttribute` with `ElementType: types.StringType` | `github.com/hashicorp/terraform-plugin-framework/types` |
| `integer[]` | `types.List` | `schema.ListAttribute` with `ElementType: types.Int64Type` | `github.com/hashicorp/terraform-plugin-framework/types` |
| `boolean[]` | `types.List` | `schema.ListAttribute` with `ElementType: types.BoolType` | `github.com/hashicorp/terraform-plugin-framework/types` |
| `number[]` | `types.List` | `schema.ListAttribute` with `ElementType: types.Float64Type` | `github.com/hashicorp/terraform-plugin-framework/types` |

