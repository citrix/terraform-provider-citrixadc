package vpnglobal_intranetip6_binding

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
var _ resource.Resource = &VpnglobalIntranetip6BindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalIntranetip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalIntranetip6BindingResource)(nil)

func NewVpnglobalIntranetip6BindingResource() resource.Resource {
	return &VpnglobalIntranetip6BindingResource{}
}

// VpnglobalIntranetip6BindingResource defines the resource implementation.
type VpnglobalIntranetip6BindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalIntranetip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalIntranetip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_intranetip6_binding"
}

func (r *VpnglobalIntranetip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalIntranetip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalIntranetip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_intranetip6_binding resource")
	vpnglobal_intranetip6_binding := vpnglobal_intranetip6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_intranetip6_binding.Type(), &vpnglobal_intranetip6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_intranetip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_intranetip6_binding resource")

	// Set ID for the resource before reading state.
	// Legacy single-key plain-value ID (intranetip6) for backward compatibility
	// with SDK v2 state and resource_id_mapping.json ("intranetip6").
	data.Id = types.StringValue(data.Intranetip6.ValueString())

	// Read the updated state back
	r.readVpnglobalIntranetip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalIntranetip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalIntranetip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_intranetip6_binding resource")

	r.readVpnglobalIntranetip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalIntranetip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalIntranetip6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnglobal_intranetip6_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnglobal_intranetip6_binding := vpnglobal_intranetip6_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnglobal_intranetip6_binding.Type(), &vpnglobal_intranetip6_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_intranetip6_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnglobal_intranetip6_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnglobal_intranetip6_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnglobalIntranetip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalIntranetip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalIntranetip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_intranetip6_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name.
	// Single-key plain-value ID (intranetip6); numaddr (from state) disambiguates,
	// mirroring the SDK v2 delete behavior. URL-encode values that may contain
	// slashy/special characters (intranetip6 ranges).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"intranetip6"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}
	intranetip6 := idMap["intranetip6"]
	if intranetip6 == "" {
		intranetip6 = data.Id.ValueString()
	}

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip6:%s", url.QueryEscape(intranetip6)))
	if !data.Numaddr.IsNull() && !data.Numaddr.IsUnknown() {
		args = append(args, fmt.Sprintf("numaddr:%d", data.Numaddr.ValueInt64()))
	}

	err = r.client.DeleteResourceWithArgs(service.Vpnglobal_intranetip6_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_intranetip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_intranetip6_binding binding")
}

// Helper function to read vpnglobal_intranetip6_binding data from API
func (r *VpnglobalIntranetip6BindingResource) readVpnglobalIntranetip6BindingFromApi(ctx context.Context, data *VpnglobalIntranetip6BindingResourceModel, diags *diag.Diagnostics) {

	// Single-key resource: ID is the plain intranetip6 value (matches SDK v2 and
	// resource_id_mapping.json). ParseIdString also accepts the legacy/new forms,
	// but the filter key here is simply the ID value.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"intranetip6"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}
	intranetip6 := idMap["intranetip6"]
	if intranetip6 == "" {
		intranetip6 = data.Id.ValueString()
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_intranetip6_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_intranetip6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vpnglobal_intranetip6_binding returned empty array")
		return
	}

	// Iterate through results to find the one matching intranetip6 (SDK v2 parity)
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["intranetip6"].(string); ok && val == intranetip6 {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "vpnglobal_intranetip6_binding not found with the provided ID attributes")
		return
	}

	vpnglobal_intranetip6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
