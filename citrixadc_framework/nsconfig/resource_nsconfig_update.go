package nsconfig

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsconfigUpdateResource{}
var _ resource.ResourceWithConfigure = (*NsconfigUpdateResource)(nil)
var _ resource.ResourceWithImportState = (*NsconfigUpdateResource)(nil)

func NewNsconfigUpdateResource() resource.Resource {
	return &NsconfigUpdateResource{}
}

// NsconfigUpdateResource defines the resource implementation.
type NsconfigUpdateResource struct {
	client *service.NitroClient
}

// NsconfigUpdateResourceModel describes the resource data model.
// Mirrors the SDK v2 `citrixadc_nsconfig_update` resource: it applies a subset of
// settable nsconfig params via the NITRO `set ns config` (PUT) call and reads them
// back. The ID is a synthetic constant since nsconfig is an unnamed singleton.
type NsconfigUpdateResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Ipaddress               types.String `tfsdk:"ipaddress"`
	Netmask                 types.String `tfsdk:"netmask"`
	Nsvlan                  types.Int64  `tfsdk:"nsvlan"`
	Ifnum                   types.Set    `tfsdk:"ifnum"`
	Tagged                  types.String `tfsdk:"tagged"`
	Httpport                types.Set    `tfsdk:"httpport"`
	Maxconn                 types.Int64  `tfsdk:"maxconn"`
	Maxreq                  types.Int64  `tfsdk:"maxreq"`
	Cip                     types.String `tfsdk:"cip"`
	Cipheader               types.String `tfsdk:"cipheader"`
	Cookieversion           types.String `tfsdk:"cookieversion"`
	Securecookie            types.String `tfsdk:"securecookie"`
	Pmtumin                 types.Int64  `tfsdk:"pmtumin"`
	Pmtutimeout             types.Int64  `tfsdk:"pmtutimeout"`
	Ftpportrange            types.String `tfsdk:"ftpportrange"`
	Crportrange             types.String `tfsdk:"crportrange"`
	Timezone                types.String `tfsdk:"timezone"`
	Grantquotamaxclient     types.Int64  `tfsdk:"grantquotamaxclient"`
	Exclusivequotamaxclient types.Int64  `tfsdk:"exclusivequotamaxclient"`
	Grantquotaspillover     types.Int64  `tfsdk:"grantquotaspillover"`
	Exclusivequotaspillover types.Int64  `tfsdk:"exclusivequotaspillover"`
	Securemanagementtraffic types.String `tfsdk:"securemanagementtraffic"`
	Securemanagementtd      types.Int64  `tfsdk:"securemanagementtd"`
}

func (r *NsconfigUpdateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsconfig_update"
}

func (r *NsconfigUpdateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The synthetic ID of the nsconfig_update resource.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the Citrix ADC (NSIP address).",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Netmask corresponding to the IP address.",
			},
			"nsvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN (NSVLAN) for the subnet on which the IP address resides.",
			},
			"ifnum": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Interfaces of the appliance that must be bound to the NSVLAN.",
			},
			"tagged": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies that the interfaces will be added as 802.1q tagged interfaces.",
			},
			"httpport": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The HTTP ports on the Web server. Allows the system to perform connection off-load for any client request whose destination port matches one of these configured ports.",
			},
			"maxconn": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					utils.UnsetOnRemoveOrKeepDefaultInt64{DefaultValue: 0},
				},
				Description: "The maximum number of connections that will be made from the system to the web server(s) attached to it. Applied globally to all attached servers.",
			},
			"maxreq": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					utils.UnsetOnRemoveOrKeepDefaultInt64{DefaultValue: 0},
				},
				Description: "The maximum number of requests that the system can pass on a particular connection between the system and a server. Setting this value to 0 allows an unlimited number of requests.",
			},
			"cip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "DISABLED"},
				},
				Description: "Control (enable or disable) the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers. Possible values: [ ENABLED, DISABLED ]",
			},
			"cipheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: ""},
				},
				Description: "The text that will be used as the client IP header.",
			},
			"cookieversion": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "0"},
				},
				Description: "The version of the cookie inserted by the system. Possible values: [ 0, 1 ]",
			},
			"securecookie": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "ENABLED"},
				},
				Description: "Enable or disable the secure flag for the persistence cookie. Possible values: [ ENABLED, DISABLED ]",
			},
			"pmtumin": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					utils.UnsetOnRemoveOrKeepDefaultInt64{DefaultValue: 576},
				},
				Description: "The minimum Path MTU.",
			},
			"pmtutimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					utils.UnsetOnRemoveOrKeepDefaultInt64{DefaultValue: 10},
				},
				Description: "The Path MTU timeout value in minutes.",
			},
			"ftpportrange": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: ""},
				},
				Description: "Port range configured for FTP services.",
			},
			"crportrange": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: ""},
				},
				Description: "Port range for cache redirection services.",
			},
			"timezone": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "CoordinatedUniversalTime"},
				},
				Description: "Name of the timezone.",
			},
			"grantquotamaxclient": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					utils.UnsetOnRemoveOrKeepDefaultInt64{DefaultValue: 10},
				},
				Description: "The percentage of shared quota to be granted at a time for maxClient.",
			},
			"exclusivequotamaxclient": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					utils.UnsetOnRemoveOrKeepDefaultInt64{DefaultValue: 80},
				},
				Description: "The percentage of maxClient to be given to PEs.",
			},
			"grantquotaspillover": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					utils.UnsetOnRemoveOrKeepDefaultInt64{DefaultValue: 10},
				},
				Description: "The percentage of shared quota to be granted at a time for spillover.",
			},
			"exclusivequotaspillover": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					utils.UnsetOnRemoveOrKeepDefaultInt64{DefaultValue: 80},
				},
				Description: "The percentage of spillover threshold to be given to PEs.",
			},
			"securemanagementtraffic": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "DISABLED"},
				},
				Description: "Enable secure management traffic handling. Possible values: [ ENABLED, DISABLED ]",
			},
			"securemanagementtd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer that identifies the Management traffic domain. If not specified, defaults to 4094.",
			},
		},
	}
}

func (r *NsconfigUpdateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsconfigUpdateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// getPayload builds the ns.Nsconfig PUT payload from the model.
func (r *NsconfigUpdateResource) getPayload(ctx context.Context, data *NsconfigUpdateResourceModel) ns.Nsconfig {
	nsconfig := ns.Nsconfig{}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		nsconfig.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		nsconfig.Netmask = data.Netmask.ValueString()
	}
	if !data.Nsvlan.IsNull() && !data.Nsvlan.IsUnknown() {
		nsconfig.Nsvlan = utils.IntPtr(int(data.Nsvlan.ValueInt64()))
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		var ifnumList []string
		data.Ifnum.ElementsAs(ctx, &ifnumList, false)
		nsconfig.Ifnum = ifnumList
	}
	if !data.Tagged.IsNull() && !data.Tagged.IsUnknown() {
		nsconfig.Tagged = data.Tagged.ValueString()
	}
	if !data.Httpport.IsNull() && !data.Httpport.IsUnknown() {
		var httpportList []int64
		data.Httpport.ElementsAs(ctx, &httpportList, false)
		httpportInts := make([]int, 0, len(httpportList))
		for _, v := range httpportList {
			httpportInts = append(httpportInts, int(v))
		}
		nsconfig.Httpport = httpportInts
	}
	if !data.Maxconn.IsNull() && !data.Maxconn.IsUnknown() {
		nsconfig.Maxconn = utils.IntPtr(int(data.Maxconn.ValueInt64()))
	}
	if !data.Maxreq.IsNull() && !data.Maxreq.IsUnknown() {
		nsconfig.Maxreq = utils.IntPtr(int(data.Maxreq.ValueInt64()))
	}
	if !data.Cip.IsNull() && !data.Cip.IsUnknown() {
		nsconfig.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		nsconfig.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Cookieversion.IsNull() && !data.Cookieversion.IsUnknown() {
		nsconfig.Cookieversion = data.Cookieversion.ValueString()
	}
	if !data.Securecookie.IsNull() && !data.Securecookie.IsUnknown() {
		nsconfig.Securecookie = data.Securecookie.ValueString()
	}
	if !data.Pmtumin.IsNull() && !data.Pmtumin.IsUnknown() {
		nsconfig.Pmtumin = utils.IntPtr(int(data.Pmtumin.ValueInt64()))
	}
	if !data.Pmtutimeout.IsNull() && !data.Pmtutimeout.IsUnknown() {
		nsconfig.Pmtutimeout = utils.IntPtr(int(data.Pmtutimeout.ValueInt64()))
	}
	if !data.Ftpportrange.IsNull() && !data.Ftpportrange.IsUnknown() {
		nsconfig.Ftpportrange = data.Ftpportrange.ValueString()
	}
	if !data.Crportrange.IsNull() && !data.Crportrange.IsUnknown() {
		nsconfig.Crportrange = data.Crportrange.ValueString()
	}
	if !data.Timezone.IsNull() && !data.Timezone.IsUnknown() {
		nsconfig.Timezone = data.Timezone.ValueString()
	}
	if !data.Grantquotamaxclient.IsNull() && !data.Grantquotamaxclient.IsUnknown() {
		nsconfig.Grantquotamaxclient = utils.IntPtr(int(data.Grantquotamaxclient.ValueInt64()))
	}
	if !data.Exclusivequotamaxclient.IsNull() && !data.Exclusivequotamaxclient.IsUnknown() {
		nsconfig.Exclusivequotamaxclient = utils.IntPtr(int(data.Exclusivequotamaxclient.ValueInt64()))
	}
	if !data.Grantquotaspillover.IsNull() && !data.Grantquotaspillover.IsUnknown() {
		nsconfig.Grantquotaspillover = utils.IntPtr(int(data.Grantquotaspillover.ValueInt64()))
	}
	if !data.Exclusivequotaspillover.IsNull() && !data.Exclusivequotaspillover.IsUnknown() {
		nsconfig.Exclusivequotaspillover = utils.IntPtr(int(data.Exclusivequotaspillover.ValueInt64()))
	}
	if !data.Securemanagementtraffic.IsNull() && !data.Securemanagementtraffic.IsUnknown() {
		nsconfig.Securemanagementtraffic = data.Securemanagementtraffic.ValueString()
	}
	if !data.Securemanagementtd.IsNull() && !data.Securemanagementtd.IsUnknown() {
		nsconfig.Securemanagementtd = utils.IntPtr(int(data.Securemanagementtd.ValueInt64()))
	}
	return nsconfig
}

func (r *NsconfigUpdateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsconfigUpdateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsconfig_update resource (set ns config)")
	nsconfig := r.getPayload(ctx, &data)
	if err := r.client.UpdateUnnamedResource(service.Nsconfig.Type(), &nsconfig); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ns config, got error: %s", err))
		return
	}

	data.Id = types.StringValue("nsconfig-update-config")

	r.readFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigUpdateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsconfigUpdateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Reading nsconfig_update resource")
	r.readFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigUpdateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state, config NsconfigUpdateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsconfig_update resource")

	hasChange := false
	// Behavioral parameters removed from config are reset to their appliance
	// default via a single batched ?action=unset (see the UnsetOnRemoveOrKeepDefault
	// plan modifiers on these attributes). The addressing/connectivity parameters
	// (ipaddress/netmask/nsvlan/ifnum/tagged), httpport, and securemanagementtd are
	// set-only: they are never auto-unset.
	attributesToUnset := []string{}

	// --- set-only parameters ---
	if !data.Ipaddress.Equal(state.Ipaddress) {
		hasChange = true
	}
	if !data.Netmask.Equal(state.Netmask) {
		hasChange = true
	}
	if !data.Nsvlan.Equal(state.Nsvlan) {
		hasChange = true
	}
	if !data.Ifnum.Equal(state.Ifnum) {
		hasChange = true
	}
	if !data.Tagged.Equal(state.Tagged) {
		hasChange = true
	}
	if !data.Httpport.Equal(state.Httpport) {
		hasChange = true
	}
	if !data.Securemanagementtd.Equal(state.Securemanagementtd) {
		hasChange = true
	}

	// --- behavioral parameters (unset on removal) ---
	if !data.Maxconn.Equal(state.Maxconn) {
		if config.Maxconn.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxconn")
		} else {
			hasChange = true
		}
	}
	if !data.Maxreq.Equal(state.Maxreq) {
		if config.Maxreq.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxreq")
		} else {
			hasChange = true
		}
	}
	if !data.Cip.Equal(state.Cip) {
		if config.Cip.IsNull() {
			attributesToUnset = append(attributesToUnset, "cip")
		} else {
			hasChange = true
		}
	}
	if !data.Cipheader.Equal(state.Cipheader) {
		if config.Cipheader.IsNull() {
			attributesToUnset = append(attributesToUnset, "cipheader")
		} else {
			hasChange = true
		}
	}
	if !data.Cookieversion.Equal(state.Cookieversion) {
		if config.Cookieversion.IsNull() {
			attributesToUnset = append(attributesToUnset, "cookieversion")
		} else {
			hasChange = true
		}
	}
	if !data.Securecookie.Equal(state.Securecookie) {
		if config.Securecookie.IsNull() {
			attributesToUnset = append(attributesToUnset, "securecookie")
		} else {
			hasChange = true
		}
	}
	if !data.Pmtumin.Equal(state.Pmtumin) {
		if config.Pmtumin.IsNull() {
			attributesToUnset = append(attributesToUnset, "pmtumin")
		} else {
			hasChange = true
		}
	}
	if !data.Pmtutimeout.Equal(state.Pmtutimeout) {
		if config.Pmtutimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "pmtutimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Ftpportrange.Equal(state.Ftpportrange) {
		if config.Ftpportrange.IsNull() {
			attributesToUnset = append(attributesToUnset, "ftpportrange")
		} else {
			hasChange = true
		}
	}
	if !data.Crportrange.Equal(state.Crportrange) {
		if config.Crportrange.IsNull() {
			attributesToUnset = append(attributesToUnset, "crportrange")
		} else {
			hasChange = true
		}
	}
	if !data.Timezone.Equal(state.Timezone) {
		if config.Timezone.IsNull() {
			attributesToUnset = append(attributesToUnset, "timezone")
		} else {
			hasChange = true
		}
	}
	if !data.Grantquotamaxclient.Equal(state.Grantquotamaxclient) {
		if config.Grantquotamaxclient.IsNull() {
			attributesToUnset = append(attributesToUnset, "grantquotamaxclient")
		} else {
			hasChange = true
		}
	}
	if !data.Exclusivequotamaxclient.Equal(state.Exclusivequotamaxclient) {
		if config.Exclusivequotamaxclient.IsNull() {
			attributesToUnset = append(attributesToUnset, "exclusivequotamaxclient")
		} else {
			hasChange = true
		}
	}
	if !data.Grantquotaspillover.Equal(state.Grantquotaspillover) {
		if config.Grantquotaspillover.IsNull() {
			attributesToUnset = append(attributesToUnset, "grantquotaspillover")
		} else {
			hasChange = true
		}
	}
	if !data.Exclusivequotaspillover.Equal(state.Exclusivequotaspillover) {
		if config.Exclusivequotaspillover.IsNull() {
			attributesToUnset = append(attributesToUnset, "exclusivequotaspillover")
		} else {
			hasChange = true
		}
	}
	if !data.Securemanagementtraffic.Equal(state.Securemanagementtraffic) {
		if config.Securemanagementtraffic.IsNull() {
			attributesToUnset = append(attributesToUnset, "securemanagementtraffic")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		nsconfig := r.getPayload(ctx, &data)
		if err := r.client.UpdateUnnamedResource(service.Nsconfig.Type(), &nsconfig); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ns config, got error: %s", err))
			return
		}
	} else {
		tflog.Debug(ctx, "No changes detected for nsconfig_update resource, skipping update")
	}

	// Reset behavioral parameters removed from config to their appliance defaults.
	if err := utils.ExecuteUnset(r.client, service.Nsconfig.Type(), map[string]interface{}{}, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset ns config attributes, got error: %s", err))
		return
	}

	r.readFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigUpdateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Mirrors SDK v2 schema.Noop delete: nsconfig has no delete API. Just drop state.
	tflog.Debug(ctx, "Deleting nsconfig_update: no delete API, removing from state only")
}

// readFromApi reads the live nsconfig and populates the model. On read failure it
// clears the ID (mirrors SDK v2 readNsconfigUpdateFunc behavior).
func (r *NsconfigUpdateResource) readFromApi(ctx context.Context, data *NsconfigUpdateResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsconfig.Type(), "")
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Clearing nsconfig_update state, got error: %s", err))
		data.Id = types.StringNull()
		return
	}

	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["nsvlan"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Nsvlan = types.Int64Value(intVal)
		}
	} else {
		data.Nsvlan = types.Int64Null()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			setValue, _ := types.SetValueFrom(ctx, types.StringType, stringList)
			data.Ifnum = setValue
		} else {
			data.Ifnum = types.SetNull(types.StringType)
		}
	} else {
		data.Ifnum = types.SetNull(types.StringType)
	}
	if val, ok := getResponseData["tagged"]; ok && val != nil {
		data.Tagged = types.StringValue(val.(string))
	} else {
		data.Tagged = types.StringNull()
	}
	if val, ok := getResponseData["httpport"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			var httpportList []int64
			for _, e := range sliceVal {
				if iv, cerr := utils.ConvertToInt64(e); cerr == nil {
					httpportList = append(httpportList, iv)
				}
			}
			setValue, _ := types.SetValueFrom(ctx, types.Int64Type, httpportList)
			data.Httpport = setValue
		} else {
			data.Httpport = types.SetNull(types.Int64Type)
		}
	} else {
		data.Httpport = types.SetNull(types.Int64Type)
	}
	if val, ok := getResponseData["maxconn"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Maxconn = types.Int64Value(intVal)
		}
	} else {
		data.Maxconn = types.Int64Null()
	}
	if val, ok := getResponseData["maxreq"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Maxreq = types.Int64Value(intVal)
		}
	} else {
		data.Maxreq = types.Int64Null()
	}
	if val, ok := getResponseData["cip"]; ok && val != nil {
		data.Cip = types.StringValue(val.(string))
	} else {
		data.Cip = types.StringNull()
	}
	if val, ok := getResponseData["cipheader"]; ok && val != nil {
		data.Cipheader = types.StringValue(val.(string))
	} else {
		data.Cipheader = types.StringNull()
	}
	if val, ok := getResponseData["cookieversion"]; ok && val != nil {
		data.Cookieversion = types.StringValue(val.(string))
	} else {
		data.Cookieversion = types.StringNull()
	}
	if val, ok := getResponseData["securecookie"]; ok && val != nil {
		data.Securecookie = types.StringValue(val.(string))
	} else {
		data.Securecookie = types.StringNull()
	}
	if val, ok := getResponseData["pmtumin"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Pmtumin = types.Int64Value(intVal)
		}
	} else {
		data.Pmtumin = types.Int64Null()
	}
	if val, ok := getResponseData["pmtutimeout"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Pmtutimeout = types.Int64Value(intVal)
		}
	} else {
		data.Pmtutimeout = types.Int64Null()
	}
	if val, ok := getResponseData["ftpportrange"]; ok && val != nil {
		data.Ftpportrange = types.StringValue(val.(string))
	} else {
		data.Ftpportrange = types.StringNull()
	}
	if val, ok := getResponseData["crportrange"]; ok && val != nil {
		data.Crportrange = types.StringValue(val.(string))
	} else {
		data.Crportrange = types.StringNull()
	}
	if val, ok := getResponseData["timezone"]; ok && val != nil {
		data.Timezone = types.StringValue(val.(string))
	} else {
		data.Timezone = types.StringNull()
	}
	if val, ok := getResponseData["grantquotamaxclient"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Grantquotamaxclient = types.Int64Value(intVal)
		}
	} else {
		data.Grantquotamaxclient = types.Int64Null()
	}
	if val, ok := getResponseData["exclusivequotamaxclient"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Exclusivequotamaxclient = types.Int64Value(intVal)
		}
	} else {
		data.Exclusivequotamaxclient = types.Int64Null()
	}
	if val, ok := getResponseData["grantquotaspillover"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Grantquotaspillover = types.Int64Value(intVal)
		}
	} else {
		data.Grantquotaspillover = types.Int64Null()
	}
	if val, ok := getResponseData["exclusivequotaspillover"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Exclusivequotaspillover = types.Int64Value(intVal)
		}
	} else {
		data.Exclusivequotaspillover = types.Int64Null()
	}
	if val, ok := getResponseData["securemanagementtraffic"]; ok && val != nil {
		data.Securemanagementtraffic = types.StringValue(val.(string))
	} else {
		data.Securemanagementtraffic = types.StringNull()
	}
	if val, ok := getResponseData["securemanagementtd"]; ok && val != nil {
		if intVal, cerr := utils.ConvertToInt64(val); cerr == nil {
			data.Securemanagementtd = types.Int64Value(intVal)
		}
	} else {
		data.Securemanagementtd = types.Int64Null()
	}
}
