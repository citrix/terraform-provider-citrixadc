package sslprofile_sslciphersuite_binding

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

// SslprofileSslciphersuiteBindingResourceModel describes the resource data model.
type SslprofileSslciphersuiteBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Ciphername     types.String `tfsdk:"ciphername"`
	Cipherpriority types.Int64  `tfsdk:"cipherpriority"`
	Description    types.String `tfsdk:"description"`
	Name           types.String `tfsdk:"name"`
}

func (r *SslprofileSslciphersuiteBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslprofile_sslciphersuite_binding resource.",
			},
			"ciphername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The cipher group/alias/individual cipher configuration",
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
				// Pattern 15: read-only attribute returned by GET; not accepted in the
				// bind payload, so surface it as Computed-only.
				Computed:    true,
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

func sslprofile_sslciphersuite_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslprofileSslciphersuiteBindingResourceModel) ssl.Sslprofilesslciphersuitebinding {
	tflog.Debug(ctx, "In sslprofile_sslciphersuite_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslprofile_sslciphersuite_binding := ssl.Sslprofilesslciphersuitebinding{}
	if !data.Ciphername.IsNull() && !data.Ciphername.IsUnknown() {
		sslprofile_sslciphersuite_binding.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Cipherpriority.IsNull() && !data.Cipherpriority.IsUnknown() {
		sslprofile_sslciphersuite_binding.Cipherpriority = utils.IntPtr(int(data.Cipherpriority.ValueInt64()))
	}
	// Pattern 15: description is a read-only attribute; excluded from the bind payload.
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslprofile_sslciphersuite_binding.Name = data.Name.ValueString()
	}

	return sslprofile_sslciphersuite_binding
}

func sslprofile_sslciphersuite_bindingSetAttrFromGet(ctx context.Context, data *SslprofileSslciphersuiteBindingResourceModel, getResponseData map[string]interface{}) *SslprofileSslciphersuiteBindingResourceModel {
	tflog.Debug(ctx, "In sslprofile_sslciphersuite_bindingSetAttrFromGet Function")

	// Convert API response to model
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

	// Pattern 6: ID is set exactly once in Create; do not recompute it here.
	return data
}

// sslprofile_sslciphersuite_bindingSetAttrFromGetForDatasource faithfully copies the
// GET response into the model and sets the composite ID, because the datasource has no
// Create step to establish it (Pattern 7 split).
func sslprofile_sslciphersuite_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SslprofileSslciphersuiteBindingResourceModel, getResponseData map[string]interface{}) *SslprofileSslciphersuiteBindingResourceModel {
	tflog.Debug(ctx, "In sslprofile_sslciphersuite_bindingSetAttrFromGetForDatasource Function")

	sslprofile_sslciphersuite_bindingSetAttrFromGet(ctx, data, getResponseData)

	// Set ID for the datasource: composite key:UrlEncode(value) pairs. ID = name,ciphername
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
