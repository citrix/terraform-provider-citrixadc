package dnssuffix

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
var _ resource.Resource = &DnssuffixResource{}
var _ resource.ResourceWithConfigure = (*DnssuffixResource)(nil)
var _ resource.ResourceWithImportState = (*DnssuffixResource)(nil)

func NewDnssuffixResource() resource.Resource {
	return &DnssuffixResource{}
}

// DnssuffixResource defines the resource implementation.
type DnssuffixResource struct {
	client *service.NitroClient
}

func (r *DnssuffixResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnssuffixResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssuffix"
}

func (r *DnssuffixResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnssuffixResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnssuffixResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnssuffix resource")

	dnssuffix := dnssuffixGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	dnssuffixName := data.Dnssuffix.ValueString()
	_, err := r.client.AddResource(service.Dnssuffix.Type(), dnssuffixName, &dnssuffix)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnssuffix, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnssuffix resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(dnssuffix))
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Dnssuffix.ValueString()))

	// Read the updated state back
	if !r.readDnssuffixFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnssuffix not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssuffixResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnssuffixResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnssuffix resource")

	found := r.readDnssuffixFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnssuffixResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnssuffixResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnssuffix resource")

	// dnssuffix has no NITRO-updatable attributes: the only configurable attribute is
	// "dnssuffix" itself, which is the primary key and carries RequiresReplace (matching
	// SDK v2 ForceNew). Any change forces a destroy/recreate, so Update never issues a
	// NITRO write call. Just read the current state back.

	// Read the updated state back
	if !r.readDnssuffixFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnssuffix not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssuffixResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnssuffixResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnssuffix resource")

	// Named resource - delete using DeleteResource (matches SDK v2 DeleteResource(type, d.Id()))
	dnssuffixName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dnssuffix.Type(), dnssuffixName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnssuffix, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnssuffix resource")
}

// Helper function to read dnssuffix data from API
func (r *DnssuffixResource) readDnssuffixFromApi(ctx context.Context, data *DnssuffixResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (the dnssuffix name)
	dnssuffixName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Dnssuffix.Type(), dnssuffixName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnssuffix, got error: %s", err))
		return false
	}

	dnssuffixSetAttrFromGet(ctx, data, getResponseData)

	return true
}
