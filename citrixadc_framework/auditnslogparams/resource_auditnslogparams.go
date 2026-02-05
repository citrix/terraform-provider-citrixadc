package auditnslogparams

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
var _ resource.Resource = &AuditnslogparamsResource{}
var _ resource.ResourceWithConfigure = (*AuditnslogparamsResource)(nil)
var _ resource.ResourceWithImportState = (*AuditnslogparamsResource)(nil)

func NewAuditnslogparamsResource() resource.Resource {
	return &AuditnslogparamsResource{}
}

// AuditnslogparamsResource defines the resource implementation.
type AuditnslogparamsResource struct {
	client *service.NitroClient
}

func (r *AuditnslogparamsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditnslogparamsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditnslogparams"
}

func (r *AuditnslogparamsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditnslogparamsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuditnslogparamsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditnslogparams resource")

	// auditnslogparams := auditnslogparamsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Auditnslogparams.Type(), &auditnslogparams)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditnslogparams, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("auditnslogparams-config")

	tflog.Trace(ctx, "Created auditnslogparams resource")

	// Read the updated state back
	r.readAuditnslogparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogparamsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditnslogparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditnslogparams resource")

	r.readAuditnslogparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogparamsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuditnslogparamsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating auditnslogparams resource")

	// Create API request body from the model
	// auditnslogparams := auditnslogparamsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Auditnslogparams.Type(), &auditnslogparams)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditnslogparams, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated auditnslogparams resource")

	// Read the updated state back
	r.readAuditnslogparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogparamsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditnslogparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditnslogparams resource")

	// For auditnslogparams, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted auditnslogparams resource from state")
}

// Helper function to read auditnslogparams data from API
func (r *AuditnslogparamsResource) readAuditnslogparamsFromApi(ctx context.Context, data *AuditnslogparamsResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Auditnslogparams.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditnslogparams, got error: %s", err))
		return
	}

	auditnslogparamsSetAttrFromGet(ctx, data, getResponseData)

}
