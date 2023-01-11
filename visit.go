package track

import (
	cond "github.com/vela-security/vela-cond"
	"github.com/vela-security/vela-process"
)

func ByPid(pid int32, cnd *cond.Cond) *track {
	return newTrackByPid(pid, cnd)
}

func ByProcess(p *process.Process, cnd *cond.Cond) *track {
	return newTrackByPid(int32(p.Pid), cnd)
}

func ByName(name string, cnd *cond.Cond) *tracks {
	return newTrackByName(name, cnd)
}
