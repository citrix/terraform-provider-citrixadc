package sslhsmkey

import (
	"context"
	"fmt"

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

// SslhsmkeyResourceModel describes the resource data model.
type SslhsmkeyResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Hsmkeyname        types.String `tfsdk:"hsmkeyname"`
	Hsmtype           types.String `tfsdk:"hsmtype"`
	Key               types.String `tfsdk:"key"`
	Keystore          types.String `tfsdk:"keystore"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Serialnum         types.String `tfsdk:"serialnum"`
}

func (r *SslhsmkeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslhsmkey resource.",
			},
			"hsmkeyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"hsmtype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of HSM.",
			},
			"key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the key. optionally, for Thales, path to the HSM key file; /var/opt/nfast/kmdata/local/ is the default path. Applies when HSMTYPE is THALES or KEYVAULT.",
			},
			"keystore": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of keystore object representing HSM where key is stored. For example, name of keyvault object or azurekeyvault authentication object. Applies only to KEYVAULT type HSM.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password for a partition. Applies only to SafeNet HSM.",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password for a partition. Applies only to SafeNet HSM.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Increment this version to signal a password_wo update.",
			},
			"serialnum": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Serial number of the partition on which the key is present. Applies only to SafeNet HSM.",
			},
		},
	}
}

func sslhsmkeyGetThePayloadFromthePlan(ctx context.Context, data *SslhsmkeyResourceModel) ssl.Sslhsmkey {
	tflog.Debug(ctx, "In sslhsmkeyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslhsmkey := ssl.Sslhsmkey{}
	if !data.Hsmkeyname.IsNull() && !data.Hsmkeyname.IsUnknown() {
		sslhsmkey.Hsmkeyname = data.Hsmkeyname.ValueString()
	}
	if !data.Hsmtype.IsNull() && !data.Hsmtype.IsUnknown() {
		sslhsmkey.Hsmtype = data.Hsmtype.ValueString()
	}
	if !data.Key.IsNull() && !data.Key.IsUnknown() {
		sslhsmkey.Key = data.Key.ValueString()
	}
	if !data.Keystore.IsNull() && !data.Keystore.IsUnknown() {
		sslhsmkey.Keystore = data.Keystore.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslhsmkey.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Serialnum.IsNull() && !data.Serialnum.IsUnknown() {
		sslhsmkey.Serialnum = data.Serialnum.ValueString()
	}

	return sslhsmkey
}

func sslhsmkeyGetThePayloadFromtheConfig(ctx context.Context, data *SslhsmkeyResourceModel, payload *ssl.Sslhsmkey) {
	tflog.Debug(ctx, "In sslhsmkeyGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func sslhsmkeySetAttrFromGet(ctx context.Context, data *SslhsmkeyResourceModel, getResponseData map[string]interface{}) *SslhsmkeyResourceModel {
	tflog.Debug(ctx, "In sslhsmkeySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["hsmkeyname"]; ok && val != nil {
		data.Hsmkeyname = types.StringValue(val.(string))
	} else {
		data.Hsmkeyname = types.StringNull()
	}
	if val, ok := getResponseData["hsmtype"]; ok && val != nil {
		data.Hsmtype = types.StringValue(val.(string))
	} else {
		data.Hsmtype = types.StringNull()
	}
	if val, ok := getResponseData["key"]; ok && val != nil {
		data.Key = types.StringValue(val.(string))
	} else {
		data.Key = types.StringNull()
	}
	if val, ok := getResponseData["keystore"]; ok && val != nil {
		data.Keystore = types.StringValue(val.(string))
	} else {
		data.Keystore = types.StringNull()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["serialnum"]; ok && val != nil {
		data.Serialnum = types.StringValue(val.(string))
	} else {
		data.Serialnum = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Hsmkeyname.ValueString()))

	return data
}
