package nsservicepath

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsservicepathResourceModel describes the resource data model.
type NsservicepathResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Servicepathname types.String `tfsdk:"servicepathname"`
}

func (r *NsservicepathResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsservicepath resource.",
			},
			"servicepathname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must\n      contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-)\n      characters.",
			},
		},
	}
}

func nsservicepathGetThePayloadFromtheConfig(ctx context.Context, data *NsservicepathResourceModel) ns.Nsservicepath {
	tflog.Debug(ctx, "In nsservicepathGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsservicepath := ns.Nsservicepath{}
	if !data.Servicepathname.IsNull() {
		nsservicepath.Servicepathname = data.Servicepathname.ValueString()
	}

	return nsservicepath
}

func nsservicepathSetAttrFromGet(ctx context.Context, data *NsservicepathResourceModel, getResponseData map[string]interface{}) *NsservicepathResourceModel {
	tflog.Debug(ctx, "In nsservicepathSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["servicepathname"]; ok && val != nil {
		data.Servicepathname = types.StringValue(val.(string))
	} else {
		data.Servicepathname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Servicepathname.ValueString())

	return data
}
