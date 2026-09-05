package nssimpleacl

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
var _ resource.Resource = &NssimpleaclResource{}
var _ resource.ResourceWithConfigure = (*NssimpleaclResource)(nil)
var _ resource.ResourceWithImportState = (*NssimpleaclResource)(nil)

func NewNssimpleaclResource() resource.Resource {
	return &NssimpleaclResource{}
}

// NssimpleaclResource defines the resource implementation.
type NssimpleaclResource struct {
	client *service.NitroClient
}

func (r *NssimpleaclResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NssimpleaclResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nssimpleacl"
}

func (r *NssimpleaclResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NssimpleaclResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NssimpleaclResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nssimpleacl resource")

	nssimpleacl := nssimpleaclGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource (NITRO add is HTTP POST)
	aclnameValue := data.Aclname.ValueString()
	_, err := r.client.AddResource(service.Nssimpleacl.Type(), aclnameValue, &nssimpleacl)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nssimpleacl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nssimpleacl resource")

	// Set ID for the resource before reading state (single unique attribute)
	data.Id = types.StringValue(aclnameValue)

	// Read the updated state back
	if !r.readNssimpleaclFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nssimpleacl not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssimpleaclResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NssimpleaclResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nssimpleacl resource")

	found := r.readNssimpleaclFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NssimpleaclResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NssimpleaclResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// nssimpleacl exposes no NITRO update endpoint and every attribute is
	// ForceNew/RequiresReplace (matching the SDK v2 resource, which had no
	// UpdateContext). Any attribute change forces recreation, so Update is a
	// documented no-op that simply reads the current state back.
	tflog.Debug(ctx, "Update is a no-op for nssimpleacl; all attributes are RequiresReplace")

	r.readNssimpleaclFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssimpleaclResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NssimpleaclResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nssimpleacl resource")

	// Named resource - delete using DeleteResource (NITRO exposes DELETE /nssimpleacl/{aclname})
	aclnameValue := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nssimpleacl.Type(), aclnameValue)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nssimpleacl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nssimpleacl resource")
}

// Helper function to read nssimpleacl data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *NssimpleaclResource) readNssimpleaclFromApi(ctx context.Context, data *NssimpleaclResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (aclname)
	aclnameName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nssimpleacl.Type(), aclnameName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nssimpleacl, got error: %s", err))
		return false
	}

	nssimpleaclSetAttrFromGet(ctx, data, getResponseData)

	return true
}
