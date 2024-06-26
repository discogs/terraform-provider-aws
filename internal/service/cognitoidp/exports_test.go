// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoidp

// Exports for use in tests only.
var (
	ResourceUserGroup             = resourceUserGroup
	ResourceUserPoolClient        = newResourceUserPoolClient
	ResourceManagedUserPoolClient = newResourceManagedUserPoolClient

	FindGroupByTwoPartKey = findGroupByTwoPartKey
)
