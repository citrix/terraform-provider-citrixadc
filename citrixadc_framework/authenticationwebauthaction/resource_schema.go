package authenticationwebauthaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationwebauthactionResourceModel describes the resource data model.
type AuthenticationwebauthactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
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
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Fullreqexpr                types.String `tfsdk:"fullreqexpr"`
	Name                       types.String `tfsdk:"name"`
	Scheme                     types.String `tfsdk:"scheme"`
	Serverip                   types.String `tfsdk:"serverip"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Successrule                types.String `tfsdk:"successrule"`
}

func (r *AuthenticationwebauthactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationwebauthaction resource.",
			},
			"attribute1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute1 from the webauth response",
			},
			"attribute10": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute10 from the webauth response",
			},
			"attribute11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute11 from the webauth response",
			},
			"attribute12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute12 from the webauth response",
			},
			"attribute13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute13 from the webauth response",
			},
			"attribute14": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute14 from the webauth response",
			},
			"attribute15": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute15 from the webauth response",
			},
			"attribute16": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute16 from the webauth response",
			},
			"attribute2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute2 from the webauth response",
			},
			"attribute3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute3 from the webauth response",
			},
			"attribute4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute4 from the webauth response",
			},
			"attribute5": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute5 from the webauth response",
			},
			"attribute6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute6 from the webauth response",
			},
			"attribute7": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute7 from the webauth response",
			},
			"attribute8": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute8 from the webauth response",
			},
			"attribute9": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute9 from the webauth response",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"fullreqexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exact HTTP request, in the form of an expression, which the Citrix ADC sends to the authentication server.\nThe Citrix ADC does not check the validity of this request. One must manually validate the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Web Authentication action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"scheme": schema.StringAttribute{
				Required:    true,
				Description: "Type of scheme for the web server.",
			},
			"serverip": schema.StringAttribute{
				Required:    true,
				Description: "IP address of the web server to be used for authentication.",
			},
			"serverport": schema.Int64Attribute{
				Required:    true,
				Description: "Port on which the web server accepts connections.",
			},
			"successrule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, that checks to see if authentication is successful.",
			},
		},
	}
}

func authenticationwebauthactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationwebauthactionResourceModel) authentication.Authenticationwebauthaction {
	tflog.Debug(ctx, "In authenticationwebauthactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationwebauthaction := authentication.Authenticationwebauthaction{}
	if !data.Attribute1.IsNull() {
		authenticationwebauthaction.Attribute1 = data.Attribute1.ValueString()
	}
	if !data.Attribute10.IsNull() {
		authenticationwebauthaction.Attribute10 = data.Attribute10.ValueString()
	}
	if !data.Attribute11.IsNull() {
		authenticationwebauthaction.Attribute11 = data.Attribute11.ValueString()
	}
	if !data.Attribute12.IsNull() {
		authenticationwebauthaction.Attribute12 = data.Attribute12.ValueString()
	}
	if !data.Attribute13.IsNull() {
		authenticationwebauthaction.Attribute13 = data.Attribute13.ValueString()
	}
	if !data.Attribute14.IsNull() {
		authenticationwebauthaction.Attribute14 = data.Attribute14.ValueString()
	}
	if !data.Attribute15.IsNull() {
		authenticationwebauthaction.Attribute15 = data.Attribute15.ValueString()
	}
	if !data.Attribute16.IsNull() {
		authenticationwebauthaction.Attribute16 = data.Attribute16.ValueString()
	}
	if !data.Attribute2.IsNull() {
		authenticationwebauthaction.Attribute2 = data.Attribute2.ValueString()
	}
	if !data.Attribute3.IsNull() {
		authenticationwebauthaction.Attribute3 = data.Attribute3.ValueString()
	}
	if !data.Attribute4.IsNull() {
		authenticationwebauthaction.Attribute4 = data.Attribute4.ValueString()
	}
	if !data.Attribute5.IsNull() {
		authenticationwebauthaction.Attribute5 = data.Attribute5.ValueString()
	}
	if !data.Attribute6.IsNull() {
		authenticationwebauthaction.Attribute6 = data.Attribute6.ValueString()
	}
	if !data.Attribute7.IsNull() {
		authenticationwebauthaction.Attribute7 = data.Attribute7.ValueString()
	}
	if !data.Attribute8.IsNull() {
		authenticationwebauthaction.Attribute8 = data.Attribute8.ValueString()
	}
	if !data.Attribute9.IsNull() {
		authenticationwebauthaction.Attribute9 = data.Attribute9.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationwebauthaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Fullreqexpr.IsNull() {
		authenticationwebauthaction.Fullreqexpr = data.Fullreqexpr.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationwebauthaction.Name = data.Name.ValueString()
	}
	if !data.Scheme.IsNull() {
		authenticationwebauthaction.Scheme = data.Scheme.ValueString()
	}
	if !data.Serverip.IsNull() {
		authenticationwebauthaction.Serverip = data.Serverip.ValueString()
	}
	if !data.Serverport.IsNull() {
		authenticationwebauthaction.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Successrule.IsNull() {
		authenticationwebauthaction.Successrule = data.Successrule.ValueString()
	}

	return authenticationwebauthaction
}

func authenticationwebauthactionSetAttrFromGet(ctx context.Context, data *AuthenticationwebauthactionResourceModel, getResponseData map[string]interface{}) *AuthenticationwebauthactionResourceModel {
	tflog.Debug(ctx, "In authenticationwebauthactionSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["fullreqexpr"]; ok && val != nil {
		data.Fullreqexpr = types.StringValue(val.(string))
	} else {
		data.Fullreqexpr = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["scheme"]; ok && val != nil {
		data.Scheme = types.StringValue(val.(string))
	} else {
		data.Scheme = types.StringNull()
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
	if val, ok := getResponseData["successrule"]; ok && val != nil {
		data.Successrule = types.StringValue(val.(string))
	} else {
		data.Successrule = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
