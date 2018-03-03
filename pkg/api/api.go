package api

type Guest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Admin struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Smtp struct {
	Server string `json:"server,omitempty"`
	Port   string `json:"port,omitempty"`
}

type Config struct {
	Guests  []Guest  `json:"guests,omitempty"`
	Words   []string `json:"words,omitempty"`
	Admin   Admin    `json:"admin,omitempty"`
	Smtp    Smtp     `json:"smtp,omitempty"`
	Body    string   `json:"body,omitempty"`
	Subject string   `json:"subject,omitempty"`
}

type Logger interface {
	Info(m string)
	Debug(m string)
	Error(m string)
}

type Sender interface {
	Send(body string) error
}

type Reader interface {
	Read(filename string) []byte
}
