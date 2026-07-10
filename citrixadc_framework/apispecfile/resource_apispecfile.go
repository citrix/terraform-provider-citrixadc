package apispecfile

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
var _ resource.Resource = &ApispecfileResource{}
var _ resource.ResourceWithConfigure = (*ApispecfileResource)(nil)
var _ resource.ResourceWithImportState = (*ApispecfileResource)(nil)

func NewApispecfileResource() resource.Resource {
	return &ApispecfileResource{}
}

// ApispecfileResource defines the resource implementation.
type ApispecfileResource struct {
	client *service.NitroClient
}

func (r *ApispecfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ApispecfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apispecfile"
}

func (r *ApispecfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ApispecfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApispecfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating apispecfile resource")
	apispecfile := apispecfileGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes apispecfile create only via POST ?action=Import (no `add`).
	// Use ActOnResource with the case-sensitive "Import" verb.
	err := r.client.ActOnResource(service.Apispecfile.Type(), &apispecfile, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create apispecfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created apispecfile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readApispecfileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "apispecfile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApispecfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApispecfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading apispecfile resource")

	found := r.readApispecfileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ApispecfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for apispecfile (only Import, delete, get).
	// Every schema attribute is marked RequiresReplace, so Terraform will never
	// actually invoke Update with field changes — any plan change forces destroy
	// + recreate. This body is therefore a documented no-op that preserves the
	// prior ID and re-reads state for consistency.
	var data, state ApispecfileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for apispecfile; NITRO has no update endpoint and all attributes are RequiresReplace")

	if !r.readApispecfileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "apispecfile not found immediately after update")
		}
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApispecfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApispecfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting apispecfile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Apispecfile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete apispecfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted apispecfile resource")
}

// Helper function to read apispecfile data from API
func (r *ApispecfileResource) readApispecfileFromApi(ctx context.Context, data *ApispecfileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Apispecfile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read apispecfile, got error: %s", err))
		return false
	}

	apispecfileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
