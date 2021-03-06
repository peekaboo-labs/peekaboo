package user

import (
	"fmt"

	"github.com/ake-persson/peekaboo/pkg/pb/v1/resources"
)

var Headers = []string{
	"username",
	"comment",
	"uid",
	"gid",
	"uid_signed",
	"gid_signed",
	"directory",
	"shell",
}

func StringSlice(u *resources.User) []string {
	return []string{
		u.Username,
		u.Comment,
		fmt.Sprint(u.Uid),
		fmt.Sprint(u.Gid),
		fmt.Sprint(u.UidSigned),
		fmt.Sprint(u.GidSigned),
		fmt.Sprint(u.Directory),
		fmt.Sprint(u.Shell),
	}
}
