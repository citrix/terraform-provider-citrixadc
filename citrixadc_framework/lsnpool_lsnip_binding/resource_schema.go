package lsnpool_lsnip_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsnpoolLsnipBindingResourceModel describes the resource data model.
type LsnpoolLsnipBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Lsnip     types.String `tfsdk:"lsnip"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
	Poolname  types.String `tfsdk:"poolname"`
}

func (r *LsnpoolLsnipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnpool_lsnip_binding resource.",
			},
			"lsnip": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv4 address or a range of IPv4 addresses to be used as NAT IP address(es) for LSN.\nAfter the pool is created, these IPv4 addresses are added to the Citrix ADC as Citrix ADC owned IP address of type LSN. A maximum of 4096 IP addresses can be bound to an LSN pool. An LSN IP address associated with an LSN pool cannot be shared with other LSN pools. IP addresses specified for this parameter must not already exist on the Citrix ADC as any Citrix ADC owned IP addresses. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189. You can later remove some or all the LSN IP addresses from the pool, and add IP addresses to the LSN pool.\nBy default , arp is enabled on LSN IP address but, you can disable it using command - \"set ns ip\"",
			},
			"ownernode": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID(s) of cluster node(s) on which command is to be executed",
			},
			"poolname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the LSN pool. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN pool is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn pool1\" or 'lsn pool1').",
			},
		},
	}
}

func lsnpool_lsnip_bindingGetThePayloadFromthePlan(ctx context.Context, data *LsnpoolLsnipBindingResourceModel) lsn.Lsnpoollsnipbinding {
	tflog.Debug(ctx, "In lsnpool_lsnip_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lsnpool_lsnip_binding := lsn.Lsnpoollsnipbinding{}
	if !data.Lsnip.IsNull() && !data.Lsnip.IsUnknown() {
		lsnpool_lsnip_binding.Lsnip = data.Lsnip.ValueString()
	}
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		lsnpool_lsnip_binding.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Poolname.IsNull() && !data.Poolname.IsUnknown() {
		lsnpool_lsnip_binding.Poolname = data.Poolname.ValueString()
	}

	return lsnpool_lsnip_binding
}

// lsnpool_lsnip_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied values (poolname, lsnip, ownernode are all RequiresReplace
// identity attrs) and does NOT recompute the ID, which is set exactly once in Create.
func lsnpool_lsnip_bindingSetAttrFromGet(ctx context.Context, data *LsnpoolLsnipBindingResourceModel, getResponseData map[string]interface{}) *LsnpoolLsnipBindingResourceModel {
	tflog.Debug(ctx, "In lsnpool_lsnip_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["lsnip"]; ok && val != nil {
		data.Lsnip = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["poolname"]; ok && val != nil {
		data.Poolname = types.StringValue(val.(string))
	}

	return data
}

// lsnpool_lsnip_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the ID, since the datasource has no Create to seed those values.
func lsnpool_lsnip_bindingSetAttrFromGetForDatasource(ctx context.Context, data *LsnpoolLsnipBindingResourceModel, getResponseData map[string]interface{}) *LsnpoolLsnipBindingResourceModel {
	tflog.Debug(ctx, "In lsnpool_lsnip_bindingSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["lsnip"]; ok && val != nil {
		data.Lsnip = types.StringValue(val.(string))
	} else {
		data.Lsnip = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["poolname"]; ok && val != nil {
		data.Poolname = types.StringValue(val.(string))
	} else {
		data.Poolname = types.StringNull()
	}

	// Set ID for the datasource
	// Composite key: poolname,lsnip (ownernode is cluster-only, not part of the unique key)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("poolname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Poolname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("lsnip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Lsnip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
