// Code generated by ent, DO NOT EDIT.

package cleanuser

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/viewcomposite/ent/predicate"
)

// ID applies equality check predicate on the "id" field. It's identical to IDEQ.
func ID(v int) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldEQ(FieldID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldEQ(FieldName, v))
}

// PublicInfo applies equality check predicate on the "public_info" field. It's identical to PublicInfoEQ.
func PublicInfo(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldEQ(FieldPublicInfo, v))
}

// IDEQ applies the EQ predicate on the "id" field.
func IDEQ(v int) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldEQ(FieldID, v))
}

// IDNEQ applies the NEQ predicate on the "id" field.
func IDNEQ(v int) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldNEQ(FieldID, v))
}

// IDIn applies the In predicate on the "id" field.
func IDIn(vs ...int) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldIn(FieldID, vs...))
}

// IDNotIn applies the NotIn predicate on the "id" field.
func IDNotIn(vs ...int) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldNotIn(FieldID, vs...))
}

// IDGT applies the GT predicate on the "id" field.
func IDGT(v int) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldGT(FieldID, v))
}

// IDGTE applies the GTE predicate on the "id" field.
func IDGTE(v int) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldGTE(FieldID, v))
}

// IDLT applies the LT predicate on the "id" field.
func IDLT(v int) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldLT(FieldID, v))
}

// IDLTE applies the LTE predicate on the "id" field.
func IDLTE(v int) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldLTE(FieldID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldContainsFold(FieldName, v))
}

// PublicInfoEQ applies the EQ predicate on the "public_info" field.
func PublicInfoEQ(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldEQ(FieldPublicInfo, v))
}

// PublicInfoNEQ applies the NEQ predicate on the "public_info" field.
func PublicInfoNEQ(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldNEQ(FieldPublicInfo, v))
}

// PublicInfoIn applies the In predicate on the "public_info" field.
func PublicInfoIn(vs ...string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldIn(FieldPublicInfo, vs...))
}

// PublicInfoNotIn applies the NotIn predicate on the "public_info" field.
func PublicInfoNotIn(vs ...string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldNotIn(FieldPublicInfo, vs...))
}

// PublicInfoGT applies the GT predicate on the "public_info" field.
func PublicInfoGT(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldGT(FieldPublicInfo, v))
}

// PublicInfoGTE applies the GTE predicate on the "public_info" field.
func PublicInfoGTE(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldGTE(FieldPublicInfo, v))
}

// PublicInfoLT applies the LT predicate on the "public_info" field.
func PublicInfoLT(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldLT(FieldPublicInfo, v))
}

// PublicInfoLTE applies the LTE predicate on the "public_info" field.
func PublicInfoLTE(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldLTE(FieldPublicInfo, v))
}

// PublicInfoContains applies the Contains predicate on the "public_info" field.
func PublicInfoContains(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldContains(FieldPublicInfo, v))
}

// PublicInfoHasPrefix applies the HasPrefix predicate on the "public_info" field.
func PublicInfoHasPrefix(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldHasPrefix(FieldPublicInfo, v))
}

// PublicInfoHasSuffix applies the HasSuffix predicate on the "public_info" field.
func PublicInfoHasSuffix(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldHasSuffix(FieldPublicInfo, v))
}

// PublicInfoEqualFold applies the EqualFold predicate on the "public_info" field.
func PublicInfoEqualFold(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldEqualFold(FieldPublicInfo, v))
}

// PublicInfoContainsFold applies the ContainsFold predicate on the "public_info" field.
func PublicInfoContainsFold(v string) predicate.CleanUser {
	return predicate.CleanUser(sql.FieldContainsFold(FieldPublicInfo, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CleanUser) predicate.CleanUser {
	return predicate.CleanUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CleanUser) predicate.CleanUser {
	return predicate.CleanUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CleanUser) predicate.CleanUser {
	return predicate.CleanUser(sql.NotPredicates(p))
}