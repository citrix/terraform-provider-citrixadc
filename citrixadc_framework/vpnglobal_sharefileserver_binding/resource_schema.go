package vpnglobal_sharefileserver_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalSharefileserverBindingResourceModel describes the resource data model.
type VpnglobalSharefileserverBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Sharefile              types.String `tfsdk:"sharefile"`
}

func (r *VpnglobalSharefileserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_sharefileserver_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"sharefile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configured Sharefile server, in the format IP:PORT / FQDN:PORT",
			},
		},
	}
}

func vpnglobal_sharefileserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalSharefileserverBindingResourceModel) vpn.Vpnglobalsharefileserverbinding {
	tflog.Debug(ctx, "In vpnglobal_sharefileserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_sharefileserver_binding := vpn.Vpnglobalsharefileserverbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_sharefileserver_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Sharefile.IsNull() {
		vpnglobal_sharefileserver_binding.Sharefile = data.Sharefile.ValueString()
	}

	return vpnglobal_sharefileserver_binding
}

func vpnglobal_sharefileserver_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalSharefileserverBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalSharefileserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_sharefileserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["sharefile"]; ok && val != nil {
		data.Sharefile = types.StringValue(val.(string))
	} else {
		data.Sharefile = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Sharefile.ValueString())

	return data
}
