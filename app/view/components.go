package view

import (
"github.com/pejman-hkh/gdp/gox"
)

type Component struct {}
var react gox.Gox = gox.Gox{Component{}}

func (r Component) Link(props map[string]string, childrens string) string {
return react.Run("a", map[string]string{`href` :props["to"]}, []string{``,childrens})
}

func (r Component) SideNav(props map[string]string, childrens string) string {
return react.Run("nav", map[string]string{}, []string{``,childrens})
}

func (r Component) Content(props map[string]string, childrens string) string {
return react.Run("main", map[string]string{}, []string{`
	`, react.Run("div", map[string]string{`class` :`mx-auto max-w-7xl py-6 sm:px-6 lg:px-8`}, []string{`		`,childrens}), `
`})
}

func (r Component) Header(props map[string]string, childrens string) string {
return react.Run("f", map[string]string{}, []string{`
	`, react.Run("header", map[string]string{}, []string{`
		`, react.Run("div", map[string]string{`class` :``}, []string{`
			`, react.Run("nav", map[string]string{`class` :`bg-gray-800`}, []string{`
				`, react.Run("div", map[string]string{`class` :`mx-auto max-w-7xl px-4 sm:px-6 lg:px-8`}, []string{`
					`, react.Run("div", map[string]string{`class` :`flex h-16 items-center justify-between`}, []string{`
						`, react.Run("div", map[string]string{`class` :`flex items-center`}, []string{`
							`, react.Run("div", map[string]string{`class` :`flex-shrink-0`}, []string{`
								`, react.Run("img", map[string]string{`class` :`h-8 w-8`,`src` :`https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500`,`alt` :`Your Company`}, []string{}), `
							`}), `
							`, react.Run("div", map[string]string{`class` :`hidden md:block`}, []string{`
								`, react.Run("div", map[string]string{`class` :`ml-10 flex items-baseline space-x-4`}, []string{`
									`, react.Run("comment", map[string]string{}, []string{}), `
									`, react.Run("a", map[string]string{`href` :`#`,`class` :`bg-gray-900 text-white rounded-md px-3 py-2 text-sm font-medium`,`aria-current` :`page`}, []string{`Dashboard`}), `
									`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium`}, []string{`Team`}), `
									`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium`}, []string{`Projects`}), `
									`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium`}, []string{`Calendar`}), `
									`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium`}, []string{`Reports`}), `
								`}), `
							`}), `
						`}), `
						`, react.Run("div", map[string]string{`class` :`hidden md:block`}, []string{`
							`, react.Run("div", map[string]string{`class` :`ml-4 flex items-center md:ml-6`}, []string{`
								`, react.Run("button", map[string]string{`type` :`button`,`class` :`relative rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800`}, []string{`
									`, react.Run("span", map[string]string{`class` :`absolute -inset-1.5`}, []string{}), `
									`, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`View notifications`}), `
									`, react.Run("svg", map[string]string{`viewBox` :`0 0 24 24`,`stroke-width` :`1.5`,`stroke` :`currentColor`,`aria-hidden` :`true`,`class` :`h-6 w-6`,`fill` :`none`}, []string{`
										`, react.Run("path", map[string]string{`stroke-linecap` :`round`,`stroke-linejoin` :`round`,`d` :`M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0`}, []string{}), `
									`}), `
								`}), `

								`, react.Run("comment", map[string]string{}, []string{}), `
								`, react.Run("div", map[string]string{`class` :`relative ml-3`}, []string{`
									`, react.Run("div", map[string]string{}, []string{`
										`, react.Run("button", map[string]string{`class` :`relative flex max-w-xs items-center rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800`,`id` :`user-menu-button`,`aria-expanded` :`false`,`aria-haspopup` :`true`,`type` :`button`}, []string{`
											`, react.Run("span", map[string]string{`class` :`absolute -inset-1.5`}, []string{}), `
											`, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`Open user menu`}), `
											`, react.Run("img", map[string]string{`class` :`h-8 w-8 rounded-full`,`src` :`https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80`,`alt` :``}, []string{}), `
										`}), `
									`}), `


									`, react.Run("div", map[string]string{`class` :`absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none`,`role` :`menu`,`aria-orientation` :`vertical`,`aria-labelledby` :`user-menu-button`,`tabindex` :`-1`}, []string{`

										`, react.Run("a", map[string]string{`href` :`#`,`class` :`block px-4 py-2 text-sm text-gray-700`,`role` :`menuitem`,`tabindex` :`-1`,`id` :`user-menu-item-0`}, []string{`Your Profile`}), `
										`, react.Run("a", map[string]string{`tabindex` :`-1`,`id` :`user-menu-item-1`,`href` :`#`,`class` :`block px-4 py-2 text-sm text-gray-700`,`role` :`menuitem`}, []string{`Settings`}), `
										`, react.Run("a", map[string]string{`href` :`#`,`class` :`block px-4 py-2 text-sm text-gray-700`,`role` :`menuitem`,`tabindex` :`-1`,`id` :`user-menu-item-2`}, []string{`Sign out`}), `
									`}), `
								`}), `
							`}), `
						`}), `
						`, react.Run("div", map[string]string{`class` :`-mr-2 flex md:hidden`}, []string{`

							`, react.Run("button", map[string]string{`type` :`button`,`class` :`relative inline-flex items-center justify-center rounded-md bg-gray-800 p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800`,`aria-controls` :`mobile-menu`,`aria-expanded` :`false`}, []string{`
								`, react.Run("span", map[string]string{`class` :`absolute -inset-0.5`}, []string{}), `
								`, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`Open main menu`}), `

								`, react.Run("svg", map[string]string{`class` :`block h-6 w-6`,`fill` :`none`,`viewBox` :`0 0 24 24`,`stroke-width` :`1.5`,`stroke` :`currentColor`,`aria-hidden` :`true`}, []string{`
									`, react.Run("path", map[string]string{`stroke-linecap` :`round`,`stroke-linejoin` :`round`,`d` :`M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5`}, []string{}), `
								`}), `

								`, react.Run("svg", map[string]string{`class` :`hidden h-6 w-6`,`fill` :`none`,`viewBox` :`0 0 24 24`,`stroke-width` :`1.5`,`stroke` :`currentColor`,`aria-hidden` :`true`}, []string{`
									`, react.Run("path", map[string]string{`stroke-linecap` :`round`,`stroke-linejoin` :`round`,`d` :`M6 18L18 6M6 6l12 12`}, []string{}), `
								`}), `
							`}), `
						`}), `
					`}), `
				`}), `


				`, react.Run("div", map[string]string{`class` :`md:hidden`,`id` :`mobile-menu`}, []string{`
					`, react.Run("div", map[string]string{`class` :`space-y-1 px-2 pb-3 pt-2 sm:px-3`}, []string{`

						`, react.Run("a", map[string]string{`href` :`#`,`class` :`bg-gray-900 text-white block rounded-md px-3 py-2 text-base font-medium`,`aria-current` :`page`}, []string{`Dashboard`}), `
						`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium`}, []string{`Team`}), `
						`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium`}, []string{`Projects`}), `
						`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium`}, []string{`Calendar`}), `
						`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium`}, []string{`Reports`}), `
					`}), `
					`, react.Run("div", map[string]string{`class` :`border-t border-gray-700 pb-3 pt-4`}, []string{`
						`, react.Run("div", map[string]string{`class` :`flex items-center px-5`}, []string{`
							`, react.Run("div", map[string]string{`class` :`flex-shrink-0`}, []string{`
								`, react.Run("img", map[string]string{`class` :`h-10 w-10 rounded-full`,`src` :`https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80`,`alt` :``}, []string{}), `
							`}), `
							`, react.Run("div", map[string]string{`class` :`ml-3`}, []string{`
								`, react.Run("div", map[string]string{`class` :`text-base font-medium leading-none text-white`}, []string{`Tom Cook`}), `
								`, react.Run("div", map[string]string{`class` :`text-sm font-medium leading-none text-gray-400`}, []string{`tom@example.com`}), `
							`}), `
							`, react.Run("button", map[string]string{`type` :`button`,`class` :`relative ml-auto flex-shrink-0 rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800`}, []string{`
								`, react.Run("span", map[string]string{`class` :`absolute -inset-1.5`}, []string{}), `
								`, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`View notifications`}), `
								`, react.Run("svg", map[string]string{`class` :`h-6 w-6`,`fill` :`none`,`viewBox` :`0 0 24 24`,`stroke-width` :`1.5`,`stroke` :`currentColor`,`aria-hidden` :`true`}, []string{`
									`, react.Run("path", map[string]string{`stroke-linecap` :`round`,`stroke-linejoin` :`round`,`d` :`M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0`}, []string{}), `
								`}), `
							`}), `
						`}), `
						`, react.Run("div", map[string]string{`class` :`mt-3 space-y-1 px-2`}, []string{`
							`, react.Run("a", map[string]string{`href` :`#`,`class` :`block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white`}, []string{`Your
								Profile`}), `
							`, react.Run("a", map[string]string{`href` :`#`,`class` :`block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white`}, []string{`Settings`}), `
							`, react.Run("a", map[string]string{`href` :`#`,`class` :`block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white`}, []string{`Sign
								out`}), `
						`}), `
					`}), `
				`}), `
			`}), `

			`, react.Run("header", map[string]string{`class` :`bg-white shadow`}, []string{`
				`, react.Run("div", map[string]string{`class` :`mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8`}, []string{`
					`, react.Run("h1", map[string]string{`class` :`text-3xl font-bold tracking-tight text-gray-900`}, []string{`Dashboard`}), `
				`}), `
			`}), `

		`}), `		`,childrens}), `
	`, react.Run("h1", map[string]string{}, []string{``,props["title"]}), `
`})
}

func (r Component) Footer(props map[string]string, childrens string) string {
return react.Run("footer", map[string]string{}, []string{`
	`, react.Run("comment", map[string]string{}, []string{}), `
`, react.Run("div", map[string]string{`class` :`bg-white`}, []string{`
	`, react.Run("comment", map[string]string{}, []string{}), `
	`, react.Run("div", map[string]string{`role` :`dialog`,`aria-modal` :`true`,`class` :`relative z-40 lg:hidden`}, []string{`
	  `, react.Run("comment", map[string]string{}, []string{}), `
	  `, react.Run("div", map[string]string{`class` :`fixed inset-0 bg-black bg-opacity-25`}, []string{}), `
  
	  `, react.Run("div", map[string]string{`class` :`fixed inset-0 z-40 flex`}, []string{`
		`, react.Run("comment", map[string]string{}, []string{}), `
		`, react.Run("div", map[string]string{`class` :`relative flex w-full max-w-xs flex-col overflow-y-auto bg-white pb-12 shadow-xl`}, []string{`
		  `, react.Run("div", map[string]string{`class` :`flex px-4 pb-2 pt-5`}, []string{`
			`, react.Run("button", map[string]string{`type` :`button`,`class` :`relative -m-2 inline-flex items-center justify-center rounded-md p-2 text-gray-400`}, []string{`
			  `, react.Run("span", map[string]string{`class` :`absolute -inset-0.5`}, []string{}), `
			  `, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`Close menu`}), `
			  `, react.Run("svg", map[string]string{`class` :`h-6 w-6`,`fill` :`none`,`viewBox` :`0 0 24 24`,`stroke-width` :`1.5`,`stroke` :`currentColor`,`aria-hidden` :`true`}, []string{`
				`, react.Run("path", map[string]string{`stroke-linecap` :`round`,`stroke-linejoin` :`round`,`d` :`M6 18L18 6M6 6l12 12`}, []string{}), `
			  `}), `
			`}), `
		  `}), `
  
		  `, react.Run("comment", map[string]string{}, []string{}), `
		  `, react.Run("div", map[string]string{`class` :`mt-2`}, []string{`
			`, react.Run("div", map[string]string{`class` :`border-b border-gray-200`}, []string{`
			  `, react.Run("div", map[string]string{`role` :`tablist`,`class` :`-mb-px flex space-x-8 px-4`,`aria-orientation` :`horizontal`}, []string{`
				`, react.Run("comment", map[string]string{}, []string{}), `
				`, react.Run("button", map[string]string{`id` :`tabs-1-tab-1`,`class` :`border-transparent text-gray-900 flex-1 whitespace-nowrap border-b-2 px-1 py-4 text-base font-medium`,`aria-controls` :`tabs-1-panel-1`,`role` :`tab`,`type` :`button`}, []string{`Women`}), `
				`, react.Run("comment", map[string]string{}, []string{}), `
				`, react.Run("button", map[string]string{`role` :`tab`,`type` :`button`,`id` :`tabs-1-tab-2`,`class` :`border-transparent text-gray-900 flex-1 whitespace-nowrap border-b-2 px-1 py-4 text-base font-medium`,`aria-controls` :`tabs-1-panel-2`}, []string{`Men`}), `
			  `}), `
			`}), `
  
			`, react.Run("comment", map[string]string{}, []string{}), `
			`, react.Run("div", map[string]string{`class` :`space-y-10 px-4 pb-8 pt-10`,`aria-labelledby` :`tabs-1-tab-1`,`role` :`tabpanel`,`tabindex` :`0`,`id` :`tabs-1-panel-1`}, []string{`
			  `, react.Run("div", map[string]string{`class` :`grid grid-cols-2 gap-x-4`}, []string{`
				`, react.Run("div", map[string]string{`class` :`group relative text-sm`}, []string{`
				  `, react.Run("div", map[string]string{`class` :`aspect-h-1 aspect-w-1 overflow-hidden rounded-lg bg-gray-100 group-hover:opacity-75`}, []string{`
					`, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/ecommerce-images/mega-menu-category-01.jpg`,`alt` :`Models sitting back to back, wearing Basic Tee in black and bone.`,`class` :`object-cover object-center`}, []string{}), `
				  `}), `
				  `, react.Run("a", map[string]string{`href` :`#`,`class` :`mt-6 block font-medium text-gray-900`}, []string{`
					`, react.Run("span", map[string]string{`class` :`absolute inset-0 z-10`,`aria-hidden` :`true`}, []string{}), `
					New Arrivals
				  `}), `
				  `, react.Run("p", map[string]string{`aria-hidden` :`true`,`class` :`mt-1`}, []string{`Shop now`}), `
				`}), `
				`, react.Run("div", map[string]string{`class` :`group relative text-sm`}, []string{`
				  `, react.Run("div", map[string]string{`class` :`aspect-h-1 aspect-w-1 overflow-hidden rounded-lg bg-gray-100 group-hover:opacity-75`}, []string{`
					`, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/ecommerce-images/mega-menu-category-02.jpg`,`alt` :`Close up of Basic Tee fall bundle with off-white, ochre, olive, and black tees.`,`class` :`object-cover object-center`}, []string{}), `
				  `}), `
				  `, react.Run("a", map[string]string{`href` :`#`,`class` :`mt-6 block font-medium text-gray-900`}, []string{`
					`, react.Run("span", map[string]string{`class` :`absolute inset-0 z-10`,`aria-hidden` :`true`}, []string{}), `
					Basic Tees
				  `}), `
				  `, react.Run("p", map[string]string{`aria-hidden` :`true`,`class` :`mt-1`}, []string{`Shop now`}), `
				`}), `
			  `}), `
			  `, react.Run("div", map[string]string{}, []string{`
				`, react.Run("p", map[string]string{`id` :`women-clothing-heading-mobile`,`class` :`font-medium text-gray-900`}, []string{`Clothing`}), `
				`, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`women-clothing-heading-mobile`,`class` :`mt-6 flex flex-col space-y-6`}, []string{`
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Tops`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Dresses`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Pants`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Denim`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Sweaters`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`T-Shirts`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`class` :`-m-2 block p-2 text-gray-500`,`href` :`#`}, []string{`Jackets`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Activewear`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Browse All`}), `
				  `}), `
				`}), `
			  `}), `
			  `, react.Run("div", map[string]string{}, []string{`
				`, react.Run("p", map[string]string{`id` :`women-accessories-heading-mobile`,`class` :`font-medium text-gray-900`}, []string{`Accessories`}), `
				`, react.Run("ul", map[string]string{`class` :`mt-6 flex flex-col space-y-6`,`role` :`list`,`aria-labelledby` :`women-accessories-heading-mobile`}, []string{`
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Watches`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Wallets`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Bags`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Sunglasses`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Hats`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Belts`}), `
				  `}), `
				`}), `
			  `}), `
			  `, react.Run("div", map[string]string{}, []string{`
				`, react.Run("p", map[string]string{`id` :`women-brands-heading-mobile`,`class` :`font-medium text-gray-900`}, []string{`Brands`}), `
				`, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`women-brands-heading-mobile`,`class` :`mt-6 flex flex-col space-y-6`}, []string{`
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Full Nelson`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`My Way`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Re-Arranged`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Counterfeit`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`class` :`-m-2 block p-2 text-gray-500`,`href` :`#`}, []string{`Significant Other`}), `
				  `}), `
				`}), `
			  `}), `
			`}), `
			`, react.Run("comment", map[string]string{}, []string{}), `
			`, react.Run("div", map[string]string{`id` :`tabs-1-panel-2`,`class` :`space-y-10 px-4 pb-8 pt-10`,`aria-labelledby` :`tabs-1-tab-2`,`role` :`tabpanel`,`tabindex` :`0`}, []string{`
			  `, react.Run("div", map[string]string{`class` :`grid grid-cols-2 gap-x-4`}, []string{`
				`, react.Run("div", map[string]string{`class` :`group relative text-sm`}, []string{`
				  `, react.Run("div", map[string]string{`class` :`aspect-h-1 aspect-w-1 overflow-hidden rounded-lg bg-gray-100 group-hover:opacity-75`}, []string{`
					`, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/ecommerce-images/product-page-04-detail-product-shot-01.jpg`,`alt` :`Drawstring top with elastic loop closure and textured interior padding.`,`class` :`object-cover object-center`}, []string{}), `
				  `}), `
				  `, react.Run("a", map[string]string{`href` :`#`,`class` :`mt-6 block font-medium text-gray-900`}, []string{`
					`, react.Run("span", map[string]string{`class` :`absolute inset-0 z-10`,`aria-hidden` :`true`}, []string{}), `
					New Arrivals
				  `}), `
				  `, react.Run("p", map[string]string{`aria-hidden` :`true`,`class` :`mt-1`}, []string{`Shop now`}), `
				`}), `
				`, react.Run("div", map[string]string{`class` :`group relative text-sm`}, []string{`
				  `, react.Run("div", map[string]string{`class` :`aspect-h-1 aspect-w-1 overflow-hidden rounded-lg bg-gray-100 group-hover:opacity-75`}, []string{`
					`, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/ecommerce-images/category-page-02-image-card-06.jpg`,`alt` :`Three shirts in gray, white, and blue arranged on table with same line drawing of hands and shapes overlapping on front of shirt.`,`class` :`object-cover object-center`}, []string{}), `
				  `}), `
				  `, react.Run("a", map[string]string{`class` :`mt-6 block font-medium text-gray-900`,`href` :`#`}, []string{`
					`, react.Run("span", map[string]string{`class` :`absolute inset-0 z-10`,`aria-hidden` :`true`}, []string{}), `
					Artwork Tees
				  `}), `
				  `, react.Run("p", map[string]string{`aria-hidden` :`true`,`class` :`mt-1`}, []string{`Shop now`}), `
				`}), `
			  `}), `
			  `, react.Run("div", map[string]string{}, []string{`
				`, react.Run("p", map[string]string{`id` :`men-clothing-heading-mobile`,`class` :`font-medium text-gray-900`}, []string{`Clothing`}), `
				`, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`men-clothing-heading-mobile`,`class` :`mt-6 flex flex-col space-y-6`}, []string{`
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`class` :`-m-2 block p-2 text-gray-500`,`href` :`#`}, []string{`Tops`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Pants`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Sweaters`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`class` :`-m-2 block p-2 text-gray-500`,`href` :`#`}, []string{`T-Shirts`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Jackets`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Activewear`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Browse All`}), `
				  `}), `
				`}), `
			  `}), `
			  `, react.Run("div", map[string]string{}, []string{`
				`, react.Run("p", map[string]string{`id` :`men-accessories-heading-mobile`,`class` :`font-medium text-gray-900`}, []string{`Accessories`}), `
				`, react.Run("ul", map[string]string{`class` :`mt-6 flex flex-col space-y-6`,`role` :`list`,`aria-labelledby` :`men-accessories-heading-mobile`}, []string{`
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Watches`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Wallets`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Bags`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Sunglasses`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Hats`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Belts`}), `
				  `}), `
				`}), `
			  `}), `
			  `, react.Run("div", map[string]string{}, []string{`
				`, react.Run("p", map[string]string{`id` :`men-brands-heading-mobile`,`class` :`font-medium text-gray-900`}, []string{`Brands`}), `
				`, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`men-brands-heading-mobile`,`class` :`mt-6 flex flex-col space-y-6`}, []string{`
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`class` :`-m-2 block p-2 text-gray-500`,`href` :`#`}, []string{`Re-Arranged`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`Counterfeit`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`class` :`-m-2 block p-2 text-gray-500`,`href` :`#`}, []string{`Full Nelson`}), `
				  `}), `
				  `, react.Run("li", map[string]string{`class` :`flow-root`}, []string{`
					`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 text-gray-500`}, []string{`My Way`}), `
				  `}), `
				`}), `
			  `}), `
			`}), `
		  `}), `
  
		  `, react.Run("div", map[string]string{`class` :`space-y-6 border-t border-gray-200 px-4 py-6`}, []string{`
			`, react.Run("div", map[string]string{`class` :`flow-root`}, []string{`
			  `, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 font-medium text-gray-900`}, []string{`Company`}), `
			`}), `
			`, react.Run("div", map[string]string{`class` :`flow-root`}, []string{`
			  `, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 font-medium text-gray-900`}, []string{`Stores`}), `
			`}), `
		  `}), `
  
		  `, react.Run("div", map[string]string{`class` :`space-y-6 border-t border-gray-200 px-4 py-6`}, []string{`
			`, react.Run("div", map[string]string{`class` :`flow-root`}, []string{`
			  `, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 block p-2 font-medium text-gray-900`}, []string{`Sign in`}), `
			`}), `
			`, react.Run("div", map[string]string{`class` :`flow-root`}, []string{`
			  `, react.Run("a", map[string]string{`class` :`-m-2 block p-2 font-medium text-gray-900`,`href` :`#`}, []string{`Create account`}), `
			`}), `
		  `}), `
  
		  `, react.Run("div", map[string]string{`class` :`border-t border-gray-200 px-4 py-6`}, []string{`
			`, react.Run("a", map[string]string{`href` :`#`,`class` :`-m-2 flex items-center p-2`}, []string{`
			  `, react.Run("img", map[string]string{`alt` :``,`class` :`block h-auto w-5 flex-shrink-0`,`src` :`https://tailwindui.com/img/flags/flag-canada.svg`}, []string{}), `
			  `, react.Run("span", map[string]string{`class` :`ml-3 block text-base font-medium text-gray-900`}, []string{`CAD`}), `
			  `, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`, change currency`}), `
			`}), `
		  `}), `
		`}), `
	  `}), `
	`}), `
  
	`, react.Run("header", map[string]string{`class` :`relative bg-white`}, []string{`
	  `, react.Run("p", map[string]string{`class` :`flex h-10 items-center justify-center bg-indigo-600 px-4 text-sm font-medium text-white sm:px-6 lg:px-8`}, []string{`Get free delivery on orders over $100`}), `
  
	  `, react.Run("nav", map[string]string{`class` :`mx-auto max-w-7xl px-4 sm:px-6 lg:px-8`,`aria-label` :`Top`}, []string{`
		`, react.Run("div", map[string]string{`class` :`border-b border-gray-200`}, []string{`
		  `, react.Run("div", map[string]string{`class` :`flex h-16 items-center`}, []string{`
			`, react.Run("comment", map[string]string{}, []string{}), `
			`, react.Run("button", map[string]string{`type` :`button`,`class` :`relative rounded-md bg-white p-2 text-gray-400 lg:hidden`}, []string{`
			  `, react.Run("span", map[string]string{`class` :`absolute -inset-0.5`}, []string{}), `
			  `, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`Open menu`}), `
			  `, react.Run("svg", map[string]string{`fill` :`none`,`viewBox` :`0 0 24 24`,`stroke-width` :`1.5`,`stroke` :`currentColor`,`aria-hidden` :`true`,`class` :`h-6 w-6`}, []string{`
				`, react.Run("path", map[string]string{`stroke-linecap` :`round`,`stroke-linejoin` :`round`,`d` :`M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5`}, []string{}), `
			  `}), `
			`}), `
  
			`, react.Run("comment", map[string]string{}, []string{}), `
			`, react.Run("div", map[string]string{`class` :`ml-4 flex lg:ml-0`}, []string{`
			  `, react.Run("a", map[string]string{`href` :`#`}, []string{`
				`, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`Your Company`}), `
				`, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600`,`alt` :``,`class` :`h-8 w-auto`}, []string{}), `
			  `}), `
			`}), `
  
			`, react.Run("comment", map[string]string{}, []string{}), `
			`, react.Run("div", map[string]string{`class` :`hidden lg:ml-8 lg:block lg:self-stretch`}, []string{`
			  `, react.Run("div", map[string]string{`class` :`flex h-full space-x-8`}, []string{`
				`, react.Run("div", map[string]string{`class` :`flex`}, []string{`
				  `, react.Run("div", map[string]string{`class` :`relative flex`}, []string{`
					`, react.Run("comment", map[string]string{}, []string{}), `
					`, react.Run("button", map[string]string{`aria-expanded` :`false`,`type` :`button`,`class` :`border-transparent text-gray-700 hover:text-gray-800 relative z-10 -mb-px flex items-center border-b-2 pt-px text-sm font-medium transition-colors duration-200 ease-out`}, []string{`Women`}), `
				  `}), `
  
				  `, react.Run("comment", map[string]string{}, []string{}), `
				  `, react.Run("div", map[string]string{`class` :`absolute inset-x-0 top-full text-sm text-gray-500`}, []string{`
					`, react.Run("comment", map[string]string{}, []string{}), `
					`, react.Run("div", map[string]string{`class` :`absolute inset-0 top-1/2 bg-white shadow`,`aria-hidden` :`true`}, []string{}), `
  
					`, react.Run("div", map[string]string{`class` :`relative bg-white`}, []string{`
					  `, react.Run("div", map[string]string{`class` :`mx-auto max-w-7xl px-8`}, []string{`
						`, react.Run("div", map[string]string{`class` :`grid grid-cols-2 gap-x-8 gap-y-10 py-16`}, []string{`
						  `, react.Run("div", map[string]string{`class` :`col-start-2 grid grid-cols-2 gap-x-8`}, []string{`
							`, react.Run("div", map[string]string{`class` :`group relative text-base sm:text-sm`}, []string{`
							  `, react.Run("div", map[string]string{`class` :`aspect-h-1 aspect-w-1 overflow-hidden rounded-lg bg-gray-100 group-hover:opacity-75`}, []string{`
								`, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/ecommerce-images/mega-menu-category-01.jpg`,`alt` :`Models sitting back to back, wearing Basic Tee in black and bone.`,`class` :`object-cover object-center`}, []string{}), `
							  `}), `
							  `, react.Run("a", map[string]string{`class` :`mt-6 block font-medium text-gray-900`,`href` :`#`}, []string{`
								`, react.Run("span", map[string]string{`class` :`absolute inset-0 z-10`,`aria-hidden` :`true`}, []string{}), `
								New Arrivals
							  `}), `
							  `, react.Run("p", map[string]string{`aria-hidden` :`true`,`class` :`mt-1`}, []string{`Shop now`}), `
							`}), `
							`, react.Run("div", map[string]string{`class` :`group relative text-base sm:text-sm`}, []string{`
							  `, react.Run("div", map[string]string{`class` :`aspect-h-1 aspect-w-1 overflow-hidden rounded-lg bg-gray-100 group-hover:opacity-75`}, []string{`
								`, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/ecommerce-images/mega-menu-category-02.jpg`,`alt` :`Close up of Basic Tee fall bundle with off-white, ochre, olive, and black tees.`,`class` :`object-cover object-center`}, []string{}), `
							  `}), `
							  `, react.Run("a", map[string]string{`href` :`#`,`class` :`mt-6 block font-medium text-gray-900`}, []string{`
								`, react.Run("span", map[string]string{`class` :`absolute inset-0 z-10`,`aria-hidden` :`true`}, []string{}), `
								Basic Tees
							  `}), `
							  `, react.Run("p", map[string]string{`aria-hidden` :`true`,`class` :`mt-1`}, []string{`Shop now`}), `
							`}), `
						  `}), `
						  `, react.Run("div", map[string]string{`class` :`row-start-1 grid grid-cols-3 gap-x-8 gap-y-10 text-sm`}, []string{`
							`, react.Run("div", map[string]string{}, []string{`
							  `, react.Run("p", map[string]string{`id` :`Clothing-heading`,`class` :`font-medium text-gray-900`}, []string{`Clothing`}), `
							  `, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`Clothing-heading`,`class` :`mt-6 space-y-6 sm:mt-4 sm:space-y-4`}, []string{`
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Tops`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`class` :`hover:text-gray-800`,`href` :`#`}, []string{`Dresses`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Pants`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Denim`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Sweaters`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`T-Shirts`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`class` :`hover:text-gray-800`,`href` :`#`}, []string{`Jackets`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Activewear`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Browse All`}), `
								`}), `
							  `}), `
							`}), `
							`, react.Run("div", map[string]string{}, []string{`
							  `, react.Run("p", map[string]string{`id` :`Accessories-heading`,`class` :`font-medium text-gray-900`}, []string{`Accessories`}), `
							  `, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`Accessories-heading`,`class` :`mt-6 space-y-6 sm:mt-4 sm:space-y-4`}, []string{`
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Watches`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Wallets`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Bags`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Sunglasses`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Hats`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Belts`}), `
								`}), `
							  `}), `
							`}), `
							`, react.Run("div", map[string]string{}, []string{`
							  `, react.Run("p", map[string]string{`id` :`Brands-heading`,`class` :`font-medium text-gray-900`}, []string{`Brands`}), `
							  `, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`Brands-heading`,`class` :`mt-6 space-y-6 sm:mt-4 sm:space-y-4`}, []string{`
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Full Nelson`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`My Way`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Re-Arranged`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Counterfeit`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Significant Other`}), `
								`}), `
							  `}), `
							`}), `
						  `}), `
						`}), `
					  `}), `
					`}), `
				  `}), `
				`}), `
				`, react.Run("div", map[string]string{`class` :`flex`}, []string{`
				  `, react.Run("div", map[string]string{`class` :`relative flex`}, []string{`
					`, react.Run("comment", map[string]string{}, []string{}), `
					`, react.Run("button", map[string]string{`aria-expanded` :`false`,`type` :`button`,`class` :`border-transparent text-gray-700 hover:text-gray-800 relative z-10 -mb-px flex items-center border-b-2 pt-px text-sm font-medium transition-colors duration-200 ease-out`}, []string{`Men`}), `
				  `}), `
  
				  `, react.Run("comment", map[string]string{}, []string{}), `
				  `, react.Run("div", map[string]string{`class` :`absolute inset-x-0 top-full text-sm text-gray-500`}, []string{`
					`, react.Run("comment", map[string]string{}, []string{}), `
					`, react.Run("div", map[string]string{`class` :`absolute inset-0 top-1/2 bg-white shadow`,`aria-hidden` :`true`}, []string{}), `
  
					`, react.Run("div", map[string]string{`class` :`relative bg-white`}, []string{`
					  `, react.Run("div", map[string]string{`class` :`mx-auto max-w-7xl px-8`}, []string{`
						`, react.Run("div", map[string]string{`class` :`grid grid-cols-2 gap-x-8 gap-y-10 py-16`}, []string{`
						  `, react.Run("div", map[string]string{`class` :`col-start-2 grid grid-cols-2 gap-x-8`}, []string{`
							`, react.Run("div", map[string]string{`class` :`group relative text-base sm:text-sm`}, []string{`
							  `, react.Run("div", map[string]string{`class` :`aspect-h-1 aspect-w-1 overflow-hidden rounded-lg bg-gray-100 group-hover:opacity-75`}, []string{`
								`, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/ecommerce-images/product-page-04-detail-product-shot-01.jpg`,`alt` :`Drawstring top with elastic loop closure and textured interior padding.`,`class` :`object-cover object-center`}, []string{}), `
							  `}), `
							  `, react.Run("a", map[string]string{`href` :`#`,`class` :`mt-6 block font-medium text-gray-900`}, []string{`
								`, react.Run("span", map[string]string{`class` :`absolute inset-0 z-10`,`aria-hidden` :`true`}, []string{}), `
								New Arrivals
							  `}), `
							  `, react.Run("p", map[string]string{`class` :`mt-1`,`aria-hidden` :`true`}, []string{`Shop now`}), `
							`}), `
							`, react.Run("div", map[string]string{`class` :`group relative text-base sm:text-sm`}, []string{`
							  `, react.Run("div", map[string]string{`class` :`aspect-h-1 aspect-w-1 overflow-hidden rounded-lg bg-gray-100 group-hover:opacity-75`}, []string{`
								`, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/ecommerce-images/category-page-02-image-card-06.jpg`,`alt` :`Three shirts in gray, white, and blue arranged on table with same line drawing of hands and shapes overlapping on front of shirt.`,`class` :`object-cover object-center`}, []string{}), `
							  `}), `
							  `, react.Run("a", map[string]string{`class` :`mt-6 block font-medium text-gray-900`,`href` :`#`}, []string{`
								`, react.Run("span", map[string]string{`class` :`absolute inset-0 z-10`,`aria-hidden` :`true`}, []string{}), `
								Artwork Tees
							  `}), `
							  `, react.Run("p", map[string]string{`aria-hidden` :`true`,`class` :`mt-1`}, []string{`Shop now`}), `
							`}), `
						  `}), `
						  `, react.Run("div", map[string]string{`class` :`row-start-1 grid grid-cols-3 gap-x-8 gap-y-10 text-sm`}, []string{`
							`, react.Run("div", map[string]string{}, []string{`
							  `, react.Run("p", map[string]string{`id` :`Clothing-heading`,`class` :`font-medium text-gray-900`}, []string{`Clothing`}), `
							  `, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`Clothing-heading`,`class` :`mt-6 space-y-6 sm:mt-4 sm:space-y-4`}, []string{`
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Tops`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Pants`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Sweaters`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`T-Shirts`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Jackets`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Activewear`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Browse All`}), `
								`}), `
							  `}), `
							`}), `
							`, react.Run("div", map[string]string{}, []string{`
							  `, react.Run("p", map[string]string{`id` :`Accessories-heading`,`class` :`font-medium text-gray-900`}, []string{`Accessories`}), `
							  `, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`Accessories-heading`,`class` :`mt-6 space-y-6 sm:mt-4 sm:space-y-4`}, []string{`
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Watches`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Wallets`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Bags`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Sunglasses`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Hats`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Belts`}), `
								`}), `
							  `}), `
							`}), `
							`, react.Run("div", map[string]string{}, []string{`
							  `, react.Run("p", map[string]string{`class` :`font-medium text-gray-900`,`id` :`Brands-heading`}, []string{`Brands`}), `
							  `, react.Run("ul", map[string]string{`role` :`list`,`aria-labelledby` :`Brands-heading`,`class` :`mt-6 space-y-6 sm:mt-4 sm:space-y-4`}, []string{`
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Re-Arranged`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Counterfeit`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`Full Nelson`}), `
								`}), `
								`, react.Run("li", map[string]string{`class` :`flex`}, []string{`
								  `, react.Run("a", map[string]string{`href` :`#`,`class` :`hover:text-gray-800`}, []string{`My Way`}), `
								`}), `
							  `}), `
							`}), `
						  `}), `
						`}), `
					  `}), `
					`}), `
				  `}), `
				`}), `
  
				`, react.Run("a", map[string]string{`href` :`#`,`class` :`flex items-center text-sm font-medium text-gray-700 hover:text-gray-800`}, []string{`Company`}), `
				`, react.Run("a", map[string]string{`href` :`#`,`class` :`flex items-center text-sm font-medium text-gray-700 hover:text-gray-800`}, []string{`Stores`}), `
			  `}), `
			`}), `
  
			`, react.Run("div", map[string]string{`class` :`ml-auto flex items-center`}, []string{`
			  `, react.Run("div", map[string]string{`class` :`hidden lg:flex lg:flex-1 lg:items-center lg:justify-end lg:space-x-6`}, []string{`
				`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-sm font-medium text-gray-700 hover:text-gray-800`}, []string{`Sign in`}), `
				`, react.Run("span", map[string]string{`class` :`h-6 w-px bg-gray-200`,`aria-hidden` :`true`}, []string{}), `
				`, react.Run("a", map[string]string{`href` :`#`,`class` :`text-sm font-medium text-gray-700 hover:text-gray-800`}, []string{`Create account`}), `
			  `}), `
  
			  `, react.Run("div", map[string]string{`class` :`hidden lg:ml-8 lg:flex`}, []string{`
				`, react.Run("a", map[string]string{`href` :`#`,`class` :`flex items-center text-gray-700 hover:text-gray-800`}, []string{`
				  `, react.Run("img", map[string]string{`src` :`https://tailwindui.com/img/flags/flag-canada.svg`,`alt` :``,`class` :`block h-auto w-5 flex-shrink-0`}, []string{}), `
				  `, react.Run("span", map[string]string{`class` :`ml-3 block text-sm font-medium`}, []string{`CAD`}), `
				  `, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`, change currency`}), `
				`}), `
			  `}), `
  
			  `, react.Run("comment", map[string]string{}, []string{}), `
			  `, react.Run("div", map[string]string{`class` :`flex lg:ml-6`}, []string{`
				`, react.Run("a", map[string]string{`href` :`#`,`class` :`p-2 text-gray-400 hover:text-gray-500`}, []string{`
				  `, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`Search`}), `
				  `, react.Run("svg", map[string]string{`fill` :`none`,`viewBox` :`0 0 24 24`,`stroke-width` :`1.5`,`stroke` :`currentColor`,`aria-hidden` :`true`,`class` :`h-6 w-6`}, []string{`
					`, react.Run("path", map[string]string{`stroke-linecap` :`round`,`stroke-linejoin` :`round`,`d` :`M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z`}, []string{}), `
				  `}), `
				`}), `
			  `}), `
  
			  `, react.Run("comment", map[string]string{}, []string{}), `
			  `, react.Run("div", map[string]string{`class` :`ml-4 flow-root lg:ml-6`}, []string{`
				`, react.Run("a", map[string]string{`class` :`group -m-2 flex items-center p-2`,`href` :`#`}, []string{`
				  `, react.Run("svg", map[string]string{`viewBox` :`0 0 24 24`,`stroke-width` :`1.5`,`stroke` :`currentColor`,`aria-hidden` :`true`,`class` :`h-6 w-6 flex-shrink-0 text-gray-400 group-hover:text-gray-500`,`fill` :`none`}, []string{`
					`, react.Run("path", map[string]string{`stroke-linecap` :`round`,`stroke-linejoin` :`round`,`d` :`M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z`}, []string{}), `
				  `}), `
				  `, react.Run("span", map[string]string{`class` :`ml-2 text-sm font-medium text-gray-700 group-hover:text-gray-800`}, []string{`0`}), `
				  `, react.Run("span", map[string]string{`class` :`sr-only`}, []string{`items in cart, view bag`}), `
				`}), `
			  `}), `
			`}), `
		  `}), `
		`}), `
	  `}), `
	`}), `
  `}), `
  
`})
}


func (r Component) Layout(props map[string]string, childrens string) string {
return react.Run("f", map[string]string{}, []string{`
	`, react.Run("html", map[string]string{}, []string{`

	`, react.Run("head", map[string]string{}, []string{`
		`, react.Run("link", map[string]string{`rel` :`stylesheet`,`href` :`/assets/css/dist/main.css`}, []string{}), `
	`}), `
	`, react.Run("Header", map[string]string{}, []string{`
	`}), `

	`, react.Run("Content", map[string]string{}, []string{`		`,childrens}), `
	`, react.Run("Footer", map[string]string{}, []string{`

	`}), `
	`}), `
`})
}