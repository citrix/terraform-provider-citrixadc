package dnspolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnspolicylabelResource{}
var _ resource.ResourceWithConfigure = (*DnspolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*DnspolicylabelResource)(nil)

func NewDnspolicylabelResource() resource.Resource {
	return &DnspolicylabelResource{}
}

// DnspolicylabelResource defines the resource implementation.
type DnspolicylabelResource struct {
	client *service.NitroClient
}

func (r *DnspolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnspolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnspolicylabel"
}

func (r *DnspolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnspolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnspolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnspolicylabel resource")
	dnspolicylabel := dnspolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	labelname_value := data.Labelname.ValueString()
	_, err := r.client.AddResource(service.Dnspolicylabel.Type(), labelname_value, &dnspolicylabel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnspolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnspolicylabel resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	// Read the updated state back
	r.readDnspolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnspolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnspolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnspolicylabel resource")

	r.readDnspolicylabelFromApi(ctx, &data, &resp.Diagnostics)

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

func (r *DnspolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnspolicylabelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Rename support: every user-facing schema attribute (labelname, transform)
	// uses RequiresReplace to preserve the SDK v2 ForceNew contract, so Terraform
	// recreates the resource on any of those changes and never reaches here for them.
	// The ONLY change that lands in Update is `newname`, which drives an in-place
	// rename via the NITRO ?action=rename endpoint.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Labelname (which stays pinned to the originally configured value and
		// would point at a stale name on a second rename). The live name is whatever
		// the prior rename set the ID to (== labelname before any rename).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming dnspolicylabel from %q to %q", oldName, newName))

		renamePayload := dns.Dnspolicylabel{
			Labelname: oldName,
			Newname:   newName,
		}
		if err := r.client.ActOnResource(service.Dnspolicylabel.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename dnspolicylabel, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the read below
		// (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Read the current state back. Capture the plan's labelname/newname and restore
	// them after the read so a rename does not clobber the user-facing values or
	// trigger an inconsistent-result / perpetual diff.
	planLabelname := data.Labelname
	planNewname := data.Newname
	r.readDnspolicylabelFromApi(ctx, &data, &resp.Diagnostics)
	data.Labelname = planLabelname
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnspolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnspolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnspolicylabel resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== labelname at create, == newname after a rename), so we delete by
	// data.Id, NOT data.Labelname (which stays at the originally configured value and
	// would target a non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dnspolicylabel.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnspolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnspolicylabel resource")
}

// Helper function to read dnspolicylabel data from API
func (r *DnspolicylabelResource) readDnspolicylabelFromApi(ctx context.Context, data *DnspolicylabelResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	labelname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Dnspolicylabel.Type(), labelname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnspolicylabel, got error: %s", err))
		return
	}

	dnspolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
