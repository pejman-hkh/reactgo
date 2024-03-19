package view

type IndexRenderer struct {

}

func( r *IndexRenderer ) Index() string {
    return react.Run("Layout", map[string]string{}, []string{`
       Home Page		
	`})
}