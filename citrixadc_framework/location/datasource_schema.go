package location

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LocationDataSourceModel is the data-source-specific model, decoupled from
// LocationResourceModel. A data source is a pure read surface, so it exposes the
// existing datasource attributes as Computed outputs PLUS the read-only
// attributes the NITRO GET returns (zion73x_readonly/location.json) that the
// resource intentionally omits (the q1label..q6label location qualifiers).
type LocationDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Ipfrom            types.String `tfsdk:"ipfrom"` // Required lookup key
	Ipto              types.String `tfsdk:"ipto"`
	Latitude          types.Int64  `tfsdk:"latitude"`
	Longitude         types.Int64  `tfsdk:"longitude"`
	Preferredlocation types.String `tfsdk:"preferredlocation"`

	// Read-only (GET-only) location qualifiers from the NITRO read-only set.
	Q1label types.String `tfsdk:"q1label"`
	Q2label types.String `tfsdk:"q2label"`
	Q3label types.String `tfsdk:"q3label"`
	Q4label types.String `tfsdk:"q4label"`
	Q5label types.String `tfsdk:"q5label"`
	Q6label types.String `tfsdk:"q6label"`
}

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

			// Read-only (GET-only) location qualifiers surfaced by the data source.
			"q1label": schema.StringAttribute{
				Computed:    true,
				Description: "Least specific location qualifier.",
			},
			"q2label": schema.StringAttribute{
				Computed:    true,
				Description: "Location qualifier 2.",
			},
			"q3label": schema.StringAttribute{
				Computed:    true,
				Description: "Location qualifier 3.",
			},
			"q4label": schema.StringAttribute{
				Computed:    true,
				Description: "Location qualifier 4.",
			},
			"q5label": schema.StringAttribute{
				Computed:    true,
				Description: "Location qualifier 5.",
			},
			"q6label": schema.StringAttribute{
				Computed:    true,
				Description: "Most specific location qualifier.",
			},
		},
	}
}

// locationDataSourceSetAttrFromGet projects a NITRO location GET response onto
// the data-source model. Unlike the resource setter it copies every attribute
// from the GET (including the ADC-normalized preferredlocation and the read-only
// q1label..q6label qualifiers) and sets the datasource ID from ipfrom. Absent
// attributes resolve to Null.
func locationDataSourceSetAttrFromGet(ctx context.Context, data *LocationDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In locationDataSourceSetAttrFromGet Function")

	data.Ipfrom = utils.MapGetString(g, "ipfrom")
	data.Ipto = utils.MapGetString(g, "ipto")
	data.Latitude = utils.MapGetInt64(g, "latitude")
	data.Longitude = utils.MapGetInt64(g, "longitude")
	data.Preferredlocation = utils.MapGetString(g, "preferredlocation")

	// Read-only location qualifiers.
	data.Q1label = utils.MapGetString(g, "q1label")
	data.Q2label = utils.MapGetString(g, "q2label")
	data.Q3label = utils.MapGetString(g, "q3label")
	data.Q4label = utils.MapGetString(g, "q4label")
	data.Q5label = utils.MapGetString(g, "q5label")
	data.Q6label = utils.MapGetString(g, "q6label")

	// Set ID for the datasource (single unique attr: ipfrom).
	if v, ok := g["ipfrom"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	}
}
