package nschannelparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NschannelparamResourceModel describes the resource data model.
type NschannelparamResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Vfautorecover types.String `tfsdk:"vfautorecover"`
}

func (r *NschannelparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nschannelparam resource.",
			},
			"vfautorecover": schema.StringAttribute{
				Required:    true,
				Description: "VF autorecover mode",
			},
		},
	}
}

func nschannelparamGetThePayloadFromthePlan(ctx context.Context, data *NschannelparamResourceModel) ns.Nschannelparam {
	tflog.Debug(ctx, "In nschannelparamGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nschannelparam := ns.Nschannelparam{}
	if !data.Vfautorecover.IsNull() && !data.Vfautorecover.IsUnknown() {
		nschannelparam.Vfautorecover = data.Vfautorecover.ValueString()
	}

	return nschannelparam
}

// nschannelparamSetAttrFromGet populates the resource model from the GET response.
// vfautorecover is Required and always echoed by GET, so we copy it when present.
// The synthetic ID is set exactly once in Create (Pattern 6), so it is NOT recomputed
// here.
func nschannelparamSetAttrFromGet(ctx context.Context, data *NschannelparamResourceModel, getResponseData map[string]interface{}) *NschannelparamResourceModel {
	tflog.Debug(ctx, "In nschannelparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["vfautorecover"]; ok && val != nil {
		data.Vfautorecover = types.StringValue(val.(string))
	}

	return data
}

// nschannelparamSetAttrFromGetForDatasource faithfully copies every field from the GET
// response and sets the synthetic ID, because the datasource never calls Create
// (Pattern 7).
func nschannelparamSetAttrFromGetForDatasource(ctx context.Context, data *NschannelparamResourceModel, getResponseData map[string]interface{}) *NschannelparamResourceModel {
	tflog.Debug(ctx, "In nschannelparamSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["vfautorecover"]; ok && val != nil {
		data.Vfautorecover = types.StringValue(val.(string))
	} else {
		data.Vfautorecover = types.StringNull()
	}

	// Datasource has no Create, so set the synthetic ID here.
	data.Id = types.StringValue("nschannelparam-config")

	return data
}
