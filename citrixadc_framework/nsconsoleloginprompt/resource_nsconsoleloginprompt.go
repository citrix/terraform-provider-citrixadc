package nsconsoleloginprompt

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
var _ resource.Resource = &NsconsoleloginpromptResource{}
var _ resource.ResourceWithConfigure = (*NsconsoleloginpromptResource)(nil)
var _ resource.ResourceWithImportState = (*NsconsoleloginpromptResource)(nil)

func NewNsconsoleloginpromptResource() resource.Resource {
	return &NsconsoleloginpromptResource{}
}

// NsconsoleloginpromptResource defines the resource implementation.
type NsconsoleloginpromptResource struct {
	client *service.NitroClient
}

func (r *NsconsoleloginpromptResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsconsoleloginpromptResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsconsoleloginprompt"
}

func (r *NsconsoleloginpromptResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsconsoleloginpromptResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsconsoleloginpromptResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsconsoleloginprompt resource")

	// nsconsoleloginprompt := nsconsoleloginpromptGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsconsoleloginprompt.Type(), &nsconsoleloginprompt)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsconsoleloginprompt, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nsconsoleloginprompt-config")

	tflog.Trace(ctx, "Created nsconsoleloginprompt resource")

	// Read the updated state back
	r.readNsconsoleloginpromptFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconsoleloginpromptResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsconsoleloginpromptResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsconsoleloginprompt resource")

	r.readNsconsoleloginpromptFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconsoleloginpromptResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NsconsoleloginpromptResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsconsoleloginprompt resource")

	// Create API request body from the model
	// nsconsoleloginprompt := nsconsoleloginpromptGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsconsoleloginprompt.Type(), &nsconsoleloginprompt)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsconsoleloginprompt, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nsconsoleloginprompt resource")

	// Read the updated state back
	r.readNsconsoleloginpromptFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconsoleloginpromptResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsconsoleloginpromptResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsconsoleloginprompt resource")

	// For nsconsoleloginprompt, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nsconsoleloginprompt resource from state")
}

// Helper function to read nsconsoleloginprompt data from API
func (r *NsconsoleloginpromptResource) readNsconsoleloginpromptFromApi(ctx context.Context, data *NsconsoleloginpromptResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsconsoleloginprompt.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsconsoleloginprompt, got error: %s", err))
		return
	}

	nsconsoleloginpromptSetAttrFromGet(ctx, data, getResponseData)

}
