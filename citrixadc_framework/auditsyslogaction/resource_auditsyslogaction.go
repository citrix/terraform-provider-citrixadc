package auditsyslogaction

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
var _ resource.Resource = &AuditsyslogactionResource{}
var _ resource.ResourceWithConfigure = (*AuditsyslogactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuditsyslogactionResource)(nil)

func NewAuditsyslogactionResource() resource.Resource {
	return &AuditsyslogactionResource{}
}

// AuditsyslogactionResource defines the resource implementation.
type AuditsyslogactionResource struct {
	client *service.NitroClient
}

func (r *AuditsyslogactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditsyslogactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditsyslogaction"
}

func (r *AuditsyslogactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditsyslogactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuditsyslogactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditsyslogaction resource")

	// auditsyslogaction := auditsyslogactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Auditsyslogaction.Type(), &auditsyslogaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditsyslogaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("auditsyslogaction-config")

	tflog.Trace(ctx, "Created auditsyslogaction resource")

	// Read the updated state back
	r.readAuditsyslogactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditsyslogactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditsyslogaction resource")

	r.readAuditsyslogactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuditsyslogactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating auditsyslogaction resource")

	// Create API request body from the model
	// auditsyslogaction := auditsyslogactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Auditsyslogaction.Type(), &auditsyslogaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditsyslogaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated auditsyslogaction resource")

	// Read the updated state back
	r.readAuditsyslogactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditsyslogactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditsyslogaction resource")

	// For auditsyslogaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted auditsyslogaction resource from state")
}

// Helper function to read auditsyslogaction data from API
func (r *AuditsyslogactionResource) readAuditsyslogactionFromApi(ctx context.Context, data *AuditsyslogactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Auditsyslogaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditsyslogaction, got error: %s", err))
		return
	}

	auditsyslogactionSetAttrFromGet(ctx, data, getResponseData)

}
