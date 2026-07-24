package authenticationpushservice

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationpushserviceResourceModel describes the resource data model.
type AuthenticationpushserviceResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Clientid              types.String `tfsdk:"clientid"`
	Clientsecret          types.String `tfsdk:"clientsecret"`
	ClientsecretWo        types.String `tfsdk:"clientsecret_wo"`
	ClientsecretWoVersion types.Int64  `tfsdk:"clientsecret_wo_version"`
	Customerid            types.String `tfsdk:"customerid"`
	Name                  types.String `tfsdk:"name"`
	Refreshinterval       types.Int64  `tfsdk:"refreshinterval"`
}

func (r *AuthenticationpushserviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationpushservice resource.",
			},
			"clientid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique identity for communicating with Citrix Push server in cloud.",
			},
			"clientsecret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Unique secret for communicating with Citrix Push server in cloud.",
			},
			"clientsecret_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Unique secret for communicating with Citrix Push server in cloud.",
			},
			"clientsecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a clientsecret_wo update.",
			},
			"customerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Customer id/name of the account in cloud that is used to create clientid/secret pair.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the push service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my push service\" or 'my push service').",
			},
			"refreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(50),
				Description: "Interval at which certificates or idtoken is refreshed.",
			},
		},
	}
}

func authenticationpushserviceGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationpushserviceResourceModel) authentication.Authenticationpushservice {
	tflog.Debug(ctx, "In authenticationpushserviceGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationpushservice := authentication.Authenticationpushservice{}
	if !data.Clientid.IsNull() && !data.Clientid.IsUnknown() {
		authenticationpushservice.Clientid = data.Clientid.ValueString()
	}
	if !data.Clientsecret.IsNull() && !data.Clientsecret.IsUnknown() {
		authenticationpushservice.Clientsecret = data.Clientsecret.ValueString()
	}
	// Skip write-only attribute: clientsecret_wo
	// Skip version tracker attribute: clientsecret_wo_version
	if !data.Customerid.IsNull() && !data.Customerid.IsUnknown() {
		authenticationpushservice.Customerid = data.Customerid.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationpushservice.Name = data.Name.ValueString()
	}
	if !data.Refreshinterval.IsNull() && !data.Refreshinterval.IsUnknown() {
		authenticationpushservice.Refreshinterval = utils.IntPtr(int(data.Refreshinterval.ValueInt64()))
	}

	return authenticationpushservice
}

func authenticationpushserviceGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationpushserviceResourceModel, payload *authentication.Authenticationpushservice) {
	tflog.Debug(ctx, "In authenticationpushserviceGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: clientsecret_wo -> clientsecret
	if !data.ClientsecretWo.IsNull() {
		clientsecretWo := data.ClientsecretWo.ValueString()
		if clientsecretWo != "" {
			payload.Clientsecret = clientsecretWo
		}
	}
}

func authenticationpushserviceSetAttrFromGet(ctx context.Context, data *AuthenticationpushserviceResourceModel, getResponseData map[string]interface{}) *AuthenticationpushserviceResourceModel {
	tflog.Debug(ctx, "In authenticationpushserviceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientid"]; ok && val != nil {
		data.Clientid = types.StringValue(val.(string))
	} else {
		data.Clientid = types.StringNull()
	}
	// clientsecret is not returned by NITRO API (secret/ephemeral) - retain from config
	// clientsecret_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// clientsecret_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["customerid"]; ok && val != nil {
		data.Customerid = types.StringValue(val.(string))
	} else {
		data.Customerid = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["refreshinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Refreshinterval = types.Int64Value(intVal)
		}
	} else {
		data.Refreshinterval = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
