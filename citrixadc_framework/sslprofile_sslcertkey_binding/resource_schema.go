package sslprofile_sslcertkey_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslprofileSslcertkeyBindingResourceModel describes the resource data model.
type SslprofileSslcertkeyBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Cipherpriority types.Int64  `tfsdk:"cipherpriority"`
	Name           types.String `tfsdk:"name"`
	Sslicacertkey  types.String `tfsdk:"sslicacertkey"`
}

func (r *SslprofileSslcertkeyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslprofile_sslcertkey_binding resource.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the cipher binding",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL profile.",
			},
			"sslicacertkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The certkey (CA certificate + private key) to be used for SSL interception.",
			},
		},
	}
}

func sslprofile_sslcertkey_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslprofileSslcertkeyBindingResourceModel) ssl.Sslprofilesslcertkeybinding {
	tflog.Debug(ctx, "In sslprofile_sslcertkey_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslprofile_sslcertkey_binding := ssl.Sslprofilesslcertkeybinding{}
	if !data.Cipherpriority.IsNull() {
		sslprofile_sslcertkey_binding.Cipherpriority = utils.IntPtr(int(data.Cipherpriority.ValueInt64()))
	}
	if !data.Name.IsNull() {
		sslprofile_sslcertkey_binding.Name = data.Name.ValueString()
	}
	if !data.Sslicacertkey.IsNull() {
		sslprofile_sslcertkey_binding.Sslicacertkey = data.Sslicacertkey.ValueString()
	}

	return sslprofile_sslcertkey_binding
}

func sslprofile_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *SslprofileSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *SslprofileSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In sslprofile_sslcertkey_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cipherpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherpriority = types.Int64Value(intVal)
		}
	} else {
		data.Cipherpriority = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["sslicacertkey"]; ok && val != nil {
		data.Sslicacertkey = types.StringValue(val.(string))
	} else {
		data.Sslicacertkey = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sslicacertkey:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Sslicacertkey.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
