package snmpuser

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpuserDataSourceModel is the data-source-specific model, decoupled from
// SnmpuserResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (engineid, storagetype, status). Every non-key attribute is Computed; the
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type SnmpuserDataSourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Authpasswd          types.String `tfsdk:"authpasswd"`
	AuthpasswdWo        types.String `tfsdk:"authpasswd_wo"`
	AuthpasswdWoVersion types.Int64  `tfsdk:"authpasswd_wo_version"`
	Authtype            types.String `tfsdk:"authtype"`
	Group               types.String `tfsdk:"group"`
	Name                types.String `tfsdk:"name"` // Required lookup key
	Privpasswd          types.String `tfsdk:"privpasswd"`
	PrivpasswdWo        types.String `tfsdk:"privpasswd_wo"`
	PrivpasswdWoVersion types.Int64  `tfsdk:"privpasswd_wo_version"`
	Privtype            types.String `tfsdk:"privtype"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/snmpuser.json). Never settable; populated from GET.
	Engineid    types.String `tfsdk:"engineid"`
	Storagetype types.String `tfsdk:"storagetype"`
	Status      types.String `tfsdk:"status"`
}

func SnmpuserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"authpasswd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the pass phrase includes one or more spaces, enclose it in double or single quotation marks (for example, \"my phrase\" or 'my phrase').",
			},
			"authpasswd_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the pass phrase includes one or more spaces, enclose it in double or single quotation marks (for example, \"my phrase\" or 'my phrase').",
			},
			"authpasswd_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a authpasswd_wo update.",
			},
			"authtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication algorithm used by the Citrix ADC and the SNMPv3 user for authenticating the communication between them. You must specify the same authentication algorithm when you configure the SNMPv3 user in the SNMP manager.",
			},
			"group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured SNMPv3 group to which to bind this SNMPv3 user. The access rights (bound SNMPv3 views) and security level set for this group are assigned to this user.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SNMPv3 user. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my user\" or 'my user').",
			},
			"privpasswd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the key includes one or more spaces, enclose it in double or single quotation marks (for example, \"my key\" or 'my key').",
			},
			"privpasswd_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the key includes one or more spaces, enclose it in double or single quotation marks (for example, \"my key\" or 'my key').",
			},
			"privpasswd_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a privpasswd_wo update.",
			},
			"privtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encryption algorithm used by the Citrix ADC and the SNMPv3 user for encrypting the communication between them. You must specify the same encryption algorithm when you configure the SNMPv3 user in the SNMP manager.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"engineid": schema.StringAttribute{
				Computed:    true,
				Description: "The context engine ID of the user.",
			},
			"storagetype": schema.StringAttribute{
				Computed:    true,
				Description: "The storage type for this user. Possible values: [ volatile, nonVolatile ].",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "The status of this user. Possible values: [ active ].",
			},
		},
	}
}

// snmpuserDataSourceSetAttrFromGet projects a NITRO snmpuser GET response onto
// the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) — no unknown->null resolution or plan preservation is
// required. The shared utils.MapGet* helpers implement that projection.
func snmpuserDataSourceSetAttrFromGet(ctx context.Context, data *SnmpuserDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In snmpuserDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Authtype = utils.MapGetString(g, "authtype")
	data.Group = utils.MapGetString(g, "group")
	data.Privtype = utils.MapGetString(g, "privtype")

	// authpasswd / privpasswd (+ their _wo and _wo_version trackers) are
	// write-only/secret or action-only inputs the GET never returns -> Null.
	data.Authpasswd = types.StringNull()
	data.AuthpasswdWo = types.StringNull()
	data.AuthpasswdWoVersion = types.Int64Null()
	data.Privpasswd = types.StringNull()
	data.PrivpasswdWo = types.StringNull()
	data.PrivpasswdWoVersion = types.Int64Null()

	// Read-only metadata.
	data.Engineid = utils.MapGetString(g, "engineid")
	data.Storagetype = utils.MapGetString(g, "storagetype")
	data.Status = utils.MapGetString(g, "status")
}
