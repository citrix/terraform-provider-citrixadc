package citrixadc_framework

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LbParameterResource{}

func NewLbParameterResource() resource.Resource {
	return &LbParameterResource{}
}

// LbParameterResource defines the resource implementation.
type LbParameterResource struct {
	providerData *ProviderData
}

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
	RetainServiceState            types.String `tfsdk:"retainservicestate"`
	StartupRrFactor               types.Int64  `tfsdk:"startuprrfactor"`
	StoreMqttClientIdAndUsername  types.String `tfsdk:"storemqttclientidandusername"`
	SessionsThreshold             types.Int64  `tfsdk:"sessionsthreshold"`
	UseEncryptedPersistenceCookie types.String `tfsdk:"useencryptedpersistencecookie"`
	UsePortForHashLb              types.String `tfsdk:"useportforhashlb"`
	UseSecuredPersistenceCookie   types.String `tfsdk:"usesecuredpersistencecookie"`
	VserverSpecificMac            types.String `tfsdk:"vserverspecificmac"`
	LbHashAlgorithm               types.String `tfsdk:"lbhashalgorithm"`
	LbHashFingers                 types.Int64  `tfsdk:"lbhashfingers"`
}

type Lbparameter struct {
	Adccookieattributewarningmsg  string      `json:"adccookieattributewarningmsg,omitempty"`
	Allowboundsvcremoval          string      `json:"allowboundsvcremoval,omitempty"`
	Builtin                       interface{} `json:"builtin,omitempty"`
	Computedadccookieattribute    string      `json:"computedadccookieattribute,omitempty"`
	Consolidatedlconn             string      `json:"consolidatedlconn,omitempty"`
	Cookiepassphrase              string      `json:"cookiepassphrase,omitempty"`
	Dbsttl                        *int        `json:"dbsttl,omitempty"`
	Dropmqttjumbomessage          string      `json:"dropmqttjumbomessage,omitempty"`
	Feature                       string      `json:"feature,omitempty"`
	Httponlycookieflag            string      `json:"httponlycookieflag,omitempty"`
	Literaladccookieattribute     string      `json:"literaladccookieattribute,omitempty"`
	Maxpipelinenat                *int        `json:"maxpipelinenat,omitempty"`
	Monitorconnectionclose        string      `json:"monitorconnectionclose,omitempty"`
	Monitorskipmaxclient          string      `json:"monitorskipmaxclient,omitempty"`
	Preferdirectroute             string      `json:"preferdirectroute,omitempty"`
	Retainservicestate            string      `json:"retainservicestate,omitempty"`
	Sessionsthreshold             *int        `json:"sessionsthreshold,omitempty"`
	Startuprrfactor               *int        `json:"startuprrfactor,omitempty"`
	Storemqttclientidandusername  string      `json:"storemqttclientidandusername,omitempty"`
	Useencryptedpersistencecookie string      `json:"useencryptedpersistencecookie,omitempty"`
	Useportforhashlb              string      `json:"useportforhashlb,omitempty"`
	Usesecuredpersistencecookie   string      `json:"usesecuredpersistencecookie,omitempty"`
	Vserverspecificmac            string      `json:"vserverspecificmac,omitempty"`
	Lbhashalgorithm               string      `json:"lbhashalgorithm,omitempty"`
	Lbhashfingers                 *int        `json:"lbhashfingers,omitempty"`
}

func (r *LbParameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbparameter"
}

func (r *LbParameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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

func (r *LbParameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*ProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.providerData = providerData
}

func (r *LbParameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbParameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbparameter resource")

	// Create API request body from the model
	lbparameter := Lbparameter{}

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
		lbparameter.Dbsttl = intPtr(int(data.DbsTtl.ValueInt64()))
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
		lbparameter.Maxpipelinenat = intPtr(int(data.MaxPipelineNat.ValueInt64()))
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
	if !data.RetainServiceState.IsNull() {
		lbparameter.Retainservicestate = data.RetainServiceState.ValueString()
	}
	if !data.StartupRrFactor.IsNull() {
		lbparameter.Startuprrfactor = intPtr(int(data.StartupRrFactor.ValueInt64()))
	}
	if !data.StoreMqttClientIdAndUsername.IsNull() {
		lbparameter.Storemqttclientidandusername = data.StoreMqttClientIdAndUsername.ValueString()
	}
	if !data.SessionsThreshold.IsNull() {
		lbparameter.Sessionsthreshold = intPtr(int(data.SessionsThreshold.ValueInt64()))
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
		lbparameter.Lbhashfingers = intPtr(int(data.LbHashFingers.ValueInt64()))
	}

	// Make API call
	err := r.providerData.Client.UpdateUnnamedResource(service.Lbparameter.Type(), &lbparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbparameter-config")

	tflog.Trace(ctx, "Created lbparameter resource")

	// Read the updated state back
	r.readLbParameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbParameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbParameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbparameter resource")

	r.readLbParameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbParameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbParameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbparameter resource")

	// Create API request body from the model
	lbparameter := Lbparameter{}

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
		lbparameter.Dbsttl = intPtr(int(data.DbsTtl.ValueInt64()))
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
		lbparameter.Maxpipelinenat = intPtr(int(data.MaxPipelineNat.ValueInt64()))
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
	if !data.RetainServiceState.IsNull() {
		lbparameter.Retainservicestate = data.RetainServiceState.ValueString()
	}
	if !data.StartupRrFactor.IsNull() {
		lbparameter.Startuprrfactor = intPtr(int(data.StartupRrFactor.ValueInt64()))
	}
	if !data.StoreMqttClientIdAndUsername.IsNull() {
		lbparameter.Storemqttclientidandusername = data.StoreMqttClientIdAndUsername.ValueString()
	}
	if !data.SessionsThreshold.IsNull() {
		lbparameter.Sessionsthreshold = intPtr(int(data.SessionsThreshold.ValueInt64()))
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
		lbparameter.Lbhashfingers = intPtr(int(data.LbHashFingers.ValueInt64()))
	}

	// Make API call
	err := r.providerData.Client.UpdateUnnamedResource(service.Lbparameter.Type(), &lbparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated lbparameter resource")

	// Read the updated state back
	r.readLbParameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbParameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbParameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbparameter resource")

	// For lbparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbparameter resource from state")
}

// Helper function to read lbparameter data from API
func (r *LbParameterResource) readLbParameterFromApi(ctx context.Context, data *LbParameterResourceModel, diags *diag.Diagnostics) {
	result, err := r.providerData.Client.FindResource(service.Lbparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbparameter, got error: %s", err))
		return
	}

	// Set ID for the resource
	data.Id = types.StringValue("lbparameter-config")

	// Convert API response to model
	if val, ok := result["allowboundsvcremoval"]; ok && val != nil {
		data.AllowBoundSvcRemoval = types.StringValue(val.(string))
	} else {
		data.AllowBoundSvcRemoval = types.StringNull()
	}
	if val, ok := result["computedadccookieattribute"]; ok && val != nil {
		data.ComputedAdcCookieAttribute = types.StringValue(val.(string))
	} else {
		data.ComputedAdcCookieAttribute = types.StringNull()
	}
	if val, ok := result["consolidatedlconn"]; ok && val != nil {
		data.ConsolidatedLconn = types.StringValue(val.(string))
	} else {
		data.ConsolidatedLconn = types.StringNull()
	}
	if val, ok := result["cookiepassphrase"]; ok && val != nil {
		data.CookiePassphrase = types.StringValue(val.(string))
	} else {
		data.CookiePassphrase = types.StringNull()
	}
	if val, ok := result["dbsttl"]; ok && val != nil {
		if intVal, err := convertToInt64(val); err == nil {
			data.DbsTtl = types.Int64Value(intVal)
		}
	} else {
		data.DbsTtl = types.Int64Null()
	}
	if val, ok := result["dropmqttjumbomessage"]; ok && val != nil {
		data.DropMqttJumboMessage = types.StringValue(val.(string))
	} else {
		data.DropMqttJumboMessage = types.StringNull()
	}
	if val, ok := result["httponlycookieflag"]; ok && val != nil {
		data.HttpOnlyCookieFlag = types.StringValue(val.(string))
	} else {
		data.HttpOnlyCookieFlag = types.StringNull()
	}
	if val, ok := result["literaladccookieattribute"]; ok && val != nil {
		data.LiteralAdcCookieAttribute = types.StringValue(val.(string))
	} else {
		data.LiteralAdcCookieAttribute = types.StringNull()
	}
	if val, ok := result["maxpipelinenat"]; ok && val != nil {
		if intVal, err := convertToInt64(val); err == nil {
			data.MaxPipelineNat = types.Int64Value(intVal)
		}
	} else {
		data.MaxPipelineNat = types.Int64Null()
	}
	if val, ok := result["monitorconnectionclose"]; ok && val != nil {
		data.MonitorConnectionClose = types.StringValue(val.(string))
	} else {
		data.MonitorConnectionClose = types.StringNull()
	}
	if val, ok := result["monitorskipmaxclient"]; ok && val != nil {
		data.MonitorSkipMaxClient = types.StringValue(val.(string))
	} else {
		data.MonitorSkipMaxClient = types.StringNull()
	}
	if val, ok := result["preferdirectroute"]; ok && val != nil {
		data.PreferDirectRoute = types.StringValue(val.(string))
	} else {
		data.PreferDirectRoute = types.StringNull()
	}
	if val, ok := result["retainservicestate"]; ok && val != nil {
		data.RetainServiceState = types.StringValue(val.(string))
	} else {
		data.RetainServiceState = types.StringNull()
	}
	if val, ok := result["startuprrfactor"]; ok && val != nil {
		if intVal, err := convertToInt64(val); err == nil {
			data.StartupRrFactor = types.Int64Value(intVal)
		}
	} else {
		data.StartupRrFactor = types.Int64Null()
	}
	if val, ok := result["storemqttclientidandusername"]; ok && val != nil {
		data.StoreMqttClientIdAndUsername = types.StringValue(val.(string))
	} else {
		data.StoreMqttClientIdAndUsername = types.StringNull()
	}
	if val, ok := result["sessionsthreshold"]; ok && val != nil {
		if intVal, err := convertToInt64(val); err == nil {
			data.SessionsThreshold = types.Int64Value(intVal)
		}
	} else {
		data.SessionsThreshold = types.Int64Null()
	}
	if val, ok := result["useencryptedpersistencecookie"]; ok && val != nil {
		data.UseEncryptedPersistenceCookie = types.StringValue(val.(string))
	} else {
		data.UseEncryptedPersistenceCookie = types.StringNull()
	}
	if val, ok := result["useportforhashlb"]; ok && val != nil {
		data.UsePortForHashLb = types.StringValue(val.(string))
	} else {
		data.UsePortForHashLb = types.StringNull()
	}
	if val, ok := result["usesecuredpersistencecookie"]; ok && val != nil {
		data.UseSecuredPersistenceCookie = types.StringValue(val.(string))
	} else {
		data.UseSecuredPersistenceCookie = types.StringNull()
	}
	if val, ok := result["vserverspecificmac"]; ok && val != nil {
		data.VserverSpecificMac = types.StringValue(val.(string))
	} else {
		data.VserverSpecificMac = types.StringNull()
	}
	if val, ok := result["lbhashalgorithm"]; ok && val != nil {
		data.LbHashAlgorithm = types.StringValue(val.(string))
	} else {
		data.LbHashAlgorithm = types.StringNull()
	}
	if val, ok := result["lbhashfingers"]; ok && val != nil {
		if intVal, err := convertToInt64(val); err == nil {
			data.LbHashFingers = types.Int64Value(intVal)
		}
	} else {
		data.LbHashFingers = types.Int64Null()
	}
}

// Helper function to convert interface{} to int64
func convertToInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("cannot convert %T to int64", value)
	}
}
