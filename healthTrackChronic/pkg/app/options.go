package app

import (
	cliflag "github.com/maxiaolu1981/base/component-base/pkg/cli/flag"
)

type CliOptions interface {
	Flags() (fss cliflag.NamedFlagSets)
	Validate() []error
}

// CompleteableOptions abstracts options which can be completed.
type CompleteableOptions interface {
	Complete() error
}

type PrintableOptions interface {
	String() string
}
