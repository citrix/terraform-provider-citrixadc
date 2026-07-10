package vpnglobal_staserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalStaserverBindingResourceModel describes the resource data model.
type VpnglobalStaserverBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Staaddresstype         types.String `tfsdk:"staaddresstype"`
	Staserver              types.String `tfsdk:"staserver"`
}

func (r *VpnglobalStaserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_staserver_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// Pattern 13: NITRO never echoes this field back in GET, so it cannot be
				// Computed (Terraform would require a resolved value after apply that the
				// provider can never supply). Keep it Optional-only.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"staaddresstype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the STA server address(ipv4/v6).",
			},
			"staserver": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Configured Secure Ticketing Authority (STA) server.",
			},
		},
	}
}

func vpnglobal_staserver_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalStaserverBindingResourceModel) vpn.Vpnglobalstaserverbinding {
	tflog.Debug(ctx, "In vpnglobal_staserver_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_staserver_binding := vpn.Vpnglobalstaserverbinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_staserver_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Staaddresstype.IsNull() && !data.Staaddresstype.IsUnknown() {
		vpnglobal_staserver_binding.Staaddresstype = data.Staaddresstype.ValueString()
	}
	if !data.Staserver.IsNull() && !data.Staserver.IsUnknown() {
		vpnglobal_staserver_binding.Staserver = data.Staserver.ValueString()
	}

	return vpnglobal_staserver_binding
}

func vpnglobal_staserver_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalStaserverBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalStaserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_staserver_bindingSetAttrFromGet Function")

	// Convert API response to model.
	// gotopriorityexpression is a write-only input that NITRO does NOT echo back in
	// the GET response (Pattern 7): preserve the existing plan/state value rather than
	// nulling it, which would cause an "inconsistent result after apply" / perpetual diff.
	if val, ok := getResponseData["staaddresstype"]; ok && val != nil {
		data.Staaddresstype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["staserver"]; ok && val != nil {
		data.Staserver = types.StringValue(val.(string))
	}

	// ID is set once in Create (single unique attribute - plain value); do not recompute here.

	return data
}

// Datasource-only setter (Pattern 7): faithfully copies every field from the GET
// response and sets the ID, since a datasource has no prior plan/state to preserve
// and never calls Create.
func vpnglobal_staserver_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalStaserverBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalStaserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_staserver_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["staaddresstype"]; ok && val != nil {
		data.Staaddresstype = types.StringValue(val.(string))
	} else {
		data.Staaddresstype = types.StringNull()
	}
	if val, ok := getResponseData["staserver"]; ok && val != nil {
		data.Staserver = types.StringValue(val.(string))
	} else {
		data.Staserver = types.StringNull()
	}

	// Set ID for the datasource (single unique attribute - plain value).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Staserver.ValueString()))

	return data
}
