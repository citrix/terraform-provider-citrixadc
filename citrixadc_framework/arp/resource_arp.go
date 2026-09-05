package arp

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ArpResource{}
var _ resource.ResourceWithConfigure = (*ArpResource)(nil)
var _ resource.ResourceWithImportState = (*ArpResource)(nil)

func NewArpResource() resource.Resource {
	return &ArpResource{}
}

// ArpResource defines the resource implementation.
type ArpResource struct {
	client *service.NitroClient
}

func (r *ArpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ArpResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_arp"
}

func (r *ArpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ArpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ArpResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating arp resource")

	// Build the NITRO add payload from the plan
	arp := arpGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource keyed on ipaddress
	arpName := data.Ipaddress.ValueString()
	_, err := r.client.AddResource(service.Arp.Type(), arpName, &arp)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create arp, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created arp resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(arpName))
	data.Id = types.StringValue(arpName)

	// Read the updated state back
	if !r.readArpFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "arp not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ArpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ArpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading arp resource")

	found := r.readArpFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ArpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ArpResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating arp resource")

	// arp has no NITRO-updatable attributes (all configurable fields are
	// ForceNew/RequiresReplace); there is nothing to push to the appliance.
	// Simply refresh state from the API.
	if !r.readArpFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "arp not found immediately after update")
		}
		return
	}

	tflog.Trace(ctx, "Updated arp resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ArpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ArpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting arp resource")

	arpName := data.Id.ValueString()

	// Build the delete args exactly as SDK v2 deleteArpFunc did (only non-zero /
	// true values are appended, matching the old d.GetOk semantics).
	args := make([]string, 0)
	if !data.Td.IsNull() && data.Td.ValueInt64() != 0 {
		args = append(args, fmt.Sprintf("td:%v", data.Td.ValueInt64()))
	}
	if !data.All.IsNull() && data.All.ValueBool() {
		args = append(args, fmt.Sprintf("all:%s", strconv.FormatBool(data.All.ValueBool())))
	}
	if !data.Ownernode.IsNull() && data.Ownernode.ValueInt64() != 0 {
		args = append(args, fmt.Sprintf("ownernode:%v", data.Ownernode.ValueInt64()))
	}

	var err error
	if len(args) == 0 {
		err = r.client.DeleteResource(service.Arp.Type(), arpName)
	} else {
		err = r.client.DeleteResourceWithArgs(service.Arp.Type(), arpName, args)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete arp, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted arp resource")
}

// Helper function to read arp data from API. Returns false when the resource no
// longer exists on the appliance.
func (r *ArpResource) readArpFromApi(ctx context.Context, data *ArpResourceModel, diags *diag.Diagnostics) bool {
	arpName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Arp.Type(), arpName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read arp, got error: %s", err))
		return false
	}

	arpSetAttrFromGet(ctx, data, getResponseData)

	return true
}
