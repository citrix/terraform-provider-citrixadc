package mapdomain

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MapdomainResourceModel describes the resource data model.
type MapdomainResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Mapdmrname types.String `tfsdk:"mapdmrname"`
	Name       types.String `tfsdk:"name"`
}

func (r *MapdomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the mapdomain resource.",
			},
			"mapdmrname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Default Mapping rule name.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the MAP Domain. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Domain is created . The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"add network MapDomain map1\").",
			},
		},
	}
}

func mapdomainGetThePayloadFromtheConfig(ctx context.Context, data *MapdomainResourceModel) network.Mapdomain {
	tflog.Debug(ctx, "In mapdomainGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	mapdomain := network.Mapdomain{}
	if !data.Mapdmrname.IsNull() {
		mapdomain.Mapdmrname = data.Mapdmrname.ValueString()
	}
	if !data.Name.IsNull() {
		mapdomain.Name = data.Name.ValueString()
	}

	return mapdomain
}

func mapdomainSetAttrFromGet(ctx context.Context, data *MapdomainResourceModel, getResponseData map[string]interface{}) *MapdomainResourceModel {
	tflog.Debug(ctx, "In mapdomainSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["mapdmrname"]; ok && val != nil {
		data.Mapdmrname = types.StringValue(val.(string))
	} else {
		data.Mapdmrname = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
