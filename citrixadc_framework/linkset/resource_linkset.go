package linkset

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LinksetResource{}
var _ resource.ResourceWithConfigure = (*LinksetResource)(nil)
var _ resource.ResourceWithImportState = (*LinksetResource)(nil)

func NewLinksetResource() resource.Resource {
	return &LinksetResource{}
}

// LinksetResource defines the resource implementation.
type LinksetResource struct {
	client *service.NitroClient
}

func (r *LinksetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LinksetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_linkset"
}

func (r *LinksetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LinksetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LinksetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating linkset resource")

	linkset := linksetGetThePayloadFromthePlan(ctx, &data)
	linksetName := data.Linksetid.ValueString()

	// Named resource - use AddResource (POST /config/linkset)
	_, err := r.client.AddResource(service.Linkset.Type(), "", &linkset)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create linkset, got error: %s", err))
		return
	}

	// Set ID for the resource before reading state
	data.Id = types.StringValue(linksetName)

	// Handle interfacebinding convenience block - bind the configured interfaces.
	if err := r.updateLinksetInterfaceBindings(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to bind interfaces to linkset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created linkset resource")

	// Read the updated state back
	if !r.readLinksetFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "linkset not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LinksetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading linkset resource")

	found := r.readLinksetFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LinksetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// linkset has no NITRO update endpoint and every schema attribute is
	// RequiresReplace, so Terraform never actually invokes Update with changes.
	// This documented no-op preserves the ID and re-reads live state.
	var data, state LinksetResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for linkset; all attributes are RequiresReplace")

	if !r.readLinksetFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "linkset not found during update read-back")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LinksetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting linkset resource")

	// Named resource - delete using DeleteResource (DELETE /config/linkset/<id>).
	// Deleting the linkset removes its interface bindings as well.
	linksetName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Linkset.Type(), linksetName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete linkset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted linkset resource")
}

// updateLinksetInterfaceBindings binds the interfaces configured in the
// interfacebinding set to the linkset. Called from Create; because the set is
// RequiresReplace there is never an in-place diff to reconcile.
func (r *LinksetResource) updateLinksetInterfaceBindings(ctx context.Context, data *LinksetResourceModel) error {
	tflog.Debug(ctx, "In updateLinksetInterfaceBindings")

	if data.Interfacebinding.IsNull() || data.Interfacebinding.IsUnknown() {
		return nil
	}

	var ifnums []string
	diags := data.Interfacebinding.ElementsAs(ctx, &ifnums, false)
	if diags.HasError() {
		return fmt.Errorf("unable to read interfacebinding set")
	}

	linksetName := data.Linksetid.ValueString()
	for _, ifnum := range ifnums {
		bindingStruct := network.Linksetinterfacebinding{}
		bindingStruct.Id = linksetName
		bindingStruct.Ifnum = ifnum
		// NITRO expects a PUT here, hence UpdateUnnamedResource.
		if err := r.client.UpdateUnnamedResource(service.Linkset_interface_binding.Type(), bindingStruct); err != nil {
			return err
		}
	}

	return nil
}

// Helper function to read linkset data from API
func (r *LinksetResource) readLinksetFromApi(ctx context.Context, data *LinksetResourceModel, diags *diag.Diagnostics) bool {
	linksetName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Linkset.Type(), linksetName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read linkset, got error: %s", err))
		return false
	}

	linksetSetAttrFromGet(ctx, data, getResponseData)

	// Populate the interfacebinding convenience block from the ADC.
	if err := linksetReadInterfaceBindings(ctx, r.client, data, linksetName); err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read linkset interface bindings, got error: %s", err))
		return false
	}

	return true
}
