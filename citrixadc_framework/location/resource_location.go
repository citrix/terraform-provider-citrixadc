package location

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LocationResource{}
var _ resource.ResourceWithConfigure = (*LocationResource)(nil)
var _ resource.ResourceWithImportState = (*LocationResource)(nil)

func NewLocationResource() resource.Resource {
	return &LocationResource{}
}

// LocationResource defines the resource implementation.
type LocationResource struct {
	client *service.NitroClient
}

func (r *LocationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LocationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_location"
}

func (r *LocationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LocationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LocationResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating location resource")

	location := locationGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource. The resource name is ipfrom.
	ipfrom_value := data.Ipfrom.ValueString()
	_, err := r.client.AddResource(service.Location.Type(), ipfrom_value, &location)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create location, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created location resource")

	// Set ID for the resource before reading state (single unique attr: ipfrom)
	data.Id = types.StringValue(ipfrom_value)

	// Read the updated state back
	if !r.readLocationFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "location not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LocationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LocationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading location resource")

	found := r.readLocationFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LocationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// location has no NITRO update endpoint and every attribute is ForceNew
	// (RequiresReplace / RequiresReplaceIfConfigured), so Update is never
	// invoked by Terraform. Kept defensively: preserve the ID and re-read.
	var data, state LocationResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating location resource")

	// Read the current state back (no update API call for location)
	if !r.readLocationFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "location not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LocationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LocationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting location resource")

	// Only the ipfrom and ipto properties are required for deletion (matches SDK v2).
	argsMap := make(map[string]string)
	argsMap["ipfrom"] = url.QueryEscape(data.Ipfrom.ValueString())
	argsMap["ipto"] = url.QueryEscape(data.Ipto.ValueString())
	err := r.client.DeleteResourceWithArgsMap(service.Location.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete location, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted location resource")
}

// Helper function to read location data from API.
// Returns false (without an error) when the resource no longer exists.
func (r *LocationResource) readLocationFromApi(ctx context.Context, data *LocationResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (ipfrom)
	ipfrom_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Location.Type(), ipfrom_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read location, got error: %s", err))
		return false
	}

	locationSetAttrFromGet(ctx, data, getResponseData)

	return true
}
