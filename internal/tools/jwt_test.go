package tools

import "testing"

func TestValidateAndParseJwtToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImVhODZiMGZhLTk0ODQtNGI1MC1hNjQxLWIzZTkyZTdiZTMzMCIsImlzcyI6ImtvYnJhIiwic3ViIjoidXNlcnMiLCJhdWQiOlsiY3VpaGU1MDAiXSwiZXhwIjoxNzEyNjc5MTc3LCJuYmYiOjE3MTI2NjgzNzcsImlhdCI6MTcxMjY2ODM3N30.hVnPkWfdI8z9B1yKJkIMCk_Z1EodoeQGnlTrAM5oo1w"
	jwtToken, err := ValidateAndParseJwtToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(jwtToken.UUID)
}
