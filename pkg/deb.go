// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package pkg

import "github.com/kless/osutil"

type deb packageSystem

func (p deb) Install(name ...string) error {
	args := []string{"install", "-y"}

	return osutil.ExecSudo("/usr/bin/apt-get", append(args, name...)...)
}

func (p deb) Remove(name ...string) error {
	args := []string{"remove", "-y"}

	return osutil.ExecSudo("/usr/bin/apt-get", append(args, name...)...)
}

func (p deb) RemoveMeta(name ...string) error {
	if err := p.Remove(name...); err != nil {
		return err
	}
	return osutil.ExecSudo("/usr/bin/apt-get", "autoremove", "-y")
}

func (p deb) Purge(name ...string) error {
	args := []string{"purge", "-y"}

	return osutil.ExecSudo("/usr/bin/apt-get", append(args, name...)...)
}

func (p deb) PurgeMeta(name ...string) error {
	if err := p.Purge(name...); err != nil {
		return err
	}
	return osutil.ExecSudo("/usr/bin/apt-get", "autoremove", "--purge", "-y")
}

func (p deb) Update() error {
	return osutil.ExecSudo("/usr/bin/apt-get", "update")
}

func (p deb) Upgrade() error {
	return osutil.ExecSudo("/usr/bin/apt-get", "upgrade")
}

func (p deb) Clean() error {
	return osutil.ExecSudo("/usr/bin/apt-get", "clean")
}
