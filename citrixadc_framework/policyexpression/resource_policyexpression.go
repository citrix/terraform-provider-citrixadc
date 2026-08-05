package policyexpression

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
var _ resource.Resource = &PolicyexpressionResource{}
var _ resource.ResourceWithConfigure = (*PolicyexpressionResource)(nil)
var _ resource.ResourceWithImportState = (*PolicyexpressionResource)(nil)

func NewPolicyexpressionResource() resource.Resource {
	return &PolicyexpressionResource{}
}

// PolicyexpressionResource defines the resource implementation.
type PolicyexpressionResource struct {
	client *service.NitroClient
}

func (r *PolicyexpressionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicyexpressionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policyexpression"
}

func (r *PolicyexpressionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicyexpressionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicyexpressionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policyexpression resource")

	policyexpression := policyexpressionGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	policyexpressionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Policyexpression.Type(), policyexpressionName, &policyexpression)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policyexpression, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created policyexpression resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readPolicyexpressionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policyexpression not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyexpressionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicyexpressionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policyexpression resource")

	found := r.readPolicyexpressionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *PolicyexpressionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state PolicyexpressionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (name is RequiresReplace, so it never changes here)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating policyexpression resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Clientsecuritymessage.Equal(state.Clientsecuritymessage) {
		tflog.Debug(ctx, "clientsecuritymessage has changed for policyexpression")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for policyexpression")
		hasChange = true
	}
	if !data.Value.Equal(state.Value) {
		tflog.Debug(ctx, "value has changed for policyexpression")
		hasChange = true
	}

	if hasChange {
		policyexpression := policyexpressionGetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource
		policyexpressionName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Policyexpression.Type(), policyexpressionName, &policyexpression)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policyexpression, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated policyexpression resource")
	} else {
		tflog.Debug(ctx, "No changes detected for policyexpression resource, skipping update")
	}

	// Read the updated state back
	if !r.readPolicyexpressionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policyexpression not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyexpressionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicyexpressionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policyexpression resource")

	// Named resource - delete using DeleteResource (keyed by the live name held in the ID)
	policyexpressionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Policyexpression.Type(), policyexpressionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete policyexpression, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted policyexpression resource")
}

// Helper function to read policyexpression data from API. Returns false when the
// resource no longer exists on the ADC.
func (r *PolicyexpressionResource) readPolicyexpressionFromApi(ctx context.Context, data *PolicyexpressionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	policyexpressionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Policyexpression.Type(), policyexpressionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policyexpression, got error: %s", err))
		return false
	}

	policyexpressionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
