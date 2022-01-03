package modelos

import "time"

// Usuario representacao da entidade no banco de dados
type Usuario struct {
	ID       int64     `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	Status   int       `json:"status,omitempty"`
	CriadoEm time.Time `json:"ciadoEm,omitempty"`
}
