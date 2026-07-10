package vpnglobal_domain_binding

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

// VpnglobalDomainBindingResourceModel describes the resource data model.
type VpnglobalDomainBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Intranetdomain         types.String `tfsdk:"intranetdomain"`
}

func (r *VpnglobalDomainBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_domain_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetdomain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The conflicting intranet domain name.",
			},
		},
	}
}

func vpnglobal_domain_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalDomainBindingResourceModel) vpn.Vpnglobaldomainbinding {
	tflog.Debug(ctx, "In vpnglobal_domain_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_domain_binding := vpn.Vpnglobaldomainbinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_domain_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetdomain.IsNull() && !data.Intranetdomain.IsUnknown() {
		vpnglobal_domain_binding.Intranetdomain = data.Intranetdomain.ValueString()
	}

	return vpnglobal_domain_binding
}

// vpnglobal_domain_bindingSetAttrFromGet is the RESOURCE-side setter. The NITRO
// GET response for this binding does NOT echo back gotopriorityexpression
// (it only returns intranetdomain + stateflag), so we must preserve the existing
// plan/state value for gotopriorityexpression instead of nulling it out (Pattern 7).
// The ID is set once in Create and must not be recomputed here.
func vpnglobal_domain_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalDomainBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalDomainBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_domain_bindingSetAttrFromGet Function")

	// gotopriorityexpression is not returned by NITRO GET.
	// It is Optional+Computed: when the user omitted it, the planned value is
	// unknown ("known after apply"); since GET never echoes it back we must
	// resolve that unknown to a concrete value (null) so the apply does not fail
	// with "still indicated an unknown value". When the user DID set it, the
	// value is known and we preserve it (do not null it out).
	if data.Gotopriorityexpression.IsUnknown() {
		data.Gotopriorityexpression = types.StringNull()
	}

	if val, ok := getResponseData["intranetdomain"]; ok && val != nil {
		data.Intranetdomain = types.StringValue(val.(string))
	}

	return data
}

// vpnglobal_domain_bindingSetAttrFromGetForDatasource is the DATASOURCE-side setter.
// A datasource has no prior plan/state to preserve, so it faithfully copies every
// field from the GET response and sets the ID itself (Pattern 7 datasource split).
func vpnglobal_domain_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalDomainBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalDomainBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_domain_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["intranetdomain"]; ok && val != nil {
		data.Intranetdomain = types.StringValue(val.(string))
	} else {
		data.Intranetdomain = types.StringNull()
	}

	// Set ID for the datasource (no Create to set it).
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Intranetdomain.ValueString()))

	return data
}
