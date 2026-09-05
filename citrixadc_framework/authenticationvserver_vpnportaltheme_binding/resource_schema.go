package authenticationvserver_vpnportaltheme_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationvserverVpnportalthemeBindingResourceModel describes the resource data model.
type AuthenticationvserverVpnportalthemeBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Portaltheme types.String `tfsdk:"portaltheme"`
}

func (r *AuthenticationvserverVpnportalthemeBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationvserver_vpnportaltheme_binding resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the authentication virtual server to which to bind the policy.",
			},
			"portaltheme": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Theme for Authentication virtual server Login portal",
			},
		},
	}
}

func authenticationvserver_vpnportaltheme_bindingGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationvserverVpnportalthemeBindingResourceModel) authentication.Authenticationvservervpnportalthemebinding {
	tflog.Debug(ctx, "In authenticationvserver_vpnportaltheme_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationvserver_vpnportaltheme_binding := authentication.Authenticationvservervpnportalthemebinding{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationvserver_vpnportaltheme_binding.Name = data.Name.ValueString()
	}
	if !data.Portaltheme.IsNull() && !data.Portaltheme.IsUnknown() {
		authenticationvserver_vpnportaltheme_binding.Portaltheme = data.Portaltheme.ValueString()
	}

	return authenticationvserver_vpnportaltheme_binding
}

func authenticationvserver_vpnportaltheme_bindingSetAttrFromGet(ctx context.Context, data *AuthenticationvserverVpnportalthemeBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverVpnportalthemeBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_vpnportaltheme_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["portaltheme"]; ok && val != nil {
		data.Portaltheme = types.StringValue(val.(string))
	} else {
		data.Portaltheme = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("portaltheme:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Portaltheme.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
