package lbmonitor_servicegroup_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LbmonitorServicegroupBindingResource{}
var _ resource.ResourceWithConfigure = (*LbmonitorServicegroupBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbmonitorServicegroupBindingResource)(nil)

func NewLbmonitorServicegroupBindingResource() resource.Resource {
	return &LbmonitorServicegroupBindingResource{}
}

// LbmonitorServicegroupBindingResource defines the resource implementation.
type LbmonitorServicegroupBindingResource struct {
	client *service.NitroClient
}

func (r *LbmonitorServicegroupBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbmonitorServicegroupBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmonitor_servicegroup_binding"
}

func (r *LbmonitorServicegroupBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbmonitorServicegroupBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbmonitorServicegroupBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbmonitor_servicegroup_binding resource")
	lbmonitor_servicegroup_binding := lbmonitor_servicegroup_bindingGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes the binding "add" as an HTTP PUT (unnamed resource).
	err := r.client.UpdateUnnamedResource(service.Lbmonitor_servicegroup_binding.Type(), &lbmonitor_servicegroup_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbmonitor_servicegroup_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbmonitor_servicegroup_binding resource")

	// Compose the resource ID: monitorname,servicegroupname[,servicename]
	idParts := []string{
		"monitorname:" + utils.UrlEncode(data.Monitorname.ValueString()),
		"servicegroupname:" + utils.UrlEncode(data.Servicegroupname.ValueString()),
	}
	if !data.Servicename.IsNull() {
		idParts = append(idParts, "servicename:"+utils.UrlEncode(data.Servicename.ValueString()))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// No GET endpoint on NITRO side; persist the planned values directly.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorServicegroupBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbmonitorServicegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for lbmonitor_servicegroup_binding; NITRO exposes no
	// get/get(all)/count endpoint, so prior state is preserved unchanged.
	tflog.Debug(ctx, "Read is a no-op for lbmonitor_servicegroup_binding; no GET endpoint on NITRO side")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorServicegroupBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbmonitorServicegroupBindingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for lbmonitor_servicegroup_binding; all attributes are
	// RequiresReplace and NITRO exposes no update endpoint.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for lbmonitor_servicegroup_binding; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorServicegroupBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbmonitorServicegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbmonitor_servicegroup_binding resource")

	// Recover keys from the composite ID.
	idMap, _, _ := utils.ParseIdString(data.Id.ValueString(), nil, nil)

	monitorname := data.Monitorname.ValueString()
	if v, ok := idMap["monitorname"]; ok {
		monitorname = v
	}

	args := []string{}
	servicegroupname := data.Servicegroupname.ValueString()
	if v, ok := idMap["servicegroupname"]; ok {
		servicegroupname = v
	}
	if servicegroupname != "" {
		args = append(args, "servicegroupname:"+utils.UrlEncode(servicegroupname))
	}

	servicename := data.Servicename.ValueString()
	if v, ok := idMap["servicename"]; ok {
		servicename = v
	}
	if servicename != "" {
		args = append(args, "servicename:"+utils.UrlEncode(servicename))
	}

	err := r.client.DeleteResourceWithArgs(service.Lbmonitor_servicegroup_binding.Type(), monitorname, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbmonitor_servicegroup_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Removed lbmonitor_servicegroup_binding from Terraform state")
}
