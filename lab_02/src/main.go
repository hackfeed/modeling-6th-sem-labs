package main

import (
	"fmt"
	"lab_02/circuit"
)

func main() {
	var (
		I      = circuit.Params.I0
		Uc     = circuit.Params.Uc0
		h      = 1e-6
		IRes   circuit.FArr64
		RpRes  circuit.FArr64
		UcRes  circuit.FArr64
		T0Res  circuit.FArr64
		IRpRes circuit.FArr64
		tRes   circuit.FArr64
	)

	for _, t := range circuit.Arange(0, 0.0008, h) {
		T0 := circuit.GetT0(I)
		Rp := circuit.GetRp(I, T0, circuit.GetM(I))
		I, Uc = circuit.GetRungeKutta(t, I, Uc, h, Rp)

		if t > h {
			tRes = append(tRes, t)
			IRes = append(IRes, I)
			RpRes = append(RpRes, Rp)
			UcRes = append(UcRes, Uc)
			T0Res = append(T0Res, T0)
			IRpRes = append(IRpRes, I*Rp)
		}

		fmt.Printf("-- DEBUG -- Rp: %v -- I: %v -- Uc -- %v T0: %v -- Rk: %v --\n", Rp, I, Uc, T0, circuit.Params.Rk)
	}

	circuit.DrawPlot(tRes, IRes, "I(t)", "t", "I", "data/it.png")
	circuit.DrawPlot(tRes, UcRes, "U(t)", "t", "U", "data/ut.png")
	circuit.DrawPlot(tRes, RpRes, "Rp(t)", "t", "Rp", "data/rpt.png")
	circuit.DrawPlot(tRes, T0Res, "T0(t)", "t", "T0", "data/t0t.png")
	circuit.DrawPlot(tRes, IRpRes, "I(t) * Rp(t)", "t", "I * Rp", "data/irpt.png")
}
