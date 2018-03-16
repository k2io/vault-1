package activedirectory

import (
	"github.com/go-ldap/ldap"
	"strings"
	log "github.com/mgutz/logxi/v1"
)

// Entry is an Active Directory-specific construct
// to make knowing and grabbing fields more convenient,
// while retaining all original information.
func NewEntry(ldapEntry *ldap.Entry) *Entry {
	m := make(map[Field][]string)
	for _, attribute := range ldapEntry.Attributes {
		field, err := Parse(attribute.Name)
		if err != nil {
			log.Warn("no field exists in the ActiveDirectory registry for %s, ignoring it")
			continue
		}
		m[field] = attribute.Values
	}
	return &Entry{m: m, Entry: ldapEntry}
}

type Entry struct {
	*ldap.Entry
	m map[Field][]string
}

func (e *Entry) Get(field Field) ([]string, bool) {
	values, found := e.m[field]
	return values, found
}

func (e *Entry) GetJoined(field Field) (string, bool) {
	values, found := e.Get(field)
	if !found {
		return "", false
	}
	return strings.Join(values, ","), true
}

func (e *Entry) AsMap() map[Field][]string {
	return e.m
}