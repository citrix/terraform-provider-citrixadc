package authenticationazurekeyvault

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func AuthenticationazurekeyvaultDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If authentication is disabled, otp checks are not performed after azure vault keys are obtained. This is useful to distinguish whether user has registered devices.",
			},
			"clientid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique identity of the relying party requesting for authentication.",
			},
			"clientsecret": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Unique secret string to authorize relying party at authorization server.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new Azure Key Vault profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"pushservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service used to send push notifications",
			},
			"refreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval at which access token in obtained.",
			},
			"servicekeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Friendly name of the Key to be used to compute signature.",
			},
			"signaturealg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm to be used to sign/verify transactions",
			},
			"tenantid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TenantID of the application. This is usually specific to providers such as Microsoft and usually refers to the deployment identifier.",
			},
			"tokenendpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL endpoint on relying party to which the OAuth token is to be sent.",
			},
			"vaultname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Azure vault account as configured in azure portal.",
			},
		},
	}
}

type AuthenticationazurekeyvaultDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Authentication             types.String `tfsdk:"authentication"`
	Clientid                   types.String `tfsdk:"clientid"`
	Clientsecret               types.String `tfsdk:"clientsecret"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Name                       types.String `tfsdk:"name"`
	Pushservice                types.String `tfsdk:"pushservice"`
	Refreshinterval            types.Int64  `tfsdk:"refreshinterval"`
	Servicekeyname             types.String `tfsdk:"servicekeyname"`
	Signaturealg               types.String `tfsdk:"signaturealg"`
	Tenantid                   types.String `tfsdk:"tenantid"`
	Tokenendpoint              types.String `tfsdk:"tokenendpoint"`
	Vaultname                  types.String `tfsdk:"vaultname"`
}

func authenticationazurekeyvaultDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationazurekeyvaultDataSourceModel, getResponseData map[string]interface{}) *AuthenticationazurekeyvaultDataSourceModel {
	tflog.Debug(ctx, "In authenticationazurekeyvaultDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["clientid"]; ok && val != nil {
		data.Clientid = types.StringValue(val.(string))
	} else {
		data.Clientid = types.StringNull()
	}
	// clientsecret is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["pushservice"]; ok && val != nil {
		data.Pushservice = types.StringValue(val.(string))
	} else {
		data.Pushservice = types.StringNull()
	}
	if val, ok := getResponseData["refreshinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Refreshinterval = types.Int64Value(intVal)
		}
	} else {
		data.Refreshinterval = types.Int64Null()
	}
	if val, ok := getResponseData["servicekeyname"]; ok && val != nil {
		data.Servicekeyname = types.StringValue(val.(string))
	} else {
		data.Servicekeyname = types.StringNull()
	}
	if val, ok := getResponseData["signaturealg"]; ok && val != nil {
		data.Signaturealg = types.StringValue(val.(string))
	} else {
		data.Signaturealg = types.StringNull()
	}
	if val, ok := getResponseData["tenantid"]; ok && val != nil {
		data.Tenantid = types.StringValue(val.(string))
	} else {
		data.Tenantid = types.StringNull()
	}
	if val, ok := getResponseData["tokenendpoint"]; ok && val != nil {
		data.Tokenendpoint = types.StringValue(val.(string))
	} else {
		data.Tokenendpoint = types.StringNull()
	}
	if val, ok := getResponseData["vaultname"]; ok && val != nil {
		data.Vaultname = types.StringValue(val.(string))
	} else {
		data.Vaultname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
