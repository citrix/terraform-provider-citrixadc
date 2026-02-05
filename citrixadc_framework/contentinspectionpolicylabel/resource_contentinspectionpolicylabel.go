package contentinspectionpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ContentinspectionpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionpolicylabelResource)(nil)

func NewContentinspectionpolicylabelResource() resource.Resource {
	return &ContentinspectionpolicylabelResource{}
}

// ContentinspectionpolicylabelResource defines the resource implementation.
type ContentinspectionpolicylabelResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionpolicylabel"
}

func (r *ContentinspectionpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionpolicylabel resource")

	// contentinspectionpolicylabel := contentinspectionpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Contentinspectionpolicylabel.Type(), &contentinspectionpolicylabel)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionpolicylabel, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("contentinspectionpolicylabel-config")

	tflog.Trace(ctx, "Created contentinspectionpolicylabel resource")

	// Read the updated state back
	r.readContentinspectionpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionpolicylabel resource")

	r.readContentinspectionpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ContentinspectionpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating contentinspectionpolicylabel resource")

	// Create API request body from the model
	// contentinspectionpolicylabel := contentinspectionpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Contentinspectionpolicylabel.Type(), &contentinspectionpolicylabel)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionpolicylabel, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated contentinspectionpolicylabel resource")

	// Read the updated state back
	r.readContentinspectionpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionpolicylabel resource")

	// For contentinspectionpolicylabel, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted contentinspectionpolicylabel resource from state")
}

// Helper function to read contentinspectionpolicylabel data from API
func (r *ContentinspectionpolicylabelResource) readContentinspectionpolicylabelFromApi(ctx context.Context, data *ContentinspectionpolicylabelResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Contentinspectionpolicylabel.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionpolicylabel, got error: %s", err))
		return
	}

	contentinspectionpolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
