package cacheobject

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CacheobjectDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Computed:    true,
				Description: "Not applicable for reads (action-only on the resource). Present for model compatibility.",
			},
			"group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the content group whose objects should be listed.",
			},
			"groupname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the content group to which the object belongs. It will display only the objects belonging to the specified content group. You must also set the Host parameter.",
			},
			"host": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host name of the object. Parameter \"url\" must be specified.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP request method that caused the object to be stored.",
			},
			"httpstatus": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP status of the object.",
			},
			"ignoremarkerobjects": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore marker objects. Marker objects are created when a response exceeds the maximum or minimum response size for the content group or has not yet received the minimum number of hits for the content group.",
			},
			"includenotreadyobjects": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include responses that have not yet reached a minimum number of hits before being cached.",
			},
			"locator": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the cached object.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Host port of the object. You must also set the Host parameter.",
			},
			"tosecondary": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Object will be saved onto Secondary.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the particular object whose details is required. Parameter \"host\" must be specified along with the URL.",
			},
		},
	}
}
