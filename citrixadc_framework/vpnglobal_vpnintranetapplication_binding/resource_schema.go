package vpnglobal_vpnintranetapplication_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalVpnintranetapplicationBindingResourceModel describes the resource data model.
type VpnglobalVpnintranetapplicationBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Intranetapplication    types.String `tfsdk:"intranetapplication"`
}

func (r *VpnglobalVpnintranetapplicationBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_vpnintranetapplication_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetapplication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet vpn application.",
			},
		},
	}
}

func vpnglobal_vpnintranetapplication_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalVpnintranetapplicationBindingResourceModel) vpn.Vpnglobalvpnintranetapplicationbinding {
	tflog.Debug(ctx, "In vpnglobal_vpnintranetapplication_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_vpnintranetapplication_binding := vpn.Vpnglobalvpnintranetapplicationbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_vpnintranetapplication_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetapplication.IsNull() {
		vpnglobal_vpnintranetapplication_binding.Intranetapplication = data.Intranetapplication.ValueString()
	}

	return vpnglobal_vpnintranetapplication_binding
}

func vpnglobal_vpnintranetapplication_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalVpnintranetapplicationBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnintranetapplicationBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnintranetapplication_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["intranetapplication"]; ok && val != nil {
		data.Intranetapplication = types.StringValue(val.(string))
	} else {
		data.Intranetapplication = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Intranetapplication.ValueString())

	return data
}
