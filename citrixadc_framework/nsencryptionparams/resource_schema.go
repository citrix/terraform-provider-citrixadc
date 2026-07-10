package nsencryptionparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsencryptionparamsResourceModel describes the resource data model.
type NsencryptionparamsResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Keyvalue          types.String `tfsdk:"keyvalue"`
	KeyvalueWo        types.String `tfsdk:"keyvalue_wo"`
	KeyvalueWoVersion types.Int64  `tfsdk:"keyvalue_wo_version"`
	Method            types.String `tfsdk:"method"`
}

func (r *NsencryptionparamsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsencryptionparams resource.",
			},
			"keyvalue": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "The base64-encoded key generation number, method, and key value.\nNote:\n* Do not include this argument if you are changing the encryption method.\n* To generate a new key value for the current encryption method, specify an empty string \\(\"\"\\) as the value of this parameter. The parameter is passed implicitly, with its automatically generated value, to the Citrix ADC packet engines even when it is not included in the command. Passing the parameter to the packet engines enables the appliance to save the key value to the configuration file and to propagate the key value to the secondary appliance in a high availability setup.",
			},
			"keyvalue_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "The base64-encoded key generation number, method, and key value.\nNote:\n* Do not include this argument if you are changing the encryption method.\n* To generate a new key value for the current encryption method, specify an empty string \\(\"\"\\) as the value of this parameter. The parameter is passed implicitly, with its automatically generated value, to the Citrix ADC packet engines even when it is not included in the command. Passing the parameter to the packet engines enables the appliance to save the key value to the configuration file and to propagate the key value to the secondary appliance in a high availability setup.",
			},
			"keyvalue_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a keyvalue_wo update.",
			},
			"method": schema.StringAttribute{
				Required:    true,
				Description: "Cipher method (and key length) to be used to encrypt and decrypt content. The default value is AES256.",
			},
		},
	}
}

func nsencryptionparamsGetThePayloadFromthePlan(ctx context.Context, data *NsencryptionparamsResourceModel) ns.Nsencryptionparams {
	tflog.Debug(ctx, "In nsencryptionparamsGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nsencryptionparams := ns.Nsencryptionparams{}
	if !data.Keyvalue.IsNull() && !data.Keyvalue.IsUnknown() {
		nsencryptionparams.Keyvalue = data.Keyvalue.ValueString()
	}
	// Skip write-only attribute: keyvalue_wo
	// Skip version tracker attribute: keyvalue_wo_version
	if !data.Method.IsNull() && !data.Method.IsUnknown() {
		nsencryptionparams.Method = data.Method.ValueString()
	}

	return nsencryptionparams
}

func nsencryptionparamsGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *NsencryptionparamsResourceModel) ns.Nsencryptionparams {
	tflog.Debug(ctx, "In nsencryptionparamsGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model
	nsencryptionparams := ns.Nsencryptionparams{}
	if !data.Keyvalue.IsNull() && !data.Keyvalue.IsUnknown() {
		nsencryptionparams.Keyvalue = data.Keyvalue.ValueString()
	}
	if !data.Method.IsNull() && !data.Method.IsUnknown() {
		nsencryptionparams.Method = data.Method.ValueString()
	}

	return nsencryptionparams
}

func nsencryptionparamsGetThePayloadFromtheConfig(ctx context.Context, data *NsencryptionparamsResourceModel, payload *ns.Nsencryptionparams) {
	tflog.Debug(ctx, "In nsencryptionparamsGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: keyvalue_wo -> keyvalue
	if !data.KeyvalueWo.IsNull() {
		keyvalueWo := data.KeyvalueWo.ValueString()
		if keyvalueWo != "" {
			payload.Keyvalue = keyvalueWo
		}
	}
}

func nsencryptionparamsSetAttrFromGet(ctx context.Context, data *NsencryptionparamsResourceModel, getResponseData map[string]interface{}) *NsencryptionparamsResourceModel {
	tflog.Debug(ctx, "In nsencryptionparamsSetAttrFromGet Function")

	// Convert API response to model
	// keyvalue is not returned by NITRO API (secret/ephemeral) - retain from config
	// keyvalue_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// keyvalue_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["method"]; ok && val != nil {
		data.Method = types.StringValue(val.(string))
	} else {
		data.Method = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsencryptionparams-config")

	return data
}
