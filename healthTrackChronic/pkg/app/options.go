package app

import cliflag "github.com/maxiaolu1981/component-base/pkg/cli/flag"

type CliOptions interface {
	Flags() (fss cliflag.NamedFlagSets)
	Validate() []error
}
