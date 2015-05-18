package output

import (
	"testing"
	"time"

	"github.com/registrobr/rdap-client/protocol"
)

var TestAsToTextOutput = `aut-num:     a_123456-NICBR
owner:       (name)
ownerid:     (CPF/CNPJ)
responsible: (name)
address:     
address:     
country:     BR
phone:       
owner-c:     (handle)
routing-c:   (handle)
abuse-c:     (handle)
created:     2015-03-01T12:00:00Z
changed:     2015-03-10T14:00:00Z

inetnum:     (ip networks)


nic-hdl-br: XXXX
person: Joe User
e-mail: joe.user@example.com
address:  Av Naçoes Unidas 11541 7 andar Sao Paulo SP 04578-000 BR
phone: tel:+55-11-5509-3506;ext=3506
created: 2015-03-01T12:00:00Z
changed: 2015-03-10T14:00:00Z

nic-hdl-br: YYYY
person: Joe User 2
e-mail: joe.user2@example.com
address:  Av Naçoes Unidas 11541 7 andar Sao Paulo SP 04578-000 BR
phone: tel:+55-11-5509-3506;ext=3507
created: 2015-03-01T12:00:00Z
changed: 2015-03-10T14:00:00Z
`

type WriterMock struct {
	Content []byte
	Err     error

	MockWrite func(p []byte) (n int, err error)
}

func (w *WriterMock) Write(p []byte) (n int, err error) {
	if w.MockWrite != nil {
		return w.MockWrite(p)
	}

	w.Content = append(w.Content, p...)
	return len(p), w.Err
}

func TestASToText(t *testing.T) {
	asResponse := protocol.ASResponse{
		ObjectClassName: "autnum",
		Handle:          "a_123456-NICBR",
		StartAutnum:     123456,
		EndAutnum:       123456,
		Type:            "DIRECT ALLOCATION",
		Country:         "BR",
		Links: []protocol.Link{
			{
				Value: "https://rdap.registro.br/autnum/123456",
				Rel:   "self",
				Href:  "https://rdap.registro.br/autnum/123456",
				Type:  "application/rdap+json",
			},
		},
		Entities: []protocol.Entity{
			{
				ObjectClassName: "entity",
				Handle:          "XXXX",
				VCardArray: []interface{}{
					"vcard",
					[]interface{}{
						[]interface{}{"version", struct{}{}, "text", "4.0"},
						[]interface{}{"fn", struct{}{}, "text", "Joe User"},
						[]interface{}{"kind", struct{}{}, "text", "individual"},
						[]interface{}{"email", struct{ Type string }{Type: "work"}, "text", "joe.user@example.com"},
						[]interface{}{"lang", struct{ Pref string }{Pref: "1"}, "language-tag", "pt"},
						[]interface{}{"adr", struct{ Type string }{Type: "work"}, "text",
							[]string{
								"Av Naçoes Unidas", "11541", "7 andar", "Sao Paulo", "SP", "04578-000", "BR",
							},
						},
						[]interface{}{"tel", struct{ Type string }{Type: "work"}, "uri", "tel:+55-11-5509-3506;ext=3506"},
					},
				},
				Events: []protocol.Event{
					protocol.Event{Action: protocol.EventActionRegistration, Actor: "", Date: time.Date(2015, 03, 01, 12, 00, 00, 00, time.UTC)},
					protocol.Event{Action: protocol.EventActionLastChanged, Actor: "", Date: time.Date(2015, 03, 10, 14, 00, 00, 00, time.UTC)},
				},
			},
			{
				ObjectClassName: "entity",
				Handle:          "YYYY",
				VCardArray: []interface{}{
					"vcard",
					[]interface{}{
						[]interface{}{"version", struct{}{}, "text", "4.0"},
						[]interface{}{"fn", struct{}{}, "text", "Joe User 2"},
						[]interface{}{"kind", struct{}{}, "text", "individual"},
						[]interface{}{"email", struct{ Type string }{Type: "work"}, "text", "joe.user2@example.com"},
						[]interface{}{"lang", struct{ Pref string }{Pref: "1"}, "language-tag", "pt"},
						[]interface{}{"adr", struct{ Type string }{Type: "work"}, "text",
							[]string{
								"Av Naçoes Unidas", "11541", "7 andar", "Sao Paulo", "SP", "04578-000", "BR",
							},
						},
						[]interface{}{"tel", struct{ Type string }{Type: "work"}, "uri", "tel:+55-11-5509-3506;ext=3507"},
					},
				},
				Events: []protocol.Event{
					protocol.Event{Action: protocol.EventActionRegistration, Actor: "", Date: time.Date(2015, 03, 01, 12, 00, 00, 00, time.UTC)},
					protocol.Event{Action: protocol.EventActionLastChanged, Actor: "", Date: time.Date(2015, 03, 10, 14, 00, 00, 00, time.UTC)},
				},
			},
		},
		Events: []protocol.Event{
			protocol.Event{Action: protocol.EventActionRegistration, Actor: "", Date: time.Date(2015, 03, 01, 12, 00, 00, 00, time.UTC)},
			protocol.Event{Action: protocol.EventActionLastChanged, Actor: "", Date: time.Date(2015, 03, 10, 14, 00, 00, 00, time.UTC)},
		},
	}

	asOutput := AS{AS: &asResponse}

	var w WriterMock
	if err := asOutput.ToText(&w); err != nil {
		t.Fatal(err)
	}

	if string(w.Content) != TestAsToTextOutput {
		t.Error("Wrong output")
		t.Log(string(w.Content))
		return
	}
}
