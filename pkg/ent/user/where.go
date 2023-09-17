// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wtkeqrf0/restService/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Surname applies equality check predicate on the "surname" field. It's identical to SurnameEQ.
func Surname(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSurname, v))
}

// Patronymic applies equality check predicate on the "patronymic" field. It's identical to PatronymicEQ.
func Patronymic(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPatronymic, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAge, v))
}

// Gender applies equality check predicate on the "gender" field. It's identical to GenderEQ.
func Gender(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGender, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCountry, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// SurnameEQ applies the EQ predicate on the "surname" field.
func SurnameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSurname, v))
}

// SurnameNEQ applies the NEQ predicate on the "surname" field.
func SurnameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSurname, v))
}

// SurnameIn applies the In predicate on the "surname" field.
func SurnameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldSurname, vs...))
}

// SurnameNotIn applies the NotIn predicate on the "surname" field.
func SurnameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSurname, vs...))
}

// SurnameGT applies the GT predicate on the "surname" field.
func SurnameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldSurname, v))
}

// SurnameGTE applies the GTE predicate on the "surname" field.
func SurnameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldSurname, v))
}

// SurnameLT applies the LT predicate on the "surname" field.
func SurnameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldSurname, v))
}

// SurnameLTE applies the LTE predicate on the "surname" field.
func SurnameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldSurname, v))
}

// SurnameContains applies the Contains predicate on the "surname" field.
func SurnameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldSurname, v))
}

// SurnameHasPrefix applies the HasPrefix predicate on the "surname" field.
func SurnameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldSurname, v))
}

// SurnameHasSuffix applies the HasSuffix predicate on the "surname" field.
func SurnameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldSurname, v))
}

// SurnameEqualFold applies the EqualFold predicate on the "surname" field.
func SurnameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldSurname, v))
}

// SurnameContainsFold applies the ContainsFold predicate on the "surname" field.
func SurnameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldSurname, v))
}

// PatronymicEQ applies the EQ predicate on the "patronymic" field.
func PatronymicEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPatronymic, v))
}

// PatronymicNEQ applies the NEQ predicate on the "patronymic" field.
func PatronymicNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPatronymic, v))
}

// PatronymicIn applies the In predicate on the "patronymic" field.
func PatronymicIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPatronymic, vs...))
}

// PatronymicNotIn applies the NotIn predicate on the "patronymic" field.
func PatronymicNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPatronymic, vs...))
}

// PatronymicGT applies the GT predicate on the "patronymic" field.
func PatronymicGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPatronymic, v))
}

// PatronymicGTE applies the GTE predicate on the "patronymic" field.
func PatronymicGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPatronymic, v))
}

// PatronymicLT applies the LT predicate on the "patronymic" field.
func PatronymicLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPatronymic, v))
}

// PatronymicLTE applies the LTE predicate on the "patronymic" field.
func PatronymicLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPatronymic, v))
}

// PatronymicContains applies the Contains predicate on the "patronymic" field.
func PatronymicContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPatronymic, v))
}

// PatronymicHasPrefix applies the HasPrefix predicate on the "patronymic" field.
func PatronymicHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPatronymic, v))
}

// PatronymicHasSuffix applies the HasSuffix predicate on the "patronymic" field.
func PatronymicHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPatronymic, v))
}

// PatronymicIsNil applies the IsNil predicate on the "patronymic" field.
func PatronymicIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPatronymic))
}

// PatronymicNotNil applies the NotNil predicate on the "patronymic" field.
func PatronymicNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPatronymic))
}

// PatronymicEqualFold applies the EqualFold predicate on the "patronymic" field.
func PatronymicEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPatronymic, v))
}

// PatronymicContainsFold applies the ContainsFold predicate on the "patronymic" field.
func PatronymicContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPatronymic, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAge, v))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGender, vs...))
}

// GenderGT applies the GT predicate on the "gender" field.
func GenderGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGender, v))
}

// GenderGTE applies the GTE predicate on the "gender" field.
func GenderGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGender, v))
}

// GenderLT applies the LT predicate on the "gender" field.
func GenderLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGender, v))
}

// GenderLTE applies the LTE predicate on the "gender" field.
func GenderLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGender, v))
}

// GenderContains applies the Contains predicate on the "gender" field.
func GenderContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGender, v))
}

// GenderHasPrefix applies the HasPrefix predicate on the "gender" field.
func GenderHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGender, v))
}

// GenderHasSuffix applies the HasSuffix predicate on the "gender" field.
func GenderHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGender, v))
}

// GenderEqualFold applies the EqualFold predicate on the "gender" field.
func GenderEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGender, v))
}

// GenderContainsFold applies the ContainsFold predicate on the "gender" field.
func GenderContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGender, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldCountry, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}