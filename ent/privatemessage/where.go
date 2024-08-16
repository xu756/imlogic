// Code generated by ent, DO NOT EDIT.

package privatemessage

import (
	"imlogic/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLTE(FieldID, id))
}

// MsgType applies equality check predicate on the "msg_type" field. It's identical to MsgTypeEQ.
func MsgType(v int32) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldMsgType, v))
}

// MsgID applies equality check predicate on the "msg_id" field. It's identical to MsgIDEQ.
func MsgID(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldMsgID, v))
}

// ChatID applies equality check predicate on the "chat_id" field. It's identical to ChatIDEQ.
func ChatID(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldChatID, v))
}

// SenderID applies equality check predicate on the "sender_id" field. It's identical to SenderIDEQ.
func SenderID(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldSenderID, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldTimestamp, v))
}

// MsgTypeEQ applies the EQ predicate on the "msg_type" field.
func MsgTypeEQ(v int32) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldMsgType, v))
}

// MsgTypeNEQ applies the NEQ predicate on the "msg_type" field.
func MsgTypeNEQ(v int32) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNEQ(FieldMsgType, v))
}

// MsgTypeIn applies the In predicate on the "msg_type" field.
func MsgTypeIn(vs ...int32) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldIn(FieldMsgType, vs...))
}

// MsgTypeNotIn applies the NotIn predicate on the "msg_type" field.
func MsgTypeNotIn(vs ...int32) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNotIn(FieldMsgType, vs...))
}

// MsgTypeGT applies the GT predicate on the "msg_type" field.
func MsgTypeGT(v int32) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGT(FieldMsgType, v))
}

// MsgTypeGTE applies the GTE predicate on the "msg_type" field.
func MsgTypeGTE(v int32) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGTE(FieldMsgType, v))
}

// MsgTypeLT applies the LT predicate on the "msg_type" field.
func MsgTypeLT(v int32) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLT(FieldMsgType, v))
}

// MsgTypeLTE applies the LTE predicate on the "msg_type" field.
func MsgTypeLTE(v int32) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLTE(FieldMsgType, v))
}

// MsgIDEQ applies the EQ predicate on the "msg_id" field.
func MsgIDEQ(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldMsgID, v))
}

// MsgIDNEQ applies the NEQ predicate on the "msg_id" field.
func MsgIDNEQ(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNEQ(FieldMsgID, v))
}

// MsgIDIn applies the In predicate on the "msg_id" field.
func MsgIDIn(vs ...string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldIn(FieldMsgID, vs...))
}

// MsgIDNotIn applies the NotIn predicate on the "msg_id" field.
func MsgIDNotIn(vs ...string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNotIn(FieldMsgID, vs...))
}

// MsgIDGT applies the GT predicate on the "msg_id" field.
func MsgIDGT(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGT(FieldMsgID, v))
}

// MsgIDGTE applies the GTE predicate on the "msg_id" field.
func MsgIDGTE(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGTE(FieldMsgID, v))
}

// MsgIDLT applies the LT predicate on the "msg_id" field.
func MsgIDLT(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLT(FieldMsgID, v))
}

// MsgIDLTE applies the LTE predicate on the "msg_id" field.
func MsgIDLTE(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLTE(FieldMsgID, v))
}

// MsgIDContains applies the Contains predicate on the "msg_id" field.
func MsgIDContains(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldContains(FieldMsgID, v))
}

// MsgIDHasPrefix applies the HasPrefix predicate on the "msg_id" field.
func MsgIDHasPrefix(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldHasPrefix(FieldMsgID, v))
}

// MsgIDHasSuffix applies the HasSuffix predicate on the "msg_id" field.
func MsgIDHasSuffix(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldHasSuffix(FieldMsgID, v))
}

// MsgIDEqualFold applies the EqualFold predicate on the "msg_id" field.
func MsgIDEqualFold(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEqualFold(FieldMsgID, v))
}

// MsgIDContainsFold applies the ContainsFold predicate on the "msg_id" field.
func MsgIDContainsFold(v string) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldContainsFold(FieldMsgID, v))
}

// ChatIDEQ applies the EQ predicate on the "chat_id" field.
func ChatIDEQ(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldChatID, v))
}

// ChatIDNEQ applies the NEQ predicate on the "chat_id" field.
func ChatIDNEQ(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNEQ(FieldChatID, v))
}

// ChatIDIn applies the In predicate on the "chat_id" field.
func ChatIDIn(vs ...int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldIn(FieldChatID, vs...))
}

// ChatIDNotIn applies the NotIn predicate on the "chat_id" field.
func ChatIDNotIn(vs ...int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNotIn(FieldChatID, vs...))
}

// ChatIDGT applies the GT predicate on the "chat_id" field.
func ChatIDGT(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGT(FieldChatID, v))
}

// ChatIDGTE applies the GTE predicate on the "chat_id" field.
func ChatIDGTE(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGTE(FieldChatID, v))
}

// ChatIDLT applies the LT predicate on the "chat_id" field.
func ChatIDLT(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLT(FieldChatID, v))
}

// ChatIDLTE applies the LTE predicate on the "chat_id" field.
func ChatIDLTE(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLTE(FieldChatID, v))
}

// SenderIDEQ applies the EQ predicate on the "sender_id" field.
func SenderIDEQ(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldSenderID, v))
}

// SenderIDNEQ applies the NEQ predicate on the "sender_id" field.
func SenderIDNEQ(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNEQ(FieldSenderID, v))
}

// SenderIDIn applies the In predicate on the "sender_id" field.
func SenderIDIn(vs ...int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldIn(FieldSenderID, vs...))
}

// SenderIDNotIn applies the NotIn predicate on the "sender_id" field.
func SenderIDNotIn(vs ...int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNotIn(FieldSenderID, vs...))
}

// SenderIDGT applies the GT predicate on the "sender_id" field.
func SenderIDGT(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGT(FieldSenderID, v))
}

// SenderIDGTE applies the GTE predicate on the "sender_id" field.
func SenderIDGTE(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGTE(FieldSenderID, v))
}

// SenderIDLT applies the LT predicate on the "sender_id" field.
func SenderIDLT(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLT(FieldSenderID, v))
}

// SenderIDLTE applies the LTE predicate on the "sender_id" field.
func SenderIDLTE(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLTE(FieldSenderID, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v int64) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.FieldLTE(FieldTimestamp, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PrivateMessage) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PrivateMessage) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PrivateMessage) predicate.PrivateMessage {
	return predicate.PrivateMessage(sql.NotPredicates(p))
}