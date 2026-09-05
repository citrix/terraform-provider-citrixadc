package tmsessionpolicy

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
var _ resource.Resource = &TmsessionpolicyResource{}
var _ resource.ResourceWithConfigure = (*TmsessionpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*TmsessionpolicyResource)(nil)

func NewTmsessionpolicyResource() resource.Resource {
	return &TmsessionpolicyResource{}
}

// TmsessionpolicyResource defines the resource implementation.
type TmsessionpolicyResource struct {
	client *service.NitroClient
}

func (r *TmsessionpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmsessionpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmsessionpolicy"
}

func (r *TmsessionpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmsessionpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmsessionpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmsessionpolicy resource")

	// Create API request body from the model
	tmsessionpolicy := tmsessionpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	tmsessionpolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Tmsessionpolicy.Type(), tmsessionpolicyName, &tmsessionpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmsessionpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created tmsessionpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readTmsessionpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmsessionpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsessionpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmsessionpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmsessionpolicy resource")

	found := r.readTmsessionpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TmsessionpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state TmsessionpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (name is ForceNew, so it never changes here)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating tmsessionpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for tmsessionpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for tmsessionpolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		tmsessionpolicy := tmsessionpolicyGetThePayloadFromthePlan(ctx, &data)
		// Make API call - matches SDK v2 semantics (UpdateUnnamedResource with name in body)
		err := r.client.UpdateUnnamedResource(service.Tmsessionpolicy.Type(), &tmsessionpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmsessionpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated tmsessionpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for tmsessionpolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readTmsessionpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmsessionpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsessionpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmsessionpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmsessionpolicy resource")
	// Named resource - delete using DeleteResource
	tmsessionpolicyName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Tmsessionpolicy.Type(), tmsessionpolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete tmsessionpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted tmsessionpolicy resource")
}

// Helper function to read tmsessionpolicy data from API
func (r *TmsessionpolicyResource) readTmsessionpolicyFromApi(ctx context.Context, data *TmsessionpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	tmsessionpolicyName := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Tmsessionpolicy.Type(), tmsessionpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmsessionpolicy, got error: %s", err))
		return false
	}

	tmsessionpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
