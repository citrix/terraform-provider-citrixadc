package sslhpkekey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslhpkekeyResource{}
var _ resource.ResourceWithConfigure = (*SslhpkekeyResource)(nil)
var _ resource.ResourceWithImportState = (*SslhpkekeyResource)(nil)

func NewSslhpkekeyResource() resource.Resource {
	return &SslhpkekeyResource{}
}

// SslhpkekeyResource defines the resource implementation.
type SslhpkekeyResource struct {
	client *service.NitroClient
}

func (r *SslhpkekeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslhpkekeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslhpkekey"
}

func (r *SslhpkekeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslhpkekeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslhpkekeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslhpkekey resource")
	sslhpkekey := sslhpkekeyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	hpkekeyname_value := data.Hpkekeyname.ValueString()
	_, err := r.client.AddResource(service.Sslhpkekey.Type(), hpkekeyname_value, &sslhpkekey)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslhpkekey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslhpkekey resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Hpkekeyname.ValueString()))

	// Read the updated state back
	r.readSslhpkekeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslhpkekeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslhpkekeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslhpkekey resource")

	r.readSslhpkekeyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the object was deleted out-of-band, remove it from state so a
	// subsequent apply re-creates it instead of erroring.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslhpkekeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslhpkekeyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslhpkekey: NITRO exposes no update endpoint
	// (only add/delete/get); every schema attribute is RequiresReplace, so
	// changes force recreation and Update is never reached for an attribute change.
	tflog.Debug(ctx, "Update is a no-op for sslhpkekey; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslhpkekeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslhpkekeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslhpkekeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslhpkekey resource")
	// Named resource - delete using DeleteResource
	hpkekeyname_value := data.Hpkekeyname.ValueString()
	err := r.client.DeleteResource(service.Sslhpkekey.Type(), hpkekeyname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslhpkekey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslhpkekey resource")
}

// Helper function to read sslhpkekey data from API
func (r *SslhpkekeyResource) readSslhpkekeyFromApi(ctx context.Context, data *SslhpkekeyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	hpkekeyname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslhpkekey.Type(), hpkekeyname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Object is gone out-of-band; signal removal via null Id.
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslhpkekey, got error: %s", err))
		return
	}

	sslhpkekeySetAttrFromGet(ctx, data, getResponseData)

}
