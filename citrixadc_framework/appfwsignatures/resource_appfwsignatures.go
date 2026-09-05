package appfwsignatures

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppfwsignaturesResource{}
var _ resource.ResourceWithConfigure = (*AppfwsignaturesResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwsignaturesResource)(nil)

func NewAppfwsignaturesResource() resource.Resource {
	return &AppfwsignaturesResource{}
}

// AppfwsignaturesResource defines the resource implementation.
type AppfwsignaturesResource struct {
	client *service.NitroClient
}

func (r *AppfwsignaturesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwsignaturesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwsignatures"
}

func (r *AppfwsignaturesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwsignaturesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwsignaturesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwsignatures resource")

	// appfwsignatures is imported via the "Import" action, then finalized with the
	// "update" action (name + mergedefault), mirroring the SDK v2 resource.
	appfwsignatures := appfwsignaturesGetThePayloadFromthePlan(ctx, &data)

	err := r.client.ActOnResource(service.Appfwsignatures.Type(), &appfwsignatures, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwsignatures, got error: %s", err))
		return
	}

	updateObj := appfw.Appfwsignatures{
		Name: data.Name.ValueString(),
	}
	if !data.Mergedefault.IsNull() && !data.Mergedefault.IsUnknown() {
		updateObj.Mergedefault = data.Mergedefault.ValueBool()
	}
	err = r.client.ActOnResource(service.Appfwsignatures.Type(), &updateObj, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwsignatures, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwsignatures resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(name)).
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readAppfwsignaturesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwsignatures not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwsignaturesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwsignaturesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwsignatures resource")

	found := r.readAppfwsignaturesFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwsignaturesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwsignaturesResourceModel

	// Read Terraform prior state to detect changes and preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwsignatures resource")

	// Base import payload always carries the key + src + overwrite, mirroring SDK v2.
	importPayload := appfw.Appfwsignatures{
		Name: data.Name.ValueString(),
		Src:  data.Src.ValueString(),
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		importPayload.Overwrite = data.Overwrite.ValueBool()
	}

	updateObj := appfw.Appfwsignatures{
		Name: data.Name.ValueString(),
	}
	if !data.Mergedefault.IsNull() && !data.Mergedefault.IsUnknown() {
		updateObj.Mergedefault = data.Mergedefault.ValueBool()
	}

	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		importPayload.Comment = data.Comment.ValueString()
		hasChange = true
	}
	if !data.Merge.Equal(state.Merge) {
		importPayload.Merge = data.Merge.ValueBool()
		hasChange = true
	}
	if !data.Mergedefault.Equal(state.Mergedefault) {
		importPayload.Mergedefault = data.Mergedefault.ValueBool()
		hasChange = true
	}
	if !data.Preservedefactions.Equal(state.Preservedefactions) {
		importPayload.Preservedefactions = data.Preservedefactions.ValueBool()
		hasChange = true
	}
	if !data.Sha1.Equal(state.Sha1) {
		importPayload.Sha1 = data.Sha1.ValueString()
		hasChange = true
	}
	if !data.Vendortype.Equal(state.Vendortype) {
		importPayload.Vendortype = data.Vendortype.ValueString()
		hasChange = true
	}
	if !data.Xslt.Equal(state.Xslt) {
		importPayload.Xslt = data.Xslt.ValueString()
		hasChange = true
	}
	if !data.Autoenablenewsignatures.Equal(state.Autoenablenewsignatures) {
		importPayload.Autoenablenewsignatures = data.Autoenablenewsignatures.ValueString()
		hasChange = true
	}
	if !data.Ruleid.Equal(state.Ruleid) {
		var ruleidList []int
		data.Ruleid.ElementsAs(ctx, &ruleidList, false)
		importPayload.Ruleid = ruleidList
		hasChange = true
	}
	if !data.Category.Equal(state.Category) {
		importPayload.Category = data.Category.ValueString()
		hasChange = true
	}
	if !data.Enabled.Equal(state.Enabled) {
		importPayload.Enabled = data.Enabled.ValueString()
		hasChange = true
	}
	if !data.Action.Equal(state.Action) {
		var actionList []string
		data.Action.ElementsAs(ctx, &actionList, false)
		importPayload.Action = actionList
		hasChange = true
	}

	if hasChange {
		err := r.client.ActOnResource(service.Appfwsignatures.Type(), &importPayload, "Import")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwsignatures, got error: %s", err))
			return
		}

		err = r.client.ActOnResource(service.Appfwsignatures.Type(), &updateObj, "update")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwsignatures, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwsignatures resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwsignatures resource, skipping update")
	}

	// Read the updated state back
	if !r.readAppfwsignaturesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwsignatures not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwsignaturesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwsignaturesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwsignatures resource")
	// Named resource - delete using DeleteResource (matches SDK v2).
	appfwsignaturesName := data.Name.ValueString()
	err := r.client.DeleteResource(service.Appfwsignatures.Type(), appfwsignaturesName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwsignatures, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwsignatures resource")
}

// Helper function to read appfwsignatures data from API
func (r *AppfwsignaturesResource) readAppfwsignaturesFromApi(ctx context.Context, data *AppfwsignaturesResourceModel, diags *diag.Diagnostics) bool {

	// Single ID attribute - ID is the plain name value.
	appfwsignaturesName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appfwsignatures.Type(), appfwsignaturesName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwsignatures, got error: %s", err))
		return false
	}

	appfwsignaturesSetAttrFromGet(ctx, data, getResponseData)

	return true
}
