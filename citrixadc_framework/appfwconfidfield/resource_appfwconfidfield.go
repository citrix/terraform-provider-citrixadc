package appfwconfidfield

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppfwconfidfieldResource{}
var _ resource.ResourceWithConfigure = (*AppfwconfidfieldResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwconfidfieldResource)(nil)

func NewAppfwconfidfieldResource() resource.Resource {
	return &AppfwconfidfieldResource{}
}

// AppfwconfidfieldResource defines the resource implementation.
type AppfwconfidfieldResource struct {
	client *service.NitroClient
}

func (r *AppfwconfidfieldResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwconfidfieldResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwconfidfield"
}

func (r *AppfwconfidfieldResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwconfidfieldResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwconfidfieldResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwconfidfield resource")

	// Get payload from plan
	appfwconfidfield := appfwconfidfieldGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	fieldname_value := data.Fieldname.ValueString()
	_, err := r.client.AddResource(service.Appfwconfidfield.Type(), fieldname_value, &appfwconfidfield)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwconfidfield, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwconfidfield resource")

	// Set ID for the resource before reading state
	// Composite ID matches SDK v2 d.SetId: "fieldname,url"
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Fieldname.ValueString(), data.Url.ValueString()))

	// Read the updated state back
	if !r.readAppfwconfidfieldFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwconfidfield not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwconfidfieldResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwconfidfieldResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwconfidfield resource")

	found := r.readAppfwconfidfieldFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwconfidfieldResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwconfidfieldResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwconfidfield resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for appfwconfidfield")
		hasChange = true
	}
	if !data.Isregex.Equal(state.Isregex) {
		tflog.Debug(ctx, "isregex has changed for appfwconfidfield")
		hasChange = true
	}
	if !data.State.Equal(state.State) {
		tflog.Debug(ctx, "state has changed for appfwconfidfield")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		appfwconfidfield := appfwconfidfieldGetThePayloadFromthePlan(ctx, &data)
		// Unnamed update - the composite key (fieldname,url) is carried in the payload
		err := r.client.UpdateUnnamedResource(service.Appfwconfidfield.Type(), &appfwconfidfield)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwconfidfield, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwconfidfield resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwconfidfield resource, skipping update")
	}

	// Read the updated state back
	if !r.readAppfwconfidfieldFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwconfidfield not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwconfidfieldResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwconfidfieldResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwconfidfield resource")

	// Only the fieldname and url properties are required for deletion
	argsMap := make(map[string]string)
	argsMap["fieldname"] = url.QueryEscape(data.Fieldname.ValueString())
	argsMap["url"] = url.QueryEscape(data.Url.ValueString())

	err := r.client.DeleteResourceWithArgsMap(service.Appfwconfidfield.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwconfidfield, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwconfidfield resource")
}

// Helper function to read appfwconfidfield data from API
func (r *AppfwconfidfieldResource) readAppfwconfidfieldFromApi(ctx context.Context, data *AppfwconfidfieldResourceModel, diags *diag.Diagnostics) bool {

	// Resolve the composite key (fieldname,url) from prior state, falling back to
	// the composite ID (needed on import where the attributes are not yet populated).
	fieldname_value := data.Fieldname.ValueString()
	url_value := data.Url.ValueString()
	if idParts := strings.SplitN(data.Id.ValueString(), ",", 2); len(idParts) == 2 {
		fieldname_value = idParts[0]
		url_value = idParts[1]
	}

	findParams := service.FindParams{
		ResourceType: service.Appfwconfidfield.Type(),
	}
	dataArray, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwconfidfield, got error: %s", err))
		return false
	}

	if len(dataArray) == 0 {
		return false
	}

	foundIndex := -1
	for i, item := range dataArray {
		if item["fieldname"] == fieldname_value && item["url"] == url_value {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return false
	}

	appfwconfidfieldSetAttrFromGet(ctx, data, dataArray[foundIndex])

	return true
}
