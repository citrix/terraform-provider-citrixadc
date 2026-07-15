package gslbservicegroup

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GslbservicegroupResource{}
var _ resource.ResourceWithConfigure = (*GslbservicegroupResource)(nil)
var _ resource.ResourceWithImportState = (*GslbservicegroupResource)(nil)

func NewGslbservicegroupResource() resource.Resource {
	return &GslbservicegroupResource{}
}

// GslbservicegroupResource defines the resource implementation.
type GslbservicegroupResource struct {
	client *service.NitroClient
}

func (r *GslbservicegroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbservicegroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservicegroup"
}

func (r *GslbservicegroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbservicegroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbservicegroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservicegroup resource")
	gslbservicegroup := gslbservicegroupGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (POST add)
	servicegroupname_value := data.Servicegroupname.ValueString()
	_, err := r.client.AddResource(service.Gslbservicegroup.Type(), servicegroupname_value, &gslbservicegroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservicegroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created gslbservicegroup resource")

	// Set ID for the resource before reading state (plain servicegroupname value).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))

	// Read the updated state back
	r.readGslbservicegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbservicegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservicegroup resource")

	r.readGslbservicegroupFromApi(ctx, &data, &resp.Diagnostics)

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

func (r *GslbservicegroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state GslbservicegroupResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the live name).
	data.Id = state.Id

	tflog.Debug(ctx, "Updating gslbservicegroup resource")

	// Rename support: if newname changed and is set, perform an in-place rename via
	// the NITRO ?action=rename endpoint. The rename SOURCE must be the CURRENT LIVE
	// name, tracked by state.Id (NOT state.Servicegroupname, which stays pinned to the
	// originally configured value and would be stale on a second rename).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming gslbservicegroup from %q to %q", oldName, newName))

		renamePayload := gslb.Gslbservicegroup{
			Servicegroupname: oldName,
			Newname:          newName,
		}
		if err := r.client.ActOnResource(service.Gslbservicegroup.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename gslbservicegroup, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update and
		// read below address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Apply the updateable-attribute changes via PUT (set). The update payload
	// (Pattern 9) excludes the create-only attrs (servicetype, autoscale,
	// autodelayedtrofs, state) and the non-write attrs (newname, includemembers,
	// delay, graceful).
	updatePayload := gslbservicegroupGetTheUpdatePayloadFromthePlan(ctx, &data)
	// Address the resource by its CURRENT LIVE name (data.Id), which reflects any
	// rename that just happened.
	updatePayload.Servicegroupname = data.Id.ValueString()
	_, err := r.client.UpdateResource(service.Gslbservicegroup.Type(), data.Id.ValueString(), &updatePayload)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbservicegroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated gslbservicegroup resource")

	// Read the current state back. SetAttrFromGet preserves the configured
	// servicegroupname and newname; restore the plan values (belt-and-suspenders)
	// so a rename does not surface as a spurious diff.
	planServicegroupname := data.Servicegroupname
	planNewname := data.Newname
	r.readGslbservicegroupFromApi(ctx, &data, &resp.Diagnostics)
	data.Servicegroupname = planServicegroupname
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbservicegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservicegroup resource")
	// Named resource - delete by the CURRENT LIVE name (data.Id), which reflects any
	// rename. Deleting by data.Servicegroupname would target a stale name after a
	// rename and leave the renamed object dangling.
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Gslbservicegroup.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete gslbservicegroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted gslbservicegroup resource")
}

// Helper function to read gslbservicegroup data from API
func (r *GslbservicegroupResource) readGslbservicegroupFromApi(ctx context.Context, data *GslbservicegroupResourceModel, diags *diag.Diagnostics) {

	// Named resource - read by the live name held in data.Id.
	servicegroupname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Gslbservicegroup.Type(), servicegroupname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup, got error: %s", err))
		return
	}

	gslbservicegroupSetAttrFromGet(ctx, data, getResponseData)
}
