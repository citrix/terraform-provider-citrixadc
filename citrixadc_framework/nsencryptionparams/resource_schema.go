package nsencryptionparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsencryptionparamsResourceModel describes the resource data model.
type NsencryptionparamsResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Keyvalue types.String `tfsdk:"keyvalue"`
	Method   types.String `tfsdk:"method"`
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
				Computed:    true,
				Description: "The base64-encoded key generation number, method, and key value.\nNote:\n* Do not include this argument if you are changing the encryption method.\n* To generate a new key value for the current encryption method, specify an empty string \\(\"\"\\) as the value of this parameter. The parameter is passed implicitly, with its automatically generated value, to the Citrix ADC packet engines even when it is not included in the command. Passing the parameter to the packet engines enables the appliance to save the key value to the configuration file and to propagate the key value to the secondary appliance in a high availability setup.",
			},
			"method": schema.StringAttribute{
				Required:    true,
				Description: "Cipher method (and key length) to be used to encrypt and decrypt content. The default value is AES256.",
			},
		},
	}
}

func nsencryptionparamsGetThePayloadFromtheConfig(ctx context.Context, data *NsencryptionparamsResourceModel) ns.Nsencryptionparams {
	tflog.Debug(ctx, "In nsencryptionparamsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsencryptionparams := ns.Nsencryptionparams{}
	if !data.Keyvalue.IsNull() {
		nsencryptionparams.Keyvalue = data.Keyvalue.ValueString()
	}
	if !data.Method.IsNull() {
		nsencryptionparams.Method = data.Method.ValueString()
	}

	return nsencryptionparams
}

func nsencryptionparamsSetAttrFromGet(ctx context.Context, data *NsencryptionparamsResourceModel, getResponseData map[string]interface{}) *NsencryptionparamsResourceModel {
	tflog.Debug(ctx, "In nsencryptionparamsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["keyvalue"]; ok && val != nil {
		data.Keyvalue = types.StringValue(val.(string))
	} else {
		data.Keyvalue = types.StringNull()
	}
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
