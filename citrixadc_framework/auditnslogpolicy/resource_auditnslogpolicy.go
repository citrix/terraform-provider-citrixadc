package auditnslogpolicy

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
var _ resource.Resource = &AuditnslogpolicyResource{}
var _ resource.ResourceWithConfigure = (*AuditnslogpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuditnslogpolicyResource)(nil)

func NewAuditnslogpolicyResource() resource.Resource {
	return &AuditnslogpolicyResource{}
}

// AuditnslogpolicyResource defines the resource implementation.
type AuditnslogpolicyResource struct {
	client *service.NitroClient
}

func (r *AuditnslogpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditnslogpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditnslogpolicy"
}

func (r *AuditnslogpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditnslogpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuditnslogpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditnslogpolicy resource")
	// Get payload from plan
	auditnslogpolicy := auditnslogpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Auditnslogpolicy.Type(), name_value, &auditnslogpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditnslogpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created auditnslogpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuditnslogpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "auditnslogpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditnslogpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditnslogpolicy resource")

	found := r.readAuditnslogpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuditnslogpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuditnslogpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating auditnslogpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for auditnslogpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for auditnslogpolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		auditnslogpolicy := auditnslogpolicyGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// auditnslogpolicy update is an unnamed PUT (name is carried in the payload)
		err := r.client.UpdateUnnamedResource(service.Auditnslogpolicy.Type(), &auditnslogpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditnslogpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated auditnslogpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for auditnslogpolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuditnslogpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "auditnslogpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditnslogpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditnslogpolicy resource")
	// Named resource - delete using DeleteResource
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Auditnslogpolicy.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete auditnslogpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted auditnslogpolicy resource")
}

// Helper function to read auditnslogpolicy data from API
func (r *AuditnslogpolicyResource) readAuditnslogpolicyFromApi(ctx context.Context, data *AuditnslogpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Auditnslogpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditnslogpolicy, got error: %s", err))
		return false
	}

	auditnslogpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
