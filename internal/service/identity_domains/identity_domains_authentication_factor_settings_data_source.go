// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsAuthenticationFactorSettingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsAuthenticationFactorSettings,
		Schema: map[string]*schema.Schema{
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"authentication_factor_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityDomainsAuthenticationFactorSettingResource()),
			},
			"items_per_page": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"start_index": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_results": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readIdentityDomainsAuthenticationFactorSettings(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAuthenticationFactorSettingsDataSourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	return tfresource.ReadResource(sync)
}

type IdentityDomainsAuthenticationFactorSettingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListAuthenticationFactorSettingsResponse
}

func (s *IdentityDomainsAuthenticationFactorSettingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsAuthenticationFactorSettingsDataSourceCrud) Get() error {
	request := oci_identity_domains.ListAuthenticationFactorSettingsRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.ListAuthenticationFactorSettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsAuthenticationFactorSettingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsAuthenticationFactorSettingsDataSource-", IdentityDomainsAuthenticationFactorSettingsDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, AuthenticationFactorSettingToMap(item))
	}
	s.D.Set("authentication_factor_settings", resources)

	if s.Res.ItemsPerPage != nil {
		s.D.Set("items_per_page", *s.Res.ItemsPerPage)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.StartIndex != nil {
		s.D.Set("start_index", *s.Res.StartIndex)
	}

	if s.Res.TotalResults != nil {
		s.D.Set("total_results", *s.Res.TotalResults)
	}

	return nil
}
