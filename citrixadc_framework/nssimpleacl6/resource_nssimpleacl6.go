package nssimpleacl6

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
var _ resource.Resource = &Nssimpleacl6Resource{}
var _ resource.ResourceWithConfigure = (*Nssimpleacl6Resource)(nil)
var _ resource.ResourceWithImportState = (*Nssimpleacl6Resource)(nil)

func NewNssimpleacl6Resource() resource.Resource {
	return &Nssimpleacl6Resource{}
}

// Nssimpleacl6Resource defines the resource implementation.
type Nssimpleacl6Resource struct {
	client *service.NitroClient
}

func (r *Nssimpleacl6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nssimpleacl6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nssimpleacl6"
}

func (r *Nssimpleacl6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nssimpleacl6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nssimpleacl6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nssimpleacl6 resource")

	nssimpleacl6 := nssimpleacl6GetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	aclname_value := data.Aclname.ValueString()
	_, err := r.client.AddResource(service.Nssimpleacl6.Type(), aclname_value, &nssimpleacl6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nssimpleacl6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nssimpleacl6 resource")

	// Set ID for the resource before reading state (single unique attribute -> plain value)
	data.Id = types.StringValue(aclname_value)

	// Read the updated state back
	if !r.readNssimpleacl6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nssimpleacl6 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nssimpleacl6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nssimpleacl6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nssimpleacl6 resource")

	found := r.readNssimpleacl6FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Nssimpleacl6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Nssimpleacl6ResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// All nssimpleacl6 attributes are ForceNew in SDK v2 (RequiresReplace /
	// RequiresReplaceIfConfigured), so there are no NITRO-updatable fields. Any
	// real attribute change triggers a destroy/recreate instead of reaching here.
	// This branch simply re-reads current state to stay consistent.
	tflog.Debug(ctx, "Updating nssimpleacl6 resource (no updatable attributes)")

	if !r.readNssimpleacl6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nssimpleacl6 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nssimpleacl6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nssimpleacl6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nssimpleacl6 resource")

	// Named resource - delete using DeleteResource (ID is the plain aclname value)
	err := r.client.DeleteResource(service.Nssimpleacl6.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nssimpleacl6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nssimpleacl6 resource")
}

// Helper function to read nssimpleacl6 data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *Nssimpleacl6Resource) readNssimpleacl6FromApi(ctx context.Context, data *Nssimpleacl6ResourceModel, diags *diag.Diagnostics) bool {

	// Single unique attribute - ID is the plain value
	aclname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nssimpleacl6.Type(), aclname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nssimpleacl6, got error: %s", err))
		return false
	}

	nssimpleacl6SetAttrFromGet(ctx, data, getResponseData)

	return true
}
