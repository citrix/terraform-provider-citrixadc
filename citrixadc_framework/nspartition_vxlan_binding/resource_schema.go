package nspartition_vxlan_binding

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

// NspartitionVxlanBindingResourceModel describes the resource data model.
type NspartitionVxlanBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Partitionname types.String `tfsdk:"partitionname"`
	Vxlan         types.Int64  `tfsdk:"vxlan"`
}

func (r *NspartitionVxlanBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nspartition_vxlan_binding resource.",
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Identifier of the vxlan that is assigned to this partition.",
			},
		},
	}
}

func nspartition_vxlan_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NspartitionVxlanBindingResourceModel) ns.Nspartitionvxlanbinding {
	tflog.Debug(ctx, "In nspartition_vxlan_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nspartition_vxlan_binding := ns.Nspartitionvxlanbinding{}
	if !data.Partitionname.IsNull() {
		nspartition_vxlan_binding.Partitionname = data.Partitionname.ValueString()
	}
	if !data.Vxlan.IsNull() {
		nspartition_vxlan_binding.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}

	return nspartition_vxlan_binding
}

func nspartition_vxlan_bindingSetAttrFromGet(ctx context.Context, data *NspartitionVxlanBindingResourceModel, getResponseData map[string]interface{}) *NspartitionVxlanBindingResourceModel {
	tflog.Debug(ctx, "In nspartition_vxlan_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["partitionname"]; ok && val != nil {
		data.Partitionname = types.StringValue(val.(string))
	} else {
		data.Partitionname = types.StringNull()
	}
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		}
	} else {
		data.Vxlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("partitionname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Partitionname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vxlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vxlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
