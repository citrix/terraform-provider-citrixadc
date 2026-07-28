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
	Id        types.String `tfsdk:"id"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Netmask   types.String `tfsdk:"netmask"`
	Nsvlan    types.Int64  `tfsdk:"nsvlan"`
	Ifnum     types.List   `tfsdk:"ifnum"`
	Tagged    types.String `tfsdk:"tagged"`
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
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Interfaces of the appliance that must be bound to the NSVLAN.",
			},
			"tagged": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies that the interfaces will be added as 802.1q tagged interfaces.",
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
	var data, state NsconfigUpdateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsconfig_update resource")

	hasChange := false
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

	if hasChange {
		nsconfig := r.getPayload(ctx, &data)
		if err := r.client.UpdateUnnamedResource(service.Nsconfig.Type(), &nsconfig); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ns config, got error: %s", err))
			return
		}
	} else {
		tflog.Debug(ctx, "No changes detected for nsconfig_update resource, skipping update")
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
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Ifnum = listValue
		} else {
			data.Ifnum = types.ListNull(types.StringType)
		}
	} else {
		data.Ifnum = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["tagged"]; ok && val != nil {
		data.Tagged = types.StringValue(val.(string))
	} else {
		data.Tagged = types.StringNull()
	}
}
