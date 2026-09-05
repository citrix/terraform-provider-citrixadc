package dnsnaptrrec

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
var _ resource.Resource = &DnsnaptrrecResource{}
var _ resource.ResourceWithConfigure = (*DnsnaptrrecResource)(nil)
var _ resource.ResourceWithImportState = (*DnsnaptrrecResource)(nil)

func NewDnsnaptrrecResource() resource.Resource {
	return &DnsnaptrrecResource{}
}

// DnsnaptrrecResource defines the resource implementation.
type DnsnaptrrecResource struct {
	client *service.NitroClient
}

func (r *DnsnaptrrecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsnaptrrecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsnaptrrec"
}

func (r *DnsnaptrrecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsnaptrrecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsnaptrrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsnaptrrec resource")

	dnsnaptrrec := dnsnaptrrecGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource. The domain is the resource name.
	domainName := data.Domain.ValueString()
	_, err := r.client.AddResource(service.Dnsnaptrrec.Type(), domainName, &dnsnaptrrec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsnaptrrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnsnaptrrec resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(domainName)

	// Read the created state back
	if !r.readDnsnaptrrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsnaptrrec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsnaptrrecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsnaptrrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsnaptrrec resource")

	found := r.readDnsnaptrrecFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnsnaptrrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsnaptrrecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// dnsnaptrrec has no NITRO-updatable attributes - every attribute is ForceNew
	// (RequiresReplace), so any change forces recreation and Update is never reached
	// with an actual attribute change. Simply refresh state from the API.
	tflog.Debug(ctx, "Updating dnsnaptrrec resource (no updatable attributes)")

	if !r.readDnsnaptrrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsnaptrrec not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsnaptrrecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsnaptrrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsnaptrrec resource")

	domainName := data.Id.ValueString()

	// The NAPTR record is deleted by domain name, disambiguated by the internally
	// generated recordid (matches SDK v2 behavior).
	argsMap := make(map[string]string)
	if !data.Recordid.IsNull() && !data.Recordid.IsUnknown() {
		argsMap["recordid"] = fmt.Sprintf("%d", data.Recordid.ValueInt64())
	} else {
		// Fall back to a fresh GET to obtain the recordid (e.g. imported state).
		getResponseData, err := r.client.FindResource(service.Dnsnaptrrec.Type(), domainName)
		if err != nil {
			if utils.IsNotFoundError(err) {
				// Already gone - nothing to delete.
				return
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnsnaptrrec before delete, got error: %s", err))
			return
		}
		if v, ok := getResponseData["recordid"]; ok && v != nil {
			argsMap["recordid"] = fmt.Sprintf("%v", v)
		}
	}

	err := r.client.DeleteResourceWithArgsMap(service.Dnsnaptrrec.Type(), domainName, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsnaptrrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsnaptrrec resource")
}

// Helper function to read dnsnaptrrec data from API. Returns false when the
// resource no longer exists on the ADC.
func (r *DnsnaptrrecResource) readDnsnaptrrecFromApi(ctx context.Context, data *DnsnaptrrecResourceModel, diags *diag.Diagnostics) bool {
	// Named resource - find by ID (the domain value).
	domainName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Dnsnaptrrec.Type(), domainName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsnaptrrec, got error: %s", err))
		return false
	}

	dnsnaptrrecSetAttrFromGet(ctx, data, getResponseData)

	return true
}
