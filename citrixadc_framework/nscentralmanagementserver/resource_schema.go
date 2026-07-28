package nscentralmanagementserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NscentralmanagementserverResourceModel describes the resource data model.
type NscentralmanagementserverResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Activationcode       types.String `tfsdk:"activationcode"`
	Adcpassword          types.String `tfsdk:"adcpassword"`
	AdcpasswordWo        types.String `tfsdk:"adcpassword_wo"`
	AdcpasswordWoVersion types.Int64  `tfsdk:"adcpassword_wo_version"`
	Adcusername          types.String `tfsdk:"adcusername"`
	Deviceprofilename    types.String `tfsdk:"deviceprofilename"`
	Ipaddress            types.String `tfsdk:"ipaddress"`
	Password             types.String `tfsdk:"password"`
	PasswordWo           types.String `tfsdk:"password_wo"`
	PasswordWoVersion    types.Int64  `tfsdk:"password_wo_version"`
	Servername           types.String `tfsdk:"servername"`
	Type                 types.String `tfsdk:"type"`
	Username             types.String `tfsdk:"username"`
	Validatecert         types.String `tfsdk:"validatecert"`
}

func (r *NscentralmanagementserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nscentralmanagementserver resource.",
			},
			"activationcode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Activation code is used to register to ADM service",
			},
			"adcpassword": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ADC password used to create device profile on ADM",
			},
			"adcpassword_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ADC password used to create device profile on ADM",
			},
			"adcpassword_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a adcpassword_wo update.",
			},
			"adcusername": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ADC username used to create device profile on ADM",
			},
			"deviceprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Device profile is created on ADM and contains the user name and password of the instance(s).",
			},
			"ipaddress": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Ip Address of central management server.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password for access to central management server. Required for any user account.",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password for access to central management server. Required for any user account.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a password_wo update.",
			},
			"servername": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Fully qualified domain name of the central management server or service-url to locate ADM service.",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the central management server. Must be either CLOUD or ONPREM depending on whether the server is on the cloud or on premise.",
			},
			"username": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Username for access to central management server. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or\nsingle quotation marks (for example, \"my ns centralmgmtserver\" or \"my ns centralmgmtserver\").",
			},
			"validatecert": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "validate the server certificate for secure SSL connections.",
			},
		},
	}
}

func nscentralmanagementserverGetThePayloadFromthePlan(ctx context.Context, data *NscentralmanagementserverResourceModel) ns.Nscentralmanagementserver {
	tflog.Debug(ctx, "In nscentralmanagementserverGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nscentralmanagementserver := ns.Nscentralmanagementserver{}
	if !data.Activationcode.IsNull() && !data.Activationcode.IsUnknown() {
		nscentralmanagementserver.Activationcode = data.Activationcode.ValueString()
	}
	if !data.Adcpassword.IsNull() && !data.Adcpassword.IsUnknown() {
		nscentralmanagementserver.Adcpassword = data.Adcpassword.ValueString()
	}
	// Skip write-only attribute: adcpassword_wo
	// Skip version tracker attribute: adcpassword_wo_version
	if !data.Adcusername.IsNull() && !data.Adcusername.IsUnknown() {
		nscentralmanagementserver.Adcusername = data.Adcusername.ValueString()
	}
	if !data.Deviceprofilename.IsNull() && !data.Deviceprofilename.IsUnknown() {
		nscentralmanagementserver.Deviceprofilename = data.Deviceprofilename.ValueString()
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		nscentralmanagementserver.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		nscentralmanagementserver.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		nscentralmanagementserver.Servername = data.Servername.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		nscentralmanagementserver.Type = data.Type.ValueString()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		nscentralmanagementserver.Username = data.Username.ValueString()
	}
	if !data.Validatecert.IsNull() && !data.Validatecert.IsUnknown() {
		nscentralmanagementserver.Validatecert = data.Validatecert.ValueString()
	}

	return nscentralmanagementserver
}

func nscentralmanagementserverGetThePayloadFromtheConfig(ctx context.Context, data *NscentralmanagementserverResourceModel, payload *ns.Nscentralmanagementserver) {
	tflog.Debug(ctx, "In nscentralmanagementserverGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: adcpassword_wo -> adcpassword
	if !data.AdcpasswordWo.IsNull() {
		adcpasswordWo := data.AdcpasswordWo.ValueString()
		if adcpasswordWo != "" {
			payload.Adcpassword = adcpasswordWo
		}
	}
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func nscentralmanagementserverSetAttrFromGet(ctx context.Context, data *NscentralmanagementserverResourceModel, getResponseData map[string]interface{}) *NscentralmanagementserverResourceModel {
	tflog.Debug(ctx, "In nscentralmanagementserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["activationcode"]; ok && val != nil {
		data.Activationcode = types.StringValue(val.(string))
	} else {
		data.Activationcode = types.StringNull()
	}
	// adcpassword is not returned by NITRO API (secret/ephemeral) - retain from config
	// adcpassword_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// adcpassword_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["adcusername"]; ok && val != nil {
		data.Adcusername = types.StringValue(val.(string))
	} else {
		data.Adcusername = types.StringNull()
	}
	if val, ok := getResponseData["deviceprofilename"]; ok && val != nil {
		data.Deviceprofilename = types.StringValue(val.(string))
	} else {
		data.Deviceprofilename = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
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
	if val, ok := getResponseData["validatecert"]; ok && val != nil {
		data.Validatecert = types.StringValue(val.(string))
	} else {
		data.Validatecert = types.StringNull()
	}

	// ID is set once in Create (Pattern 6); do not recompute it here.

	return data
}

// nscentralmanagementserverSetAttrFromGetForDatasource faithfully copies the GET
// response and sets the ID, since the datasource has no Create step to set it.
func nscentralmanagementserverSetAttrFromGetForDatasource(ctx context.Context, data *NscentralmanagementserverResourceModel, getResponseData map[string]interface{}) *NscentralmanagementserverResourceModel {
	tflog.Debug(ctx, "In nscentralmanagementserverSetAttrFromGetForDatasource Function")

	nscentralmanagementserverSetAttrFromGet(ctx, data, getResponseData)

	// Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Type.ValueString()))

	return data
}
