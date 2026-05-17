package sslecdsakey

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
var _ resource.Resource = &SslecdsakeyResource{}
var _ resource.ResourceWithConfigure = (*SslecdsakeyResource)(nil)
var _ resource.ResourceWithImportState = (*SslecdsakeyResource)(nil)

func NewSslecdsakeyResource() resource.Resource {
	return &SslecdsakeyResource{}
}

// SslecdsakeyResource defines the resource implementation.
type SslecdsakeyResource struct {
	client *service.NitroClient
}

func (r *SslecdsakeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslecdsakeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslecdsakey"
}

func (r *SslecdsakeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslecdsakeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslecdsakeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslecdsakey resource")
	// Get payload from plan (regular attributes)
	sslecdsakey := sslecdsakeyGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslecdsakeyGetThePayloadFromtheConfig(ctx, &config, &sslecdsakey)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.ActOnResource(service.Sslecdsakey.Type(), &sslecdsakey, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslecdsakey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslecdsakey resource")

	// Set ID for the resource
	data.Id = types.StringValue("sslecdsakey-config")

	// Resolve unknown Optional+Computed attributes to null since Read is a noop
	// (no server-side state to read back for this action-only resource)
	if data.Aes256.IsUnknown() {
		data.Aes256 = types.BoolNull()
	}
	if data.Curve.IsUnknown() {
		data.Curve = types.StringNull()
	}
	if data.Des.IsUnknown() {
		data.Des = types.BoolNull()
	}
	if data.Des3.IsUnknown() {
		data.Des3 = types.BoolNull()
	}
	if data.Keyfile.IsUnknown() {
		data.Keyfile = types.StringNull()
	}
	if data.Keyform.IsUnknown() {
		data.Keyform = types.StringNull()
	}
	if data.Pkcs8.IsUnknown() {
		data.Pkcs8 = types.BoolNull()
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslecdsakeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslecdsakeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslecdsakey resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslecdsakeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslecdsakeyResourceModel

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

	tflog.Debug(ctx, "Updating sslecdsakey resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslecdsakeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslecdsakeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslecdsakey resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed sslecdsakey from Terraform state")
}
