// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package googlecloud

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "googlecloud", asset.ModuleFieldsPri, AssetGooglecloud); err != nil {
		panic(err)
	}
}

// AssetGooglecloud returns asset data.
// This is the base64 encoded gzipped contents of module/googlecloud.
func AssetGooglecloud() string {
	return "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"
}
