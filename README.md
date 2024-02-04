# ftk
Forestry specific formulas
## Install
`go get github.com/Opeezy/ftk@latest`

## Functions
### Basal Area (per tree)
```GO
BA(dbh float64, denominator int, plotnum int, round int) (n float32, err error)
```

#### Params
`dbh - Diameter breast height of tree in inches`<br><br>
`denominator - Denominator of the acre size (e.g., 20 if acre size is 1/20)`<br><br>
`plotnum - number of plots`<br><br>
`round - Integer value to determine rounding. (0 = Round down, 1 = Round up and 2 = No rounding)`<br><br>

### Example
```GO
    //DBH = 14.5
    //Denominator = 20
    //# of plots = 2
	n, err := ftk.TPA(14.5, 20, 2, 1)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(n)
	}
```

