package vpnglobal_staserver_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"staaddresstype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the STA server address(ipv4/v6).",
			},
			"staserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configured Secure Ticketing Authority (STA) server.",
			},
		},
	}
}

func vpnglobal_staserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalStaserverBindingResourceModel) vpn.Vpnglobalstaserverbinding {
	tflog.Debug(ctx, "In vpnglobal_staserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_staserver_binding := vpn.Vpnglobalstaserverbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_staserver_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Staaddresstype.IsNull() {
		vpnglobal_staserver_binding.Staaddresstype = data.Staaddresstype.ValueString()
	}
	if !data.Staserver.IsNull() {
		vpnglobal_staserver_binding.Staserver = data.Staserver.ValueString()
	}

	return vpnglobal_staserver_binding
}

func vpnglobal_staserver_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalStaserverBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalStaserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_staserver_bindingSetAttrFromGet Function")

	// Convert API response to model
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Staserver.ValueString())

	return data
}
