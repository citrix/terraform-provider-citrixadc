package sslkeyfile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslkeyfileResourceModel describes the resource data model.
type SslkeyfileResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Src               types.String `tfsdk:"src"`
}

func (r *SslkeyfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslkeyfile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported key file. Must begin with an ASCII alphanumeric or underscore(_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
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
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL specifying the protocol, host, and path, including file name, to the key file to be imported. For example, http://www.example.com/key_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
		},
	}
}

func sslkeyfileGetThePayloadFromthePlan(ctx context.Context, data *SslkeyfileResourceModel) ssl.Sslkeyfile {
	tflog.Debug(ctx, "In sslkeyfileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslkeyfile := ssl.Sslkeyfile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslkeyfile.Name = data.Name.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslkeyfile.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		sslkeyfile.Src = data.Src.ValueString()
	}

	return sslkeyfile
}

func sslkeyfileGetThePayloadFromtheConfig(ctx context.Context, data *SslkeyfileResourceModel, payload *ssl.Sslkeyfile) {
	tflog.Debug(ctx, "In sslkeyfileGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func sslkeyfileSetAttrFromGet(ctx context.Context, data *SslkeyfileResourceModel, getResponseData map[string]interface{}) *SslkeyfileResourceModel {
	tflog.Debug(ctx, "In sslkeyfileSetAttrFromGet Function")

	// Resource setter: preserves plan/state values. `name` is the key and `src`
	// is a write-only Import input that NITRO does not faithfully echo back, so
	// do not overwrite either from the GET response. The password secret triple
	// is never returned by NITRO and is retained from config. ID is set in Create.
	if data.Name.IsNull() || data.Name.IsUnknown() {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}

	return data
}

func sslkeyfileSetAttrFromGetForDatasource(ctx context.Context, data *SslkeyfileResourceModel, getResponseData map[string]interface{}) *SslkeyfileResourceModel {
	tflog.Debug(ctx, "In sslkeyfileSetAttrFromGetForDatasource Function")

	// Datasource setter: faithfully copy the GET response (no prior state).
	// Secrets are never returned by NITRO, so password attributes stay null.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}

	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
