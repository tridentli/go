// Copyright 2012, Jeramey Crawford <jeramey@antihe.ro>
// Copyright 2013, Jonas mg
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// Package apr1_crypt implements the standard Unix MD5-crypt algorithm created
// by Poul-Henning Kamp for FreeBSD, and modified by the Apache project.
//
// The only change from MD5-crypt is the use of the magic constant "$apr1$"
// instead of "$1$". The algorithms are otherwise identical.
package apr1_crypt

import (
	"../common"
	"../md5_crypt"
)

const (
	MagicPrefix   = "$apr1$"
	SaltLenMin    = 1
	SaltLenMax    = 8
	RoundsDefault = 1000
)

var md5Crypt = md5_crypt.New()

func init() {
	md5Crypt.SetSalt(crypt.Salt{
		MagicPrefix:   []byte(MagicPrefix),
		SaltLenMin:    SaltLenMin,
		SaltLenMax:    SaltLenMax,
		RoundsDefault: RoundsDefault,
	})
}

type crypter struct{ Salt crypt.Salt }

// New returns a new crypt.Crypter computing the variant "apr1" of MD5-crypt
func New() crypt.Crypter { return &crypter{crypt.Salt{}} }

func (c *crypter) Generate(key, salt []byte) (string, error) {
	return md5Crypt.Generate(key, salt)
}

func (c *crypter) Verify(hashedKey string, key []byte) error {
	return md5Crypt.Verify(hashedKey, key)
}

func (c *crypter) Cost(hashedKey string) (int, error) { return RoundsDefault, nil }

func (c *crypter) SetSalt(salt crypt.Salt) {}
