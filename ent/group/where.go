// Code generated by ent, DO NOT EDIT.

package group

import (
	"imlogic/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldUUID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldName, v))
}

// Intro applies equality check predicate on the "intro" field. It's identical to IntroEQ.
func Intro(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldIntro, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldUUID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldName, v))
}

// IntroEQ applies the EQ predicate on the "intro" field.
func IntroEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldIntro, v))
}

// IntroNEQ applies the NEQ predicate on the "intro" field.
func IntroNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldIntro, v))
}

// IntroIn applies the In predicate on the "intro" field.
func IntroIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldIntro, vs...))
}

// IntroNotIn applies the NotIn predicate on the "intro" field.
func IntroNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldIntro, vs...))
}

// IntroGT applies the GT predicate on the "intro" field.
func IntroGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldIntro, v))
}

// IntroGTE applies the GTE predicate on the "intro" field.
func IntroGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldIntro, v))
}

// IntroLT applies the LT predicate on the "intro" field.
func IntroLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldIntro, v))
}

// IntroLTE applies the LTE predicate on the "intro" field.
func IntroLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldIntro, v))
}

// IntroContains applies the Contains predicate on the "intro" field.
func IntroContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldIntro, v))
}

// IntroHasPrefix applies the HasPrefix predicate on the "intro" field.
func IntroHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldIntro, v))
}

// IntroHasSuffix applies the HasSuffix predicate on the "intro" field.
func IntroHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldIntro, v))
}

// IntroEqualFold applies the EqualFold predicate on the "intro" field.
func IntroEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldIntro, v))
}

// IntroContainsFold applies the ContainsFold predicate on the "intro" field.
func IntroContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldIntro, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(sql.NotPredicates(p))
}
