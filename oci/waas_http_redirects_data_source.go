// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v26/common"
	oci_waas "github.com/oracle/oci-go-sdk/v26/waas"
)

func init() {
	RegisterDatasource("oci_waas_http_redirects", WaasHttpRedirectsDataSource())
}

func WaasHttpRedirectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWaasHttpRedirects,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"states": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_redirects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(WaasHttpRedirectResource()),
			},
		},
	}
}

func readWaasHttpRedirects(d *schema.ResourceData, m interface{}) error {
	sync := &WaasHttpRedirectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).redirectClient()

	return ReadResource(sync)
}

type WaasHttpRedirectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.RedirectClient
	Res    *oci_waas.ListHttpRedirectsResponse
}

func (s *WaasHttpRedirectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasHttpRedirectsDataSourceCrud) Get() error {
	request := oci_waas.ListHttpRedirectsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayNames, ok := s.D.GetOkExists("display_names"); ok {
		interfaces := displayNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("display_names") {
			request.DisplayName = tmp
		}
	}

	if ids, ok := s.D.GetOkExists("ids"); ok {
		interfaces := ids.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ids") {
			request.Id = tmp
		}
	}

	if states, ok := s.D.GetOkExists("states"); ok {
		interfaces := states.([]interface{})
		tmp := make([]oci_waas.ListHttpRedirectsLifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_waas.ListHttpRedirectsLifecycleStateEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("states") {
			request.LifecycleState = tmp
		}
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "waas")

	response, err := s.Client.ListHttpRedirects(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListHttpRedirects(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WaasHttpRedirectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		httpRedirect := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			httpRedirect["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			httpRedirect["display_name"] = *r.DisplayName
		}

		if r.Domain != nil {
			httpRedirect["domain"] = *r.Domain
		}

		httpRedirect["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			httpRedirect["id"] = *r.Id
		}

		httpRedirect["response_code"] = r.ResponseCode

		httpRedirect["state"] = r.LifecycleState

		if r.Target != nil {
			httpRedirect["target"] = []interface{}{HttpRedirectTargetToMap(r.Target)}
		} else {
			httpRedirect["target"] = nil
		}

		if r.TimeCreated != nil {
			httpRedirect["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, httpRedirect)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, WaasHttpRedirectsDataSource().Schema["http_redirects"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("http_redirects", resources); err != nil {
		return err
	}

	return nil
}
