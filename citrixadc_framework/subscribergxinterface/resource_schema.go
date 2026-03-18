package subscribergxinterface

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SubscribergxinterfaceResourceModel describes the resource data model.
type SubscribergxinterfaceResourceModel struct {
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
}

func (r *SubscribergxinterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the subscribergxinterface resource.",
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
				Default:     int64default.StaticInt64(30),
				Description: "q!Healthcheck timeout, in seconds, after which the DWR will be sent in order to ensure the state of the PCRF server. Any CCR, CCA, RAR or RRA message resets the timer. !",
			},
			"holdonsubscriberabsence": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Set this setting to yes if Citrix ADC needs to Hold pakcets till subscriber session is fetched from PCRF. Else set to NO. By default set to yes. If this setting is set to NO, then till Citrix ADC fetches subscriber from PCRF, default subscriber profile will be applied to this subscriber if configured. If default subscriber profile is also not configured an undef would be raised to expressions which use Subscriber attributes.",
			},
			"idlettl": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(900),
				Description: "q!Idle Time, in seconds, after which the Gx CCR-U request will be sent after any PCRF activity on a session. Any RAR or CCA message resets the timer.\nZero value disables the idle timeout. !",
			},
			"negativettl": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(600),
				Description: "q!Negative TTL, in seconds, after which the Gx CCR-I request will be resent for sessions that have not been resolved by PCRF due to server being down or no response or failed response. Instead of polling the PCRF server constantly, negative-TTL makes Citrix ADC stick to un-resolved session. Meanwhile Citrix ADC installs a negative session to avoid going to PCRF.\nFor Negative Sessions, Netcaler inherits the attributes from default subscriber profile if default subscriber is configured. A default subscriber could be configured as 'add subscriber profile *'. Or these attributes can be inherited from Radius as well if Radius is configued.\nZero value disables the Negative Sessions. And Citrix ADC does not install Negative sessions even if subscriber session could not be fetched. !",
			},
			"negativettllimitedsuccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set this to YES if Citrix ADC should create negative session for Result-Code DIAMETER_LIMITED_SUCCESS (2002) received in CCA-I. If set to NO, regular session is created.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
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
				Default:     int64default.StaticInt64(3),
				Description: "If the request does not complete within requestTimeout time, the request is retransmitted for requestRetryAttempts time.",
			},
			"requesttimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
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
		},
	}
}

func subscribergxinterfaceGetThePayloadFromtheConfig(ctx context.Context, data *SubscribergxinterfaceResourceModel) subscriber.Subscribergxinterface {
	tflog.Debug(ctx, "In subscribergxinterfaceGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	subscribergxinterface := subscriber.Subscribergxinterface{}
	if !data.Cerrequesttimeout.IsNull() {
		subscribergxinterface.Cerrequesttimeout = utils.IntPtr(int(data.Cerrequesttimeout.ValueInt64()))
	}
	if !data.Healthcheck.IsNull() {
		subscribergxinterface.Healthcheck = data.Healthcheck.ValueString()
	}
	if !data.Healthcheckttl.IsNull() {
		subscribergxinterface.Healthcheckttl = utils.IntPtr(int(data.Healthcheckttl.ValueInt64()))
	}
	if !data.Holdonsubscriberabsence.IsNull() {
		subscribergxinterface.Holdonsubscriberabsence = data.Holdonsubscriberabsence.ValueString()
	}
	if !data.Idlettl.IsNull() {
		subscribergxinterface.Idlettl = utils.IntPtr(int(data.Idlettl.ValueInt64()))
	}
	if !data.Negativettl.IsNull() {
		subscribergxinterface.Negativettl = utils.IntPtr(int(data.Negativettl.ValueInt64()))
	}
	if !data.Negativettllimitedsuccess.IsNull() {
		subscribergxinterface.Negativettllimitedsuccess = data.Negativettllimitedsuccess.ValueString()
	}
	if !data.Nodeid.IsNull() {
		subscribergxinterface.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Pcrfrealm.IsNull() {
		subscribergxinterface.Pcrfrealm = data.Pcrfrealm.ValueString()
	}
	if !data.Purgesdbongxfailure.IsNull() {
		subscribergxinterface.Purgesdbongxfailure = data.Purgesdbongxfailure.ValueString()
	}
	if !data.Requestretryattempts.IsNull() {
		subscribergxinterface.Requestretryattempts = utils.IntPtr(int(data.Requestretryattempts.ValueInt64()))
	}
	if !data.Requesttimeout.IsNull() {
		subscribergxinterface.Requesttimeout = utils.IntPtr(int(data.Requesttimeout.ValueInt64()))
	}
	if !data.Revalidationtimeout.IsNull() {
		subscribergxinterface.Revalidationtimeout = utils.IntPtr(int(data.Revalidationtimeout.ValueInt64()))
	}
	if !data.Service.IsNull() {
		subscribergxinterface.Service = data.Service.ValueString()
	}
	if !data.Servicepathvendorid.IsNull() {
		subscribergxinterface.Servicepathvendorid = utils.IntPtr(int(data.Servicepathvendorid.ValueInt64()))
	}
	if !data.Vserver.IsNull() {
		subscribergxinterface.Vserver = data.Vserver.ValueString()
	}

	return subscribergxinterface
}

func subscribergxinterfaceSetAttrFromGet(ctx context.Context, data *SubscribergxinterfaceResourceModel, getResponseData map[string]interface{}) *SubscribergxinterfaceResourceModel {
	tflog.Debug(ctx, "In subscribergxinterfaceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cerrequesttimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cerrequesttimeout = types.Int64Value(intVal)
		}
	} else {
		data.Cerrequesttimeout = types.Int64Null()
	}
	if val, ok := getResponseData["healthcheck"]; ok && val != nil {
		data.Healthcheck = types.StringValue(val.(string))
	} else {
		data.Healthcheck = types.StringNull()
	}
	if val, ok := getResponseData["healthcheckttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Healthcheckttl = types.Int64Value(intVal)
		}
	} else {
		data.Healthcheckttl = types.Int64Null()
	}
	if val, ok := getResponseData["holdonsubscriberabsence"]; ok && val != nil {
		data.Holdonsubscriberabsence = types.StringValue(val.(string))
	} else {
		data.Holdonsubscriberabsence = types.StringNull()
	}
	if val, ok := getResponseData["idlettl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Idlettl = types.Int64Value(intVal)
		}
	} else {
		data.Idlettl = types.Int64Null()
	}
	if val, ok := getResponseData["negativettl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Negativettl = types.Int64Value(intVal)
		}
	} else {
		data.Negativettl = types.Int64Null()
	}
	if val, ok := getResponseData["negativettllimitedsuccess"]; ok && val != nil {
		data.Negativettllimitedsuccess = types.StringValue(val.(string))
	} else {
		data.Negativettllimitedsuccess = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["pcrfrealm"]; ok && val != nil {
		data.Pcrfrealm = types.StringValue(val.(string))
	} else {
		data.Pcrfrealm = types.StringNull()
	}
	if val, ok := getResponseData["purgesdbongxfailure"]; ok && val != nil {
		data.Purgesdbongxfailure = types.StringValue(val.(string))
	} else {
		data.Purgesdbongxfailure = types.StringNull()
	}
	if val, ok := getResponseData["requestretryattempts"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Requestretryattempts = types.Int64Value(intVal)
		}
	} else {
		data.Requestretryattempts = types.Int64Null()
	}
	if val, ok := getResponseData["requesttimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Requesttimeout = types.Int64Value(intVal)
		}
	} else {
		data.Requesttimeout = types.Int64Null()
	}
	if val, ok := getResponseData["revalidationtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Revalidationtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Revalidationtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["service"]; ok && val != nil {
		data.Service = types.StringValue(val.(string))
	} else {
		data.Service = types.StringNull()
	}
	if val, ok := getResponseData["servicepathvendorid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Servicepathvendorid = types.Int64Value(intVal)
		}
	} else {
		data.Servicepathvendorid = types.Int64Null()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else {
		data.Vserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("subscribergxinterface-config")

	return data
}
