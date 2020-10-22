package States

type States struct{
	Wait     bool
	Question bool
}

type Info struct{
	Trials int
	WordsToIgnore []string
}