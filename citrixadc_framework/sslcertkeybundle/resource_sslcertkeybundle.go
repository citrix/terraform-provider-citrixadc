package sslcertkeybundle

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
var _ resource.Resource = &SslcertkeybundleResource{}
var _ resource.ResourceWithConfigure = (*SslcertkeybundleResource)(nil)
var _ resource.ResourceWithImportState = (*SslcertkeybundleResource)(nil)

func NewSslcertkeybundleResource() resource.Resource {
	return &SslcertkeybundleResource{}
}

// SslcertkeybundleResource defines the resource implementation.
type SslcertkeybundleResource struct {
	client *service.NitroClient
}

func (r *SslcertkeybundleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcertkeybundleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertkeybundle"
}

func (r *SslcertkeybundleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertkeybundleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslcertkeybundleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertkeybundle resource")
	// Get payload from plan (regular attributes)
	sslcertkeybundle := sslcertkeybundleGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslcertkeybundleGetThePayloadFromtheConfig(ctx, &config, &sslcertkeybundle)

	// Make API call
	// Named resource - use AddResource
	certkeybundlename_value := data.Certkeybundlename.ValueString()
	_, err := r.client.AddResource(service.Sslcertkeybundle.Type(), certkeybundlename_value, &sslcertkeybundle)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcertkeybundle, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcertkeybundle resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Certkeybundlename.ValueString()))

	// Read the updated state back
	r.readSslcertkeybundleFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeybundleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcertkeybundleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcertkeybundle resource")

	r.readSslcertkeybundleFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeybundleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslcertkeybundleResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslcertkeybundle resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Bundlefile.Equal(state.Bundlefile) {
		tflog.Debug(ctx, fmt.Sprintf("bundlefile has changed for sslcertkeybundle"))
		hasChange = true
	}
	// Check secret attribute passplain or its version tracker
	if !data.Passplain.Equal(state.Passplain) {
		tflog.Debug(ctx, fmt.Sprintf("passplain has changed for sslcertkeybundle"))
		hasChange = true
	} else if !data.PassplainWoVersion.Equal(state.PassplainWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("passplain_wo_version has changed for sslcertkeybundle"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		sslcertkeybundle := sslcertkeybundleGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		sslcertkeybundleGetThePayloadFromtheConfig(ctx, &config, &sslcertkeybundle)
		// Make API call
		// NITRO update is the "change" action (?action=update, POST), not a plain PUT.
		err := r.client.ActOnResource(service.Sslcertkeybundle.Type(), &sslcertkeybundle, "update")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcertkeybundle, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslcertkeybundle resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslcertkeybundle resource, skipping update")
	}

	// Read the updated state back
	r.readSslcertkeybundleFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeybundleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcertkeybundleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcertkeybundle resource")
	// Named resource - delete using DeleteResource
	certkeybundlename_value := data.Certkeybundlename.ValueString()
	err := r.client.DeleteResource(service.Sslcertkeybundle.Type(), certkeybundlename_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcertkeybundle, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcertkeybundle resource")
}

// Helper function to read sslcertkeybundle data from API
func (r *SslcertkeybundleResource) readSslcertkeybundleFromApi(ctx context.Context, data *SslcertkeybundleResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	certkeybundlename_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslcertkeybundle.Type(), certkeybundlename_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcertkeybundle, got error: %s", err))
		return
	}

	sslcertkeybundleSetAttrFromGet(ctx, data, getResponseData)

}
