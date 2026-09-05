package l3param

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
var _ resource.Resource = &L3paramResource{}
var _ resource.ResourceWithConfigure = (*L3paramResource)(nil)
var _ resource.ResourceWithImportState = (*L3paramResource)(nil)

func NewL3paramResource() resource.Resource {
	return &L3paramResource{}
}

// L3paramResource defines the resource implementation.
type L3paramResource struct {
	client *service.NitroClient
}

func (r *L3paramResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *L3paramResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l3param"
}

func (r *L3paramResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *L3paramResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data L3paramResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating l3param resource")

	// Build the payload from the plan (singleton — only known/configured values).
	l3param := l3paramGetThePayloadFromtheConfig(ctx, &data)

	// Make API call. l3param is a singleton (no name) — use UpdateUnnamedResource,
	// matching the SDK v2 create semantics.
	err := r.client.UpdateUnnamedResource(service.L3param.Type(), &l3param)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create l3param, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("l3param-config")

	tflog.Trace(ctx, "Created l3param resource")

	// Read the updated state back
	r.readL3paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L3paramResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data L3paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading l3param resource")

	r.readL3paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L3paramResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state L3paramResourceModel

	// Read Terraform prior state, plan, and config into the model.
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating l3param resource")

	// Preserve the ID from prior state (singleton static ID).
	data.Id = types.StringValue("l3param-config")

	// Determine which attributes changed and, for those removed from config
	// (config value is null), collect them for an unset so the appliance reverts
	// them to their NITRO defaults.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Acllogtime.Equal(state.Acllogtime) {
		if config.Acllogtime.IsNull() {
			attributesToUnset = append(attributesToUnset, "acllogtime")
		} else {
			hasChange = true
		}
	}
	if !data.Allowclasseipv4.Equal(state.Allowclasseipv4) {
		if config.Allowclasseipv4.IsNull() {
			attributesToUnset = append(attributesToUnset, "allowclasseipv4")
		} else {
			hasChange = true
		}
	}
	if !data.Dropdfflag.Equal(state.Dropdfflag) {
		if config.Dropdfflag.IsNull() {
			attributesToUnset = append(attributesToUnset, "dropdfflag")
		} else {
			hasChange = true
		}
	}
	if !data.Dropipfragments.Equal(state.Dropipfragments) {
		if config.Dropipfragments.IsNull() {
			attributesToUnset = append(attributesToUnset, "dropipfragments")
		} else {
			hasChange = true
		}
	}
	if !data.Dynamicrouting.Equal(state.Dynamicrouting) {
		if config.Dynamicrouting.IsNull() {
			attributesToUnset = append(attributesToUnset, "dynamicrouting")
		} else {
			hasChange = true
		}
	}
	if !data.Externalloopback.Equal(state.Externalloopback) {
		if config.Externalloopback.IsNull() {
			attributesToUnset = append(attributesToUnset, "externalloopback")
		} else {
			hasChange = true
		}
	}
	if !data.Forwardicmpfragments.Equal(state.Forwardicmpfragments) {
		if config.Forwardicmpfragments.IsNull() {
			attributesToUnset = append(attributesToUnset, "forwardicmpfragments")
		} else {
			hasChange = true
		}
	}
	if !data.Icmpgenratethreshold.Equal(state.Icmpgenratethreshold) {
		if config.Icmpgenratethreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "icmpgenratethreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Implicitaclallow.Equal(state.Implicitaclallow) {
		if config.Implicitaclallow.IsNull() {
			attributesToUnset = append(attributesToUnset, "implicitaclallow")
		} else {
			hasChange = true
		}
	}
	if !data.Implicitpbr.Equal(state.Implicitpbr) {
		if config.Implicitpbr.IsNull() {
			attributesToUnset = append(attributesToUnset, "implicitpbr")
		} else {
			hasChange = true
		}
	}
	if !data.Ipv6dynamicrouting.Equal(state.Ipv6dynamicrouting) {
		if config.Ipv6dynamicrouting.IsNull() {
			attributesToUnset = append(attributesToUnset, "ipv6dynamicrouting")
		} else {
			hasChange = true
		}
	}
	if !data.Miproundrobin.Equal(state.Miproundrobin) {
		if config.Miproundrobin.IsNull() {
			attributesToUnset = append(attributesToUnset, "miproundrobin")
		} else {
			hasChange = true
		}
	}
	if !data.Overridernat.Equal(state.Overridernat) {
		if config.Overridernat.IsNull() {
			attributesToUnset = append(attributesToUnset, "overridernat")
		} else {
			hasChange = true
		}
	}
	if !data.Srcnat.Equal(state.Srcnat) {
		if config.Srcnat.IsNull() {
			attributesToUnset = append(attributesToUnset, "srcnat")
		} else {
			hasChange = true
		}
	}
	if !data.Tnlpmtuwoconn.Equal(state.Tnlpmtuwoconn) {
		if config.Tnlpmtuwoconn.IsNull() {
			attributesToUnset = append(attributesToUnset, "tnlpmtuwoconn")
		} else {
			hasChange = true
		}
	}
	if !data.Usipserverstraypkt.Equal(state.Usipserverstraypkt) {
		if config.Usipserverstraypkt.IsNull() {
			attributesToUnset = append(attributesToUnset, "usipserverstraypkt")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model (only known/configured values).
		l3param := l3paramGetThePayloadFromtheConfig(ctx, &data)

		// Make API call — singleton update via UpdateUnnamedResource, matching SDK v2.
		err := r.client.UpdateUnnamedResource(service.L3param.Type(), &l3param)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update l3param, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated l3param resource")
	} else {
		tflog.Debug(ctx, "No l3param attribute changes require an update")
	}

	// Unset attributes removed from config so the appliance reverts them to their
	// NITRO defaults. l3param is a singleton — no identifying key in the payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.L3param.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset l3param attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readL3paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L3paramResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data L3paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting l3param resource")

	// For l3param, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted l3param resource from state")
}

// Helper function to read l3param data from API
func (r *L3paramResource) readL3paramFromApi(ctx context.Context, data *L3paramResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.L3param.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read l3param, got error: %s", err))
		return
	}

	l3paramSetAttrFromGet(ctx, data, getResponseData)

}
