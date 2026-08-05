package contentinspectionpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ContentinspectionpolicyResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionpolicyResource)(nil)

func NewContentinspectionpolicyResource() resource.Resource {
	return &ContentinspectionpolicyResource{}
}

// ContentinspectionpolicyResource defines the resource implementation.
type ContentinspectionpolicyResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionpolicy"
}

func (r *ContentinspectionpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionpolicy resource")

	contentinspectionpolicy := contentinspectionpolicyGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Contentinspectionpolicy.Type(), name_value, &contentinspectionpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created contentinspectionpolicy resource")

	// Set ID for the resource before reading state (single unique attr -> plain value)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readContentinspectionpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionpolicy resource")

	r.readContentinspectionpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Resource no longer exists on the ADC - remove from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ContentinspectionpolicyResourceModel

	// Read Terraform prior state to preserve ID / detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating contentinspectionpolicy resource")

	// Rename support: contentinspectionpolicy exposes a `rename` action (NITRO
	// ?action=rename) plus a `newname` attribute. `name` uses RequiresReplace, so a
	// change to the key itself recreates the resource (matching SDK v2 ForceNew) and
	// never reaches Update. The ONLY key mutation that lands here is `newname`.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID (== name
		// before any rename, == the prior newname after one). state.Name stays pinned
		// to the originally-configured value and is stale after a rename.
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming contentinspectionpolicy from %q to %q", oldName, newName))

		renamePayload := contentinspection.Contentinspectionpolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Contentinspectionpolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename contentinspectionpolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so subsequent
		// update/read calls address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Regular attribute update (rule, action, comment, logaction, undefaction) via
	// the NITRO PUT (unnamed) endpoint, mirroring the SDK v2 UpdateUnnamedResource.
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for contentinspectionpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for contentinspectionpolicy")
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for contentinspectionpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for contentinspectionpolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for contentinspectionpolicy")
		hasChange = true
	}

	if hasChange {
		contentinspectionpolicy := contentinspectionpolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Target the live object name (post-rename aware): the PUT identifies the
		// policy by the name in the body.
		contentinspectionpolicy.Name = data.Id.ValueString()
		err := r.client.UpdateUnnamedResource(service.Contentinspectionpolicy.Type(), &contentinspectionpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated contentinspectionpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for contentinspectionpolicy resource, skipping update")
	}

	// Read the updated state back. Capture the user-facing name/newname before the
	// read and restore them afterward so a post-rename GET (which returns the live
	// new name) cannot clobber the configured values / trigger a spurious diff.
	planName := data.Name
	planNewname := data.Newname
	r.readContentinspectionpolicyFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionpolicy resource")

	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id, NOT
	// data.Name (which is stale after a rename and would dangle the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Contentinspectionpolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete contentinspectionpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted contentinspectionpolicy resource")
}

// Helper function to read contentinspectionpolicy data from API
func (r *ContentinspectionpolicyResource) readContentinspectionpolicyFromApi(ctx context.Context, data *ContentinspectionpolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain (live) name
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Contentinspectionpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Signal to callers that the resource no longer exists.
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionpolicy, got error: %s", err))
		return
	}

	contentinspectionpolicySetAttrFromGet(ctx, data, getResponseData)
}
