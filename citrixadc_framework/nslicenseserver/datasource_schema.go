package nslicenseserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NslicenseserverDataSourceModel describes the DATASOURCE data model. It mirrors
// the configurable attributes surfaced by the datasource PLUS the read-only
// license-server metadata the NITRO `nslicenseserver` GET returns. It is
// decoupled from the resource model so the data source can expose the full GET
// projection (including GET-only fields the resource intentionally omits).
type NslicenseserverDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Deviceprofilename types.String `tfsdk:"deviceprofilename"`
	Forceupdateip     types.Bool   `tfsdk:"forceupdateip"`
	Licensemode       types.String `tfsdk:"licensemode"`
	Licenseserverip   types.String `tfsdk:"licenseserverip"`
	Nodeid            types.Int64  `tfsdk:"nodeid"`
	Password          types.String `tfsdk:"password"`
	Port              types.Int64  `tfsdk:"port"`
	Servername        types.String `tfsdk:"servername"`
	Username          types.String `tfsdk:"username"`

	// Read-only (GET-only) license-server metadata from the NITRO doc read-only
	// set (zion73x_readonly/nslicenseserver.json). Never settable; from GET.
	Status     types.Int64  `tfsdk:"status"`
	Grace      types.Int64  `tfsdk:"grace"`
	Gptimeleft types.Int64  `tfsdk:"gptimeleft"`
	Type       types.String `tfsdk:"type"`
}

// nslicenseserverDataSourceSetAttrFromGet projects a NITRO nslicenseserver GET
// response onto the data-source model using the shared utils.MapGet* helpers.
// Attributes the GET omits are left Null.
func nslicenseserverDataSourceSetAttrFromGet(ctx context.Context, data *NslicenseserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nslicenseserverDataSourceSetAttrFromGet Function")

	data.Deviceprofilename = utils.MapGetString(g, "deviceprofilename")
	data.Forceupdateip = utils.MapGetBool(g, "forceupdateip")
	data.Licensemode = utils.MapGetString(g, "licensemode")
	data.Licenseserverip = utils.MapGetString(g, "licenseserverip")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Password = utils.MapGetString(g, "password")
	data.Port = utils.MapGetInt64(g, "port")
	data.Servername = utils.MapGetString(g, "servername")
	data.Username = utils.MapGetString(g, "username")

	// Read-only license-server metadata.
	data.Status = utils.MapGetInt64(g, "status")
	data.Grace = utils.MapGetInt64(g, "grace")
	data.Gptimeleft = utils.MapGetInt64(g, "gptimeleft")
	data.Type = utils.MapGetString(g, "type")

	// ID matches the resource: servername.
	data.Id = types.StringValue(data.Servername.ValueString())
}

func NslicenseserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"deviceprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Device profile is created on ADM and contains the user name and password of the instance(s). ADM will use this info to add the NS for registration",
			},
			"forceupdateip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If this flag is used while adding the licenseserver, existing config will be overwritten. Use this flag only if you are sure that the new licenseserver has the required capacity.",
			},
			"licensemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This paramter indicates type of license customer interested while configuring add/set licenseserver",
			},
			"licenseserverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the License server. Either licenseserverip or servername must be specified.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"password": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Password to use when authenticating with ADM Agent for LAS licensing.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "License server port.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the License server. Either licenseserverip or servername must be specified.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},

			// Read-only (GET-only) license-server metadata surfaced by the data source.
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "Status of the license server. Null when the appliance omits it.",
			},
			"grace": schema.Int64Attribute{
				Computed:    true,
				Description: "Grace status of the server. Null when the appliance omits it.",
			},
			"gptimeleft": schema.Int64Attribute{
				Computed:    true,
				Description: "Grace time left. Null when the appliance omits it.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "License server type (ADM/CLA). Null when the appliance omits it.",
			},
		},
	}
}
