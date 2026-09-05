package dnsview

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnsviewResource{}
var _ resource.ResourceWithConfigure = (*DnsviewResource)(nil)
var _ resource.ResourceWithImportState = (*DnsviewResource)(nil)

func NewDnsviewResource() resource.Resource {
	return &DnsviewResource{}
}

// DnsviewResource defines the resource implementation.
type DnsviewResource struct {
	client *service.NitroClient
}

func (r *DnsviewResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsviewResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsview"
}

func (r *DnsviewResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsviewResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsviewResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsview resource")

	// Create API request body from the model
	dnsview := dnsviewGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource keyed on the primary attribute (viewname)
	viewname_value := data.Viewname.ValueString()
	_, err := r.client.AddResource(service.Dnsview.Type(), viewname_value, &dnsview)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsview, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnsview resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Viewname.ValueString()))

	// Read the updated state back
	if !r.readDnsviewFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsview not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsviewResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsviewResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsview resource")

	found := r.readDnsviewFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsviewResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsviewResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// dnsview has no updatable attributes: viewname is the primary key and
	// forces replacement, so this Update path is only reached for a no-op.
	// Mirror the SDK v2 behavior (which had no Update) by simply re-reading.
	tflog.Debug(ctx, "Updating dnsview resource (no updatable attributes; re-reading state)")

	// Read the updated state back
	if !r.readDnsviewFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsview not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsviewResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsviewResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsview resource")

	// Named resource - delete using DeleteResource keyed on viewname
	viewname_value := data.Viewname.ValueString()
	err := r.client.DeleteResource(service.Dnsview.Type(), viewname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsview, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsview resource")
}

// Helper function to read dnsview data from API.
// Returns false if the resource no longer exists on the ADC.
func (r *DnsviewResource) readDnsviewFromApi(ctx context.Context, data *DnsviewResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain viewname value
	viewname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Dnsview.Type(), viewname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsview, got error: %s", err))
		return false
	}

	dnsviewSetAttrFromGet(ctx, data, getResponseData)

	return true
}
