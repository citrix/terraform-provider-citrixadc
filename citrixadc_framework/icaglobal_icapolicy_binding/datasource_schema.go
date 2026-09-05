package icaglobal_icapolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IcaglobalIcapolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from IcaglobalIcapolicyBindingResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the configurable
// attributes (as Computed outputs) AND the read-only metadata attributes the
// resource deliberately omits.
type IcaglobalIcapolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"` // Required lookup key
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/icaglobal_icapolicy_binding.json).
	Numpol   types.Int64 `tfsdk:"numpol"`
	Flowtype types.Int64 `tfsdk:"flowtype"`
}

func IcaglobalIcapolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"globalbindtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the ICA policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Global bind point for which to show detailed information about the policies bound to the bind point.",
			},

			// Read-only (GET-only) metadata surfaced by the data source. All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of policies bound to the bindpoint.",
			},
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "Flow type of the bound ICA policy.",
			},
		},
	}
}

// icaglobal_icapolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// icaglobal_icapolicy_binding GET response onto the data-source model.
func icaglobal_icapolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *IcaglobalIcapolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In icaglobal_icapolicy_bindingDataSourceSetAttrFromGet Function")

	// Preserve the config-provided lookup keys when the GET omits them.
	policyname := data.Policyname
	typ := data.Type

	if v, ok := g["policyname"]; ok && v != nil {
		data.Policyname = types.StringValue(utils.AnyToString(v))
	} else {
		data.Policyname = policyname
	}
	if v, ok := g["type"]; ok && v != nil {
		data.Type = types.StringValue(utils.AnyToString(v))
	} else {
		data.Type = typ
	}

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Priority = utils.MapGetInt64(g, "priority")

	// Read-only (GET-only) metadata.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Flowtype = utils.MapGetInt64(g, "flowtype")

	// Set the composite ID (policyname:<v>,type:<v>).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(data.Policyname.ValueString())))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(data.Type.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
