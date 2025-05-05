package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"terraform-provider-billingbox/internal/client"
)

// var _ provider.Provider = &BillingboxProvider{}
// var _ provider.ProviderWithFunctions = &BillingboxProvider{}
// var _ provider.ProviderWithEphemeralResources = &BillingboxProvider{}

// BillingboxProvider defines the provider implementation.
type BillingboxProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// BillingboxProviderModel describes the provider data model.
type BillingboxProviderModel struct {
	Url          types.String `tfsdk:"url"`
	ClientId     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
	Username     types.String `tfsdk:"username"`
	Password     types.String `tfsdk:"password"`
	Endpoint     types.String `tfsdk:"endpoint"`
}

func (p *BillingboxProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "billingbox"
	resp.Version = p.version
}

func (p *BillingboxProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				MarkdownDescription: "The URL of the Aidbox instance",
				Required:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: "The client ID for authentication",
				Required:            true,
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: "The client secret for authentication",
				Required:            true,
				Sensitive:           true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "The username for password‐grant authentication (optional)",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "The password for password‐grant authentication (optional)",
				Optional:            true,
				Sensitive:           true,
			},
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
			},
		},
	}
}

func (p *BillingboxProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Billingbox client")

	var data BillingboxProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Provider configuration", map[string]interface{}{
		"url":       data.Url.ValueString(),
		"client_id": data.ClientId.ValueString(),
	})

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client, err := client.NewClient(&client.ClientConfig{
		URL:          data.Url.ValueString(),
		ClientID:     data.ClientId.ValueString(),
		ClientSecret: data.ClientSecret.ValueString(),
		Username:     data.Username.ValueString(),
		Password:     data.Password.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating client", err.Error())
		return
	}

	tflog.Debug(ctx, "Created Billingbox client", map[string]interface{}{
		"url": data.Url.ValueString(),
	})

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *BillingboxProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUserResource,
		NewRoleResource,
		NewAccessPolicyResource,
	}
}

// func (p *BillingboxProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
//	return []func() ephemeral.EphemeralResource{
//		NewExampleEphemeralResource,
//	}
// }

func (p *BillingboxProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// func (p *BillingboxProvider) Functions(ctx context.Context) []func() function.Function {
//	return []func() function.Function{
//		NewExampleFunction,
//	}
// }

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &BillingboxProvider{
			version: version,
		}
	}
}
