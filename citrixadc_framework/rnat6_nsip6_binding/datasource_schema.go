package rnat6_nsip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Rnat6Nsip6BindingDataSourceModel is the data-source-specific model, decoupled
// from the resource model. A data source is a pure read surface (Read only; no
// plan/apply lifecycle), so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes that
// the resource deliberately omits. Every non-key attribute is Computed.
type Rnat6Nsip6BindingDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Natip6     types.String `tfsdk:"natip6"`
	Ownergroup types.String `tfsdk:"ownergroup"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/rnat6_nsip6_binding.json). Never settable; populated
	// from GET, Null when the appliance omits them.
	Td types.Int64 `tfsdk:"td"`
}

func Rnat6Nsip6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the RNAT6 rule to which to bind NAT IPs.",
			},
			"natip6": schema.StringAttribute{
				Required:    true,
				Description: "Nat IP Address.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this rnat rule.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"td": schema.Int64Attribute{
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which the entity is configured.",
			},
		},
	}
}

// rnat6_nsip6_bindingDataSourceSetAttrFromGet projects a NITRO rnat6_nsip6_binding
// GET response onto the data-source model.
func rnat6_nsip6_bindingDataSourceSetAttrFromGet(ctx context.Context, data *Rnat6Nsip6BindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In rnat6_nsip6_bindingDataSourceSetAttrFromGet Function")

	data.Name = utils.MapGetString(g, "name")
	data.Natip6 = utils.MapGetString(g, "natip6")
	data.Ownergroup = utils.MapGetString(g, "ownergroup")

	// Read-only (GET-only) attributes.
	data.Td = utils.MapGetInt64(g, "td")

	// Set composite ID. Backward-compatible with SDK v2: identity is
	// "name,natip6" (comma-separated key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("natip6:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Natip6.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
