package location

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LocationDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipfrom": schema.StringAttribute{
				Required:    true,
				Description: "First IP address in the range, in dotted decimal notation.",
			},
			"ipto": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Last IP address in the range, in dotted decimal notation.",
			},
			"latitude": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Numerical value, in degrees, specifying the latitude of the geographical location of the IP address-range.\nNote: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.",
			},
			"longitude": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Numerical value, in degrees, specifying the longitude of the geographical location of the IP address-range.\nNote: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.",
			},
			"preferredlocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String of qualifiers, in dotted notation, describing the geographical location of the IP address range. Each qualifier is more specific than the one that precedes it, as in continent.country.region.city.isp.organization. For example, \"NA.US.CA.San Jose.ATT.citrix\".\nNote: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.",
			},
		},
	}
}
