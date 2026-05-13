package sslcert

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslcertResource{}
var _ resource.ResourceWithConfigure = (*SslcertResource)(nil)
var _ resource.ResourceWithImportState = (*SslcertResource)(nil)

func NewSslcertResource() resource.Resource {
	return &SslcertResource{}
}

// SslcertResource defines the resource implementation.
type SslcertResource struct {
	client *service.NitroClient
}

func (r *SslcertResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcertResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcert"
}

func (r *SslcertResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslcertResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcert resource")
	// Get payload from plan (regular attributes)
	sslcert := sslcertGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslcertGetThePayloadFromtheConfig(ctx, &config, &sslcert)

	// Make API call
	err := r.client.ActOnResource(service.Sslcert.Type(), &sslcert, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcert, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcert resource")

	// Set ID for the resource
	data.Id = types.StringValue("sslcert-config")

	// Resolve unknown Optional+Computed attributes to null since Read is a noop
	// (no server-side state to read back for this action-only resource)
	if data.Cacert.IsUnknown() {
		data.Cacert = types.StringNull()
	}
	if data.Cacertform.IsUnknown() {
		data.Cacertform = types.StringNull()
	}
	if data.Cakey.IsUnknown() {
		data.Cakey = types.StringNull()
	}
	if data.Cakeyform.IsUnknown() {
		data.Cakeyform = types.StringNull()
	}
	if data.Caserial.IsUnknown() {
		data.Caserial = types.StringNull()
	}
	if data.Certform.IsUnknown() {
		data.Certform = types.StringNull()
	}
	if data.Days.IsUnknown() {
		data.Days = types.Int64Null()
	}
	if data.Keyfile.IsUnknown() {
		data.Keyfile = types.StringNull()
	}
	if data.Keyform.IsUnknown() {
		data.Keyform = types.StringNull()
	}
	if data.Subjectaltname.IsUnknown() {
		data.Subjectaltname = types.StringNull()
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcertResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcert resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslcertResourceModel

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

	tflog.Debug(ctx, "Updating sslcert resource")
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcertResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcert resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed sslcert from Terraform state")
}
