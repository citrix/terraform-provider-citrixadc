package authenticationvserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationvserverResourceModel describes the resource data model.
type AuthenticationvserverResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Appflowlog           types.String `tfsdk:"appflowlog"`
	Authentication       types.String `tfsdk:"authentication"`
	Authenticationdomain types.String `tfsdk:"authenticationdomain"`
	Certkeynames         types.String `tfsdk:"certkeynames"`
	Comment              types.String `tfsdk:"comment"`
	Failedlogintimeout   types.Int64  `tfsdk:"failedlogintimeout"`
	Ipv46                types.String `tfsdk:"ipv46"`
	Maxloginattempts     types.Int64  `tfsdk:"maxloginattempts"`
	Name                 types.String `tfsdk:"name"`
	Newname              types.String `tfsdk:"newname"`
	Port                 types.Int64  `tfsdk:"port"`
	Range                types.Int64  `tfsdk:"range"`
	Samesite             types.String `tfsdk:"samesite"`
	Servicetype          types.String `tfsdk:"servicetype"`
	State                types.String `tfsdk:"state"`
	Td                   types.Int64  `tfsdk:"td"`
}

func (r *AuthenticationvserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationvserver resource.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Log AppFlow flow information.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Require users to be authenticated before sending traffic through this virtual server.",
			},
			"authenticationdomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The domain of the authentication cookie set by Authentication vserver",
			},
			"certkeynames": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with this virtual server.",
			},
			"failedlogintimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of minutes an account will be locked if user exceeds maximum permissible attempts",
			},
			"ipv46": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the authentication virtual server, if a single IP address is assigned to the virtual server.",
			},
			"maxloginattempts": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum Number of login Attempts",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new authentication virtual server.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the authentication virtual server is added by using the rename authentication vserver command.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name of the authentication virtual server.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, 'my authentication policy' or \"my authentication policy\").",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "TCP port on which the virtual server accepts connections.",
			},
			"range": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "If you are creating a series of virtual servers with a range of IP addresses assigned to them, the length of the range.\nThe new range of authentication virtual servers will have IP addresses consecutively numbered, starting with the primary address specified with the IP Address parameter.",
			},
			"samesite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("SSL"),
				Description: "Protocol type of the authentication virtual server. Always SSL.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Initial state of the new virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func authenticationvserverGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationvserverResourceModel) authentication.Authenticationvserver {
	tflog.Debug(ctx, "In authenticationvserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationvserver := authentication.Authenticationvserver{}
	if !data.Appflowlog.IsNull() {
		authenticationvserver.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Authentication.IsNull() {
		authenticationvserver.Authentication = data.Authentication.ValueString()
	}
	if !data.Authenticationdomain.IsNull() {
		authenticationvserver.Authenticationdomain = data.Authenticationdomain.ValueString()
	}
	if !data.Certkeynames.IsNull() {
		authenticationvserver.Certkeynames = data.Certkeynames.ValueString()
	}
	if !data.Comment.IsNull() {
		authenticationvserver.Comment = data.Comment.ValueString()
	}
	if !data.Failedlogintimeout.IsNull() {
		authenticationvserver.Failedlogintimeout = utils.IntPtr(int(data.Failedlogintimeout.ValueInt64()))
	}
	if !data.Ipv46.IsNull() {
		authenticationvserver.Ipv46 = data.Ipv46.ValueString()
	}
	if !data.Maxloginattempts.IsNull() {
		authenticationvserver.Maxloginattempts = utils.IntPtr(int(data.Maxloginattempts.ValueInt64()))
	}
	if !data.Name.IsNull() {
		authenticationvserver.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		authenticationvserver.Newname = data.Newname.ValueString()
	}
	if !data.Port.IsNull() {
		authenticationvserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Range.IsNull() {
		authenticationvserver.Range = utils.IntPtr(int(data.Range.ValueInt64()))
	}
	if !data.Samesite.IsNull() {
		authenticationvserver.Samesite = data.Samesite.ValueString()
	}
	if !data.Servicetype.IsNull() {
		authenticationvserver.Servicetype = data.Servicetype.ValueString()
	}
	if !data.State.IsNull() {
		authenticationvserver.State = data.State.ValueString()
	}
	if !data.Td.IsNull() {
		authenticationvserver.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return authenticationvserver
}

func authenticationvserverSetAttrFromGet(ctx context.Context, data *AuthenticationvserverResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverResourceModel {
	tflog.Debug(ctx, "In authenticationvserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["authenticationdomain"]; ok && val != nil {
		data.Authenticationdomain = types.StringValue(val.(string))
	} else {
		data.Authenticationdomain = types.StringNull()
	}
	if val, ok := getResponseData["certkeynames"]; ok && val != nil {
		data.Certkeynames = types.StringValue(val.(string))
	} else {
		data.Certkeynames = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["failedlogintimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Failedlogintimeout = types.Int64Value(intVal)
		}
	} else {
		data.Failedlogintimeout = types.Int64Null()
	}
	if val, ok := getResponseData["ipv46"]; ok && val != nil {
		data.Ipv46 = types.StringValue(val.(string))
	} else {
		data.Ipv46 = types.StringNull()
	}
	if val, ok := getResponseData["maxloginattempts"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxloginattempts = types.Int64Value(intVal)
		}
	} else {
		data.Maxloginattempts = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["range"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Range = types.Int64Value(intVal)
		}
	} else {
		data.Range = types.Int64Null()
	}
	if val, ok := getResponseData["samesite"]; ok && val != nil {
		data.Samesite = types.StringValue(val.(string))
	} else {
		data.Samesite = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
