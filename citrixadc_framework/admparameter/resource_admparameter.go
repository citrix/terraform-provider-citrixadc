package admparameter

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

	sdkresource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// admparameter is an unnamed singleton resource. The vendored adc-nitro-go
// service package does not expose a service.Admparameter enum, so the NITRO
// resource token is referenced by its literal name "admparameter" exactly as
// the SDK v2 resource did.
const admparameterResourceType = "admparameter"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AdmparameterResource{}
var _ resource.ResourceWithConfigure = (*AdmparameterResource)(nil)
var _ resource.ResourceWithImportState = (*AdmparameterResource)(nil)

func NewAdmparameterResource() resource.Resource {
	return &AdmparameterResource{}
}

// AdmparameterResource defines the resource implementation.
type AdmparameterResource struct {
	client *service.NitroClient
}

func (r *AdmparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AdmparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_admparameter"
}

func (r *AdmparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AdmparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AdmparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating admparameter resource")

	// Get payload from plan
	admparameter := admparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(admparameterResourceType, &admparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create admparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created admparameter resource")

	// Set ID for the resource before reading state.
	// Mirror SDK v2: d.SetId(resource.PrefixedUniqueId("tf-admparameter-")).
	data.Id = types.StringValue(sdkresource.PrefixedUniqueId("tf-admparameter-"))

	// Read the updated state back
	if !r.readAdmparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "admparameter not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AdmparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AdmparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading admparameter resource")

	found := r.readAdmparameterFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AdmparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AdmparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating admparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Admserviceconnect.Equal(state.Admserviceconnect) {
		tflog.Debug(ctx, "admserviceconnect has changed for admparameter")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		admparameter := admparameterGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Unnamed singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(admparameterResourceType, &admparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update admparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated admparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for admparameter resource, skipping update")
	}

	// Read the updated state back
	if !r.readAdmparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "admparameter not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AdmparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AdmparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting admparameter resource")
	// admparameter does not support a DELETE operation (it is a global,
	// unnamed configuration singleton). Mirror SDK v2, which performed no NITRO
	// call and simply cleared the ID. The framework removes it from state.
	tflog.Trace(ctx, "Deleted admparameter resource from state")
}

// Helper function to read admparameter data from API
func (r *AdmparameterResource) readAdmparameterFromApi(ctx context.Context, data *AdmparameterResourceModel, diags *diag.Diagnostics) bool {

	// Unnamed singleton resource - read with empty name
	getResponseData, err := r.client.FindResource(admparameterResourceType, "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read admparameter, got error: %s", err))
		return false
	}

	admparameterSetAttrFromGet(ctx, data, getResponseData)

	return true
}
