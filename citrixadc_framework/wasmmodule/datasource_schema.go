package wasmmodule

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func WasmmoduleDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the WASM module file.",
			},
			"modulefile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File name of the WASM module.",
			},
			"signaturefile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The SHA256 file contains the hash value used to validate the WASM module.",
			},
			"settingfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The WASM module filename contains module-specific configuration settings.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this WASM module.",
			},
		},
	}
}
