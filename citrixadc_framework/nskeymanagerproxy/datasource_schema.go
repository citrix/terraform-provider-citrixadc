package nskeymanagerproxy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NskeymanagerproxyDataSourceModel is the data-source-specific model, decoupled
// from NskeymanagerproxyResourceModel.
//
// A data source is a pure read surface, so it exposes the read/write attributes
// (as Computed outputs) AND the read-only attribute the resource deliberately
// omits (status), which is populated only from a GET.
type NskeymanagerproxyDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Port       types.Int64  `tfsdk:"port"`
	Serverip   types.String `tfsdk:"serverip"`
	Servername types.String `tfsdk:"servername"`

	// Read-only (GET-only) attribute from the NITRO read-only set. Never
	// settable; populated from GET.
	Status types.Int64 `tfsdk:"status"`
}

func NskeymanagerproxyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Key Manager proxy server port.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the Key Manager proxy server.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the Key Manager proxy server.",
			},

			// Read-only (GET-only) attribute surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed.
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "Status of Key Manager proxy server connectivity.",
			},
		},
	}
}

// nskeymanagerproxyDataSourceSetAttrFromGet projects a NITRO nskeymanagerproxy
// GET response onto the data-source model. The shared utils.MapGet* helpers fill
// each attribute (or leave it Null when the GET omits it). nodeid is a cluster
// GET-context filter that the GET never returns, so its configured value is
// preserved rather than nulled. The ID is serverip when set, otherwise
// servername (both are unique lookup keys), matching the resource ID scheme.
func nskeymanagerproxyDataSourceSetAttrFromGet(ctx context.Context, data *NskeymanagerproxyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nskeymanagerproxyDataSourceSetAttrFromGet Function")

	data.Port = utils.MapGetInt64(g, "port")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Servername = utils.MapGetString(g, "servername")

	// nodeid is a cluster GET-context filter, not a returned property; preserve
	// the configured value unless the GET happens to echo it back.
	if v := utils.MapGetInt64(g, "nodeid"); !v.IsNull() {
		data.Nodeid = v
	}

	// Read-only attribute.
	data.Status = utils.MapGetInt64(g, "status")

	// ID is serverip when set, otherwise servername (both x-unique-attr).
	if !data.Serverip.IsNull() && data.Serverip.ValueString() != "" {
		data.Id = types.StringValue(data.Serverip.ValueString())
	} else {
		data.Id = types.StringValue(data.Servername.ValueString())
	}
}
