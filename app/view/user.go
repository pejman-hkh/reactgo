package view

type UserRenderer struct {

}

func( r *UserRenderer) Register() string {
    return react.Run("Layout", map[string]string{}, []string{`

    `})
}

func (r Component)LoginForm(props map[string]string, childrens string) string {
return       react.Run("form", map[string]string{`class` :`space-y-6`,`action` :`#`,`method` :`POST`}, []string{`
    `, react.Run("div", map[string]string{}, []string{`
      `, react.Run("label", map[string]string{`for` :`email`,`class` :`block text-sm font-medium leading-6 text-gray-900`}, []string{`Email address`}), `
      `, react.Run("div", map[string]string{`class` :`mt-2`}, []string{`
        `, react.Run("input", map[string]string{`id` :`email`,`name` :`email`,`type` :`email`,`autocomplete` :`email`,`required class` :`block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6`}, []string{}), `
      `}), `
    `}), `

    `, react.Run("div", map[string]string{}, []string{`
      `, react.Run("div", map[string]string{`class` :`flex items-center justify-between`}, []string{`
        `, react.Run("label", map[string]string{`for` :`password`,`class` :`block text-sm font-medium leading-6 text-gray-900`}, []string{`Password`}), `
        `, react.Run("div", map[string]string{`class` :`text-sm`}, []string{`
          `, react.Run("a", map[string]string{`href` :`#`,`class` :`font-semibold text-indigo-600 hover:text-indigo-500`}, []string{`Forgot password?`}), `
        `}), `
      `}), `
      `, react.Run("div", map[string]string{`class` :`mt-2`}, []string{`
        `, react.Run("input", map[string]string{`id` :`password`,`name` :`password`,`type` :`password`,`autocomplete` :`current-password`,`required class` :`block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6`}, []string{}), `
      `}), `
    `}), `

    `, react.Run("div", map[string]string{}, []string{`
      `, react.Run("button", map[string]string{`type` :`submit`,`class` :`flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600`}, []string{`Sign in`}), `
    `}), `
  `})
}

func( r *UserRenderer ) Login() string {
    return react.Run("Layout", map[string]string{}, []string{`

`, react.Run("div", map[string]string{`class` :`flex min-h-full flex-col justify-center px-6 py-12 lg:px-8`}, []string{`
    `, react.Run("div", map[string]string{`class` :`sm:mx-auto sm:w-full sm:max-w-sm`}, []string{`
      `, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600`,`alt` :`Your Company`,`class` :`mx-auto h-10 w-auto`}, []string{}), `
      `, react.Run("h2", map[string]string{`class` :`mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900`}, []string{`Sign in to your account`}), `
    `}), `
  
    `, react.Run("div", map[string]string{`class` :`mt-10 sm:mx-auto sm:w-full sm:max-w-sm`}, []string{`
        `, react.Run("LoginForm", map[string]string{}, []string{}), `
      `, react.Run("p", map[string]string{`class` :`mt-10 text-center text-sm text-gray-500`}, []string{`
        Not a member?
        `, react.Run("a", map[string]string{`href` :`#`,`class` :`font-semibold leading-6 text-indigo-600 hover:text-indigo-500`}, []string{`Start a 14 day free trial`}), `
      `}), `
    `}), `
  `}), `
      
	`})
}