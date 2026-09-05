package authenticationemailaction

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func AuthenticationemailactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"content": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content to be delivered to the user. \"$code\" string within the content will be replaced with the actual one-time-code to be sent.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.",
			},
			"emailaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An optional expression that yields user's email. When not configured, user's default mail address would be used. When configured, result of this expression is used as destination email address.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new email action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"password": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Password/Clientsecret to use when authenticating to the server.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Address of the server that delivers the message. It is fully qualified fqdn such as http(s):// or smtp(s):// for http and smtp protocols respectively. For SMTP, the port number is mandatory like smtps://smtp.example.com:25.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time after which the code expires.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the email action. Default type is SMTP.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username/Clientid/EmailID to be used to authenticate to the server.",
			},
		},
	}
}

type AuthenticationemailactionDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Content                    types.String `tfsdk:"content"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Emailaddress               types.String `tfsdk:"emailaddress"`
	Name                       types.String `tfsdk:"name"`
	Password                   types.String `tfsdk:"password"`
	Serverurl                  types.String `tfsdk:"serverurl"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	Type                       types.String `tfsdk:"type"`
	Username                   types.String `tfsdk:"username"`
}

func authenticationemailactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationemailactionDataSourceModel, getResponseData map[string]interface{}) *AuthenticationemailactionDataSourceModel {
	tflog.Debug(ctx, "In authenticationemailactionDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["content"]; ok && val != nil {
		data.Content = types.StringValue(val.(string))
	} else {
		data.Content = types.StringNull()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["emailaddress"]; ok && val != nil {
		data.Emailaddress = types.StringValue(val.(string))
	} else {
		data.Emailaddress = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["serverurl"]; ok && val != nil {
		data.Serverurl = types.StringValue(val.(string))
	} else {
		data.Serverurl = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
