package view

type UserRenderer struct {

}

func( r *UserRenderer ) Login() string {
    return react.Run("Layout", map[string]string{}, []string{`
        `, react.Run("form", map[string]string{}, []string{`
        `, react.Run("label", map[string]string{}, []string{`Username :`}), `
        `, react.Run("input", map[string]string{`type` :`text`}, []string{}), `
        `}), `		
	`})
}