package authenticationtacacsaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationtacacsactionResourceModel describes the resource data model.
type AuthenticationtacacsactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Accounting                 types.String `tfsdk:"accounting"`
	Attribute1                 types.String `tfsdk:"attribute1"`
	Attribute10                types.String `tfsdk:"attribute10"`
	Attribute11                types.String `tfsdk:"attribute11"`
	Attribute12                types.String `tfsdk:"attribute12"`
	Attribute13                types.String `tfsdk:"attribute13"`
	Attribute14                types.String `tfsdk:"attribute14"`
	Attribute15                types.String `tfsdk:"attribute15"`
	Attribute16                types.String `tfsdk:"attribute16"`
	Attribute2                 types.String `tfsdk:"attribute2"`
	Attribute3                 types.String `tfsdk:"attribute3"`
	Attribute4                 types.String `tfsdk:"attribute4"`
	Attribute5                 types.String `tfsdk:"attribute5"`
	Attribute6                 types.String `tfsdk:"attribute6"`
	Attribute7                 types.String `tfsdk:"attribute7"`
	Attribute8                 types.String `tfsdk:"attribute8"`
	Attribute9                 types.String `tfsdk:"attribute9"`
	Attributes                 types.String `tfsdk:"attributes"`
	Auditfailedcmds            types.String `tfsdk:"auditfailedcmds"`
	Authorization              types.String `tfsdk:"authorization"`
	Authtimeout                types.Int64  `tfsdk:"authtimeout"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Groupattrname              types.String `tfsdk:"groupattrname"`
	Name                       types.String `tfsdk:"name"`
	Serverip                   types.String `tfsdk:"serverip"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Tacacssecret               types.String `tfsdk:"tacacssecret"`
}

func (r *AuthenticationtacacsactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationtacacsaction resource.",
			},
			"accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the TACACS+ server is currently accepting accounting messages.",
			},
			"attribute1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '1' (where '1' changes for each attribute)",
			},
			"attribute10": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '10' (where '10' changes for each attribute)",
			},
			"attribute11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '11' (where '11' changes for each attribute)",
			},
			"attribute12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '12' (where '12' changes for each attribute)",
			},
			"attribute13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '13' (where '13' changes for each attribute)",
			},
			"attribute14": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '14' (where '14' changes for each attribute)",
			},
			"attribute15": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '15' (where '15' changes for each attribute)",
			},
			"attribute16": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '16' (where '16' changes for each attribute)",
			},
			"attribute2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '2' (where '2' changes for each attribute)",
			},
			"attribute3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '3' (where '3' changes for each attribute)",
			},
			"attribute4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '4' (where '4' changes for each attribute)",
			},
			"attribute5": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '5' (where '5' changes for each attribute)",
			},
			"attribute6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '6' (where '6' changes for each attribute)",
			},
			"attribute7": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '7' (where '7' changes for each attribute)",
			},
			"attribute8": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '8' (where '8' changes for each attribute)",
			},
			"attribute9": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '9' (where '9' changes for each attribute)",
			},
			"attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of attribute names separated by ',' which needs to be fetched from tacacs server.\nNote that preceeding and trailing spaces will be removed.\nAttribute name can be 127 bytes and total length of this string should not cross 2047 bytes.\nThese attributes have multi-value support separated by ',' and stored as key-value pair in AAA session",
			},
			"auditfailedcmds": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the TACACS+ server that will receive accounting messages.",
			},
			"authorization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use streaming authorization on the TACACS+ server.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Number of seconds the Citrix ADC waits for a response from the TACACS+ server.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupattrname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TACACS+ group attribute name.\nUsed for group extraction on the TACACS+ server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the TACACS+ profile (action).\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after TACACS profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'y authentication action').",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address assigned to the TACACS+ server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(49),
				Description: "Port number on which the TACACS+ server listens for connections.",
			},
			"tacacssecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Key shared between the TACACS+ server and the Citrix ADC.\nRequired for allowing the Citrix ADC to communicate with the TACACS+ server.",
			},
		},
	}
}

func authenticationtacacsactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationtacacsactionResourceModel) authentication.Authenticationtacacsaction {
	tflog.Debug(ctx, "In authenticationtacacsactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationtacacsaction := authentication.Authenticationtacacsaction{}
	if !data.Accounting.IsNull() {
		authenticationtacacsaction.Accounting = data.Accounting.ValueString()
	}
	if !data.Attribute1.IsNull() {
		authenticationtacacsaction.Attribute1 = data.Attribute1.ValueString()
	}
	if !data.Attribute10.IsNull() {
		authenticationtacacsaction.Attribute10 = data.Attribute10.ValueString()
	}
	if !data.Attribute11.IsNull() {
		authenticationtacacsaction.Attribute11 = data.Attribute11.ValueString()
	}
	if !data.Attribute12.IsNull() {
		authenticationtacacsaction.Attribute12 = data.Attribute12.ValueString()
	}
	if !data.Attribute13.IsNull() {
		authenticationtacacsaction.Attribute13 = data.Attribute13.ValueString()
	}
	if !data.Attribute14.IsNull() {
		authenticationtacacsaction.Attribute14 = data.Attribute14.ValueString()
	}
	if !data.Attribute15.IsNull() {
		authenticationtacacsaction.Attribute15 = data.Attribute15.ValueString()
	}
	if !data.Attribute16.IsNull() {
		authenticationtacacsaction.Attribute16 = data.Attribute16.ValueString()
	}
	if !data.Attribute2.IsNull() {
		authenticationtacacsaction.Attribute2 = data.Attribute2.ValueString()
	}
	if !data.Attribute3.IsNull() {
		authenticationtacacsaction.Attribute3 = data.Attribute3.ValueString()
	}
	if !data.Attribute4.IsNull() {
		authenticationtacacsaction.Attribute4 = data.Attribute4.ValueString()
	}
	if !data.Attribute5.IsNull() {
		authenticationtacacsaction.Attribute5 = data.Attribute5.ValueString()
	}
	if !data.Attribute6.IsNull() {
		authenticationtacacsaction.Attribute6 = data.Attribute6.ValueString()
	}
	if !data.Attribute7.IsNull() {
		authenticationtacacsaction.Attribute7 = data.Attribute7.ValueString()
	}
	if !data.Attribute8.IsNull() {
		authenticationtacacsaction.Attribute8 = data.Attribute8.ValueString()
	}
	if !data.Attribute9.IsNull() {
		authenticationtacacsaction.Attribute9 = data.Attribute9.ValueString()
	}
	if !data.Attributes.IsNull() {
		authenticationtacacsaction.Attributes = data.Attributes.ValueString()
	}
	if !data.Auditfailedcmds.IsNull() {
		authenticationtacacsaction.Auditfailedcmds = data.Auditfailedcmds.ValueString()
	}
	if !data.Authorization.IsNull() {
		authenticationtacacsaction.Authorization = data.Authorization.ValueString()
	}
	if !data.Authtimeout.IsNull() {
		authenticationtacacsaction.Authtimeout = utils.IntPtr(int(data.Authtimeout.ValueInt64()))
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationtacacsaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Groupattrname.IsNull() {
		authenticationtacacsaction.Groupattrname = data.Groupattrname.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationtacacsaction.Name = data.Name.ValueString()
	}
	if !data.Serverip.IsNull() {
		authenticationtacacsaction.Serverip = data.Serverip.ValueString()
	}
	if !data.Serverport.IsNull() {
		authenticationtacacsaction.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Tacacssecret.IsNull() {
		authenticationtacacsaction.Tacacssecret = data.Tacacssecret.ValueString()
	}

	return authenticationtacacsaction
}

func authenticationtacacsactionSetAttrFromGet(ctx context.Context, data *AuthenticationtacacsactionResourceModel, getResponseData map[string]interface{}) *AuthenticationtacacsactionResourceModel {
	tflog.Debug(ctx, "In authenticationtacacsactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["accounting"]; ok && val != nil {
		data.Accounting = types.StringValue(val.(string))
	} else {
		data.Accounting = types.StringNull()
	}
	if val, ok := getResponseData["attribute1"]; ok && val != nil {
		data.Attribute1 = types.StringValue(val.(string))
	} else {
		data.Attribute1 = types.StringNull()
	}
	if val, ok := getResponseData["attribute10"]; ok && val != nil {
		data.Attribute10 = types.StringValue(val.(string))
	} else {
		data.Attribute10 = types.StringNull()
	}
	if val, ok := getResponseData["attribute11"]; ok && val != nil {
		data.Attribute11 = types.StringValue(val.(string))
	} else {
		data.Attribute11 = types.StringNull()
	}
	if val, ok := getResponseData["attribute12"]; ok && val != nil {
		data.Attribute12 = types.StringValue(val.(string))
	} else {
		data.Attribute12 = types.StringNull()
	}
	if val, ok := getResponseData["attribute13"]; ok && val != nil {
		data.Attribute13 = types.StringValue(val.(string))
	} else {
		data.Attribute13 = types.StringNull()
	}
	if val, ok := getResponseData["attribute14"]; ok && val != nil {
		data.Attribute14 = types.StringValue(val.(string))
	} else {
		data.Attribute14 = types.StringNull()
	}
	if val, ok := getResponseData["attribute15"]; ok && val != nil {
		data.Attribute15 = types.StringValue(val.(string))
	} else {
		data.Attribute15 = types.StringNull()
	}
	if val, ok := getResponseData["attribute16"]; ok && val != nil {
		data.Attribute16 = types.StringValue(val.(string))
	} else {
		data.Attribute16 = types.StringNull()
	}
	if val, ok := getResponseData["attribute2"]; ok && val != nil {
		data.Attribute2 = types.StringValue(val.(string))
	} else {
		data.Attribute2 = types.StringNull()
	}
	if val, ok := getResponseData["attribute3"]; ok && val != nil {
		data.Attribute3 = types.StringValue(val.(string))
	} else {
		data.Attribute3 = types.StringNull()
	}
	if val, ok := getResponseData["attribute4"]; ok && val != nil {
		data.Attribute4 = types.StringValue(val.(string))
	} else {
		data.Attribute4 = types.StringNull()
	}
	if val, ok := getResponseData["attribute5"]; ok && val != nil {
		data.Attribute5 = types.StringValue(val.(string))
	} else {
		data.Attribute5 = types.StringNull()
	}
	if val, ok := getResponseData["attribute6"]; ok && val != nil {
		data.Attribute6 = types.StringValue(val.(string))
	} else {
		data.Attribute6 = types.StringNull()
	}
	if val, ok := getResponseData["attribute7"]; ok && val != nil {
		data.Attribute7 = types.StringValue(val.(string))
	} else {
		data.Attribute7 = types.StringNull()
	}
	if val, ok := getResponseData["attribute8"]; ok && val != nil {
		data.Attribute8 = types.StringValue(val.(string))
	} else {
		data.Attribute8 = types.StringNull()
	}
	if val, ok := getResponseData["attribute9"]; ok && val != nil {
		data.Attribute9 = types.StringValue(val.(string))
	} else {
		data.Attribute9 = types.StringNull()
	}
	if val, ok := getResponseData["attributes"]; ok && val != nil {
		data.Attributes = types.StringValue(val.(string))
	} else {
		data.Attributes = types.StringNull()
	}
	if val, ok := getResponseData["auditfailedcmds"]; ok && val != nil {
		data.Auditfailedcmds = types.StringValue(val.(string))
	} else {
		data.Auditfailedcmds = types.StringNull()
	}
	if val, ok := getResponseData["authorization"]; ok && val != nil {
		data.Authorization = types.StringValue(val.(string))
	} else {
		data.Authorization = types.StringNull()
	}
	if val, ok := getResponseData["authtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Authtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Authtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["groupattrname"]; ok && val != nil {
		data.Groupattrname = types.StringValue(val.(string))
	} else {
		data.Groupattrname = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["serverip"]; ok && val != nil {
		data.Serverip = types.StringValue(val.(string))
	} else {
		data.Serverip = types.StringNull()
	}
	if val, ok := getResponseData["serverport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverport = types.Int64Value(intVal)
		}
	} else {
		data.Serverport = types.Int64Null()
	}
	if val, ok := getResponseData["tacacssecret"]; ok && val != nil {
		data.Tacacssecret = types.StringValue(val.(string))
	} else {
		data.Tacacssecret = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
