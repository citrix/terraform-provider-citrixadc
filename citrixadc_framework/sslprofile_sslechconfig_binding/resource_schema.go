package sslprofile_sslechconfig_binding

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

// SslprofileSslechconfigBindingResourceModel describes the resource data model.
type SslprofileSslechconfigBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Cipherpriority types.Int64  `tfsdk:"cipherpriority"`
	Echconfigname  types.String `tfsdk:"echconfigname"`
	Name           types.String `tfsdk:"name"`
}

func (r *SslprofileSslechconfigBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslprofile_sslechconfig_binding resource.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Priority of the cipher binding",
			},
			"echconfigname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Configuration for the Encrypted Client Hello feature",
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

func sslprofile_sslechconfig_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslprofileSslechconfigBindingResourceModel) ssl.Sslprofilesslechconfigbinding {
	tflog.Debug(ctx, "In sslprofile_sslechconfig_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslprofile_sslechconfig_binding := ssl.Sslprofilesslechconfigbinding{}
	if !data.Cipherpriority.IsNull() && !data.Cipherpriority.IsUnknown() {
		sslprofile_sslechconfig_binding.Cipherpriority = utils.IntPtr(int(data.Cipherpriority.ValueInt64()))
	}
	if !data.Echconfigname.IsNull() && !data.Echconfigname.IsUnknown() {
		sslprofile_sslechconfig_binding.Echconfigname = data.Echconfigname.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslprofile_sslechconfig_binding.Name = data.Name.ValueString()
	}

	return sslprofile_sslechconfig_binding
}

func sslprofile_sslechconfig_bindingSetAttrFromGet(ctx context.Context, data *SslprofileSslechconfigBindingResourceModel, getResponseData map[string]interface{}) *SslprofileSslechconfigBindingResourceModel {
	tflog.Debug(ctx, "In sslprofile_sslechconfig_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cipherpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherpriority = types.Int64Value(intVal)
		}
	} else {
		data.Cipherpriority = types.Int64Null()
	}
	if val, ok := getResponseData["echconfigname"]; ok && val != nil {
		data.Echconfigname = types.StringValue(val.(string))
	} else {
		data.Echconfigname = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Pattern 6: ID is set exactly once in Create; do not recompute it here.
	return data
}

// sslprofile_sslechconfig_bindingSetAttrFromGetForDatasource faithfully copies the GET
// response into the model and sets the composite ID, because the datasource has no
// Create step to establish it (Pattern 7 split).
func sslprofile_sslechconfig_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SslprofileSslechconfigBindingResourceModel, getResponseData map[string]interface{}) *SslprofileSslechconfigBindingResourceModel {
	tflog.Debug(ctx, "In sslprofile_sslechconfig_bindingSetAttrFromGetForDatasource Function")

	sslprofile_sslechconfig_bindingSetAttrFromGet(ctx, data, getResponseData)

	// Set ID for the datasource: composite key:UrlEncode(value) pairs. ID = name,echconfigname
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("echconfigname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Echconfigname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
