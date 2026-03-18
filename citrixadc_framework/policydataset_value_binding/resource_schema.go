package policydataset_value_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PolicydatasetValueBindingResourceModel describes the resource data model.
type PolicydatasetValueBindingResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Comment  types.String `tfsdk:"comment"`
	Endrange types.String `tfsdk:"endrange"`
	Index    types.Int64  `tfsdk:"index"`
	Name     types.String `tfsdk:"name"`
	Value    types.String `tfsdk:"value"`
}

func (r *PolicydatasetValueBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policydataset_value_binding resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this dataset or a data bound to this dataset.",
			},
			"endrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The dataset entry is a range from <value> through <end_range>, inclusive. endRange cannot be used if value is an ipv4 or ipv6 subnet and endRange cannot itself be a subnet.",
			},
			"index": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index of the value (ipv4, ipv6, number) associated with the set.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dataset to which to bind the value.",
			},
			"value": schema.StringAttribute{
				Required:    true,
				Description: "Value of the specified type that is associated with the dataset. For ipv4 and ipv6, value can be a subnet using the slash notation address/n, where address is the beginning of the subnet and n is the number of left-most bits set in the subnet mask, defining the end of the subnet. The start address will be masked by the subnet mask if necessary, for example for 192.128.128.0/10, the start address will be 192.128.0.0.",
			},
		},
	}
}

func policydataset_value_bindingGetThePayloadFromtheConfig(ctx context.Context, data *PolicydatasetValueBindingResourceModel) policy.Policydatasetvaluebinding {
	tflog.Debug(ctx, "In policydataset_value_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policydataset_value_binding := policy.Policydatasetvaluebinding{}
	if !data.Comment.IsNull() {
		policydataset_value_binding.Comment = data.Comment.ValueString()
	}
	if !data.Endrange.IsNull() {
		policydataset_value_binding.Endrange = data.Endrange.ValueString()
	}
	if !data.Index.IsNull() {
		policydataset_value_binding.Index = utils.IntPtr(int(data.Index.ValueInt64()))
	}
	if !data.Name.IsNull() {
		policydataset_value_binding.Name = data.Name.ValueString()
	}
	if !data.Value.IsNull() {
		policydataset_value_binding.Value = data.Value.ValueString()
	}

	return policydataset_value_binding
}

func policydataset_value_bindingSetAttrFromGet(ctx context.Context, data *PolicydatasetValueBindingResourceModel, getResponseData map[string]interface{}) *PolicydatasetValueBindingResourceModel {
	tflog.Debug(ctx, "In policydataset_value_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["endrange"]; ok && val != nil {
		data.Endrange = types.StringValue(val.(string))
	} else {
		data.Endrange = types.StringNull()
	}
	if val, ok := getResponseData["index"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Index = types.Int64Value(intVal)
		}
	} else {
		data.Index = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["value"]; ok && val != nil {
		data.Value = types.StringValue(val.(string))
	} else {
		data.Value = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("endrange:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Endrange.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("value:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Value.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
