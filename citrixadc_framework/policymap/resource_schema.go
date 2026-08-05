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
				// SDK v2: Required + ForceNew
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the map policy. Must begin with a letter, number, or the underscore (_) character and must consist only of letters, numbers, and the hash (#), period (.), colon (:), space ( ), at (@), equals (=), hyphen (-), and underscore (_) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my map\" or 'my map').",
			},
			"sd": schema.StringAttribute{
				// SDK v2: Required + ForceNew
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Publicly known source domain name. This is the domain name with which a client request arrives at a reverse proxy virtual server for cache redirection. If you specify a source domain, you must specify a target domain.",
			},
			"su": schema.StringAttribute{
				// SDK v2: Optional + Computed + ForceNew. Keep Optional+Computed; use
				// RequiresReplaceIfConfigured so ForceNew only fires when the user
				// actually configured the value (a computed value never forces replace),
				// and UseStateForUnknown to keep an unconfigured value stable.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Source URL. Specify all or part of the source URL, in the following format: /[[prefix] [*]] [.suffix].",
			},
			"td": schema.StringAttribute{
				// SDK v2: Required + ForceNew
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Target domain name sent to the server. The source domain name is replaced with this domain name.",
			},
			"tu": schema.StringAttribute{
				// SDK v2: Optional + Computed + ForceNew. See su for rationale.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Target URL. Specify the target URL in the following format: /[[prefix] [*]][.suffix].",
			},
		},
	}
}

func policymapGetThePayloadFromthePlan(ctx context.Context, data *PolicymapResourceModel) policy.Policymap {
	tflog.Debug(ctx, "In policymapGetThePayloadFromthePlan Function")

	// Create API request body from the model
	policymap := policy.Policymap{}
	if !data.Mappolicyname.IsNull() && !data.Mappolicyname.IsUnknown() {
		policymap.Mappolicyname = data.Mappolicyname.ValueString()
	}
	if !data.Sd.IsNull() && !data.Sd.IsUnknown() {
		policymap.Sd = data.Sd.ValueString()
	}
	if !data.Su.IsNull() && !data.Su.IsUnknown() {
		policymap.Su = data.Su.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		policymap.Td = data.Td.ValueString()
	}
	if !data.Tu.IsNull() && !data.Tu.IsUnknown() {
		policymap.Tu = data.Tu.ValueString()
	}

	return policymap
}

func policymapSetAttrFromGet(ctx context.Context, data *PolicymapResourceModel, getResponseData map[string]interface{}) *PolicymapResourceModel {
	tflog.Debug(ctx, "In policymapSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["mappolicyname"]; ok && val != nil {
		data.Mappolicyname = types.StringValue(val.(string))
	} else if data.Mappolicyname.IsUnknown() {
		data.Mappolicyname = types.StringNull()
	}
	if val, ok := getResponseData["sd"]; ok && val != nil {
		data.Sd = types.StringValue(val.(string))
	} else if data.Sd.IsUnknown() {
		data.Sd = types.StringNull()
	}
	// su is Optional+Computed and may be omitted from GET; only null it when it is
	// unknown so a known/configured value is never clobbered (omit-on-default trap).
	if val, ok := getResponseData["su"]; ok && val != nil {
		data.Su = types.StringValue(val.(string))
	} else if data.Su.IsUnknown() {
		data.Su = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		data.Td = types.StringValue(val.(string))
	} else if data.Td.IsUnknown() {
		data.Td = types.StringNull()
	}
	// tu is Optional+Computed and may be omitted from GET; same guard as su.
	if val, ok := getResponseData["tu"]; ok && val != nil {
		data.Tu = types.StringValue(val.(string))
	} else if data.Tu.IsUnknown() {
		data.Tu = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Mappolicyname.ValueString())

	return data
}
