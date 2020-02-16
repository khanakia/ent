// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package spec

const (
	// Label holds the string label denoting the spec type in the database.
	Label = "spec"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"

	// Table holds the table name of the spec in the database.
	Table = "specs"
	// CardTable is the table the holds the card relation/edge. The primary key declared below.
	CardTable = "spec_card"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
)

// Columns holds all SQL columns for spec fields.
var Columns = []string{
	FieldID,
}

var (
	// CardPrimaryKey and CardColumn2 are the table columns denoting the
	// primary key for the card relation (M2M).
	CardPrimaryKey = []string{"spec_id", "card_id"}
)