package contentinspectionaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ContentinspectionactionDataSourceModel is the data-source-specific model,
// decoupled from ContentinspectionactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (reqtimeout, reqtimeoutaction, hits, referencecount, undefhits, builtin,
// feature). Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type ContentinspectionactionDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Icapprofilename types.String `tfsdk:"icapprofilename"`
	Ifserverdown    types.String `tfsdk:"ifserverdown"`
	Name            types.String `tfsdk:"name"`
	Serverip        types.String `tfsdk:"serverip"`
	Servername      types.String `tfsdk:"servername"`
	Serverport      types.Int64  `tfsdk:"serverport"`
	Type            types.String `tfsdk:"type"`
	Wasmprofilename types.String `tfsdk:"wasmprofilename"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/contentinspectionaction.json). Never settable;
	// populated from GET.
	Reqtimeout       types.Int64  `tfsdk:"reqtimeout"`
	Reqtimeoutaction types.String `tfsdk:"reqtimeoutaction"`
	Hits             types.Int64  `tfsdk:"hits"`
	Referencecount   types.Int64  `tfsdk:"referencecount"`
	Undefhits        types.Int64  `tfsdk:"undefhits"`
	Builtin          types.List   `tfsdk:"builtin"`
	Feature          types.String `tfsdk:"feature"`
}

func ContentinspectionactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"icapprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ICAP profile to be attached to the contentInspection action.",
			},
			"ifserverdown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the action to perform if the Vserver representing the remote service is not UP. This is not supported for NOINSPECTION Type. The Supported actions are:\n* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.\n* CONTINUE - It bypasses the ContentIsnpection and Continues/resumes the Traffic-Flow to Client/Server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the remote service action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of remoteService",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LB vserver or service",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port of remoteService",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of operation this action is going to perform. following actions are available to configure:\n* ICAP - forward the incoming request or response to an ICAP server for modification.\n* INLINEINSPECTION - forward the incoming or outgoing packets to IPS server for Intrusion Prevention.\n* MIRROR - Forwards cloned packets for Intrusion Detection.\n* NOINSPECTION - This does not forward incoming and outgoing packets to the Inspection device.\n* NSTRACE - capture current and further incoming packets on this transaction.",
			},
			"wasmprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the CI WASM profile to be attached to the contentInspection action.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"reqtimeout": schema.Int64Attribute{
				Computed:    true,
				Description: "Time, in seconds, within which the remote service request must complete. This is not supported for NOINSPECTION action type. If the request does not complete within this time, the specified request timeout action is executed. Zero disables the timeout.",
			},
			"reqtimeoutaction": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the action to perform if the Vserver representing the remote service does not respond. This is not supported for NOINSPECTION action type. Possible values = BYPASS, DROP, RESET.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action has been taken.",
			},
			"referencecount": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of references to the action.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action resulted in UNDEF.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether contentinspection action is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// contentinspectionactionDataSourceSetAttrFromGet projects a NITRO
// contentinspectionaction GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func contentinspectionactionDataSourceSetAttrFromGet(ctx context.Context, data *ContentinspectionactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In contentinspectionactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Icapprofilename = utils.MapGetString(g, "icapprofilename")
	data.Ifserverdown = utils.MapGetString(g, "ifserverdown")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Servername = utils.MapGetString(g, "servername")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Type = utils.MapGetString(g, "type")
	data.Wasmprofilename = utils.MapGetString(g, "wasmprofilename")

	// Read-only attributes.
	data.Reqtimeout = utils.MapGetInt64(g, "reqtimeout")
	data.Reqtimeoutaction = utils.MapGetString(g, "reqtimeoutaction")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
