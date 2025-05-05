package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &BillingboxProvider{}
var _ provider.ProviderWithFunctions = &BillingboxProvider{}
var _ provider.ProviderWithEphemeralResources = &BillingboxProvider{}

// BillingboxProvider defines the provider implementation.
type BillingboxProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// BillingboxProviderModel describes the provider data model.
type BillingboxProviderModel struct {
	// TODO implement the model for the following schema &schema.Provider{
	//	Schema: map[string]*schema.Schema{
	//		"url": {
	//			Type:        schema.TypeString,
	//			Required:    true,
	//			DefaultFunc: schema.EnvDefaultFunc("AIDBOX_URL", nil),
	//			Description: "The URL of the Aidbox instance",
	//		},
	//		"client_id": {
	//			Type:        schema.TypeString,
	//			Required:    true,
	//			DefaultFunc: schema.EnvDefaultFunc("AIDBOX_CLIENT_ID", nil),
	//			Description: "The client ID for authentication",
	//		},
	//		"client_secret": {
	//			Type:        schema.TypeString,
	//			Required:    true,
	//			DefaultFunc: schema.EnvDefaultFunc("AIDBOX_CLIENT_SECRET", nil),
	//			Description: "The client secret for authentication",
	//			Sensitive:   true,
	//		},
	//		"username": {
	//			Type:        schema.TypeString,
	//			Optional:    true,
	//			DefaultFunc: schema.EnvDefaultFunc("AIDBOX_USERNAME", nil),
	//			Description: "The username for password grant authentication (optional)",
	//		},
	//		"password": {
	//			Type:        schema.TypeString,
	//			Optional:    true,
	//			DefaultFunc: schema.EnvDefaultFunc("AIDBOX_PASSWORD", nil),
	//			Description: "The password for password grant authentication (optional)",
	//			Sensitive:   true,
	//		},
	//	},
	//	ResourcesMap: map[string]*schema.Resource{
	//		"aidbox_access_policy": aidbox.ResourceAidboxAccessPolicy(),
	//		"aidbox_client":        aidbox.ResourceAidboxClient(),
	//		"aidbox_role":          aidbox.ResourceAidboxRole(),
	//		"aidbox_user":          aidbox.ResourceAidboxUser(),
	//	},
	//	ConfigureContextFunc: providerConfigure,
	// }
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *BillingboxProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "billingbox"
	resp.Version = p.version
}

func (p *BillingboxProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
			},
		},
	}
}

func (p *BillingboxProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data BillingboxProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *BillingboxProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *BillingboxProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		NewExampleEphemeralResource,
	}
}

func (p *BillingboxProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewExampleDataSource,
	}
}

func (p *BillingboxProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewExampleFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &BillingboxProvider{
			version: version,
		}
	}
}
