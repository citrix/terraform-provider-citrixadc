package vpnglobal_vpnintranetapplication_binding

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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetapplication": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The intranet vpn application.",
			},
		},
	}
}

func vpnglobal_vpnintranetapplication_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalVpnintranetapplicationBindingResourceModel) vpn.Vpnglobalvpnintranetapplicationbinding {
	tflog.Debug(ctx, "In vpnglobal_vpnintranetapplication_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_vpnintranetapplication_binding := vpn.Vpnglobalvpnintranetapplicationbinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_vpnintranetapplication_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetapplication.IsNull() && !data.Intranetapplication.IsUnknown() {
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
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Intranetapplication.ValueString()))

	return data
}
