package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type PluginConfig struct {
	Data string `hcl:"plugin_string_data"`
}

func ConfigSpec() hcldec.Spec {
	spec := hcldec.ObjectSpec{
		"plugin_string_data": &hcldec.AttrSpec{
			Name:     "plugin_string_data",
			Type:     cty.String,
			Required: true,
		},
	}
	return &spec
}

func DecodeSpec(body hcl.Body) *PluginConfig {
	var config PluginConfig
	gohcl.DecodeBody(body, nil, &config)
	return &config
}

func ToString(body hcl.Body) string {
	return fmt.Sprintf("%+v", DecodeSpec(body))
}
