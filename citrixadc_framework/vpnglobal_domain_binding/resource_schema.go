package vpnglobal_domain_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetdomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The conflicting intranet domain name.",
			},
		},
	}
}

func vpnglobal_domain_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalDomainBindingResourceModel) vpn.Vpnglobaldomainbinding {
	tflog.Debug(ctx, "In vpnglobal_domain_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_domain_binding := vpn.Vpnglobaldomainbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_domain_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetdomain.IsNull() {
		vpnglobal_domain_binding.Intranetdomain = data.Intranetdomain.ValueString()
	}

	return vpnglobal_domain_binding
}

func vpnglobal_domain_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalDomainBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalDomainBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_domain_bindingSetAttrFromGet Function")

	// Convert API response to model
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Intranetdomain.ValueString())

	return data
}
