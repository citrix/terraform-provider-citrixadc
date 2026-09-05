package subscribergxinterface

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SubscribergxinterfaceDataSourceModel is the data-source-specific model,
// decoupled from SubscribergxinterfaceResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// The Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares.
type SubscribergxinterfaceDataSourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Cerrequesttimeout         types.Int64  `tfsdk:"cerrequesttimeout"`
	Healthcheck               types.String `tfsdk:"healthcheck"`
	Healthcheckttl            types.Int64  `tfsdk:"healthcheckttl"`
	Holdonsubscriberabsence   types.String `tfsdk:"holdonsubscriberabsence"`
	Idlettl                   types.Int64  `tfsdk:"idlettl"`
	Negativettl               types.Int64  `tfsdk:"negativettl"`
	Negativettllimitedsuccess types.String `tfsdk:"negativettllimitedsuccess"`
	Nodeid                    types.Int64  `tfsdk:"nodeid"`
	Pcrfrealm                 types.String `tfsdk:"pcrfrealm"`
	Purgesdbongxfailure       types.String `tfsdk:"purgesdbongxfailure"`
	Requestretryattempts      types.Int64  `tfsdk:"requestretryattempts"`
	Requesttimeout            types.Int64  `tfsdk:"requesttimeout"`
	Revalidationtimeout       types.Int64  `tfsdk:"revalidationtimeout"`
	Service                   types.String `tfsdk:"service"`
	Servicepathavp            types.List   `tfsdk:"servicepathavp"`
	Servicepathvendorid       types.Int64  `tfsdk:"servicepathvendorid"`
	Vserver                   types.String `tfsdk:"vserver"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/subscribergxinterface.json). Never settable; from GET.
	Svrstate                types.String `tfsdk:"svrstate"`
	Identity                types.String `tfsdk:"identity"`
	Realm                   types.String `tfsdk:"realm"`
	Status                  types.String `tfsdk:"status"`
	Servicepathinfomode     types.String `tfsdk:"servicepathinfomode"`
	Gxreportingavp1         types.List   `tfsdk:"gxreportingavp1"`
	Gxreportingavp1vendorid types.Int64  `tfsdk:"gxreportingavp1vendorid"`
	Gxreportingavp1type     types.String `tfsdk:"gxreportingavp1type"`
	Gxreportingavp2         types.List   `tfsdk:"gxreportingavp2"`
	Gxreportingavp2vendorid types.Int64  `tfsdk:"gxreportingavp2vendorid"`
	Gxreportingavp2type     types.String `tfsdk:"gxreportingavp2type"`
	Gxreportingavp3         types.List   `tfsdk:"gxreportingavp3"`
	Gxreportingavp3vendorid types.Int64  `tfsdk:"gxreportingavp3vendorid"`
	Gxreportingavp3type     types.String `tfsdk:"gxreportingavp3type"`
	Gxreportingavp4         types.List   `tfsdk:"gxreportingavp4"`
	Gxreportingavp4vendorid types.Int64  `tfsdk:"gxreportingavp4vendorid"`
	Gxreportingavp4type     types.String `tfsdk:"gxreportingavp4type"`
	Gxreportingavp5         types.List   `tfsdk:"gxreportingavp5"`
	Gxreportingavp5vendorid types.Int64  `tfsdk:"gxreportingavp5vendorid"`
	Gxreportingavp5type     types.String `tfsdk:"gxreportingavp5type"`
}

func SubscribergxinterfaceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cerrequesttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Healthcheck request timeout, in seconds, after which the Citrix ADC considers that no CCA packet received to the initiated CCR. After this time Citrix ADC should send again CCR to PCRF server. !",
			},
			"healthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Set this setting to yes if Citrix ADC should send DWR packets to PCRF server. When the session is idle, healthcheck timer expires and DWR packets are initiated in order to check that PCRF server is active. By default set to No. !",
			},
			"healthcheckttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Healthcheck timeout, in seconds, after which the DWR will be sent in order to ensure the state of the PCRF server. Any CCR, CCA, RAR or RRA message resets the timer. !",
			},
			"holdonsubscriberabsence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set this setting to yes if Citrix ADC needs to Hold pakcets till subscriber session is fetched from PCRF. Else set to NO. By default set to yes. If this setting is set to NO, then till Citrix ADC fetches subscriber from PCRF, default subscriber profile will be applied to this subscriber if configured. If default subscriber profile is also not configured an undef would be raised to expressions which use Subscriber attributes.",
			},
			"idlettl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Idle Time, in seconds, after which the Gx CCR-U request will be sent after any PCRF activity on a session. Any RAR or CCA message resets the timer.\nZero value disables the idle timeout. !",
			},
			"negativettl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Negative TTL, in seconds, after which the Gx CCR-I request will be resent for sessions that have not been resolved by PCRF due to server being down or no response or failed response. Instead of polling the PCRF server constantly, negative-TTL makes Citrix ADC stick to un-resolved session. Meanwhile Citrix ADC installs a negative session to avoid going to PCRF.\nFor Negative Sessions, Netcaler inherits the attributes from default subscriber profile if default subscriber is configured. A default subscriber could be configured as 'add subscriber profile *'. Or these attributes can be inherited from Radius as well if Radius is configued.\nZero value disables the Negative Sessions. And Citrix ADC does not install Negative sessions even if subscriber session could not be fetched. !",
			},
			"negativettllimitedsuccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set this to YES if Citrix ADC should create negative session for Result-Code DIAMETER_LIMITED_SUCCESS (2002) received in CCA-I. If set to NO, regular session is created.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"pcrfrealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRF realm is of type DiameterIdentity and contains the realm of PCRF to which the message is to be routed. This is the realm used in Destination-Realm AVP by Citrix ADC Gx client (as a Diameter node).",
			},
			"purgesdbongxfailure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set this setting to YES if needed to purge Subscriber Database in case of Gx failure. By default set to NO.",
			},
			"requestretryattempts": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "If the request does not complete within requestTimeout time, the request is retransmitted for requestRetryAttempts time.",
			},
			"requesttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Time, in seconds, within which the Gx CCR request must complete. If the request does not complete within this time, the request is retransmitted for requestRetryAttempts time. If still reuqest is not complete then default subscriber profile will be applied to this subscriber if configured. If default subscriber profile is also not configured an undef would be raised to expressions which use Subscriber attributes.\nZero disables the timeout. !",
			},
			"revalidationtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Revalidation Timeout, in seconds, after which the Gx CCR-U request will be sent after any PCRF activity on a session. Any RAR or CCA message resets the timer.\nZero value disables the idle timeout. !",
			},
			"service": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of DIAMETER/SSL_DIAMETER service corresponding to PCRF to which the Gx connection is established. The service type of the service must be DIAMETER/SSL_DIAMETER. Mutually exclusive with vserver parameter. Therefore, you cannot set both Service and the Virtual Server in the Gx Interface.",
			},
			"servicepathavp": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The AVP code in which PCRF sends service path applicable for subscriber.",
			},
			"servicepathvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The vendorid of the AVP in which PCRF sends service path for subscriber.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing, or content switching vserver to which the Gx connections are established. The service type of the virtual server must be DIAMETER/SSL_DIAMETER. Mutually exclusive with the service parameter. Therefore, you cannot set both service and the Virtual Server in the Gx Interface.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"svrstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the gx service.",
			},
			"identity": schema.StringAttribute{
				Computed:    true,
				Description: "DiameterIdentity to be used by NS. DiameterIdentity is used to identify a Diameter node uniquely.",
			},
			"realm": schema.StringAttribute{
				Computed:    true,
				Description: "Diameter Realm to be used by NS.",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "Citrix ADC PCRF connection Status. (Gx Protocol State).",
			},
			"servicepathinfomode": schema.StringAttribute{
				Computed:    true,
				Description: "The type of info in which service path is passed from PCRF.",
			},
			"gxreportingavp1": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "The AVP code to report as 1st custom AVP.",
			},
			"gxreportingavp1vendorid": schema.Int64Attribute{
				Computed:    true,
				Description: "The vendorid of the AVP which will be reported as 1st custom AVP.",
			},
			"gxreportingavp1type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of the AVP which will be reported as 1st custom AVP.",
			},
			"gxreportingavp2": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "The AVP code to report as 2nd custom AVP.",
			},
			"gxreportingavp2vendorid": schema.Int64Attribute{
				Computed:    true,
				Description: "The vendorid of the AVP which will be reported as 2nd custom AVP.",
			},
			"gxreportingavp2type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of the AVP which will be reported as 2nd custom AVP.",
			},
			"gxreportingavp3": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "The AVP code to report as 3rd custom AVP.",
			},
			"gxreportingavp3vendorid": schema.Int64Attribute{
				Computed:    true,
				Description: "The vendorid of the AVP which will be reported as 3rd custom AVP.",
			},
			"gxreportingavp3type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of the AVP which will be reported as 3rd custom AVP.",
			},
			"gxreportingavp4": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "The AVP code to report as 4th custom AVP.",
			},
			"gxreportingavp4vendorid": schema.Int64Attribute{
				Computed:    true,
				Description: "The vendorid of the AVP which will be reported as 4th custom AVP.",
			},
			"gxreportingavp4type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of the AVP which will be reported as 4th custom AVP.",
			},
			"gxreportingavp5": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "The AVP code to report as 5th custom AVP.",
			},
			"gxreportingavp5vendorid": schema.Int64Attribute{
				Computed:    true,
				Description: "The vendorid of the AVP which will be reported as 5th custom AVP.",
			},
			"gxreportingavp5type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of the AVP which will be reported as 5th custom AVP.",
			},
		},
	}
}

// subscribergxinterfaceDataSourceSetAttrFromGet projects a NITRO
// subscribergxinterface GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them). The shared utils.MapGet* helpers
// implement that projection.
func subscribergxinterfaceDataSourceSetAttrFromGet(ctx context.Context, data *SubscribergxinterfaceDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In subscribergxinterfaceDataSourceSetAttrFromGet Function")

	data.Cerrequesttimeout = utils.MapGetInt64(g, "cerrequesttimeout")
	data.Healthcheck = utils.MapGetString(g, "healthcheck")
	data.Healthcheckttl = utils.MapGetInt64(g, "healthcheckttl")
	data.Holdonsubscriberabsence = utils.MapGetString(g, "holdonsubscriberabsence")
	data.Idlettl = utils.MapGetInt64(g, "idlettl")
	data.Negativettl = utils.MapGetInt64(g, "negativettl")
	data.Negativettllimitedsuccess = utils.MapGetString(g, "negativettllimitedsuccess")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Pcrfrealm = utils.MapGetString(g, "pcrfrealm")
	data.Purgesdbongxfailure = utils.MapGetString(g, "purgesdbongxfailure")
	data.Requestretryattempts = utils.MapGetInt64(g, "requestretryattempts")
	data.Requesttimeout = utils.MapGetInt64(g, "requesttimeout")
	data.Revalidationtimeout = utils.MapGetInt64(g, "revalidationtimeout")
	data.Service = utils.MapGetString(g, "service")
	data.Servicepathvendorid = utils.MapGetInt64(g, "servicepathvendorid")
	data.Vserver = utils.MapGetString(g, "vserver")

	// servicepathavp is an Int64-typed list on the schema; convert explicitly.
	if val, ok := g["servicepathavp"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			int64List := make([]int64, 0, len(sliceVal))
			for _, item := range sliceVal {
				if intVal, err := utils.ConvertToInt64(item); err == nil {
					int64List = append(int64List, intVal)
				}
			}
			listValue, _ := types.ListValueFrom(ctx, types.Int64Type, int64List)
			data.Servicepathavp = listValue
		} else {
			data.Servicepathavp = types.ListNull(types.Int64Type)
		}
	} else {
		data.Servicepathavp = types.ListNull(types.Int64Type)
	}

	// Read-only attributes.
	data.Svrstate = utils.MapGetString(g, "svrstate")
	data.Identity = utils.MapGetString(g, "identity")
	data.Realm = utils.MapGetString(g, "realm")
	data.Status = utils.MapGetString(g, "status")
	data.Servicepathinfomode = utils.MapGetString(g, "servicepathinfomode")
	data.Gxreportingavp1 = utils.MapGetStringList(g, "gxreportingavp1")
	data.Gxreportingavp1vendorid = utils.MapGetInt64(g, "gxreportingavp1vendorid")
	data.Gxreportingavp1type = utils.MapGetString(g, "gxreportingavp1type")
	data.Gxreportingavp2 = utils.MapGetStringList(g, "gxreportingavp2")
	data.Gxreportingavp2vendorid = utils.MapGetInt64(g, "gxreportingavp2vendorid")
	data.Gxreportingavp2type = utils.MapGetString(g, "gxreportingavp2type")
	data.Gxreportingavp3 = utils.MapGetStringList(g, "gxreportingavp3")
	data.Gxreportingavp3vendorid = utils.MapGetInt64(g, "gxreportingavp3vendorid")
	data.Gxreportingavp3type = utils.MapGetString(g, "gxreportingavp3type")
	data.Gxreportingavp4 = utils.MapGetStringList(g, "gxreportingavp4")
	data.Gxreportingavp4vendorid = utils.MapGetInt64(g, "gxreportingavp4vendorid")
	data.Gxreportingavp4type = utils.MapGetString(g, "gxreportingavp4type")
	data.Gxreportingavp5 = utils.MapGetStringList(g, "gxreportingavp5")
	data.Gxreportingavp5vendorid = utils.MapGetInt64(g, "gxreportingavp5vendorid")
	data.Gxreportingavp5type = utils.MapGetString(g, "gxreportingavp5type")

	// Singleton resource: static ID.
	data.Id = types.StringValue("subscribergxinterface-config")
}
