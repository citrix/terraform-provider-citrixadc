package nslimitsessions

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NslimitsessionsDataSourceModel describes the datasource data model. The
// datasource uses its own model so the resource model can stay minimal.
type NslimitsessionsDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Limitidentifier types.String `tfsdk:"limitidentifier"`
	Detail          types.Bool   `tfsdk:"detail"`
	Timeout         types.String `tfsdk:"timeout"`
	Hits            types.String `tfsdk:"hits"`
	Drop            types.String `tfsdk:"drop"`
	Name            types.String `tfsdk:"name"`
	Unit            types.String `tfsdk:"unit"`
}

func NslimitsessionsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"limitidentifier": schema.StringAttribute{
				Required:    true,
				Description: "Name of the rate limit identifier for which to display the sessions.",
			},
			"detail": schema.BoolAttribute{
				Optional:    true,
				Description: "Show the individual hash values.",
			},
			"timeout": schema.StringAttribute{
				Computed:    true,
				Description: "The time remaining on the session before a flush can be attempted.",
			},
			"hits": schema.StringAttribute{
				Computed:    true,
				Description: "The number of times this entry was hit.",
			},
			"drop": schema.StringAttribute{
				Computed:    true,
				Description: "The number of times action was taken.",
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "The string formed by gathering selectlet values.",
			},
			"unit": schema.StringAttribute{
				Computed:    true,
				Description: "Total computed hash of the matched selectlets.",
			},
		},
	}
}
