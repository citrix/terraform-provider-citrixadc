package appfwsignatures

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AppfwsignaturesDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Signature action",
			},
			"autoenablenewsignatures": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable auto enable new signatures",
			},
			"category": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Signature category to be Enabled/Disabled",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the signatures object.",
			},
			"enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable enable signature rule IDs/Signature Category",
			},
			"merge": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Merges the existing Signature with new signature rules",
			},
			"mergedefault": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Merges signature file with default signature file.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the signature object.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing signatures object of the same name.",
			},
			"preservedefactions": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "preserves def actions of signature rules",
			},
			"ruleid": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Signature rule IDs to be Enabled/Disabled",
			},
			"sha1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File path for sha1 file to validate signature file",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and file name) for the location at which to store the imported signatures object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
			"vendortype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Third party vendor type for which WAF signatures has to be generated.",
			},
			"xslt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "XSLT file source.",
			},
		},
	}
}
