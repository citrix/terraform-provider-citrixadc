package contentinspectionaction

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
var _ resource.Resource = &ContentinspectionactionResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionactionResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionactionResource)(nil)

func NewContentinspectionactionResource() resource.Resource {
	return &ContentinspectionactionResource{}
}

// ContentinspectionactionResource defines the resource implementation.
type ContentinspectionactionResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionaction"
}

func (r *ContentinspectionactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionaction resource")

	// contentinspectionaction := contentinspectionactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Contentinspectionaction.Type(), &contentinspectionaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("contentinspectionaction-config")

	tflog.Trace(ctx, "Created contentinspectionaction resource")

	// Read the updated state back
	r.readContentinspectionactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionaction resource")

	r.readContentinspectionactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ContentinspectionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating contentinspectionaction resource")

	// Create API request body from the model
	// contentinspectionaction := contentinspectionactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Contentinspectionaction.Type(), &contentinspectionaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated contentinspectionaction resource")

	// Read the updated state back
	r.readContentinspectionactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionaction resource")

	// For contentinspectionaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted contentinspectionaction resource from state")
}

// Helper function to read contentinspectionaction data from API
func (r *ContentinspectionactionResource) readContentinspectionactionFromApi(ctx context.Context, data *ContentinspectionactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Contentinspectionaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionaction, got error: %s", err))
		return
	}

	contentinspectionactionSetAttrFromGet(ctx, data, getResponseData)

}
