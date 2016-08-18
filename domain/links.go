package domain

/*
URLTuple Is a wrapper for (url,title) tuple
*/
type URLTuple struct {
	Link  string `json:"url"`
	Title string `json:"title"`
}

/*
URLTupleArray is a container for arrays of URLTuple tuples
*/
type URLTupleArray struct {
	URLs []URLTuple `json:"links"`
}
