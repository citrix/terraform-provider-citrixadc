package nstrafficdomain

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
var _ resource.Resource = &NstrafficdomainResource{}
var _ resource.ResourceWithConfigure = (*NstrafficdomainResource)(nil)
var _ resource.ResourceWithImportState = (*NstrafficdomainResource)(nil)

func NewNstrafficdomainResource() resource.Resource {
	return &NstrafficdomainResource{}
}

// NstrafficdomainResource defines the resource implementation.
type NstrafficdomainResource struct {
	client *service.NitroClient
}

func (r *NstrafficdomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstrafficdomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstrafficdomain"
}

func (r *NstrafficdomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstrafficdomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstrafficdomainResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstrafficdomain resource")

	// Create API request body from the model
	nstrafficdomain := nstrafficdomainGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource keyed by td - use AddResource (matches SDK v2)
	td_value := fmt.Sprintf("%d", data.Td.ValueInt64())
	_, err := r.client.AddResource(service.Nstrafficdomain.Type(), td_value, &nstrafficdomain)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstrafficdomain, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nstrafficdomain resource")

	// Set ID for the resource before reading state (plain td value, matches SDK v2)
	data.Id = types.StringValue(td_value)

	// Read the updated state back
	if !r.readNstrafficdomainFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nstrafficdomain not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstrafficdomainResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstrafficdomain resource")

	found := r.readNstrafficdomainFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NstrafficdomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NstrafficdomainResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nstrafficdomain resource")

	// All attributes (td, aliasname, vmac) are ForceNew in SDK v2 and carry
	// RequiresReplace/RequiresReplaceIfConfigured plan modifiers, so there is no
	// in-place update path. Simply refresh state from the ADC.
	if !r.readNstrafficdomainFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nstrafficdomain not found immediately after update")
		}
		return
	}

	tflog.Trace(ctx, "Updated nstrafficdomain resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstrafficdomainResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstrafficdomain resource")

	// Named resource - delete using DeleteResource keyed by td (matches SDK v2)
	td_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nstrafficdomain.Type(), td_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nstrafficdomain, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nstrafficdomain resource")
}

// Helper function to read nstrafficdomain data from API
func (r *NstrafficdomainResource) readNstrafficdomainFromApi(ctx context.Context, data *NstrafficdomainResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain td value
	td_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nstrafficdomain.Type(), td_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstrafficdomain, got error: %s", err))
		return false
	}

	nstrafficdomainSetAttrFromGet(ctx, data, getResponseData)

	return true
}
