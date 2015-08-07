package crowbar

import (
	"errors"
	"log"
	"strconv"
)

type DnsNameFilter struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Matcher   string `json:"matcher"`
	Priority  int64  `json:"priority"`
	Service   string `json:"service"`
	Template  string `json:"template"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	lastJson  []byte
}

func (o *DnsNameFilter) Id() string {
	if o.ID != 0 {
		return strconv.FormatInt(o.ID, 10)
	} else if o.Name != "" {
		return o.Name
	} else {
		log.Panic("DnsNameFilter has no ID or name")
		return ""
	}
}

func (o *DnsNameFilter) SetId(s string) error {
	if o.ID != 0 || o.Name != "" {
		return errors.New("SetId can only be used on an un-IDed object")
	}
	if id, err := strconv.ParseInt(s, 10, 64); err == nil {
		o.ID = id
	} else {
		o.Name = s
	}
	return nil
}

func (o *DnsNameFilter) ApiName() string {
	return "dns_name_filters"
}

func (o *DnsNameFilter) setLastJSON(b []byte) {
	o.lastJson = make([]byte, len(b))
	copy(o.lastJson, b)
}

func (o *DnsNameFilter) lastJSON() []byte {
	return o.lastJson
}

func DnsNameFilters() (res []*DnsNameFilter, err error) {
	res = make([]*DnsNameFilter, 0)
	return res, session.list(&res, "dns_name_filters")
}
