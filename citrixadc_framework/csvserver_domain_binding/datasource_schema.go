package csvserver_domain_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsvserverDomainBindingDataSourceModel is the data-source-specific model,
// decoupled from CsvserverDomainBindingResourceModel. A data source is a pure
// read surface (Read only; no plan/apply lifecycle), so it can expose the FULL
// GET projection: the read/write attributes (as Computed outputs) AND the
// read-only attributes that the resource deliberately omits.
type CsvserverDomainBindingDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Backupip      types.String `tfsdk:"backupip"`
	Cookiedomain  types.String `tfsdk:"cookiedomain"`
	Cookietimeout types.Int64  `tfsdk:"cookietimeout"`
	Domainname    types.String `tfsdk:"domainname"`
	Name          types.String `tfsdk:"name"`
	Sitedomainttl types.Int64  `tfsdk:"sitedomainttl"`
	Ttl           types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/csvserver_domain_binding.json). Never settable;
	// populated from GET.
	Appflowlog types.String `tfsdk:"appflowlog"`
}

func CsvserverDomainBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"backupip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"cookiedomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"cookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"domainname": schema.StringAttribute{
				Required:    true,
				Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"sitedomainttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"appflowlog": schema.StringAttribute{
				Computed:    true,
				Description: "Enable logging appflow flow information.",
			},
		},
	}
}

// csvserver_domain_bindingDataSourceSetAttrFromGet projects a NITRO
// csvserver_domain_binding GET response onto the data-source model via the
// shared utils.MapGet* helpers.
func csvserver_domain_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CsvserverDomainBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In csvserver_domain_bindingDataSourceSetAttrFromGet Function")

	data.Backupip = utils.MapGetString(g, "backupip")
	data.Cookiedomain = utils.MapGetString(g, "cookiedomain")
	data.Cookietimeout = utils.MapGetInt64(g, "cookietimeout")
	data.Domainname = utils.MapGetString(g, "domainname")
	data.Name = utils.MapGetString(g, "name")
	data.Sitedomainttl = utils.MapGetInt64(g, "sitedomainttl")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only attributes.
	data.Appflowlog = utils.MapGetString(g, "appflowlog")

	// Set ID for the data source (composite key: domainname,name).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("domainname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Domainname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
