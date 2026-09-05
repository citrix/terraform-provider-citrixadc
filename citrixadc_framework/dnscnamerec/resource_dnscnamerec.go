package dnscnamerec

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnscnamerecResource{}
var _ resource.ResourceWithConfigure = (*DnscnamerecResource)(nil)
var _ resource.ResourceWithImportState = (*DnscnamerecResource)(nil)

func NewDnscnamerecResource() resource.Resource {
	return &DnscnamerecResource{}
}

// DnscnamerecResource defines the resource implementation.
type DnscnamerecResource struct {
	client *service.NitroClient
}

func (r *DnscnamerecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnscnamerecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnscnamerec"
}

func (r *DnscnamerecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnscnamerecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnscnamerecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnscnamerec resource")

	dnscnamerec := dnscnamerecGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource keyed on the primary attribute (aliasname).
	aliasname_value := data.Aliasname.ValueString()
	_, err := r.client.AddResource(service.Dnscnamerec.Type(), aliasname_value, &dnscnamerec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnscnamerec, got error: %s", err))
		return
	}

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(aliasname))
	data.Id = types.StringValue(fmt.Sprintf("%v", aliasname_value))

	tflog.Trace(ctx, "Created dnscnamerec resource")

	// Read the updated state back
	if !r.readDnscnamerecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnscnamerec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnscnamerecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnscnamerecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnscnamerec resource")

	found := r.readDnscnamerecFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnscnamerecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnscnamerecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// dnscnamerec has no NITRO-updatable attributes: every schema attribute maps to an
	// SDK v2 ForceNew field and carries a RequiresReplace plan modifier, so any change is
	// handled by Terraform via destroy+create and never reaches Update. Nothing to push;
	// just re-read to keep state consistent.
	tflog.Debug(ctx, "Updating dnscnamerec resource (no updatable attributes; re-reading state)")

	if !r.readDnscnamerecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnscnamerec not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnscnamerecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnscnamerecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnscnamerec resource")

	// Named resource - delete keyed on aliasname. If ecssubnet is set, it must be passed
	// as a delete arg (mirrors SDK v2 deleteDnscnamerecFunc).
	aliasname_value := data.Aliasname.ValueString()
	var err error
	if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() && data.Ecssubnet.ValueString() != "" {
		argsMap := make(map[string]string)
		argsMap["ecssubnet"] = url.QueryEscape(data.Ecssubnet.ValueString())
		err = r.client.DeleteResourceWithArgsMap(service.Dnscnamerec.Type(), aliasname_value, argsMap)
	} else {
		err = r.client.DeleteResource(service.Dnscnamerec.Type(), aliasname_value)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnscnamerec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnscnamerec resource")
}

// Helper function to read dnscnamerec data from API. Returns false when the resource no
// longer exists on the ADC so callers can remove it from state.
func (r *DnscnamerecResource) readDnscnamerecFromApi(ctx context.Context, data *DnscnamerecResourceModel, diags *diag.Diagnostics) bool {

	// Single unique attribute - ID is the plain aliasname value.
	aliasname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Dnscnamerec.Type(), aliasname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnscnamerec, got error: %s", err))
		return false
	}

	dnscnamerecSetAttrFromGet(ctx, data, getResponseData)

	return true
}
