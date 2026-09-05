package clusternode_routemonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ClusternodeRoutemonitorBindingDataSourceModel is the data-source-specific
// model, decoupled from ClusternodeRoutemonitorBindingResourceModel so the data
// source can expose read-only (GET-only) attributes the resource omits.
type ClusternodeRoutemonitorBindingDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Netmask      types.String `tfsdk:"netmask"`
	Nodeid       types.Int64  `tfsdk:"nodeid"`
	Routemonitor types.String `tfsdk:"routemonitor"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/clusternode_routemonitor_binding.json). Never settable;
	// populated from GET and null when the appliance omits them.
	Routemonstate types.Int64 `tfsdk:"routemonstate"`
}

func ClusternodeRoutemonitorBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"netmask": schema.StringAttribute{
				Required:    true,
				Description: "The netmask.",
			},
			"nodeid": schema.Int64Attribute{
				Required:    true,
				Description: "A number that uniquely identifies the cluster node.",
			},
			"routemonitor": schema.StringAttribute{
				Required:    true,
				Description: "The IP address (IPv4 or IPv6).",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"routemonstate": schema.Int64Attribute{
				Computed:    true,
				Description: "Current routemonstate.",
			},
		},
	}
}

// clusternode_routemonitor_bindingDataSourceSetAttrFromGet projects a NITRO
// clusternode_routemonitor_binding GET response onto the data-source model. The
// shared utils.MapGet* helpers fill each attribute from the GET (or leave it
// Null when the GET omits it).
func clusternode_routemonitor_bindingDataSourceSetAttrFromGet(ctx context.Context, data *ClusternodeRoutemonitorBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In clusternode_routemonitor_bindingDataSourceSetAttrFromGet Function")

	// Lookup keys: prefer the GET value, but preserve the configured value when
	// the appliance omits it from the binding response.
	if v := utils.MapGetString(g, "netmask"); !v.IsNull() {
		data.Netmask = v
	}
	if v := utils.MapGetInt64(g, "nodeid"); !v.IsNull() {
		data.Nodeid = v
	}
	if v := utils.MapGetString(g, "routemonitor"); !v.IsNull() {
		data.Routemonitor = v
	}

	// Read-only (GET-only) attributes.
	data.Routemonstate = utils.MapGetInt64(g, "routemonstate")

	// Composite key -> id (key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("nodeid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Nodeid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("routemonitor:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Routemonitor.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
