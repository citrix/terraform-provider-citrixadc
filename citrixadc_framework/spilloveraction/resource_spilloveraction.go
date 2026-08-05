package spilloveraction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/spillover"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SpilloveractionResource{}
var _ resource.ResourceWithConfigure = (*SpilloveractionResource)(nil)
var _ resource.ResourceWithImportState = (*SpilloveractionResource)(nil)

func NewSpilloveractionResource() resource.Resource {
	return &SpilloveractionResource{}
}

// SpilloveractionResource defines the resource implementation.
type SpilloveractionResource struct {
	client *service.NitroClient
}

func (r *SpilloveractionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SpilloveractionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_spilloveraction"
}

func (r *SpilloveractionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SpilloveractionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SpilloveractionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating spilloveraction resource")
	spilloveraction := spilloveractionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource (matches SDK v2 client.AddResource)
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Spilloveraction.Type(), name_value, &spilloveraction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create spilloveraction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created spilloveraction resource")

	// Set ID for the resource before reading state (SDK v2 parity: ID == name)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readSpilloveractionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SpilloveractionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SpilloveractionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading spilloveraction resource")

	r.readSpilloveractionFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SpilloveractionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SpilloveractionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating spilloveraction resource")

	// Rename support: spilloveraction exposes NO set/update endpoint (NITRO
	// operations are add/delete/get/rename only). The only in-place mutation is
	// the `rename` action. Every other schema attribute (name, action) uses
	// RequiresReplace / RequiresReplaceIfConfigured, so Terraform recreates the
	// resource on any of those changes and never reaches here for them. The ONLY
	// change that lands in Update is `newname`.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, which is tracked by the ID -
		// NOT state.Name. state.Name stays pinned to the originally configured value,
		// so on a SECOND rename it would point at the wrong (no longer live) name.
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming spilloveraction from %q to %q", oldName, newName))

		renamePayload := spillover.Spilloveraction{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Spilloveraction.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename spilloveraction, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the read
		// below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Read the current state back. Capture the plan values for the identity-bearing
	// attributes and restore them after the read so GET (which returns the live/new
	// name) does not clobber the user-facing configuration.
	planName := data.Name
	planNewname := data.Newname
	r.readSpilloveractionFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SpilloveractionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SpilloveractionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting spilloveraction resource")
	// Named resource - delete using DeleteResource (matches SDK v2). The ID holds
	// the CURRENT LIVE name (== name at create, == newname after a rename), so we
	// must delete by data.Id, NOT data.Name (which stays at the originally
	// configured value and would target a non-existent name after a rename).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Spilloveraction.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete spilloveraction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted spilloveraction resource")
}

// Helper function to read spilloveraction data from API
func (r *SpilloveractionResource) readSpilloveractionFromApi(ctx context.Context, data *SpilloveractionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value (the name)
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Spilloveraction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read spilloveraction, got error: %s", err))
		return
	}

	spilloveractionSetAttrFromGet(ctx, data, getResponseData)
}
