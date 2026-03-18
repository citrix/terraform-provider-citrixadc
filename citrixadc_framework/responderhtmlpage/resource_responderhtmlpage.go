package responderhtmlpage

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
var _ resource.Resource = &ResponderhtmlpageResource{}
var _ resource.ResourceWithConfigure = (*ResponderhtmlpageResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderhtmlpageResource)(nil)

func NewResponderhtmlpageResource() resource.Resource {
	return &ResponderhtmlpageResource{}
}

// ResponderhtmlpageResource defines the resource implementation.
type ResponderhtmlpageResource struct {
	client *service.NitroClient
}

func (r *ResponderhtmlpageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderhtmlpageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderhtmlpage"
}

func (r *ResponderhtmlpageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderhtmlpageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderhtmlpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderhtmlpage resource")

	// responderhtmlpage := responderhtmlpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Responderhtmlpage.Type(), &responderhtmlpage)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderhtmlpage, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("responderhtmlpage-config")

	tflog.Trace(ctx, "Created responderhtmlpage resource")

	// Read the updated state back
	r.readResponderhtmlpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderhtmlpageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderhtmlpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderhtmlpage resource")

	r.readResponderhtmlpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderhtmlpageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ResponderhtmlpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating responderhtmlpage resource")

	// Create API request body from the model
	// responderhtmlpage := responderhtmlpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Responderhtmlpage.Type(), &responderhtmlpage)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update responderhtmlpage, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated responderhtmlpage resource")

	// Read the updated state back
	r.readResponderhtmlpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderhtmlpageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderhtmlpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderhtmlpage resource")

	// For responderhtmlpage, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted responderhtmlpage resource from state")
}

// Helper function to read responderhtmlpage data from API
func (r *ResponderhtmlpageResource) readResponderhtmlpageFromApi(ctx context.Context, data *ResponderhtmlpageResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Responderhtmlpage.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderhtmlpage, got error: %s", err))
		return
	}

	responderhtmlpageSetAttrFromGet(ctx, data, getResponseData)

}
