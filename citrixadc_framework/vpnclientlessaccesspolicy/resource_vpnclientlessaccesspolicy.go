package vpnclientlessaccesspolicy

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
var _ resource.Resource = &VpnclientlessaccesspolicyResource{}
var _ resource.ResourceWithConfigure = (*VpnclientlessaccesspolicyResource)(nil)
var _ resource.ResourceWithImportState = (*VpnclientlessaccesspolicyResource)(nil)

func NewVpnclientlessaccesspolicyResource() resource.Resource {
	return &VpnclientlessaccesspolicyResource{}
}

// VpnclientlessaccesspolicyResource defines the resource implementation.
type VpnclientlessaccesspolicyResource struct {
	client *service.NitroClient
}

func (r *VpnclientlessaccesspolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnclientlessaccesspolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnclientlessaccesspolicy"
}

func (r *VpnclientlessaccesspolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnclientlessaccesspolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnclientlessaccesspolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnclientlessaccesspolicy resource")

	// Create API request body from the model
	vpnclientlessaccesspolicy := vpnclientlessaccesspolicyGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource (matches SDK v2 client.AddResource)
	vpnclientlessaccesspolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnclientlessaccesspolicy.Type(), vpnclientlessaccesspolicyName, &vpnclientlessaccesspolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnclientlessaccesspolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnclientlessaccesspolicy resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(vpnclientlessaccesspolicyName)

	// Read the updated state back
	if !r.readVpnclientlessaccesspolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnclientlessaccesspolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccesspolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnclientlessaccesspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnclientlessaccesspolicy resource")

	found := r.readVpnclientlessaccesspolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnclientlessaccesspolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnclientlessaccesspolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnclientlessaccesspolicy resource")

	// Check if there are any changes in updateable attributes (name is ForceNew)
	hasChange := false
	if !data.Profilename.Equal(state.Profilename) {
		tflog.Debug(ctx, "profilename has changed for vpnclientlessaccesspolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for vpnclientlessaccesspolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		vpnclientlessaccesspolicy := vpnclientlessaccesspolicyGetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource (matches SDK v2 client.UpdateResource)
		vpnclientlessaccesspolicyName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Vpnclientlessaccesspolicy.Type(), vpnclientlessaccesspolicyName, &vpnclientlessaccesspolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnclientlessaccesspolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vpnclientlessaccesspolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnclientlessaccesspolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readVpnclientlessaccesspolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnclientlessaccesspolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccesspolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnclientlessaccesspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnclientlessaccesspolicy resource")

	// Named resource - delete using DeleteResource (matches SDK v2 client.DeleteResource)
	vpnclientlessaccesspolicyName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnclientlessaccesspolicy.Type(), vpnclientlessaccesspolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnclientlessaccesspolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnclientlessaccesspolicy resource")
}

// Helper function to read vpnclientlessaccesspolicy data from API
func (r *VpnclientlessaccesspolicyResource) readVpnclientlessaccesspolicyFromApi(ctx context.Context, data *VpnclientlessaccesspolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	vpnclientlessaccesspolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnclientlessaccesspolicy.Type(), vpnclientlessaccesspolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnclientlessaccesspolicy, got error: %s", err))
		return false
	}

	vpnclientlessaccesspolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
