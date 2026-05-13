package sslrsakey

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
var _ resource.Resource = &SslrsakeyResource{}
var _ resource.ResourceWithConfigure = (*SslrsakeyResource)(nil)
var _ resource.ResourceWithImportState = (*SslrsakeyResource)(nil)

func NewSslrsakeyResource() resource.Resource {
	return &SslrsakeyResource{}
}

// SslrsakeyResource defines the resource implementation.
type SslrsakeyResource struct {
	client *service.NitroClient
}

func (r *SslrsakeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslrsakeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslrsakey"
}

func (r *SslrsakeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslrsakeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslrsakeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslrsakey resource")
	// Get payload from plan (regular attributes)
	sslrsakey := sslrsakeyGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslrsakeyGetThePayloadFromtheConfig(ctx, &config, &sslrsakey)

	// Make API call
	// Singleton resource - use ActOnResource
	err := r.client.ActOnResource(service.Sslrsakey.Type(), &sslrsakey, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslrsakey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslrsakey resource")

	// Set ID for the resource
	data.Id = types.StringValue("sslrsakey-config")

	// Resolve unknown Optional+Computed attributes to null since Read is a noop
	if data.Aes256.IsUnknown() {
		data.Aes256 = types.BoolNull()
	}
	if data.Bits.IsUnknown() {
		data.Bits = types.Int64Null()
	}
	if data.Des.IsUnknown() {
		data.Des = types.BoolNull()
	}
	if data.Des3.IsUnknown() {
		data.Des3 = types.BoolNull()
	}
	if data.Exponent.IsUnknown() {
		data.Exponent = types.StringNull()
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

func (r *SslrsakeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslrsakeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslrsakey resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslrsakeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslrsakeyResourceModel

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

	tflog.Debug(ctx, "Updating sslrsakey resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslrsakeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslrsakeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslrsakey resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed sslrsakey from Terraform state")
}

