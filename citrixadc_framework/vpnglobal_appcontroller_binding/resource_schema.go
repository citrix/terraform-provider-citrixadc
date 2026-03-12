package vpnglobal_appcontroller_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalAppcontrollerBindingResourceModel describes the resource data model.
type VpnglobalAppcontrollerBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Appcontroller          types.String `tfsdk:"appcontroller"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
}

func (r *VpnglobalAppcontrollerBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_appcontroller_binding resource.",
			},
			"appcontroller": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configured App Controller server.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
		},
	}
}

func vpnglobal_appcontroller_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalAppcontrollerBindingResourceModel) vpn.Vpnglobalappcontrollerbinding {
	tflog.Debug(ctx, "In vpnglobal_appcontroller_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_appcontroller_binding := vpn.Vpnglobalappcontrollerbinding{}
	if !data.Appcontroller.IsNull() {
		vpnglobal_appcontroller_binding.Appcontroller = data.Appcontroller.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_appcontroller_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}

	return vpnglobal_appcontroller_binding
}

func vpnglobal_appcontroller_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalAppcontrollerBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAppcontrollerBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_appcontroller_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appcontroller"]; ok && val != nil {
		data.Appcontroller = types.StringValue(val.(string))
	} else {
		data.Appcontroller = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Appcontroller.ValueString())

	return data
}
