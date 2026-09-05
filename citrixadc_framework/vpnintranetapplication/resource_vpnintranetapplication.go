package vpnintranetapplication

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
var _ resource.Resource = &VpnintranetapplicationResource{}
var _ resource.ResourceWithConfigure = (*VpnintranetapplicationResource)(nil)
var _ resource.ResourceWithImportState = (*VpnintranetapplicationResource)(nil)

func NewVpnintranetapplicationResource() resource.Resource {
	return &VpnintranetapplicationResource{}
}

// VpnintranetapplicationResource defines the resource implementation.
type VpnintranetapplicationResource struct {
	client *service.NitroClient
}

func (r *VpnintranetapplicationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnintranetapplicationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnintranetapplication"
}

func (r *VpnintranetapplicationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnintranetapplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnintranetapplicationResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnintranetapplication resource")

	// Create API request body from the model
	vpnintranetapplication := vpnintranetapplicationGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	vpnintranetapplicationName := data.Intranetapplication.ValueString()
	_, err := r.client.AddResource(service.Vpnintranetapplication.Type(), vpnintranetapplicationName, &vpnintranetapplication)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnintranetapplication, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnintranetapplication resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(vpnintranetapplicationName)

	// Read the updated state back
	if !r.readVpnintranetapplicationFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnintranetapplication not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnintranetapplicationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnintranetapplicationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnintranetapplication resource")

	found := r.readVpnintranetapplicationFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnintranetapplicationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All attributes are ForceNew (RequiresReplace) in SDK v2, so there are no
	// in-place updatable fields. Update is effectively unreachable, but the
	// framework requires the method; preserve the ID and read state back.
	var data, state VpnintranetapplicationResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnintranetapplication resource (no updatable attributes)")

	// Read the updated state back
	if !r.readVpnintranetapplicationFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnintranetapplication not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnintranetapplicationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnintranetapplicationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnintranetapplication resource")

	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Vpnintranetapplication.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnintranetapplication, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnintranetapplication resource")
}

// Helper function to read vpnintranetapplication data from API.
// Returns false (without error) when the resource no longer exists.
func (r *VpnintranetapplicationResource) readVpnintranetapplicationFromApi(ctx context.Context, data *VpnintranetapplicationResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	vpnintranetapplicationName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnintranetapplication.Type(), vpnintranetapplicationName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnintranetapplication, got error: %s", err))
		return false
	}

	vpnintranetapplicationSetAttrFromGet(ctx, data, getResponseData)

	return true
}
