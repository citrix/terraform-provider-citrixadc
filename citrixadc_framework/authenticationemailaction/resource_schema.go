package authenticationemailaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationemailactionResourceModel describes the resource data model.
type AuthenticationemailactionResourceModel struct {
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

func (r *AuthenticationemailactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationemailaction resource.",
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
				Required:    true,
				Description: "Password/Clientsecret to use when authenticating to the server.",
			},
			"serverurl": schema.StringAttribute{
				Required:    true,
				Description: "Address of the server that delivers the message. It is fully qualified fqdn such as http(s):// or smtp(s):// for http and smtp protocols respectively. For SMTP, the port number is mandatory like smtps://smtp.example.com:25.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(180),
				Description: "Time after which the code expires.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SMTP"),
				Description: "Type of the email action. Default type is SMTP.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Username/Clientid/EmailID to be used to authenticate to the server.",
			},
		},
	}
}

func authenticationemailactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationemailactionResourceModel) authentication.Authenticationemailaction {
	tflog.Debug(ctx, "In authenticationemailactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationemailaction := authentication.Authenticationemailaction{}
	if !data.Content.IsNull() {
		authenticationemailaction.Content = data.Content.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationemailaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Emailaddress.IsNull() {
		authenticationemailaction.Emailaddress = data.Emailaddress.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationemailaction.Name = data.Name.ValueString()
	}
	if !data.Password.IsNull() {
		authenticationemailaction.Password = data.Password.ValueString()
	}
	if !data.Serverurl.IsNull() {
		authenticationemailaction.Serverurl = data.Serverurl.ValueString()
	}
	if !data.Timeout.IsNull() {
		authenticationemailaction.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Type.IsNull() {
		authenticationemailaction.Type = data.Type.ValueString()
	}
	if !data.Username.IsNull() {
		authenticationemailaction.Username = data.Username.ValueString()
	}

	return authenticationemailaction
}

func authenticationemailactionSetAttrFromGet(ctx context.Context, data *AuthenticationemailactionResourceModel, getResponseData map[string]interface{}) *AuthenticationemailactionResourceModel {
	tflog.Debug(ctx, "In authenticationemailactionSetAttrFromGet Function")

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
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["serverurl"]; ok && val != nil {
		data.Serverurl = types.StringValue(val.(string))
	} else {
		data.Serverurl = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
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
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
