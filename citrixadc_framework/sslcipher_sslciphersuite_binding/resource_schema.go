package sslcipher_sslciphersuite_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslcipherSslciphersuiteBindingResourceModel describes the resource data model.
type SslcipherSslciphersuiteBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Ciphergroupname types.String `tfsdk:"ciphergroupname"`
	Ciphername      types.String `tfsdk:"ciphername"`
	Cipheroperation types.String `tfsdk:"cipheroperation"`
	Cipherpriority  types.Int64  `tfsdk:"cipherpriority"`
	Ciphgrpals      types.String `tfsdk:"ciphgrpals"`
	Description     types.String `tfsdk:"description"`
}

func (r *SslcipherSslciphersuiteBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcipher_sslciphersuite_binding resource.",
			},
			"ciphergroupname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the user-defined cipher group.",
			},
			"ciphername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cipher name.",
			},
			"cipheroperation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The operation that is performed when adding the cipher-suite.\n\nPossible cipher operations are:\n	ADD - Appends the given cipher-suite to the existing one configured for the virtual server.\n	REM - Removes the given cipher-suite from the existing one configured for the virtual server.\n	ORD - Overrides the current configured cipher-suite for the virtual server with the given cipher-suite.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This indicates priority assigned to the particular cipher",
			},
			"ciphgrpals": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or user defined cipher-group name.",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Cipher suite description.",
			},
		},
	}
}

func sslcipher_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslcipherSslciphersuiteBindingResourceModel) ssl.Sslciphersslciphersuitebinding {
	tflog.Debug(ctx, "In sslcipher_sslciphersuite_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslcipher_sslciphersuite_binding := ssl.Sslciphersslciphersuitebinding{}
	if !data.Ciphergroupname.IsNull() {
		sslcipher_sslciphersuite_binding.Ciphergroupname = data.Ciphergroupname.ValueString()
	}
	if !data.Ciphername.IsNull() {
		sslcipher_sslciphersuite_binding.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Cipheroperation.IsNull() {
		sslcipher_sslciphersuite_binding.Cipheroperation = data.Cipheroperation.ValueString()
	}
	if !data.Cipherpriority.IsNull() {
		sslcipher_sslciphersuite_binding.Cipherpriority = utils.IntPtr(int(data.Cipherpriority.ValueInt64()))
	}
	if !data.Ciphgrpals.IsNull() {
		sslcipher_sslciphersuite_binding.Ciphgrpals = data.Ciphgrpals.ValueString()
	}
	if !data.Description.IsNull() {
		sslcipher_sslciphersuite_binding.Description = data.Description.ValueString()
	}

	return sslcipher_sslciphersuite_binding
}

func sslcipher_sslciphersuite_bindingSetAttrFromGet(ctx context.Context, data *SslcipherSslciphersuiteBindingResourceModel, getResponseData map[string]interface{}) *SslcipherSslciphersuiteBindingResourceModel {
	tflog.Debug(ctx, "In sslcipher_sslciphersuite_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ciphergroupname"]; ok && val != nil {
		data.Ciphergroupname = types.StringValue(val.(string))
	} else {
		data.Ciphergroupname = types.StringNull()
	}
	if val, ok := getResponseData["ciphername"]; ok && val != nil {
		data.Ciphername = types.StringValue(val.(string))
	} else {
		data.Ciphername = types.StringNull()
	}
	if val, ok := getResponseData["cipheroperation"]; ok && val != nil {
		data.Cipheroperation = types.StringValue(val.(string))
	} else {
		data.Cipheroperation = types.StringNull()
	}
	if val, ok := getResponseData["cipherpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherpriority = types.Int64Value(intVal)
		}
	} else {
		data.Cipherpriority = types.Int64Null()
	}
	if val, ok := getResponseData["ciphgrpals"]; ok && val != nil {
		data.Ciphgrpals = types.StringValue(val.(string))
	} else {
		data.Ciphgrpals = types.StringNull()
	}
	if val, ok := getResponseData["description"]; ok && val != nil {
		data.Description = types.StringValue(val.(string))
	} else {
		data.Description = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphergroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphergroupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
