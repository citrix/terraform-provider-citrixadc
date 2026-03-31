package sslprofile_ecccurve_binding

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

// SslprofileEcccurveBindingResourceModel describes the resource data model.
type SslprofileEcccurveBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Cipherpriority types.Int64  `tfsdk:"cipherpriority"`
	Ecccurvename   types.String `tfsdk:"ecccurvename"`
	Name           types.String `tfsdk:"name"`
}

func (r *SslprofileEcccurveBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslprofile_ecccurve_binding resource.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the cipher binding",
			},
			"ecccurvename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Named ECC curve bound to vserver/service.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL profile.",
			},
		},
	}
}

func sslprofile_ecccurve_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslprofileEcccurveBindingResourceModel) ssl.Sslprofileecccurvebinding {
	tflog.Debug(ctx, "In sslprofile_ecccurve_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslprofile_ecccurve_binding := ssl.Sslprofileecccurvebinding{}
	if !data.Cipherpriority.IsNull() {
		sslprofile_ecccurve_binding.Cipherpriority = utils.IntPtr(int(data.Cipherpriority.ValueInt64()))
	}
	if !data.Ecccurvename.IsNull() {
		sslprofile_ecccurve_binding.Ecccurvename = data.Ecccurvename.ValueString()
	}
	if !data.Name.IsNull() {
		sslprofile_ecccurve_binding.Name = data.Name.ValueString()
	}

	return sslprofile_ecccurve_binding
}

func sslprofile_ecccurve_bindingSetAttrFromGet(ctx context.Context, data *SslprofileEcccurveBindingResourceModel, getResponseData map[string]interface{}) *SslprofileEcccurveBindingResourceModel {
	tflog.Debug(ctx, "In sslprofile_ecccurve_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cipherpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherpriority = types.Int64Value(intVal)
		}
	} else {
		data.Cipherpriority = types.Int64Null()
	}
	if val, ok := getResponseData["ecccurvename"]; ok && val != nil {
		data.Ecccurvename = types.StringValue(val.(string))
	} else {
		data.Ecccurvename = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ecccurvename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ecccurvename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
