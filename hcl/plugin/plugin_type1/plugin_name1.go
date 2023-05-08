package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type BlockTest struct {
	Test string `hcl:"test"`
}

type PluginConfig struct {
	Data1     string    `hcl:"plugin_string_data"`
	Data2     int       `hcl:"plugin_int_data"`
	BlockData BlockTest `hcl:"plugin_block_data,block"`
}

func ConfigSpec() hcldec.Spec {
	spec := hcldec.ObjectSpec{
		"plugin_string_data": &hcldec.AttrSpec{
			Name:     "plugin_string_data",
			Type:     cty.String,
			Required: true,
		},
		"plugin_int_data": &hcldec.AttrSpec{
			Name:     "plugin_int_data",
			Type:     cty.Number,
			Required: true,
		},
		"plugin_block_data": &hcldec.BlockAttrsSpec{
			TypeName:    "plugin_block_data",
			ElementType: cty.String,
			Required:    true,
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
