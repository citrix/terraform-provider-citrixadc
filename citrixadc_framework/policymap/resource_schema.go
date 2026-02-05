package policymap

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicymapResourceModel describes the resource data model.
type PolicymapResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Mappolicyname types.String `tfsdk:"mappolicyname"`
	Sd            types.String `tfsdk:"sd"`
	Su            types.String `tfsdk:"su"`
	Td            types.String `tfsdk:"td"`
	Tu            types.String `tfsdk:"tu"`
}

func (r *PolicymapResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policymap resource.",
			},
			"mappolicyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the map policy. Must begin with a letter, number, or the underscore (_) character and must consist only of letters, numbers, and the hash (#), period (.), colon (:), space ( ), at (@), equals (=), hyphen (-), and underscore (_) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my map\" or 'my map').",
			},
			"sd": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Publicly known source domain name. This is the domain name with which a client request arrives at a reverse proxy virtual server for cache redirection. If you specify a source domain, you must specify a target domain.",
			},
			"su": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Source URL. Specify all or part of the source URL, in the following format: /[[prefix] [*]] [.suffix].",
			},
			"td": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Target domain name sent to the server. The source domain name is replaced with this domain name.",
			},
			"tu": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Target URL. Specify the target URL in the following format: /[[prefix] [*]][.suffix].",
			},
		},
	}
}

func policymapGetThePayloadFromtheConfig(ctx context.Context, data *PolicymapResourceModel) policy.Policymap {
	tflog.Debug(ctx, "In policymapGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policymap := policy.Policymap{}
	if !data.Mappolicyname.IsNull() {
		policymap.Mappolicyname = data.Mappolicyname.ValueString()
	}
	if !data.Sd.IsNull() {
		policymap.Sd = data.Sd.ValueString()
	}
	if !data.Su.IsNull() {
		policymap.Su = data.Su.ValueString()
	}
	if !data.Td.IsNull() {
		policymap.Td = data.Td.ValueString()
	}
	if !data.Tu.IsNull() {
		policymap.Tu = data.Tu.ValueString()
	}

	return policymap
}

func policymapSetAttrFromGet(ctx context.Context, data *PolicymapResourceModel, getResponseData map[string]interface{}) *PolicymapResourceModel {
	tflog.Debug(ctx, "In policymapSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["mappolicyname"]; ok && val != nil {
		data.Mappolicyname = types.StringValue(val.(string))
	} else {
		data.Mappolicyname = types.StringNull()
	}
	if val, ok := getResponseData["sd"]; ok && val != nil {
		data.Sd = types.StringValue(val.(string))
	} else {
		data.Sd = types.StringNull()
	}
	if val, ok := getResponseData["su"]; ok && val != nil {
		data.Su = types.StringValue(val.(string))
	} else {
		data.Su = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		data.Td = types.StringValue(val.(string))
	} else {
		data.Td = types.StringNull()
	}
	if val, ok := getResponseData["tu"]; ok && val != nil {
		data.Tu = types.StringValue(val.(string))
	} else {
		data.Tu = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Mappolicyname.ValueString())

	return data
}
