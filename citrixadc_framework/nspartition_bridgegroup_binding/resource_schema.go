package nspartition_bridgegroup_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NspartitionBridgegroupBindingResourceModel describes the resource data model.
type NspartitionBridgegroupBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Bridgegroup   types.Int64  `tfsdk:"bridgegroup"`
	Partitionname types.String `tfsdk:"partitionname"`
}

func (r *NspartitionBridgegroupBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nspartition_bridgegroup_binding resource.",
			},
			"bridgegroup": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Identifier of the bridge group that is assigned to this partition.",
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
		},
	}
}

func nspartition_bridgegroup_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NspartitionBridgegroupBindingResourceModel) ns.Nspartitionbridgegroupbinding {
	tflog.Debug(ctx, "In nspartition_bridgegroup_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nspartition_bridgegroup_binding := ns.Nspartitionbridgegroupbinding{}
	if !data.Bridgegroup.IsNull() {
		nspartition_bridgegroup_binding.Bridgegroup = utils.IntPtr(int(data.Bridgegroup.ValueInt64()))
	}
	if !data.Partitionname.IsNull() {
		nspartition_bridgegroup_binding.Partitionname = data.Partitionname.ValueString()
	}

	return nspartition_bridgegroup_binding
}

func nspartition_bridgegroup_bindingSetAttrFromGet(ctx context.Context, data *NspartitionBridgegroupBindingResourceModel, getResponseData map[string]interface{}) *NspartitionBridgegroupBindingResourceModel {
	tflog.Debug(ctx, "In nspartition_bridgegroup_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bridgegroup"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bridgegroup = types.Int64Value(intVal)
		}
	} else {
		data.Bridgegroup = types.Int64Null()
	}
	if val, ok := getResponseData["partitionname"]; ok && val != nil {
		data.Partitionname = types.StringValue(val.(string))
	} else {
		data.Partitionname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bridgegroup:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Bridgegroup.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("partitionname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Partitionname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
