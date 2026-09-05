package contentinspectionpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ContentinspectionpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionpolicylabelResource)(nil)

func NewContentinspectionpolicylabelResource() resource.Resource {
	return &ContentinspectionpolicylabelResource{}
}

// ContentinspectionpolicylabelResource defines the resource implementation.
type ContentinspectionpolicylabelResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionpolicylabel"
}

func (r *ContentinspectionpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionpolicylabel resource")
	contentinspectionpolicylabel := contentinspectionpolicylabelGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	labelname_value := data.Labelname.ValueString()
	_, err := r.client.AddResource(service.Contentinspectionpolicylabel.Type(), labelname_value, &contentinspectionpolicylabel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created contentinspectionpolicylabel resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	// Read the updated state back
	r.readContentinspectionpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionpolicylabel resource")

	r.readContentinspectionpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

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

func (r *ContentinspectionpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ContentinspectionpolicylabelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Rename support: contentinspectionpolicylabel exposes NO set/update endpoint.
	// The only in-place mutation NITRO offers is the `rename` action. Every other
	// schema attribute (labelname, type, comment) uses RequiresReplace, so Terraform
	// recreates the resource on any of those changes and never reaches here for them.
	// The ONLY change that lands in Update is `newname`.
	//
	// Mirrors the SDK v2 convention (see citrixadc/resource_citrixadc_appfwpolicy.go):
	// on a newname change, POST {labelname, newname} to ?action=rename, then point
	// the resource ID at the new name so subsequent reads address the live object.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, which is tracked by the ID -
		// NOT state.Labelname. state.Labelname stays pinned to the originally
		// configured value, so on a SECOND rename it would point at the wrong (no
		// longer live) name. The live name is whatever the prior rename set the ID to
		// (== labelname before any rename, == the prior newname after one).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming contentinspectionpolicylabel from %q to %q", oldName, newName))

		renamePayload := contentinspection.Contentinspectionpolicylabel{
			Labelname: oldName,
			Newname:   newName,
		}
		if err := r.client.ActOnResource(service.Contentinspectionpolicylabel.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename contentinspectionpolicylabel, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the read
		// below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Read the current state back. SetAttrFromGet only adopts labelname from GET
	// when the model does not already carry one, so it will not clobber the user's
	// configured labelname. Capture the plan values and restore them after the read
	// to avoid an inconsistent-result / perpetual diff.
	planLabelname := data.Labelname
	planNewname := data.Newname
	r.readContentinspectionpolicylabelFromApi(ctx, &data, &resp.Diagnostics)
	data.Labelname = planLabelname
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionpolicylabel resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== labelname at create, == newname after a rename), so we must delete
	// by data.Id, NOT data.Labelname (which stays at the originally configured value
	// and would target a non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Contentinspectionpolicylabel.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete contentinspectionpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted contentinspectionpolicylabel resource")
}

// Helper function to read contentinspectionpolicylabel data from API
func (r *ContentinspectionpolicylabelResource) readContentinspectionpolicylabelFromApi(ctx context.Context, data *ContentinspectionpolicylabelResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	labelname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Contentinspectionpolicylabel.Type(), labelname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionpolicylabel, got error: %s", err))
		return
	}

	contentinspectionpolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
