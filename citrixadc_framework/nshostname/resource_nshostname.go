package nshostname

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
var _ resource.Resource = &NshostnameResource{}
var _ resource.ResourceWithConfigure = (*NshostnameResource)(nil)
var _ resource.ResourceWithImportState = (*NshostnameResource)(nil)

func NewNshostnameResource() resource.Resource {
	return &NshostnameResource{}
}

// NshostnameResource defines the resource implementation.
type NshostnameResource struct {
	client *service.NitroClient
}

func (r *NshostnameResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NshostnameResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nshostname"
}

func (r *NshostnameResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NshostnameResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NshostnameResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nshostname resource")

	nshostname := nshostnameGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nshostname.Type(), &nshostname)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nshostname, got error: %s", err))
		return
	}

	// Static singleton ID
	data.Id = types.StringValue("nshostname-config")

	tflog.Trace(ctx, "Created nshostname resource")

	// Read the updated state back
	if !r.readNshostnameFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nshostname not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshostnameResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NshostnameResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nshostname resource")

	found := r.readNshostnameFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NshostnameResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NshostnameResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nshostname resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Hostname.Equal(state.Hostname) {
		tflog.Debug(ctx, "hostname has changed for nshostname")
		hasChange = true
	}
	if !data.Ownernode.IsUnknown() && !data.Ownernode.Equal(state.Ownernode) {
		tflog.Debug(ctx, "ownernode has changed for nshostname")
		hasChange = true
	}

	if hasChange {
		nshostname := nshostnameGetThePayloadFromtheConfig(ctx, &data)
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Nshostname.Type(), &nshostname)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nshostname, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nshostname resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nshostname resource, skipping update")
	}

	// Read the updated state back
	if !r.readNshostnameFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nshostname not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshostnameResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NshostnameResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nshostname resource")
	// nshostname is a singleton and does not support delete on the ADC
	// (matches SDK v2 behavior) - just remove it from Terraform state.
	tflog.Trace(ctx, "Removed nshostname from Terraform state")
}

// Helper function to read nshostname data from API
func (r *NshostnameResource) readNshostnameFromApi(ctx context.Context, data *NshostnameResourceModel, diags *diag.Diagnostics) bool {
	// Singleton resource - simple find without ID
	getResponseData, err := r.client.FindResource(service.Nshostname.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nshostname, got error: %s", err))
		return false
	}

	nshostnameSetAttrFromGet(ctx, data, getResponseData)

	return true
}
