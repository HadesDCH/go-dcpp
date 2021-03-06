// Code generated by hublist_test.go. DO NOT EDIT.

package hublist

type Hub struct {
	Address     string `xml:"Address,attr"`
	Bots        int    `xml:"Bots,attr"`
	Country     string `xml:"Country,attr"`
	Description string `xml:"Description,attr"`
	Email       string `xml:"Email,attr"`
	Encoding    string `xml:"Encoding,attr"`
	Failover    string `xml:"Failover,attr"`
	Icon        string `xml:"Icon,attr"`
	Infected    int    `xml:"Infected,attr"`
	Logo        string `xml:"Logo,attr"`
	Maxhubs     int    `xml:"Maxhubs,attr"`
	Maxusers    int    `xml:"Maxusers,attr"`
	Minshare    uint64 `xml:"Minshare,attr"`
	Minslots    int    `xml:"Minslots,attr"`
	Name        string `xml:"Name,attr"`
	Network     string `xml:"Network,attr"`
	Operators   int    `xml:"Operators,attr"`
	Rating      string `xml:"Rating,attr"`
	Reliability string `xml:"Reliability,attr"`
	Shared      uint64 `xml:"Shared,attr"`
	Software    string `xml:"Software,attr"`
	Status      string `xml:"Status,attr"`
	Users       int    `xml:"Users,attr"`
	Website     string `xml:"Website,attr"`
}
