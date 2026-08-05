package csaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CsactionResource{}
var _ resource.ResourceWithConfigure = (*CsactionResource)(nil)
var _ resource.ResourceWithImportState = (*CsactionResource)(nil)

func NewCsactionResource() resource.Resource {
	return &CsactionResource{}
}

// CsactionResource defines the resource implementation.
type CsactionResource struct {
	client *service.NitroClient
}

func (r *CsactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csaction"
}

func (r *CsactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csaction resource")

	csaction := csactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Csaction.Type(), name_value, &csaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created csaction resource")

	// Set ID for the resource before reading state (single unique attr -> plain value)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readCsactionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "csaction not found immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csaction resource")

	r.readCsactionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Resource no longer exists on the ADC - drop it from state.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CsactionResourceModel

	// Read Terraform prior state to preserve ID / live name
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating csaction resource")

	// Rename support: csaction exposes a NITRO `rename` action. The primary key
	// (name) uses RequiresReplace, so a name change recreates the resource and never
	// reaches Update. The only rename-in-place path is a change to `newname`.
	// Mirrors the SDK v2 convention (see citrixadc/resource_citrixadc_appfwpolicy.go).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Name (which stays pinned to the originally configured value and would
		// point at a no-longer-live name on a second rename).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming csaction from %q to %q", oldName, newName))

		renamePayload := cs.Csaction{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Csaction.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename csaction, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the read below
		// (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for csaction")
		hasChange = true
	}
	if !data.Targetlbvserver.Equal(state.Targetlbvserver) {
		tflog.Debug(ctx, "targetlbvserver has changed for csaction")
		hasChange = true
	}
	if !data.Targetvserver.Equal(state.Targetvserver) {
		tflog.Debug(ctx, "targetvserver has changed for csaction")
		hasChange = true
	}
	if !data.Targetvserverexpr.Equal(state.Targetvserverexpr) {
		tflog.Debug(ctx, "targetvserverexpr has changed for csaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model. The update PUT is keyed on the
		// current live name, so use the ID (post-rename), not the configured name.
		csaction := csactionGetThePayloadFromthePlan(ctx, &data)
		csaction.Name = data.Id.ValueString()
		// Make API call - named resource, use UpdateResource
		_, err := r.client.UpdateResource(service.Csaction.Type(), data.Id.ValueString(), &csaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated csaction resource")
	} else {
		tflog.Debug(ctx, "No updateable changes detected for csaction resource, skipping update")
	}

	// Read the updated state back. Preserve the plan's user-facing key/newname across
	// the read so a rename does not clobber the configured values.
	planName := data.Name
	planNewname := data.Newname
	r.readCsactionFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csaction resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id, NOT
	// data.Name (which would be stale after a rename and dangle the renamed object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Csaction.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete csaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted csaction resource")
}

// Helper function to read csaction data from API
func (r *CsactionResource) readCsactionFromApi(ctx context.Context, data *CsactionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value (live name)
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Csaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csaction, got error: %s", err))
		return
	}

	csactionSetAttrFromGet(ctx, data, getResponseData)
}
