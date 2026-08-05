package responderpolicy

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

	sdkv2resource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ResponderpolicyResource{}
var _ resource.ResourceWithConfigure = (*ResponderpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderpolicyResource)(nil)

func NewResponderpolicyResource() resource.Resource {
	return &ResponderpolicyResource{}
}

// ResponderpolicyResource defines the resource implementation.
type ResponderpolicyResource struct {
	client *service.NitroClient
}

func (r *ResponderpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderpolicy"
}

func (r *ResponderpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderpolicy resource")

	// Backward-compatible with the SDK v2 resource: name is optional. When the user
	// does not supply a name, generate a unique one.
	responderpolicyName := data.Name.ValueString()
	if data.Name.IsNull() || data.Name.IsUnknown() || responderpolicyName == "" {
		responderpolicyName = sdkv2resource.PrefixedUniqueId("tf-responderpolicy-")
		data.Name = types.StringValue(responderpolicyName)
	}

	// Create API request body from the model
	responderpolicy := responderpolicyGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is HTTP POST)
	_, err := r.client.AddResource(service.Responderpolicy.Type(), responderpolicyName, &responderpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created responderpolicy resource")

	// Set ID for the resource before applying bindings / reading state
	data.Id = types.StringValue(responderpolicyName)

	// Apply the convenience-block bindings (global / lbvserver / csvserver).
	r.applyBindingsOnCreate(ctx, responderpolicyName, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the updated state back
	if !r.readResponderpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "responderpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderpolicy resource")

	found := r.readResponderpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ResponderpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ResponderpolicyResourceModel

	// Read Terraform prior state to preserve ID and diff the bindings.
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id
	responderpolicyName := state.Id.ValueString()

	tflog.Debug(ctx, "Updating responderpolicy resource")

	// Check for changes in the base (scalar) attributes.
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for responderpolicy")
		hasChange = true
	}
	if !data.Appflowaction.Equal(state.Appflowaction) {
		tflog.Debug(ctx, "appflowaction has changed for responderpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for responderpolicy")
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for responderpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for responderpolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for responderpolicy")
		hasChange = true
	}

	if hasChange {
		responderpolicy := responderpolicyGetThePayloadFromthePlan(ctx, &data)
		responderpolicy.Name = responderpolicyName
		_, err := r.client.UpdateResource(service.Responderpolicy.Type(), responderpolicyName, &responderpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update responderpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated responderpolicy resource")
	} else {
		tflog.Debug(ctx, "No base attribute changes detected for responderpolicy resource, skipping update")
	}

	// Reconcile the convenience-block bindings against prior state.
	r.applyBindingsOnUpdate(ctx, responderpolicyName, &data, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the updated state back
	if !r.readResponderpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "responderpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderpolicy resource")

	responderpolicyName := data.Id.ValueString()

	// Delete all bindings prior to deleting the responder policy.
	r.deleteAllBindings(ctx, responderpolicyName, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Responderpolicy.Type(), responderpolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete responderpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted responderpolicy resource")
}

// Helper function to read responderpolicy data from API. Returns false when the
// resource no longer exists on the appliance.
func (r *ResponderpolicyResource) readResponderpolicyFromApi(ctx context.Context, data *ResponderpolicyResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (the name).
	responderpolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Responderpolicy.Type(), responderpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderpolicy, got error: %s", err))
		return false
	}

	responderpolicySetAttrFromGet(ctx, data, getResponseData)

	// Refresh the managed convenience-block bindings.
	r.readBindings(ctx, responderpolicyName, data)

	return true
}
