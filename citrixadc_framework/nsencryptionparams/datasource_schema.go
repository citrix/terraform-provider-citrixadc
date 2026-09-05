package nsencryptionparams

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NsencryptionparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"keyvalue": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "The base64-encoded key generation number, method, and key value.\nNote:\n* Do not include this argument if you are changing the encryption method.\n* To generate a new key value for the current encryption method, specify an empty string \\(\"\"\\) as the value of this parameter. The parameter is passed implicitly, with its automatically generated value, to the Citrix ADC packet engines even when it is not included in the command. Passing the parameter to the packet engines enables the appliance to save the key value to the configuration file and to propagate the key value to the secondary appliance in a high availability setup.",
			},
			"method": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cipher method (and key length) to be used to encrypt and decrypt content. The default value is AES256.",
			},
		},
	}
}

type NsencryptionparamsDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Keyvalue types.String `tfsdk:"keyvalue"`
	Method   types.String `tfsdk:"method"`
}

func nsencryptionparamsDataSourceSetAttrFromGet(ctx context.Context, data *NsencryptionparamsDataSourceModel, getResponseData map[string]interface{}) *NsencryptionparamsDataSourceModel {
	tflog.Debug(ctx, "In nsencryptionparamsDataSourceSetAttrFromGet Function")

	// Convert API response to model
	// keyvalue is not returned by NITRO API (secret/ephemeral) - retain from config
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
