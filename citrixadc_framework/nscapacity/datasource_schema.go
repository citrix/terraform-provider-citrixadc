package nscapacity

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NscapacityDataSourceModel is the data-source-specific model, decoupled from
// NscapacityResourceModel. nscapacity is a singleton, so this data source
// exposes the read/write attributes (as Computed outputs) AND the read-only
// capacity/license metadata the resource deliberately omits.
type NscapacityDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Bandwidth    types.Int64  `tfsdk:"bandwidth"`
	Edition      types.String `tfsdk:"edition"`
	Ignoreexpiry types.Bool   `tfsdk:"ignoreexpiry"`
	Nodeid       types.Int64  `tfsdk:"nodeid"`
	Password     types.String `tfsdk:"password"`
	Platform     types.String `tfsdk:"platform"`
	Unit         types.String `tfsdk:"unit"`
	Username     types.String `tfsdk:"username"`
	Vcpu         types.Bool   `tfsdk:"vcpu"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nscapacity.json). Never settable; populated from GET.
	Actualbandwidth  types.Int64 `tfsdk:"actualbandwidth"`
	Vcpucount        types.Int64 `tfsdk:"vcpucount"`
	Maxvcpucount     types.Int64 `tfsdk:"maxvcpucount"`
	Maxbandwidth     types.Int64 `tfsdk:"maxbandwidth"`
	Minbandwidth     types.Int64 `tfsdk:"minbandwidth"`
	Instancecount    types.Int64 `tfsdk:"instancecount"`
	Daystoexpiration types.Int64 `tfsdk:"daystoexpiration"`
}

func NscapacityDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "System bandwidth limit.",
			},
			"edition": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Product edition.",
			},
			"ignoreexpiry": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value to mention if days to expire data needs to be fetched or not.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Password to use when authenticating with ADM Agent for LAS licensing.",
			},
			"platform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "appliance platform type.",
			},
			"unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bandwidth unit.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"vcpu": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "licensed using vcpu pool.",
			},

			// Read-only (GET-only) capacity/license metadata surfaced by the data source.
			"actualbandwidth": schema.Int64Attribute{
				Computed:    true,
				Description: "Bandwith in MBPS.",
			},
			"vcpucount": schema.Int64Attribute{
				Computed:    true,
				Description: "number of vCPUs licensed.",
			},
			"maxvcpucount": schema.Int64Attribute{
				Computed:    true,
				Description: "number of max vCPUs.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Computed:    true,
				Description: "Maximum Bandwidth.",
			},
			"minbandwidth": schema.Int64Attribute{
				Computed:    true,
				Description: "Minimum Bandwidth.",
			},
			"instancecount": schema.Int64Attribute{
				Computed:    true,
				Description: "VPX will consume one instance and MPX will consume zero instance.",
			},
			"daystoexpiration": schema.Int64Attribute{
				Computed:    true,
				Description: "Days to expire.",
			},
		},
	}
}

// nscapacityDataSourceSetAttrFromGet projects a NITRO nscapacity GET response
// onto the data-source model. nscapacity is a singleton, so the ID is static.
// Attributes are filled from the GET (or left Null when the GET omits them) via
// the shared utils.MapGet* helpers.
func nscapacityDataSourceSetAttrFromGet(ctx context.Context, data *NscapacityDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nscapacityDataSourceSetAttrFromGet Function")

	// Singleton resource - static ID.
	data.Id = types.StringValue("nscapacity-config")

	data.Bandwidth = utils.MapGetInt64(g, "bandwidth")
	data.Edition = utils.MapGetString(g, "edition")
	data.Ignoreexpiry = utils.MapGetBool(g, "ignoreexpiry")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Platform = utils.MapGetString(g, "platform")
	data.Unit = utils.MapGetString(g, "unit")
	data.Username = utils.MapGetString(g, "username")

	// password is a write-only input the GET never returns -> Null.
	data.Password = types.StringNull()

	// vCPU license is derived from the presence of the read-only "vcpucount" key
	// (the ADC never returns a "vcpu" field on GET).
	if _, ok := g["vcpucount"]; ok {
		data.Vcpu = types.BoolValue(true)
	} else {
		data.Vcpu = types.BoolValue(false)
	}

	// Read-only capacity/license metadata.
	data.Actualbandwidth = utils.MapGetInt64(g, "actualbandwidth")
	data.Vcpucount = utils.MapGetInt64(g, "vcpucount")
	data.Maxvcpucount = utils.MapGetInt64(g, "maxvcpucount")
	data.Maxbandwidth = utils.MapGetInt64(g, "maxbandwidth")
	data.Minbandwidth = utils.MapGetInt64(g, "minbandwidth")
	data.Instancecount = utils.MapGetInt64(g, "instancecount")
	data.Daystoexpiration = utils.MapGetInt64(g, "daystoexpiration")
}
