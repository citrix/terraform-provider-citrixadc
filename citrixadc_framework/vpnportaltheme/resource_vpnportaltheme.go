package vpnportaltheme

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
var _ resource.Resource = &VpnportalthemeResource{}
var _ resource.ResourceWithConfigure = (*VpnportalthemeResource)(nil)
var _ resource.ResourceWithImportState = (*VpnportalthemeResource)(nil)

func NewVpnportalthemeResource() resource.Resource {
	return &VpnportalthemeResource{}
}

// VpnportalthemeResource defines the resource implementation.
type VpnportalthemeResource struct {
	client *service.NitroClient
}

func (r *VpnportalthemeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnportalthemeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnportaltheme"
}

func (r *VpnportalthemeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnportalthemeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnportalthemeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnportaltheme resource")

	vpnportaltheme := vpnportalthemeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource (SDK v2: client.AddResource)
	vpnportalthemeName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnportaltheme.Type(), vpnportalthemeName, &vpnportaltheme)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnportaltheme, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnportaltheme resource")

	// Set ID for the resource before reading state (SDK v2: d.SetId(name))
	data.Id = types.StringValue(vpnportalthemeName)

	// Read the updated state back
	if !r.readVpnportalthemeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnportaltheme not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnportalthemeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnportalthemeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnportaltheme resource")

	found := r.readVpnportalthemeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnportalthemeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnportalthemeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for vpnportaltheme; all attributes (name, basetheme) are
	// RequiresReplace (SDK v2 ForceNew), so Terraform never reaches Update with a
	// real change and NITRO exposes no update endpoint for this resource.
	tflog.Debug(ctx, "Update is a no-op for vpnportaltheme; all attributes are RequiresReplace")

	// Read the updated state back
	if !r.readVpnportalthemeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnportaltheme not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnportalthemeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnportalthemeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnportaltheme resource")
	// Named resource - delete using DeleteResource (SDK v2: client.DeleteResource)
	vpnportalthemeName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnportaltheme.Type(), vpnportalthemeName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnportaltheme, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnportaltheme resource")
}

// Helper function to read vpnportaltheme data from API
func (r *VpnportalthemeResource) readVpnportalthemeFromApi(ctx context.Context, data *VpnportalthemeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	vpnportalthemeName := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Vpnportaltheme.Type(), vpnportalthemeName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnportaltheme, got error: %s", err))
		return false
	}

	vpnportalthemeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
