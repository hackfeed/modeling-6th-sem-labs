package main

import (
	"lab_04/emission"
)

func main() {
	{
		emission.Params = emission.Emission{
			0.0134, 1, 4.35e-4, 1, 2.049, 0.563e-3, 0.528e5,
			1, 0.05, 0.01, 2, 300, 0.5, 50, 1e-3, 1, 1e-2,
		}
		res, ti := emission.SimpleIters()
		x := emission.Arange(0, emission.Params.L, emission.Params.H)
		ptsx := emission.FMat64{}
		ptsy := emission.FMat64{}

		for i, v := range res {
			if i%2 == 0 {
				ptsx = append(ptsx, x)
				ptsy = append(ptsy, v[:len(v)-1])
			}
		}
		ptsx = append(ptsx, x)
		ptsy = append(ptsy, res[len(res)-1][:len(res[0])-1])

		emission.DrawPlot(ptsx, ptsy, "T(x)", "x", "T", "data/tx.png")

		t := emission.Arange(0, ti, emission.Params.T)
		s := emission.Arange(0, emission.Params.L, 0.05)
		ptsx = emission.FMat64{}
		ptsy = emission.FMat64{}
		for _, v := range s {
			r := emission.FArr64{}
			for _, vv := range res {
				r = append(r, vv[int(v/emission.Params.H)])
			}
			ptsx = append(ptsx, t)
			ptsy = append(ptsy, r[:len(r)-1])
		}

		emission.DrawPlot(ptsx, ptsy, "T(t)", "t", "T", "data/tx1.png")
	}
	{
		emission.Params = emission.Emission{
			0.0134, 1, 4.35e-4, 1, 2.049, 0.563e-3, 0.528e5,
			1, 0.05, 0.01, 2, 300, 0.5, -9, 1e-3, 1, 1e-2,
		}
		res, ti := emission.SimpleIters()
		x := emission.Arange(0, emission.Params.L, emission.Params.H)
		ptsx := emission.FMat64{}
		ptsy := emission.FMat64{}

		for i, v := range res {
			if i%2 == 0 {
				ptsx = append(ptsx, x)
				ptsy = append(ptsy, v[:len(v)-1])
			}
		}
		ptsx = append(ptsx, x)
		ptsy = append(ptsy, res[len(res)-1][:len(res[0])-1])

		emission.DrawPlot(ptsx, ptsy, "T(x)", "x", "T", "data/tx2.png")

		t := emission.Arange(0, ti, emission.Params.T)
		s := emission.Arange(0, emission.Params.L, 0.05)
		ptsx = emission.FMat64{}
		ptsy = emission.FMat64{}
		for _, v := range s {
			r := emission.FArr64{}
			for _, vv := range res {
				r = append(r, vv[int(v/emission.Params.H)])
			}
			ptsx = append(ptsx, t)
			ptsy = append(ptsy, r[:len(r)-1])
		}

		emission.DrawPlot(ptsx, ptsy, "T(t)", "t", "T", "data/tx3.png")
	}
	{
		emission.Params = emission.Emission{
			0.0134, 1, 4.35e-4, 1, 2.049, 0.563e-3, 0.528e5,
			1, 0.05, 0.01, 2, 300, 0.5, 0, 1e-3, 1, 1e-2,
		}
		res, _ := emission.SimpleIters()
		x := emission.Arange(0, emission.Params.L, emission.Params.H)
		ptsx := emission.FMat64{}
		ptsy := emission.FMat64{}

		for i, v := range res {
			if i%2 == 0 {
				ptsx = append(ptsx, x)
				ptsy = append(ptsy, v[:len(v)-1])
			}
		}
		ptsx = append(ptsx, x)
		ptsy = append(ptsy, res[len(res)-1][:len(res[0])-1])

		emission.DrawPlot(ptsx, ptsy, "T(x)", "x", "T", "data/tx4.png")
	}
}
