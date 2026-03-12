package nspartition_vlan_binding

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

// NspartitionVlanBindingResourceModel describes the resource data model.
type NspartitionVlanBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Partitionname types.String `tfsdk:"partitionname"`
	Vlan          types.Int64  `tfsdk:"vlan"`
}

func (r *NspartitionVlanBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nspartition_vlan_binding resource.",
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Identifier of the vlan that is assigned to this partition.",
			},
		},
	}
}

func nspartition_vlan_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NspartitionVlanBindingResourceModel) ns.Nspartitionvlanbinding {
	tflog.Debug(ctx, "In nspartition_vlan_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nspartition_vlan_binding := ns.Nspartitionvlanbinding{}
	if !data.Partitionname.IsNull() {
		nspartition_vlan_binding.Partitionname = data.Partitionname.ValueString()
	}
	if !data.Vlan.IsNull() {
		nspartition_vlan_binding.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return nspartition_vlan_binding
}

func nspartition_vlan_bindingSetAttrFromGet(ctx context.Context, data *NspartitionVlanBindingResourceModel, getResponseData map[string]interface{}) *NspartitionVlanBindingResourceModel {
	tflog.Debug(ctx, "In nspartition_vlan_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["partitionname"]; ok && val != nil {
		data.Partitionname = types.StringValue(val.(string))
	} else {
		data.Partitionname = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("partitionname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Partitionname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
