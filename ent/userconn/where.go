// Code generated by ent, DO NOT EDIT.

package userconn

import (
	"imlogic/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserConn {
	return predicate.UserConn(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserConn {
	return predicate.UserConn(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserConn {
	return predicate.UserConn(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserConn {
	return predicate.UserConn(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserConn {
	return predicate.UserConn(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserConn {
	return predicate.UserConn(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserConn {
	return predicate.UserConn(sql.FieldLTE(FieldID, id))
}

// LinkID applies equality check predicate on the "link_id" field. It's identical to LinkIDEQ.
func LinkID(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldLinkID, v))
}

// LinkTime applies equality check predicate on the "link_time" field. It's identical to LinkTimeEQ.
func LinkTime(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldLinkTime, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldUserID, v))
}

// HostName applies equality check predicate on the "host_name" field. It's identical to HostNameEQ.
func HostName(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldHostName, v))
}

// Device applies equality check predicate on the "device" field. It's identical to DeviceEQ.
func Device(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldDevice, v))
}

// LastHeartbeatTime applies equality check predicate on the "last_heartbeat_time" field. It's identical to LastHeartbeatTimeEQ.
func LastHeartbeatTime(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldLastHeartbeatTime, v))
}

// LinkIDEQ applies the EQ predicate on the "link_id" field.
func LinkIDEQ(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldLinkID, v))
}

// LinkIDNEQ applies the NEQ predicate on the "link_id" field.
func LinkIDNEQ(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldNEQ(FieldLinkID, v))
}

// LinkIDIn applies the In predicate on the "link_id" field.
func LinkIDIn(vs ...string) predicate.UserConn {
	return predicate.UserConn(sql.FieldIn(FieldLinkID, vs...))
}

// LinkIDNotIn applies the NotIn predicate on the "link_id" field.
func LinkIDNotIn(vs ...string) predicate.UserConn {
	return predicate.UserConn(sql.FieldNotIn(FieldLinkID, vs...))
}

// LinkIDGT applies the GT predicate on the "link_id" field.
func LinkIDGT(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldGT(FieldLinkID, v))
}

// LinkIDGTE applies the GTE predicate on the "link_id" field.
func LinkIDGTE(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldGTE(FieldLinkID, v))
}

// LinkIDLT applies the LT predicate on the "link_id" field.
func LinkIDLT(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldLT(FieldLinkID, v))
}

// LinkIDLTE applies the LTE predicate on the "link_id" field.
func LinkIDLTE(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldLTE(FieldLinkID, v))
}

// LinkIDContains applies the Contains predicate on the "link_id" field.
func LinkIDContains(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldContains(FieldLinkID, v))
}

// LinkIDHasPrefix applies the HasPrefix predicate on the "link_id" field.
func LinkIDHasPrefix(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldHasPrefix(FieldLinkID, v))
}

// LinkIDHasSuffix applies the HasSuffix predicate on the "link_id" field.
func LinkIDHasSuffix(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldHasSuffix(FieldLinkID, v))
}

// LinkIDEqualFold applies the EqualFold predicate on the "link_id" field.
func LinkIDEqualFold(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldEqualFold(FieldLinkID, v))
}

// LinkIDContainsFold applies the ContainsFold predicate on the "link_id" field.
func LinkIDContainsFold(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldContainsFold(FieldLinkID, v))
}

// LinkTimeEQ applies the EQ predicate on the "link_time" field.
func LinkTimeEQ(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldLinkTime, v))
}

// LinkTimeNEQ applies the NEQ predicate on the "link_time" field.
func LinkTimeNEQ(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldNEQ(FieldLinkTime, v))
}

// LinkTimeIn applies the In predicate on the "link_time" field.
func LinkTimeIn(vs ...time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldIn(FieldLinkTime, vs...))
}

// LinkTimeNotIn applies the NotIn predicate on the "link_time" field.
func LinkTimeNotIn(vs ...time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldNotIn(FieldLinkTime, vs...))
}

// LinkTimeGT applies the GT predicate on the "link_time" field.
func LinkTimeGT(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldGT(FieldLinkTime, v))
}

// LinkTimeGTE applies the GTE predicate on the "link_time" field.
func LinkTimeGTE(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldGTE(FieldLinkTime, v))
}

// LinkTimeLT applies the LT predicate on the "link_time" field.
func LinkTimeLT(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldLT(FieldLinkTime, v))
}

// LinkTimeLTE applies the LTE predicate on the "link_time" field.
func LinkTimeLTE(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldLTE(FieldLinkTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.UserConn {
	return predicate.UserConn(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.UserConn {
	return predicate.UserConn(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.UserConn {
	return predicate.UserConn(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.UserConn {
	return predicate.UserConn(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.UserConn {
	return predicate.UserConn(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.UserConn {
	return predicate.UserConn(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.UserConn {
	return predicate.UserConn(sql.FieldLTE(FieldUserID, v))
}

// HostNameEQ applies the EQ predicate on the "host_name" field.
func HostNameEQ(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldHostName, v))
}

// HostNameNEQ applies the NEQ predicate on the "host_name" field.
func HostNameNEQ(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldNEQ(FieldHostName, v))
}

// HostNameIn applies the In predicate on the "host_name" field.
func HostNameIn(vs ...string) predicate.UserConn {
	return predicate.UserConn(sql.FieldIn(FieldHostName, vs...))
}

// HostNameNotIn applies the NotIn predicate on the "host_name" field.
func HostNameNotIn(vs ...string) predicate.UserConn {
	return predicate.UserConn(sql.FieldNotIn(FieldHostName, vs...))
}

// HostNameGT applies the GT predicate on the "host_name" field.
func HostNameGT(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldGT(FieldHostName, v))
}

// HostNameGTE applies the GTE predicate on the "host_name" field.
func HostNameGTE(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldGTE(FieldHostName, v))
}

// HostNameLT applies the LT predicate on the "host_name" field.
func HostNameLT(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldLT(FieldHostName, v))
}

// HostNameLTE applies the LTE predicate on the "host_name" field.
func HostNameLTE(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldLTE(FieldHostName, v))
}

// HostNameContains applies the Contains predicate on the "host_name" field.
func HostNameContains(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldContains(FieldHostName, v))
}

// HostNameHasPrefix applies the HasPrefix predicate on the "host_name" field.
func HostNameHasPrefix(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldHasPrefix(FieldHostName, v))
}

// HostNameHasSuffix applies the HasSuffix predicate on the "host_name" field.
func HostNameHasSuffix(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldHasSuffix(FieldHostName, v))
}

// HostNameEqualFold applies the EqualFold predicate on the "host_name" field.
func HostNameEqualFold(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldEqualFold(FieldHostName, v))
}

// HostNameContainsFold applies the ContainsFold predicate on the "host_name" field.
func HostNameContainsFold(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldContainsFold(FieldHostName, v))
}

// DeviceEQ applies the EQ predicate on the "device" field.
func DeviceEQ(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldDevice, v))
}

// DeviceNEQ applies the NEQ predicate on the "device" field.
func DeviceNEQ(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldNEQ(FieldDevice, v))
}

// DeviceIn applies the In predicate on the "device" field.
func DeviceIn(vs ...string) predicate.UserConn {
	return predicate.UserConn(sql.FieldIn(FieldDevice, vs...))
}

// DeviceNotIn applies the NotIn predicate on the "device" field.
func DeviceNotIn(vs ...string) predicate.UserConn {
	return predicate.UserConn(sql.FieldNotIn(FieldDevice, vs...))
}

// DeviceGT applies the GT predicate on the "device" field.
func DeviceGT(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldGT(FieldDevice, v))
}

// DeviceGTE applies the GTE predicate on the "device" field.
func DeviceGTE(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldGTE(FieldDevice, v))
}

// DeviceLT applies the LT predicate on the "device" field.
func DeviceLT(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldLT(FieldDevice, v))
}

// DeviceLTE applies the LTE predicate on the "device" field.
func DeviceLTE(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldLTE(FieldDevice, v))
}

// DeviceContains applies the Contains predicate on the "device" field.
func DeviceContains(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldContains(FieldDevice, v))
}

// DeviceHasPrefix applies the HasPrefix predicate on the "device" field.
func DeviceHasPrefix(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldHasPrefix(FieldDevice, v))
}

// DeviceHasSuffix applies the HasSuffix predicate on the "device" field.
func DeviceHasSuffix(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldHasSuffix(FieldDevice, v))
}

// DeviceEqualFold applies the EqualFold predicate on the "device" field.
func DeviceEqualFold(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldEqualFold(FieldDevice, v))
}

// DeviceContainsFold applies the ContainsFold predicate on the "device" field.
func DeviceContainsFold(v string) predicate.UserConn {
	return predicate.UserConn(sql.FieldContainsFold(FieldDevice, v))
}

// LastHeartbeatTimeEQ applies the EQ predicate on the "last_heartbeat_time" field.
func LastHeartbeatTimeEQ(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldEQ(FieldLastHeartbeatTime, v))
}

// LastHeartbeatTimeNEQ applies the NEQ predicate on the "last_heartbeat_time" field.
func LastHeartbeatTimeNEQ(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldNEQ(FieldLastHeartbeatTime, v))
}

// LastHeartbeatTimeIn applies the In predicate on the "last_heartbeat_time" field.
func LastHeartbeatTimeIn(vs ...time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldIn(FieldLastHeartbeatTime, vs...))
}

// LastHeartbeatTimeNotIn applies the NotIn predicate on the "last_heartbeat_time" field.
func LastHeartbeatTimeNotIn(vs ...time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldNotIn(FieldLastHeartbeatTime, vs...))
}

// LastHeartbeatTimeGT applies the GT predicate on the "last_heartbeat_time" field.
func LastHeartbeatTimeGT(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldGT(FieldLastHeartbeatTime, v))
}

// LastHeartbeatTimeGTE applies the GTE predicate on the "last_heartbeat_time" field.
func LastHeartbeatTimeGTE(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldGTE(FieldLastHeartbeatTime, v))
}

// LastHeartbeatTimeLT applies the LT predicate on the "last_heartbeat_time" field.
func LastHeartbeatTimeLT(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldLT(FieldLastHeartbeatTime, v))
}

// LastHeartbeatTimeLTE applies the LTE predicate on the "last_heartbeat_time" field.
func LastHeartbeatTimeLTE(v time.Time) predicate.UserConn {
	return predicate.UserConn(sql.FieldLTE(FieldLastHeartbeatTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserConn) predicate.UserConn {
	return predicate.UserConn(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserConn) predicate.UserConn {
	return predicate.UserConn(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserConn) predicate.UserConn {
	return predicate.UserConn(sql.NotPredicates(p))
}
