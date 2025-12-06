package lbparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbParameterResourceModel describes the resource data model.
type LbParameterResourceModel struct {
	Id                            types.String `tfsdk:"id"`
	AllowBoundSvcRemoval          types.String `tfsdk:"allowboundsvcremoval"`
	ComputedAdcCookieAttribute    types.String `tfsdk:"computedadccookieattribute"`
	ConsolidatedLconn             types.String `tfsdk:"consolidatedlconn"`
	CookiePassphrase              types.String `tfsdk:"cookiepassphrase"`
	DbsTtl                        types.Int64  `tfsdk:"dbsttl"`
	DropMqttJumboMessage          types.String `tfsdk:"dropmqttjumbomessage"`
	HttpOnlyCookieFlag            types.String `tfsdk:"httponlycookieflag"`
	LiteralAdcCookieAttribute     types.String `tfsdk:"literaladccookieattribute"`
	MaxPipelineNat                types.Int64  `tfsdk:"maxpipelinenat"`
	MonitorConnectionClose        types.String `tfsdk:"monitorconnectionclose"`
	MonitorSkipMaxClient          types.String `tfsdk:"monitorskipmaxclient"`
	PreferDirectRoute             types.String `tfsdk:"preferdirectroute"`
	ProximityFromSelf             types.String `tfsdk:"proximityfromself"`
	RetainServiceState            types.String `tfsdk:"retainservicestate"`
	StartupRrFactor               types.Int64  `tfsdk:"startuprrfactor"`
	StoreMqttClientIdAndUsername  types.String `tfsdk:"storemqttclientidandusername"`
	SessionsThreshold             types.Int64  `tfsdk:"sessionsthreshold"`
	UndefAction                   types.String `tfsdk:"undefaction"`
	UseEncryptedPersistenceCookie types.String `tfsdk:"useencryptedpersistencecookie"`
	UsePortForHashLb              types.String `tfsdk:"useportforhashlb"`
	UseSecuredPersistenceCookie   types.String `tfsdk:"usesecuredpersistencecookie"`
	VserverSpecificMac            types.String `tfsdk:"vserverspecificmac"`
	LbHashAlgorithm               types.String `tfsdk:"lbhashalgorithm"`
	LbHashFingers                 types.Int64  `tfsdk:"lbhashfingers"`
}

func (r *LbParameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allowboundsvcremoval": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"computedadccookieattribute": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"consolidatedlconn": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cookiepassphrase": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"dbsttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"dropmqttjumbomessage": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"httponlycookieflag": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"literaladccookieattribute": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"maxpipelinenat": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"monitorconnectionclose": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"monitorskipmaxclient": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"preferdirectroute": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"proximityfromself": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"retainservicestate": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"startuprrfactor": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"storemqttclientidandusername": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"sessionsthreshold": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"undefaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"useencryptedpersistencecookie": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"useportforhashlb": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"usesecuredpersistencecookie": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"vserverspecificmac": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"lbhashalgorithm": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"lbhashfingers": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

func lbparameterGetThePayloadFromtheConfig(ctx context.Context, data *LbParameterResourceModel) lb.Lbparameter {
	tflog.Debug(ctx, "In lbparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbparameter := lb.Lbparameter{}

	if !data.AllowBoundSvcRemoval.IsNull() {
		lbparameter.Allowboundsvcremoval = data.AllowBoundSvcRemoval.ValueString()
	}
	if !data.ComputedAdcCookieAttribute.IsNull() {
		lbparameter.Computedadccookieattribute = data.ComputedAdcCookieAttribute.ValueString()
	}
	if !data.ConsolidatedLconn.IsNull() {
		lbparameter.Consolidatedlconn = data.ConsolidatedLconn.ValueString()
	}
	if !data.CookiePassphrase.IsNull() {
		lbparameter.Cookiepassphrase = data.CookiePassphrase.ValueString()
	}
	if !data.DbsTtl.IsNull() {
		lbparameter.Dbsttl = utils.IntPtr(int(data.DbsTtl.ValueInt64()))
	}
	if !data.DropMqttJumboMessage.IsNull() {
		lbparameter.Dropmqttjumbomessage = data.DropMqttJumboMessage.ValueString()
	}
	if !data.HttpOnlyCookieFlag.IsNull() {
		lbparameter.Httponlycookieflag = data.HttpOnlyCookieFlag.ValueString()
	}
	if !data.LiteralAdcCookieAttribute.IsNull() {
		lbparameter.Literaladccookieattribute = data.LiteralAdcCookieAttribute.ValueString()
	}
	if !data.MaxPipelineNat.IsNull() {
		lbparameter.Maxpipelinenat = utils.IntPtr(int(data.MaxPipelineNat.ValueInt64()))
	}
	if !data.MonitorConnectionClose.IsNull() {
		lbparameter.Monitorconnectionclose = data.MonitorConnectionClose.ValueString()
	}
	if !data.MonitorSkipMaxClient.IsNull() {
		lbparameter.Monitorskipmaxclient = data.MonitorSkipMaxClient.ValueString()
	}
	if !data.PreferDirectRoute.IsNull() {
		lbparameter.Preferdirectroute = data.PreferDirectRoute.ValueString()
	}
	if !data.ProximityFromSelf.IsNull() {
		lbparameter.Proximityfromself = data.ProximityFromSelf.ValueString()
	}
	if !data.RetainServiceState.IsNull() {
		lbparameter.Retainservicestate = data.RetainServiceState.ValueString()
	}
	if !data.StartupRrFactor.IsNull() {
		lbparameter.Startuprrfactor = utils.IntPtr(int(data.StartupRrFactor.ValueInt64()))
	}
	if !data.StoreMqttClientIdAndUsername.IsNull() {
		lbparameter.Storemqttclientidandusername = data.StoreMqttClientIdAndUsername.ValueString()
	}
	if !data.UndefAction.IsNull() {
		lbparameter.Undefaction = data.UndefAction.ValueString()
	}
	if !data.UseEncryptedPersistenceCookie.IsNull() {
		lbparameter.Useencryptedpersistencecookie = data.UseEncryptedPersistenceCookie.ValueString()
	}
	if !data.UsePortForHashLb.IsNull() {
		lbparameter.Useportforhashlb = data.UsePortForHashLb.ValueString()
	}
	if !data.UseSecuredPersistenceCookie.IsNull() {
		lbparameter.Usesecuredpersistencecookie = data.UseSecuredPersistenceCookie.ValueString()
	}
	if !data.VserverSpecificMac.IsNull() {
		lbparameter.Vserverspecificmac = data.VserverSpecificMac.ValueString()
	}
	if !data.LbHashAlgorithm.IsNull() {
		lbparameter.Lbhashalgorithm = data.LbHashAlgorithm.ValueString()
	}
	if !data.LbHashFingers.IsNull() {
		lbparameter.Lbhashfingers = utils.IntPtr(int(data.LbHashFingers.ValueInt64()))
	}

	return lbparameter
}

func lbparameterSetAttrFromGet(ctx context.Context, data *LbParameterResourceModel, getResponseData map[string]interface{}) *LbParameterResourceModel {
	tflog.Debug(ctx, "In lbparameterSetAttrFromGet Function")

	// Set ID for the resource
	data.Id = types.StringValue("lbparameter-config")

	// Convert API response to model
	if val, ok := getResponseData["allowboundsvcremoval"]; ok && val != nil {
		data.AllowBoundSvcRemoval = types.StringValue(val.(string))
	} else {
		data.AllowBoundSvcRemoval = types.StringNull()
	}
	if val, ok := getResponseData["computedadccookieattribute"]; ok && val != nil {
		data.ComputedAdcCookieAttribute = types.StringValue(val.(string))
	} else {
		data.ComputedAdcCookieAttribute = types.StringNull()
	}
	if val, ok := getResponseData["consolidatedlconn"]; ok && val != nil {
		data.ConsolidatedLconn = types.StringValue(val.(string))
	} else {
		data.ConsolidatedLconn = types.StringNull()
	}
	if val, ok := getResponseData["cookiepassphrase"]; ok && val != nil {
		data.CookiePassphrase = types.StringValue(val.(string))
	} else {
		data.CookiePassphrase = types.StringNull()
	}
	if val, ok := getResponseData["dbsttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.DbsTtl = types.Int64Value(intVal)
		}
	} else {
		data.DbsTtl = types.Int64Null()
	}
	if val, ok := getResponseData["dropmqttjumbomessage"]; ok && val != nil {
		data.DropMqttJumboMessage = types.StringValue(val.(string))
	} else {
		data.DropMqttJumboMessage = types.StringNull()
	}
	if val, ok := getResponseData["httponlycookieflag"]; ok && val != nil {
		data.HttpOnlyCookieFlag = types.StringValue(val.(string))
	} else {
		data.HttpOnlyCookieFlag = types.StringNull()
	}
	if val, ok := getResponseData["literaladccookieattribute"]; ok && val != nil {
		data.LiteralAdcCookieAttribute = types.StringValue(val.(string))
	} else {
		data.LiteralAdcCookieAttribute = types.StringNull()
	}
	if val, ok := getResponseData["maxpipelinenat"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.MaxPipelineNat = types.Int64Value(intVal)
		}
	} else {
		data.MaxPipelineNat = types.Int64Null()
	}
	if val, ok := getResponseData["monitorconnectionclose"]; ok && val != nil {
		data.MonitorConnectionClose = types.StringValue(val.(string))
	} else {
		data.MonitorConnectionClose = types.StringNull()
	}
	if val, ok := getResponseData["monitorskipmaxclient"]; ok && val != nil {
		data.MonitorSkipMaxClient = types.StringValue(val.(string))
	} else {
		data.MonitorSkipMaxClient = types.StringNull()
	}
	if val, ok := getResponseData["preferdirectroute"]; ok && val != nil {
		data.PreferDirectRoute = types.StringValue(val.(string))
	} else {
		data.PreferDirectRoute = types.StringNull()
	}
	if val, ok := getResponseData["proximityfromself"]; ok && val != nil {
		data.ProximityFromSelf = types.StringValue(val.(string))
	} else {
		data.ProximityFromSelf = types.StringNull()
	}
	if val, ok := getResponseData["retainservicestate"]; ok && val != nil {
		data.RetainServiceState = types.StringValue(val.(string))
	} else {
		data.RetainServiceState = types.StringNull()
	}
	if val, ok := getResponseData["startuprrfactor"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.StartupRrFactor = types.Int64Value(intVal)
		}
	} else {
		data.StartupRrFactor = types.Int64Null()
	}
	if val, ok := getResponseData["storemqttclientidandusername"]; ok && val != nil {
		data.StoreMqttClientIdAndUsername = types.StringValue(val.(string))
	} else {
		data.StoreMqttClientIdAndUsername = types.StringNull()
	}
	if val, ok := getResponseData["sessionsthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.SessionsThreshold = types.Int64Value(intVal)
		}
	} else {
		data.SessionsThreshold = types.Int64Null()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.UndefAction = types.StringValue(val.(string))
	} else {
		data.UndefAction = types.StringNull()
	}
	if val, ok := getResponseData["useencryptedpersistencecookie"]; ok && val != nil {
		data.UseEncryptedPersistenceCookie = types.StringValue(val.(string))
	} else {
		data.UseEncryptedPersistenceCookie = types.StringNull()
	}
	if val, ok := getResponseData["useportforhashlb"]; ok && val != nil {
		data.UsePortForHashLb = types.StringValue(val.(string))
	} else {
		data.UsePortForHashLb = types.StringNull()
	}
	if val, ok := getResponseData["usesecuredpersistencecookie"]; ok && val != nil {
		data.UseSecuredPersistenceCookie = types.StringValue(val.(string))
	} else {
		data.UseSecuredPersistenceCookie = types.StringNull()
	}
	if val, ok := getResponseData["vserverspecificmac"]; ok && val != nil {
		data.VserverSpecificMac = types.StringValue(val.(string))
	} else {
		data.VserverSpecificMac = types.StringNull()
	}
	if val, ok := getResponseData["lbhashalgorithm"]; ok && val != nil {
		data.LbHashAlgorithm = types.StringValue(val.(string))
	} else {
		data.LbHashAlgorithm = types.StringNull()
	}
	if val, ok := getResponseData["lbhashfingers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.LbHashFingers = types.Int64Value(intVal)
		}
	} else {
		data.LbHashFingers = types.Int64Null()
	}

	return data
}
