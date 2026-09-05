package icaaccessprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IcaaccessprofileDataSourceModel is the data-source-specific model, decoupled
// from IcaaccessprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only NITRO attributes the resource deliberately omits
// (refcnt, builtin, feature, isdefault). Every non-key attribute is Computed;
// the Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares, which is why
// it cannot reuse the resource model.
type IcaaccessprofileDataSourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Clientaudioredirection       types.String `tfsdk:"clientaudioredirection"`
	Clientclipboardredirection   types.String `tfsdk:"clientclipboardredirection"`
	Clientcomportredirection     types.String `tfsdk:"clientcomportredirection"`
	Clientdriveredirection       types.String `tfsdk:"clientdriveredirection"`
	Clientprinterredirection     types.String `tfsdk:"clientprinterredirection"`
	Clienttwaindeviceredirection types.String `tfsdk:"clienttwaindeviceredirection"`
	Clientusbdriveredirection    types.String `tfsdk:"clientusbdriveredirection"`
	Connectclientlptports        types.String `tfsdk:"connectclientlptports"`
	Draganddrop                  types.String `tfsdk:"draganddrop"`
	Fido2redirection             types.String `tfsdk:"fido2redirection"`
	Localremotedatasharing       types.String `tfsdk:"localremotedatasharing"`
	Multistream                  types.String `tfsdk:"multistream"`
	Name                         types.String `tfsdk:"name"` // Required lookup key
	Smartcardredirection         types.String `tfsdk:"smartcardredirection"`
	Wiaredirection               types.String `tfsdk:"wiaredirection"`

	// Read-only (GET-only) NITRO attributes from the read-only set
	// (zion73x_readonly/icaaccessprofile.json). Never settable; populated from GET.
	Refcnt    types.Int64  `tfsdk:"refcnt"`
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
	Isdefault types.Bool   `tfsdk:"isdefault"`
}

func IcaaccessprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientaudioredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable applications hosted on the server to play sounds through a sound device installed on the client computer, also allows or prevents users to record audio input",
			},
			"clientclipboardredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable the clipboard on the client device to be mapped to the clipboard on the server",
			},
			"clientcomportredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable COM port redirection to and from the client",
			},
			"clientdriveredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disables drive redirection to and from the client",
			},
			"clientprinterredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable client printers to be mapped to a server when a user logs on to a session",
			},
			"clienttwaindeviceredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable TWAIN devices, such as digital cameras or scanners, on the client device from published image processing applications",
			},
			"clientusbdriveredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable the redirection of USB devices to and from the client",
			},
			"connectclientlptports": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable automatic connection of LPT ports from the client when the user logs on",
			},
			"draganddrop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable drag and drop between client and remote applications and desktops",
			},
			"fido2redirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable FIDO2 redirection",
			},
			"localremotedatasharing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable file/data sharing via the Receiver for HTML5",
			},
			"multistream": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable the multistream feature for the specified users",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the ICA accessprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and\nthe hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA accessprofile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica accessprofile\" or 'my ica accessprofile').\n\nEach of the features can be configured as DEFAULT/DISABLED.\nHere, DISABLED means that the policy settings on the backend XenApp/XenDesktop server are overridden and the Citrix ADC makes the decision to deny access. Whereas DEFAULT means that the Citrix ADC allows the request to reach the XenApp/XenDesktop that takes the decision to allow/deny access based on the policy configured on it. For example, if ClientAudioRedirection is enabled on the backend XenApp/XenDesktop server, and the configured profile has ClientAudioRedirection as DISABLED, the Citrix ADC makes the decision to deny the request irrespective of the configuration on the backend. If the configured profile has ClientAudioRedirection as DEFAULT, then the Citrix ADC forwards the requests to the backend XenApp/XenDesktop server.It then makes the decision to allow/deny access based on the policy configured on it.",
			},
			"smartcardredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable smart card redirection. Smart card virtual channel is always allowed in CVAD",
			},
			"wiaredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable WIA scanner redirection",
			},

			// Read-only (GET-only) NITRO attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of entities using this accessprofile.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that the ICA accessprofile is a built-in (SYSTEM INTERNAL) type (MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL).",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default accessprofile.",
			},
		},
	}
}

// icaaccessprofileDataSourceSetAttrFromGet projects a NITRO icaaccessprofile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) — no unknown->null resolution or plan preservation is
// required. The shared utils.MapGet* helpers implement that projection.
func icaaccessprofileDataSourceSetAttrFromGet(ctx context.Context, data *IcaaccessprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In icaaccessprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Clientaudioredirection = utils.MapGetString(g, "clientaudioredirection")
	data.Clientclipboardredirection = utils.MapGetString(g, "clientclipboardredirection")
	data.Clientcomportredirection = utils.MapGetString(g, "clientcomportredirection")
	data.Clientdriveredirection = utils.MapGetString(g, "clientdriveredirection")
	data.Clientprinterredirection = utils.MapGetString(g, "clientprinterredirection")
	data.Clienttwaindeviceredirection = utils.MapGetString(g, "clienttwaindeviceredirection")
	data.Clientusbdriveredirection = utils.MapGetString(g, "clientusbdriveredirection")
	data.Connectclientlptports = utils.MapGetString(g, "connectclientlptports")
	data.Draganddrop = utils.MapGetString(g, "draganddrop")
	data.Fido2redirection = utils.MapGetString(g, "fido2redirection")
	data.Localremotedatasharing = utils.MapGetString(g, "localremotedatasharing")
	data.Multistream = utils.MapGetString(g, "multistream")
	data.Smartcardredirection = utils.MapGetString(g, "smartcardredirection")
	data.Wiaredirection = utils.MapGetString(g, "wiaredirection")

	// Read-only NITRO attributes.
	data.Refcnt = utils.MapGetInt64(g, "refcnt")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
}
