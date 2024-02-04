# ftk
Forestry specific formulas
## Install
`go get github.com/Opeezy/ftk@latest`

## Functions
### Basal Area (per tree)
`BA(dbh float64, denominator int, plotnum int, round int) (n float32, err error)`<br>

#### Params
`dbh - Diameter breast height of tree in inches`<br>
`denominator - Denominator of the acre size (e.g., 20 if acre size is 1/20)`<br>
`plotnum - number of plots`<br>
`round - Integer value to determine rounding. (0 = Round down, 1 = Round up and 2 = No rounding)`<br>
