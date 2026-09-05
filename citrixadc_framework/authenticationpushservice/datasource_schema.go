package authenticationpushservice

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationpushserviceDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationpushserviceResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (namespace, hubname, servicekey, servicekeyname, certendpoint,
// pushservicestatus, trustservice, pushcloudserverstatus, signingkeyname,
// signingkey). Every non-key attribute is Computed.
type AuthenticationpushserviceDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Clientid        types.String `tfsdk:"clientid"`
	Clientsecret    types.String `tfsdk:"clientsecret"`
	Customerid      types.String `tfsdk:"customerid"`
	Name            types.String `tfsdk:"name"`
	Refreshinterval types.Int64  `tfsdk:"refreshinterval"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationpushservice.json). Never settable;
	// populated from GET. NOTE: NITRO returns the namespace key capitalized
	// ("Namespace"); the Terraform attribute is lowercased ("namespace").
	Namespace             types.String `tfsdk:"namespace"`
	Hubname               types.String `tfsdk:"hubname"`
	Servicekey            types.String `tfsdk:"servicekey"`
	Servicekeyname        types.String `tfsdk:"servicekeyname"`
	Certendpoint          types.String `tfsdk:"certendpoint"`
	Pushservicestatus     types.String `tfsdk:"pushservicestatus"`
	Trustservice          types.String `tfsdk:"trustservice"`
	Pushcloudserverstatus types.String `tfsdk:"pushcloudserverstatus"`
	Signingkeyname        types.String `tfsdk:"signingkeyname"`
	Signingkey            types.String `tfsdk:"signingkey"`
}

func AuthenticationpushserviceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique identity for communicating with Citrix Push server in cloud.",
			},
			"clientsecret": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Unique secret for communicating with Citrix Push server in cloud.",
			},
			"customerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Customer id/name of the account in cloud that is used to create clientid/secret pair.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the push service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my push service\" or 'my push service').",
			},
			"refreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval at which certificates or idtoken is refreshed.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"namespace": schema.StringAttribute{
				Computed:    true,
				Description: "Fully qualified domain name of the notification service in the cloud. If omitted, namespace defaults to https://mfa.cloud.com/.",
			},
			"hubname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the hub within a namespace. This is used to classify different identities within a namespace.",
			},
			"servicekey": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "Key to be used to compute signature necessary for registering to notification service.",
			},
			"servicekeyname": schema.StringAttribute{
				Computed:    true,
				Description: "Friendly name of the Key to be used to compute signature necessary for registering to notification service.",
			},
			"certendpoint": schema.StringAttribute{
				Computed:    true,
				Description: "URL of the endpoint that contains JWKs (Json Web Key) for JWT (Json Web Token) verification. This is used at cloud instance that offers push service.",
			},
			"pushservicestatus": schema.StringAttribute{
				Computed:    true,
				Description: "Describes status of push service. Possible values: INIT, CERTFETCH, CCTOKEN, COMPLETE.",
			},
			"trustservice": schema.StringAttribute{
				Computed:    true,
				Description: "URL of the service that generates tokens for cloud access.",
			},
			"pushcloudserverstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Describes status of the cloud service that does push. Possible values: UP, DOWN.",
			},
			"signingkeyname": schema.StringAttribute{
				Computed:    true,
				Description: "Friendly name of the Key to be used to compute signature necessary for accessing notification service.",
			},
			"signingkey": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "Key to be used to compute signature necessary for accessing notification service.",
			},
		},
	}
}

// authenticationpushserviceDataSourceSetAttrFromGet projects a NITRO
// authenticationpushservice GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func authenticationpushserviceDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationpushserviceDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationpushserviceDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Clientid = utils.MapGetString(g, "clientid")
	data.Customerid = utils.MapGetString(g, "customerid")
	data.Refreshinterval = utils.MapGetInt64(g, "refreshinterval")

	// clientsecret is a secret input the GET never returns -> Null.
	data.Clientsecret = types.StringNull()

	// Read-only metadata. NITRO returns the namespace key capitalized.
	data.Namespace = utils.MapGetString(g, "Namespace")
	data.Hubname = utils.MapGetString(g, "hubname")
	data.Servicekey = utils.MapGetString(g, "servicekey")
	data.Servicekeyname = utils.MapGetString(g, "servicekeyname")
	data.Certendpoint = utils.MapGetString(g, "certendpoint")
	data.Pushservicestatus = utils.MapGetString(g, "pushservicestatus")
	data.Trustservice = utils.MapGetString(g, "trustservice")
	data.Pushcloudserverstatus = utils.MapGetString(g, "pushcloudserverstatus")
	data.Signingkeyname = utils.MapGetString(g, "signingkeyname")
	data.Signingkey = utils.MapGetString(g, "signingkey")
}
