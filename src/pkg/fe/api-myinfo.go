// Tonika: A distributed social networking platform
// Copyright (C) 2010 Petar Maymounkov <petar@5ttt.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.


package fe

import (
	"tonika/http"
	//"tonika/sys"
)

func (fe *FrontEnd) replyAPIMyInfo(args map[string][]string) *http.Response {
	name, ok := args["n"]
	if !ok {
		return newRespBadRequest()
	}
	email, ok := args["e"]
	if !ok {
		return newRespBadRequest()
	}
	addr, ok := args["a"]
	if !ok {
		return newRespBadRequest()
	}
	if name == nil || email == nil || addr == nil {
		return newRespBadRequest()
	}
	var n,e,a string
	if len(name) > 0 {
		n = name[0]
	}
	if len(email) > 0 {
		e = email[0]
	}
	if len(addr) > 0 {
		a = addr[0]
	}

	fe.bank.SetMy("Name", n)
	fe.bank.SetMy("Email", e)
	fe.bank.SetMy("ExtAddr", a)

	fe.bank.Save()

	return buildResp("OK")
}
