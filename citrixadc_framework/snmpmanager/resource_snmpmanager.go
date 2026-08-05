package snmpmanager

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SnmpmanagerResource{}
var _ resource.ResourceWithConfigure = (*SnmpmanagerResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpmanagerResource)(nil)

func NewSnmpmanagerResource() resource.Resource {
	return &SnmpmanagerResource{}
}

// SnmpmanagerResource defines the resource implementation.
type SnmpmanagerResource struct {
	client *service.NitroClient
}

func (r *SnmpmanagerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpmanagerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpmanager"
}

func (r *SnmpmanagerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpmanagerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpmanagerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpmanager resource")

	snmpmanager := snmpmanagerGetThePayloadFromthePlan(ctx, &data)

	// Named resource keyed by ipaddress - use AddResource (matches SDK v2).
	ipaddress := data.Ipaddress.ValueString()
	_, err := r.client.AddResource(service.Snmpmanager.Type(), ipaddress, &snmpmanager)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpmanager, got error: %s", err))
		return
	}

	// ID matches SDK v2: d.SetId(ipaddress)
	data.Id = types.StringValue(ipaddress)

	tflog.Trace(ctx, "Created snmpmanager resource")

	// Read the updated state back
	if !r.readSnmpmanagerFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpmanager not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpmanagerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpmanagerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpmanager resource")

	found := r.readSnmpmanagerFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpmanagerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SnmpmanagerResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating snmpmanager resource")

	// snmpmanager is keyed by ipaddress; the update payload must carry it (matches SDK v2).
	snmpmanager := snmp.Snmpmanager{
		Ipaddress: data.Ipaddress.ValueString(),
	}
	hasChange := false
	if !data.Domainresolveretry.Equal(state.Domainresolveretry) {
		tflog.Debug(ctx, "domainresolveretry has changed for snmpmanager, starting update")
		if !data.Domainresolveretry.IsNull() && !data.Domainresolveretry.IsUnknown() {
			snmpmanager.Domainresolveretry = utils.IntPtr(int(data.Domainresolveretry.ValueInt64()))
		}
		hasChange = true
	}
	if !data.Netmask.Equal(state.Netmask) {
		tflog.Debug(ctx, "netmask has changed for snmpmanager, starting update")
		if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
			snmpmanager.Netmask = data.Netmask.ValueString()
		}
		hasChange = true
	}

	if hasChange {
		// Matches SDK v2: UpdateUnnamedResource (PUT body carries the ipaddress key).
		err := r.client.UpdateUnnamedResource(service.Snmpmanager.Type(), &snmpmanager)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpmanager %s, got error: %s", data.Ipaddress.ValueString(), err))
			return
		}
		tflog.Trace(ctx, "Updated snmpmanager resource")
	} else {
		tflog.Debug(ctx, "No changes detected for snmpmanager resource, skipping update")
	}

	// Read the updated state back
	if !r.readSnmpmanagerFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpmanager not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpmanagerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpmanagerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpmanager resource")

	// snmpmanager delete requires the netmask disambiguator (matches SDK v2).
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(data.Netmask.ValueString())))

	err := r.client.DeleteResourceWithArgs(service.Snmpmanager.Type(), data.Ipaddress.ValueString(), args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete snmpmanager, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted snmpmanager resource")
}

// readSnmpmanagerFromApi reads the snmpmanager state from the ADC. snmpmanager is
// an array resource keyed by ipaddress, so it is fetched via FindAllResources and
// matched on the ipaddress (matches SDK v2). Returns false if the manager is gone.
func (r *SnmpmanagerResource) readSnmpmanagerFromApi(ctx context.Context, data *SnmpmanagerResourceModel, diags *diag.Diagnostics) bool {
	snmpmanagerName := data.Id.ValueString()

	dataArr, err := r.client.FindAllResources(service.Snmpmanager.Type())
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpmanager, got error: %s", err))
		return false
	}

	if len(dataArr) == 0 {
		return false
	}

	foundIndex := -1
	for i, v := range dataArr {
		if ipaddress, ok := v["ipaddress"].(string); ok && ipaddress == snmpmanagerName {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return false
	}

	snmpmanagerSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
