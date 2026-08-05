package nd6ravariables

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Nd6ravariablesResource{}
var _ resource.ResourceWithConfigure = (*Nd6ravariablesResource)(nil)
var _ resource.ResourceWithImportState = (*Nd6ravariablesResource)(nil)

func NewNd6ravariablesResource() resource.Resource {
	return &Nd6ravariablesResource{}
}

// Nd6ravariablesResource defines the resource implementation.
type Nd6ravariablesResource struct {
	client *service.NitroClient
}

func (r *Nd6ravariablesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nd6ravariablesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nd6ravariables"
}

func (r *Nd6ravariablesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nd6ravariablesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nd6ravariablesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nd6ravariables resource")

	// Create API request body from the plan
	nd6ravariables := nd6ravariablesGetThePayloadFromthePlan(ctx, &data)

	// nd6ravariables is a per-vlan configuration resource pushed with the vlan
	// carried in the request body - use UpdateUnnamedResource (matches SDK v2).
	err := r.client.UpdateUnnamedResource(service.Nd6ravariables.Type(), &nd6ravariables)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nd6ravariables, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nd6ravariables resource")

	// Set ID for the resource before reading state - the vlan number, matching
	// SDK v2 d.SetId(strconv.Itoa(vlan)).
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Vlan.ValueInt64()))

	// Read the updated state back
	if !r.readNd6ravariablesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nd6ravariables not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nd6ravariablesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nd6ravariables resource")

	found := r.readNd6ravariablesFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Nd6ravariablesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Nd6ravariablesResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nd6ravariables resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Ceaserouteradv.Equal(state.Ceaserouteradv) {
		hasChange = true
	}
	if !data.Currhoplimit.Equal(state.Currhoplimit) {
		hasChange = true
	}
	if !data.Defaultlifetime.Equal(state.Defaultlifetime) {
		hasChange = true
	}
	if !data.Linkmtu.Equal(state.Linkmtu) {
		hasChange = true
	}
	if !data.Managedaddrconfig.Equal(state.Managedaddrconfig) {
		hasChange = true
	}
	if !data.Maxrtadvinterval.Equal(state.Maxrtadvinterval) {
		hasChange = true
	}
	if !data.Minrtadvinterval.Equal(state.Minrtadvinterval) {
		hasChange = true
	}
	if !data.Onlyunicastrtadvresponse.Equal(state.Onlyunicastrtadvresponse) {
		hasChange = true
	}
	if !data.Otheraddrconfig.Equal(state.Otheraddrconfig) {
		hasChange = true
	}
	if !data.Reachabletime.Equal(state.Reachabletime) {
		hasChange = true
	}
	if !data.Retranstime.Equal(state.Retranstime) {
		hasChange = true
	}
	if !data.Sendrouteradv.Equal(state.Sendrouteradv) {
		hasChange = true
	}
	if !data.Srclinklayeraddroption.Equal(state.Srclinklayeraddroption) {
		hasChange = true
	}
	if !data.Vlan.Equal(state.Vlan) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the plan
		nd6ravariables := nd6ravariablesGetThePayloadFromthePlan(ctx, &data)

		err := r.client.UpdateUnnamedResource(service.Nd6ravariables.Type(), &nd6ravariables)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nd6ravariables, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nd6ravariables resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nd6ravariables resource, skipping update")
	}

	// Read the updated state back
	if !r.readNd6ravariablesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nd6ravariables not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nd6ravariablesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nd6ravariables resource")

	// nd6ravariables does not support a NITRO DELETE operation (matches SDK v2).
	// We simply remove it from Terraform state.
	tflog.Trace(ctx, "Deleted nd6ravariables resource from state")
}

// Helper function to read nd6ravariables data from API.
// Returns false when the resource is not found on the ADC.
func (r *Nd6ravariablesResource) readNd6ravariablesFromApi(ctx context.Context, data *Nd6ravariablesResourceModel, diags *diag.Diagnostics) bool {
	// nd6ravariables is keyed by vlan - read by the vlan number (matches SDK v2
	// FindResource(type, d.Id()) where d.Id() is the vlan string).
	nd6ravariablesName := fmt.Sprintf("%d", data.Vlan.ValueInt64())

	getResponseData, err := r.client.FindResource(service.Nd6ravariables.Type(), nd6ravariablesName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nd6ravariables, got error: %s", err))
		return false
	}

	nd6ravariablesSetAttrFromGet(ctx, data, getResponseData)

	return true
}
