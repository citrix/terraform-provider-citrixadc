package responderhtmlpage

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
var _ resource.Resource = &ResponderhtmlpageResource{}
var _ resource.ResourceWithConfigure = (*ResponderhtmlpageResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderhtmlpageResource)(nil)

func NewResponderhtmlpageResource() resource.Resource {
	return &ResponderhtmlpageResource{}
}

// ResponderhtmlpageResource defines the resource implementation.
type ResponderhtmlpageResource struct {
	client *service.NitroClient
}

func (r *ResponderhtmlpageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderhtmlpageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderhtmlpage"
}

func (r *ResponderhtmlpageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderhtmlpageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderhtmlpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderhtmlpage resource")

	responderhtmlpage := responderhtmlpageGetThePayloadFromtheConfig(ctx, &data)

	// responderhtmlpage is created via the NITRO ?action=Import endpoint
	// (POST). This mirrors the SDK v2 resource, which called
	// ActOnResource(..., "Import") (note the capital "I" — NITRO action names
	// are case-sensitive).
	err := r.client.ActOnResource(service.Responderhtmlpage.Type(), &responderhtmlpage, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderhtmlpage, got error: %s", err))
		return
	}

	// Single unique attribute — ID is the plain name value (matches SDK v2 d.SetId(name)).
	data.Id = types.StringValue(data.Name.ValueString())

	tflog.Trace(ctx, "Created responderhtmlpage resource")

	// Read the updated state back
	if !r.readResponderhtmlpageFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "responderhtmlpage not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderhtmlpageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderhtmlpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderhtmlpage resource")

	found := r.readResponderhtmlpageFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// Resource deleted out-of-band — remove from state (mirrors SDK v2 d.SetId("")).
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderhtmlpageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ResponderhtmlpageResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for responderhtmlpage: every attribute is ForceNew
	// (RequiresReplace), matching the SDK v2 resource which defined no
	// UpdateContext. Any attribute change forces recreation, so this method is
	// never invoked with real changes; re-read to keep state fresh.
	tflog.Debug(ctx, "Update is a no-op for responderhtmlpage; all attributes are RequiresReplace")

	if !r.readResponderhtmlpageFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "responderhtmlpage not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderhtmlpageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderhtmlpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderhtmlpage resource")

	// Named resource — delete by name (DELETE /responderhtmlpage/{name}), matching SDK v2.
	err := r.client.DeleteResource(service.Responderhtmlpage.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete responderhtmlpage, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted responderhtmlpage resource")
}

// Helper function to read responderhtmlpage data from API.
// Returns false (without an error diagnostic) when the resource no longer exists.
func (r *ResponderhtmlpageResource) readResponderhtmlpageFromApi(ctx context.Context, data *ResponderhtmlpageResourceModel, diags *diag.Diagnostics) bool {
	// Single unique attribute — ID is the plain name value.
	name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Responderhtmlpage.Type(), name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderhtmlpage, got error: %s", err))
		return false
	}

	responderhtmlpageSetAttrFromGet(ctx, data, getResponseData)

	return true
}
