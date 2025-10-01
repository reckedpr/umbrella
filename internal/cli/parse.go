package cli

import (
	"github.com/alexflint/go-arg"
	"github.com/reckedpr/umbrella/internal/model"
)

func ParseArgs() model.Args {
	args := model.Args{}
	arg.MustParse(&args)

	if args.Units != "f" && args.Units != "c" {
		panic("Units must be 'c' or 'f'")
	}

	return args
}
