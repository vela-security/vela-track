package track

import (
	cond "github.com/vela-security/vela-cond"
	"github.com/vela-security/vela-process"
	"github.com/vela-security/vela-public/lua"
)

func (tks *tracks) scan() {
	list := process.List()

	for _, v := range list {
		pid := int32(v)
		tk := newTrack(withPid(pid), withCnd(tks.cnd), withOption(tks.opt))
		tk.lookup()
		tk.Pid()
		if tk.ok() {
			tks.append(tk.data...)
		} else {
			xEnv.Infof("vela track pid:%d fail %v", pid, tk.cause.Error())
		}
	}
}

func newLuaTrackALL(L *lua.LState) *tracks {
	tka := &tracks{cnd: cond.CheckMany(L)}
	tka.scan()
	return tka
}
