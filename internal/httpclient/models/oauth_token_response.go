// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OauthTokenResponse OauthTokenResponse The token response
// swagger:model oauthTokenResponse
type OauthTokenResponse struct {

	// The access token issued by the authorization server.
	AccessToken string `json:"access_token,omitempty"`

	// The lifetime in seconds of the access token.  For
	// example, the value "3600" denotes that the access token will
	// expire in one hour from the time the response was generated.
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// To retrieve a refresh token request the id_token scope.
	IDToken int64 `json:"id_token,omitempty"`

	// The refresh token, which can be used to obtain new
	// access tokens. To retrieve it add the scope "offline" to your access token request.
	RefreshToken string `json:"refresh_token,omitempty"`

	// The scope of the access token
	Scope int64 `json:"scope,omitempty"`

	// The type of the token issued
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this oauth token response
func (m *OauthTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OauthTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OauthTokenResponse) UnmarshalBinary(b []byte) error {
	var res OauthTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
