package dnsnameserver

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnsnameserverResource{}
var _ resource.ResourceWithConfigure = (*DnsnameserverResource)(nil)
var _ resource.ResourceWithImportState = (*DnsnameserverResource)(nil)

func NewDnsnameserverResource() resource.Resource {
	return &DnsnameserverResource{}
}

// DnsnameserverResource defines the resource implementation.
type DnsnameserverResource struct {
	client *service.NitroClient
}

func (r *DnsnameserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsnameserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsnameserver"
}

func (r *DnsnameserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsnameserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsnameserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsnameserver resource")

	dnsnameserver := dnsnameserverGetThePayloadFromthePlan(ctx, &data)

	// The primary identifier is the IP address, or the DNS vserver name when no IP is given.
	var primaryId string
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() && data.Ip.ValueString() != "" {
		primaryId = data.Ip.ValueString()
	} else if !data.Dnsvservername.IsNull() && !data.Dnsvservername.IsUnknown() && data.Dnsvservername.ValueString() != "" {
		primaryId = data.Dnsvservername.ValueString()
	}

	// Named resource - use AddResource
	_, err := r.client.AddResource(service.Dnsnameserver.Type(), primaryId, &dnsnameserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsnameserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnsnameserver resource")

	// Build the composite ID "<name>,<type>" to match the SDK v2 format.
	typeForId := "UDP"
	if !data.Type.IsNull() && !data.Type.IsUnknown() && data.Type.ValueString() != "" {
		typeForId = data.Type.ValueString()
	}
	data.Id = types.StringValue(primaryId + "," + typeForId)

	// Read the updated state back
	if !r.readDnsnameserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsnameserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsnameserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsnameserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsnameserver resource")

	found := r.readDnsnameserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnsnameserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsnameserverResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnsnameserver resource")

	// Only dnsprofilename is updateable in place; all other attributes are
	// RequiresReplace (matching SDK v2 ForceNew) and therefore never reach Update.
	hasChange := false
	if !data.Dnsprofilename.Equal(state.Dnsprofilename) {
		tflog.Debug(ctx, "dnsprofilename has changed for dnsnameserver, starting update")
		hasChange = true
	}

	if hasChange {
		idSlice := strings.SplitN(data.Id.ValueString(), ",", 2)
		name := idSlice[0]

		dnsnameserver := dns.Dnsnameserver{
			Ip:             data.Ip.ValueString(),
			Dnsvservername: data.Dnsvservername.ValueString(),
			Dnsprofilename: data.Dnsprofilename.ValueString(),
		}

		_, err := r.client.UpdateResource(service.Dnsnameserver.Type(), name, &dnsnameserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsnameserver, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dnsnameserver resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnsnameserver resource, skipping update")
	}

	// Read the updated state back
	if !r.readDnsnameserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsnameserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsnameserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsnameserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsnameserver resource")

	primaryId := data.Id.ValueString()
	idSlice := strings.SplitN(primaryId, ",", 2)
	name := idSlice[0]
	dnsType := ""
	if len(idSlice) > 1 {
		dnsType = idSlice[1]
	}

	// If the resource is keyed on a DNS vserver name, delete it directly.
	if !data.Dnsvservername.IsNull() && data.Dnsvservername.ValueString() == name {
		err := r.client.DeleteResource(service.Dnsnameserver.Type(), name)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsnameserver, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Deleted dnsnameserver resource")
		return
	}

	// IP-based name server: delete keyed on ip with the protocol as a delete arg.
	// UDP_TCP is stored as two entries (UDP and TCP) and must both be removed.
	var typesToDelete []string
	if dnsType == "UDP_TCP" {
		typesToDelete = []string{"UDP", "TCP"}
	} else {
		typesToDelete = []string{dnsType}
	}

	for _, deleteType := range typesToDelete {
		argsMap := make(map[string]string)
		argsMap["type"] = url.QueryEscape(deleteType)
		err := r.client.DeleteResourceWithArgsMap(service.Dnsnameserver.Type(), name, argsMap)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsnameserver, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "Deleted dnsnameserver resource")
}

// Helper function to read dnsnameserver data from API. Returns false when the
// resource no longer exists on the ADC.
func (r *DnsnameserverResource) readDnsnameserverFromApi(ctx context.Context, data *DnsnameserverResourceModel, diags *diag.Diagnostics) bool {
	primaryId := data.Id.ValueString()

	// Backward compatibility: SDK v2 state prior to the composite-ID change stored
	// only the name. Append the protocol (config value or the UDP default) so the
	// filter below can locate the entry.
	oldIdSlice := strings.Split(primaryId, ",")
	if len(oldIdSlice) == 1 {
		typeVal := "UDP"
		if !data.Type.IsNull() && !data.Type.IsUnknown() && data.Type.ValueString() != "" {
			typeVal = data.Type.ValueString()
		}
		primaryId = primaryId + "," + typeVal
		data.Id = types.StringValue(primaryId)
	}

	idSlice := strings.SplitN(primaryId, ",", 2)
	name := idSlice[0]
	dnsType := ""
	if len(idSlice) > 1 {
		dnsType = idSlice[1]
	}

	findParams := service.FindParams{
		ResourceType: service.Dnsnameserver.Type(),
	}
	dataArray, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsnameserver, got error: %s", err))
		return false
	}
	if len(dataArray) == 0 {
		return false
	}

	// When type is UDP_TCP, the ADC creates two separate entries (UDP and TCP);
	// match either.
	var typesToCheck []string
	if dnsType == "UDP_TCP" {
		typesToCheck = []string{"UDP", "TCP"}
	} else {
		typesToCheck = []string{dnsType}
	}

	foundIndex := -1
	for _, checkType := range typesToCheck {
		for i, ns := range dataArray {
			match := false
			if ns["ip"] == name || ns["dnsvservername"] == name {
				match = true
			}
			if match && checkType != "" && ns["type"] != checkType {
				match = false
			}
			if match {
				foundIndex = i
				break
			}
		}
		if foundIndex != -1 {
			break
		}
	}

	if foundIndex == -1 {
		return false
	}

	dnsnameserverSetAttrFromGet(ctx, data, dataArray[foundIndex], name, dnsType)

	return true
}
