package sslcacertbundle

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
var _ resource.Resource = &SslcacertbundleResource{}
var _ resource.ResourceWithConfigure = (*SslcacertbundleResource)(nil)
var _ resource.ResourceWithImportState = (*SslcacertbundleResource)(nil)

func NewSslcacertbundleResource() resource.Resource {
	return &SslcacertbundleResource{}
}

// SslcacertbundleResource defines the resource implementation.
type SslcacertbundleResource struct {
	client *service.NitroClient
}

func (r *SslcacertbundleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcacertbundleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcacertbundle"
}

func (r *SslcacertbundleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcacertbundleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcacertbundleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcacertbundle resource")
	sslcacertbundle := sslcacertbundleGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	cacertbundlename_value := data.Cacertbundlename.ValueString()
	_, err := r.client.AddResource(service.Sslcacertbundle.Type(), cacertbundlename_value, &sslcacertbundle)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcacertbundle, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcacertbundle resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Cacertbundlename.ValueString()))

	// Read the updated state back
	r.readSslcacertbundleFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcacertbundleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcacertbundleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcacertbundle resource")

	r.readSslcacertbundleFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcacertbundleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslcacertbundleResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// NITRO has no update endpoint for sslcacertbundle; all attributes are RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for sslcacertbundle; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslcacertbundleFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcacertbundleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcacertbundleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcacertbundle resource")
	// Named resource - delete using DeleteResource
	cacertbundlename_value := data.Cacertbundlename.ValueString()
	err := r.client.DeleteResource(service.Sslcacertbundle.Type(), cacertbundlename_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcacertbundle, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcacertbundle resource")
}

// Helper function to read sslcacertbundle data from API
func (r *SslcacertbundleResource) readSslcacertbundleFromApi(ctx context.Context, data *SslcacertbundleResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	cacertbundlename_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslcacertbundle.Type(), cacertbundlename_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcacertbundle, got error: %s", err))
		return
	}

	sslcacertbundleSetAttrFromGet(ctx, data, getResponseData)

}
