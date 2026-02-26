package lsngroup_lsnpool_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsngroupLsnpoolBindingResourceModel describes the resource data model.
type LsngroupLsnpoolBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Groupname types.String `tfsdk:"groupname"`
	Poolname  types.String `tfsdk:"poolname"`
}

func (r *LsngroupLsnpoolBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsngroup_lsnpool_binding resource.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn group1\" or 'lsn group1').",
			},
			"poolname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LSN pool to bind to the specified LSN group. Only LSN Pools and LSN groups with the same NAT type settings can be bound together. Multiples LSN pools can be bound to an LSN group.\n\nFor Deterministic NAT, pools bound to an LSN group cannot be bound to other LSN groups. For Dynamic NAT, pools bound to an LSN group can be bound to multiple LSN groups.",
			},
		},
	}
}

func lsngroup_lsnpool_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsngroupLsnpoolBindingResourceModel) lsn.Lsngrouplsnpoolbinding {
	tflog.Debug(ctx, "In lsngroup_lsnpool_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsngroup_lsnpool_binding := lsn.Lsngrouplsnpoolbinding{}
	if !data.Groupname.IsNull() {
		lsngroup_lsnpool_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Poolname.IsNull() {
		lsngroup_lsnpool_binding.Poolname = data.Poolname.ValueString()
	}

	return lsngroup_lsnpool_binding
}

func lsngroup_lsnpool_bindingSetAttrFromGet(ctx context.Context, data *LsngroupLsnpoolBindingResourceModel, getResponseData map[string]interface{}) *LsngroupLsnpoolBindingResourceModel {
	tflog.Debug(ctx, "In lsngroup_lsnpool_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["poolname"]; ok && val != nil {
		data.Poolname = types.StringValue(val.(string))
	} else {
		data.Poolname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("poolname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Poolname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
