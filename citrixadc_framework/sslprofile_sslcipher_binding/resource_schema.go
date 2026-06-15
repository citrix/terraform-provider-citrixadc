package sslprofile_sslcipher_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the cipher.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the SSL profile.",
			},
		},
	}
}

func sslprofile_sslcipher_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslprofileSslcipherBindingResourceModel) ssl.Sslprofilesslcipherbinding {
	tflog.Debug(ctx, "In sslprofile_sslcipher_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslprofile_sslcipher_binding := ssl.Sslprofilesslcipherbinding{}
	if !data.Cipheraliasname.IsNull() && !data.Cipheraliasname.IsUnknown() {
		sslprofile_sslcipher_binding.Cipheraliasname = data.Cipheraliasname.ValueString()
	}
	if !data.Ciphername.IsNull() && !data.Ciphername.IsUnknown() {
		sslprofile_sslcipher_binding.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Cipherpriority.IsNull() && !data.Cipherpriority.IsUnknown() {
		sslprofile_sslcipher_binding.Cipherpriority = utils.IntPtr(int(data.Cipherpriority.ValueInt64()))
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		sslprofile_sslcipher_binding.Description = data.Description.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslprofile_sslcipher_binding.Name = data.Name.ValueString()
	}

	return sslprofile_sslcipher_binding
}

// sslprofile_sslcipher_bindingComputeId builds the resource ID in the new
// key:UrlEncode(value) format. The attribute order matches the legacy SDK v2
// ID order ("name,ciphername") recorded in resource_id_mapping.json so that
// ParseIdString decodes both legacy and new IDs consistently.
func sslprofile_sslcipher_bindingComputeId(data *SslprofileSslcipherBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	return strings.Join(idParts, ",")
}

func sslprofile_sslcipher_bindingSetAttrFromGet(ctx context.Context, data *SslprofileSslcipherBindingResourceModel, getResponseData map[string]interface{}) *SslprofileSslcipherBindingResourceModel {
	tflog.Debug(ctx, "In sslprofile_sslcipher_bindingSetAttrFromGet Function")

	// The NITRO GET response echoes the bound cipher's name in the
	// "cipheraliasname" field; there is no "ciphername" key in the response.
	// "cipheraliasname" is a Computed read-back output.
	if val, ok := getResponseData["cipheraliasname"]; ok && val != nil {
		data.Cipheraliasname = types.StringValue(val.(string))
		// Preserve the user's "ciphername" input. The SDK v2 resource read
		// ciphername back from cipheraliasname (d.Set("ciphername", cipheraliasname));
		// on import (where ciphername is unset) adopt the live value so the
		// imported resource is usable.
		if data.Ciphername.IsNull() || data.Ciphername.ValueString() == "" {
			data.Ciphername = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["cipherpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherpriority = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["description"]; ok && val != nil {
		data.Description = types.StringValue(val.(string))
	} else {
		data.Description = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}

	// Set ID for the resource (new key:UrlEncode(value) format).
	data.Id = types.StringValue(sslprofile_sslcipher_bindingComputeId(data))

	return data
}

// sslprofile_sslcipher_bindingSetAttrFromGetForDatasource faithfully copies the
// GET response into the model (the datasource has no prior plan/state to
// preserve) and sets the ID itself.
func sslprofile_sslcipher_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SslprofileSslcipherBindingResourceModel, getResponseData map[string]interface{}) *SslprofileSslcipherBindingResourceModel {
	tflog.Debug(ctx, "In sslprofile_sslcipher_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["cipheraliasname"]; ok && val != nil {
		data.Cipheraliasname = types.StringValue(val.(string))
		// The cipher value is returned only as cipheraliasname; mirror it into
		// ciphername so the datasource exposes the user-facing attribute too.
		data.Ciphername = types.StringValue(val.(string))
	} else {
		data.Cipheraliasname = types.StringNull()
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
	}

	data.Id = types.StringValue(sslprofile_sslcipher_bindingComputeId(data))

	return data
}
