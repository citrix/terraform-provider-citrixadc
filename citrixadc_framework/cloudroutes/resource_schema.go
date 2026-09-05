package cloudroutes

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. Because these attributes revert to no value
// (absent from GET) after unset, marking the plan unknown also avoids a
// "provider produced inconsistent result" error, which a static Default would
// trigger.
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// CloudroutesResourceModel describes the resource data model.
type CloudroutesResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Routesvpcnetwork types.String `tfsdk:"routesvpcnetwork"`
	Vipsubnet        types.String `tfsdk:"vipsubnet"`
	Vipvpcnetwork    types.String `tfsdk:"vipvpcnetwork"`
	Clientipaddress  types.String `tfsdk:"clientipaddress"`
}

func (r *CloudroutesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudroutes resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the route.",
			},
			"routesvpcnetwork": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "client vpc network name.",
			},
			"vipsubnet": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "vip subnet in CIDR format.",
			},
			"vipvpcnetwork": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "vip vpc network name.",
			},
			"clientipaddress": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "IPv4 or IPv6 address attached to the nic interface towards vpc mentiond in vpcnetwork.",
			},
		},
	}
}

func cloudroutesGetThePayloadFromthePlan(ctx context.Context, data *CloudroutesResourceModel) cloud.Cloudroutes {
	tflog.Debug(ctx, "In cloudroutesGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudroutes := cloud.Cloudroutes{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		cloudroutes.Name = data.Name.ValueString()
	}
	if !data.Routesvpcnetwork.IsNull() && !data.Routesvpcnetwork.IsUnknown() {
		cloudroutes.Routesvpcnetwork = data.Routesvpcnetwork.ValueString()
	}
	if !data.Vipsubnet.IsNull() && !data.Vipsubnet.IsUnknown() {
		cloudroutes.Vipsubnet = data.Vipsubnet.ValueString()
	}
	if !data.Vipvpcnetwork.IsNull() && !data.Vipvpcnetwork.IsUnknown() {
		cloudroutes.Vipvpcnetwork = data.Vipvpcnetwork.ValueString()
	}
	if !data.Clientipaddress.IsNull() && !data.Clientipaddress.IsUnknown() {
		cloudroutes.Clientipaddress = data.Clientipaddress.ValueString()
	}

	return cloudroutes
}

func cloudroutesSetAttrFromGet(ctx context.Context, data *CloudroutesResourceModel, getResponseData map[string]interface{}) *CloudroutesResourceModel {
	tflog.Debug(ctx, "In cloudroutesSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["routesvpcnetwork"]; ok && val != nil {
		data.Routesvpcnetwork = types.StringValue(val.(string))
	} else {
		data.Routesvpcnetwork = types.StringNull()
	}
	if val, ok := getResponseData["vipsubnet"]; ok && val != nil {
		data.Vipsubnet = types.StringValue(val.(string))
	} else {
		data.Vipsubnet = types.StringNull()
	}
	if val, ok := getResponseData["vipvpcnetwork"]; ok && val != nil {
		data.Vipvpcnetwork = types.StringValue(val.(string))
	} else {
		data.Vipvpcnetwork = types.StringNull()
	}
	if val, ok := getResponseData["clientipaddress"]; ok && val != nil {
		data.Clientipaddress = types.StringValue(val.(string))
	} else {
		data.Clientipaddress = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
