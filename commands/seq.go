package commands

import (
	"fmt"
	"github.com/bep/gognutil/helpers"
	"github.com/bep/gognutil/numbers"
	"github.com/spf13/cobra"
)

type seqFlagsType struct {
	format, separator string
	equalWidth        bool
}

var seqFlags seqFlagsType

func seqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `seq <first> [<increment>] [<last>]`,
		Short: "Print a sequence of numbers",
		Long: `Print a sequence of numbers
		
This works very similar to the GNU counterpart, with some subtle differences:

* Go's printf format is slightly different
* GNU's seq just prints blank when one negative argument is provided.

EXAMPLE:

gognutil seq -w 10
01
02
03
04
05
06
07
08
09
10


NOTE: With negative arguments, the flags must be terminated with '--', e.g "seq -w -- 10 -2 1"
		`,
		Run: func(cmd *cobra.Command, args []string) {
			ns, err := numbers.Seq(args)
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(ns) == 0 {
				return
			}
			l := len(ns) - 1
			md := helpers.NumDigits(helpers.MaxInt(ns[0], ns[l]))
			for i, r := range ns {
				fmt.Print(fmtSeqItem(r, md, i == l))
			}
			fmt.Println("")
		},
	}
	cmd.Flags().StringVarP(&seqFlags.format, "format", "f", "", "Go printf style floating-point format")
	cmd.Flags().StringVarP(&seqFlags.separator, "separator", "s", "\n", "used to separate numbers")
	cmd.Flags().BoolVarP(&seqFlags.equalWidth, "equal-width", "w", false, "equalize width by padding with leading zeroes")
	return cmd
}

func fmtSeqItem(i, md int, last bool) string {
	var formatted string

	if seqFlags.format != "" {
		formatted = fmt.Sprintf(seqFlags.format, i)
	} else {
		var format string
		if seqFlags.equalWidth {
			format = "%" + fmt.Sprintf("0%dd", md)
		} else {
			format = "%d"
		}
		formatted = fmt.Sprintf(format, i)
	}

	if !last {
		formatted += seqFlags.separator
	}
	return formatted
}
