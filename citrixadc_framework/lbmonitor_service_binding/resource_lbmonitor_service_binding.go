package lbmonitor_service_binding

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
var _ resource.Resource = &LbmonitorServiceBindingResource{}
var _ resource.ResourceWithConfigure = (*LbmonitorServiceBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbmonitorServiceBindingResource)(nil)
var _ resource.ResourceWithValidateConfig = (*LbmonitorServiceBindingResource)(nil)

func NewLbmonitorServiceBindingResource() resource.Resource {
	return &LbmonitorServiceBindingResource{}
}

// LbmonitorServiceBindingResource defines the resource implementation.
type LbmonitorServiceBindingResource struct {
	client *service.NitroClient
}

func (r *LbmonitorServiceBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbmonitorServiceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmonitor_service_binding"
}

func (r *LbmonitorServiceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces that at least one bind target (servicename or
// servicegroupname) is supplied. Neither is marked mandatory by NITRO, but a
// binding with no target is meaningless and would be rejected by the appliance.
func (r *LbmonitorServiceBindingResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data LbmonitorServiceBindingResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Servicename.IsNull() && data.Servicegroupname.IsNull() {
		resp.Diagnostics.AddError(
			"Missing bind target",
			"At least one of \"servicename\" or \"servicegroupname\" must be specified for lbmonitor_service_binding.",
		)
	}
}

func (r *LbmonitorServiceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbmonitorServiceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbmonitor_service_binding resource")
	lbmonitor_service_binding := lbmonitor_service_bindingGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes the binding "add" as an HTTP PUT (unnamed resource).
	err := r.client.UpdateUnnamedResource(service.Lbmonitor_service_binding.Type(), &lbmonitor_service_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbmonitor_service_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbmonitor_service_binding resource")

	// Compose the resource ID: monitorname,servicename[,servicegroupname]
	idParts := []string{
		"monitorname:" + utils.UrlEncode(data.Monitorname.ValueString()),
	}
	if !data.Servicename.IsNull() {
		idParts = append(idParts, "servicename:"+utils.UrlEncode(data.Servicename.ValueString()))
	}
	if !data.Servicegroupname.IsNull() {
		idParts = append(idParts, "servicegroupname:"+utils.UrlEncode(data.Servicegroupname.ValueString()))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// No GET endpoint on NITRO side; persist the planned values directly.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorServiceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbmonitorServiceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for lbmonitor_service_binding; NITRO exposes no
	// get/get(all)/count endpoint, so prior state is preserved unchanged.
	tflog.Debug(ctx, "Read is a no-op for lbmonitor_service_binding; no GET endpoint on NITRO side")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorServiceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbmonitorServiceBindingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for lbmonitor_service_binding; all attributes are
	// RequiresReplace and NITRO exposes no update endpoint.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for lbmonitor_service_binding; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorServiceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbmonitorServiceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbmonitor_service_binding resource")

	// Recover keys from the composite ID.
	idMap, _, _ := utils.ParseIdString(data.Id.ValueString(), nil, nil)

	monitorname := data.Monitorname.ValueString()
	if v, ok := idMap["monitorname"]; ok {
		monitorname = v
	}

	args := []string{}
	servicename := data.Servicename.ValueString()
	if v, ok := idMap["servicename"]; ok {
		servicename = v
	}
	if servicename != "" {
		args = append(args, "servicename:"+utils.UrlEncode(servicename))
	}

	servicegroupname := data.Servicegroupname.ValueString()
	if v, ok := idMap["servicegroupname"]; ok {
		servicegroupname = v
	}
	if servicegroupname != "" {
		args = append(args, "servicegroupname:"+utils.UrlEncode(servicegroupname))
	}

	err := r.client.DeleteResourceWithArgs(service.Lbmonitor_service_binding.Type(), monitorname, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbmonitor_service_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Removed lbmonitor_service_binding from Terraform state")
}
