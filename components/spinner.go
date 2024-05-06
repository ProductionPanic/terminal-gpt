package components

import (
	"math/rand"

	"github.com/charmbracelet/bubbles/spinner"
)

type SpinnerOptions struct {
}

func RandomSpinner() spinner.Spinner {
	spinners := []spinner.Spinner{
		spinner.Dot,
		spinner.Ellipsis,
		spinner.Line,
		spinner.Globe,
		spinner.Moon,
		spinner.Monkey,
		spinner.Jump,
		spinner.MiniDot,
		spinner.Points,
		spinner.Pulse,
		spinner.Meter,
		spinner.Hamburger,
	}
	return spinners[rand.Intn(len(spinners))]
}

func GenerateSpinner(options *SpinnerOptions) spinner.Model {
	s := spinner.New(spinner.WithSpinner(RandomSpinner()))
	return s
}
