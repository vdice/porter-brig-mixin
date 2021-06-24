package brig

import (
	"get.porter.sh/porter/pkg/mixin"
	"get.porter.sh/porter/pkg/pkgmgmt"
	"get.porter.sh/porter/pkg/porter/version"
	"github.com/vdice/porter-brig-mixin/pkg"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "brig",
		VersionInfo: pkgmgmt.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "Vaughn Dice",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
