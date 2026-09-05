package wasmfile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func WasmfileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the WASM module/signature page object on the Citrix ADC.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Local path or URL for the file from which the WASM object was imported.",
			},
			"filetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "WASM file type. Possible values = Module, Signature, Setting.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the WASM page object.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Import-only option; not returned by GET.",
			},
		},
	}
}
