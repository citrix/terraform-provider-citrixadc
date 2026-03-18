package sslprofile_sslcipher_binding

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

// SslprofileSslcipherBindingResourceModel describes the resource data model.
type SslprofileSslcipherBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Cipheraliasname types.String `tfsdk:"cipheraliasname"`
	Ciphername      types.String `tfsdk:"ciphername"`
	Cipherpriority  types.Int64  `tfsdk:"cipherpriority"`
	Description     types.String `tfsdk:"description"`
	Name            types.String `tfsdk:"name"`
}

func (r *SslprofileSslcipherBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslprofile_sslcipher_binding resource.",
			},
			"cipheraliasname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the cipher group/alias/individual cipheri bindings.",
			},
			"ciphername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the cipher.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "cipher priority",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The cipher suite description.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL profile.",
			},
		},
	}
}

func sslprofile_sslcipher_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslprofileSslcipherBindingResourceModel) ssl.Sslprofilesslcipherbinding {
	tflog.Debug(ctx, "In sslprofile_sslcipher_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslprofile_sslcipher_binding := ssl.Sslprofilesslcipherbinding{}
	if !data.Cipheraliasname.IsNull() {
		sslprofile_sslcipher_binding.Cipheraliasname = data.Cipheraliasname.ValueString()
	}
	if !data.Ciphername.IsNull() {
		sslprofile_sslcipher_binding.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Cipherpriority.IsNull() {
		sslprofile_sslcipher_binding.Cipherpriority = utils.IntPtr(int(data.Cipherpriority.ValueInt64()))
	}
	if !data.Description.IsNull() {
		sslprofile_sslcipher_binding.Description = data.Description.ValueString()
	}
	if !data.Name.IsNull() {
		sslprofile_sslcipher_binding.Name = data.Name.ValueString()
	}

	return sslprofile_sslcipher_binding
}

func sslprofile_sslcipher_bindingSetAttrFromGet(ctx context.Context, data *SslprofileSslcipherBindingResourceModel, getResponseData map[string]interface{}) *SslprofileSslcipherBindingResourceModel {
	tflog.Debug(ctx, "In sslprofile_sslcipher_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cipheraliasname"]; ok && val != nil {
		data.Cipheraliasname = types.StringValue(val.(string))
	} else {
		data.Cipheraliasname = types.StringNull()
	}
	if val, ok := getResponseData["ciphername"]; ok && val != nil {
		data.Ciphername = types.StringValue(val.(string))
	} else {
		data.Ciphername = types.StringNull()
	}
	if val, ok := getResponseData["cipherpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherpriority = types.Int64Value(intVal)
		}
	} else {
		data.Cipherpriority = types.Int64Null()
	}
	if val, ok := getResponseData["description"]; ok && val != nil {
		data.Description = types.StringValue(val.(string))
	} else {
		data.Description = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
