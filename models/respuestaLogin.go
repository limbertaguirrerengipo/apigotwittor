package models

/*RespuestaLogin tiene el token */
type RespuestaLogin struct {
	Token string `bson:"token,omitempty" json:"token,omitempty"`
}
