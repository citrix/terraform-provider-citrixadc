package rdpclientprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RdpclientprofileDataSourceModel is the data-source-specific model, decoupled
// from RdpclientprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (builtin,
// feature). Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type RdpclientprofileDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Addusernameinrdpfile types.String `tfsdk:"addusernameinrdpfile"`
	Audiocapturemode     types.String `tfsdk:"audiocapturemode"`
	Keyboardhook         types.String `tfsdk:"keyboardhook"`
	Multimonitorsupport  types.String `tfsdk:"multimonitorsupport"`
	Name                 types.String `tfsdk:"name"` // Required lookup key
	Psk                  types.String `tfsdk:"psk"`
	PskWo                types.String `tfsdk:"psk_wo"`
	PskWoVersion         types.Int64  `tfsdk:"psk_wo_version"`
	Randomizerdpfilename types.String `tfsdk:"randomizerdpfilename"`
	Rdpcookievalidity    types.Int64  `tfsdk:"rdpcookievalidity"`
	Rdpcustomparams      types.String `tfsdk:"rdpcustomparams"`
	Rdpfilename          types.String `tfsdk:"rdpfilename"`
	Rdphost              types.String `tfsdk:"rdphost"`
	Rdplinkattribute     types.String `tfsdk:"rdplinkattribute"`
	Rdplistener          types.String `tfsdk:"rdplistener"`
	Rdpurlmaxlen         types.Int64  `tfsdk:"rdpurlmaxlen"`
	Rdpurlmaxlencheck    types.String `tfsdk:"rdpurlmaxlencheck"`
	Rdpurloverride       types.String `tfsdk:"rdpurloverride"`
	Rdpvalidateclientip  types.String `tfsdk:"rdpvalidateclientip"`
	Redirectclipboard    types.String `tfsdk:"redirectclipboard"`
	Redirectcomports     types.String `tfsdk:"redirectcomports"`
	Redirectdrives       types.String `tfsdk:"redirectdrives"`
	Redirectpnpdevices   types.String `tfsdk:"redirectpnpdevices"`
	Redirectprinters     types.String `tfsdk:"redirectprinters"`
	Videoplaybackmode    types.String `tfsdk:"videoplaybackmode"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/rdpclientprofile.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func RdpclientprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addusernameinrdpfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add username in rdp file.",
			},
			"audiocapturemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selections in the Remote audio area on the Local Resources tab under Options in RDC.",
			},
			"keyboardhook": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selection in the Keyboard drop-down list on the Local Resources tab under Options in RDC.",
			},
			"multimonitorsupport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable Multiple Monitor Support for Remote Desktop Connection (RDC).",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the rdp profile",
			},
			"psk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pre shared key value",
			},
			"psk_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Pre shared key value",
			},
			"psk_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a psk_wo update.",
			},
			"randomizerdpfilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Will generate unique filename everytime rdp file is downloaded by appending output of time() function in the format <rdpfileName>_<time>.rdp. This tries to avoid the pop-up for replacement of existing rdp file during each rdp connection launch, hence providing better end-user experience.",
			},
			"rdpcookievalidity": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RDP cookie validity period. RDP cookie validity time is applicable for new connection and also for any re-connection that might happen, mostly due to network disruption or during fail-over.",
			},
			"rdpcustomparams": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option for RDP custom parameters settings (if any). Custom params needs to be separated by '&'",
			},
			"rdpfilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RDP file name to be sent to End User",
			},
			"rdphost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully-qualified domain name (FQDN) of the RDP Listener.",
			},
			"rdplinkattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix Gateway allows the configuration of rdpLinkAttribute parameter which can be used to fetch a list of RDP servers(IP/FQDN) that a user can access, from an Authentication server attribute(Example: LDAP, SAML). Based on the list received, the RDP links will be generated and displayed to the user.\n            Note: The Attribute mentioned in the rdpLinkAttribute should be fetched through corresponding authentication method.",
			},
			"rdplistener": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address (or) Fully-qualified domain name(FQDN) of the RDP Listener with the port in the format IP:Port (or) FQDN:Port",
			},
			"rdpurlmaxlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates the permissible max length of the RDP URL. Set to 256 by default.",
			},
			"rdpurlmaxlencheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting determines whether the RDP URL max length check is enforced during RDP file generation.",
			},
			"rdpurloverride": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting determines whether the RDP parameters supplied in the vpn url override those specified in the RDP profile.",
			},
			"rdpvalidateclientip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting determines whether RDC launch is initiated by the valid client IP",
			},
			"redirectclipboard": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the Clipboard check box on the Local Resources tab under Options in RDC.",
			},
			"redirectcomports": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selections for comports under More on the Local Resources tab under Options in RDC.",
			},
			"redirectdrives": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selections for Drives under More on the Local Resources tab under Options in RDC.",
			},
			"redirectpnpdevices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selections for pnpdevices under More on the Local Resources tab under Options in RDC.",
			},
			"redirectprinters": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selection in the Printers check box on the Local Resources tab under Options in RDC.",
			},
			"videoplaybackmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting determines if Remote Desktop Connection (RDC) will use RDP efficient multimedia streaming for video playback.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// rdpclientprofileDataSourceSetAttrFromGet projects a NITRO rdpclientprofile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func rdpclientprofileDataSourceSetAttrFromGet(ctx context.Context, data *RdpclientprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In rdpclientprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Addusernameinrdpfile = utils.MapGetString(g, "addusernameinrdpfile")
	data.Audiocapturemode = utils.MapGetString(g, "audiocapturemode")
	data.Keyboardhook = utils.MapGetString(g, "keyboardhook")
	data.Multimonitorsupport = utils.MapGetString(g, "multimonitorsupport")
	data.Randomizerdpfilename = utils.MapGetString(g, "randomizerdpfilename")
	data.Rdpcookievalidity = utils.MapGetInt64(g, "rdpcookievalidity")
	data.Rdpcustomparams = utils.MapGetString(g, "rdpcustomparams")
	data.Rdpfilename = utils.MapGetString(g, "rdpfilename")
	data.Rdphost = utils.MapGetString(g, "rdphost")
	data.Rdplinkattribute = utils.MapGetString(g, "rdplinkattribute")
	data.Rdplistener = utils.MapGetString(g, "rdplistener")
	data.Rdpurlmaxlen = utils.MapGetInt64(g, "rdpurlmaxlen")
	data.Rdpurlmaxlencheck = utils.MapGetString(g, "rdpurlmaxlencheck")
	data.Rdpurloverride = utils.MapGetString(g, "rdpurloverride")
	data.Rdpvalidateclientip = utils.MapGetString(g, "rdpvalidateclientip")
	data.Redirectclipboard = utils.MapGetString(g, "redirectclipboard")
	data.Redirectcomports = utils.MapGetString(g, "redirectcomports")
	data.Redirectdrives = utils.MapGetString(g, "redirectdrives")
	data.Redirectpnpdevices = utils.MapGetString(g, "redirectpnpdevices")
	data.Redirectprinters = utils.MapGetString(g, "redirectprinters")
	data.Videoplaybackmode = utils.MapGetString(g, "videoplaybackmode")

	// psk / psk_wo / psk_wo_version are write-only secret / version-tracker inputs
	// the GET never returns -> Null.
	data.Psk = types.StringNull()
	data.PskWo = types.StringNull()
	data.PskWoVersion = types.Int64Null()

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
