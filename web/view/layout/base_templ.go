// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type Data struct {
	Title string
	Email string
}

func Base(data Data) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(data.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/layout/base.templ`, Line: 13, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script src=\"/js/htmx.min.js\"></script><script>\n        try {\n          if (\n                  localStorage.theme === \"dark\" ||\n                  (!(\"theme\" in localStorage) &&\n                          window.matchMedia(\"(prefers-color-scheme: dark)\").matches)\n          ) {\n            document.documentElement.classList.add(\"dark\");\n            document\n                    .querySelector('meta[name=\"theme-color\"]')\n                    .setAttribute(\"content\", \"#0B1120\");\n          } else {\n            document.documentElement.classList.remove(\"dark\");\n          }\n        } catch (_) {}\n      </script><link rel=\"stylesheet\" href=\"/styles/output.css\"></head><body class=\" dark:bg-neutral-800 dark:shadow-black/10\"><nav class=\"top-0 flex w-full flex-wrap items-center justify-between bg-[#FBFBFB] py-2 shadow-md shadow-black/5 dark:bg-neutral-600 dark:shadow-black/10 lg:flex-wrap lg:justify-start lg:py-4\"><div class=\"flex w-full flex-wrap items-center justify-between px-3\"><!-- Hamburger button for mobile view --><button class=\"block border-0 bg-transparent px-2 text-neutral-500 hover:no-underline hover:shadow-none focus:no-underline focus:shadow-none focus:outline-none focus:ring-0 dark:text-neutral-200 lg:hidden\" type=\"button\" data-te-collapse-init data-te-target=\"#navbarSupportedContent1\" aria-controls=\"navbarSupportedContent1\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"><!-- Hamburger icon --><span class=\"[&amp;&gt;svg]:w-7\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"h-7 w-7\"><path fill-rule=\"evenodd\" d=\"M3 6.75A.75.75 0 013.75 6h16.5a.75.75 0 010 1.5H3.75A.75.75 0 013 6.75zM3 12a.75.75 0 01.75-.75h16.5a.75.75 0 010 1.5H3.75A.75.75 0 013 12zm0 5.25a.75.75 0 01.75-.75h16.5a.75.75 0 010 1.5H3.75a.75.75 0 01-.75-.75z\" clip-rule=\"evenodd\"></path></svg></span></button><!-- Collapsible navigation container --><div class=\"!visible hidden flex-grow basis-[100%] items-center lg:!flex lg:basis-auto\" id=\"navbarSupportedContent1\" data-te-collapse-item><!-- Logo --><a class=\"mr-4 my-1 flex items-center text-neutral-900 hover:text-neutral-900 focus:text-neutral-900 lg:mb-0 lg:mt-0\" href=\"/\"><img class=\"mr-2\" src=\"https://tecdn.b-cdn.net/img/logo/te-transparent-noshadows.webp\" style=\"height: 20px\" alt=\"TE Logo\" loading=\"lazy\"> <span class=\"font-bold text-emerald-500 text-[18px]\">GosBoost</span></a><!-- Left navigation links --><ul class=\"list-style-none mr-auto flex flex-col pl-0 lg:flex-row\" data-te-navbar-nav-ref><li class=\"mb-4 lg:mb-0 lg:pr-2\" data-te-nav-item-ref><!-- Dashboard link --><a class=\"text-neutral-500 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 lg:px-2 [&amp;.active]:text-black/90 dark:[&amp;.active]:text-zinc-400\" href=\"/contact\" data-te-nav-link-ref><span class=\"text-[1.1rem] font-semibold\">Contact</span></a></li><!-- Team link --><li class=\"mb-4 lg:mb-0 lg:pr-2\" data-te-nav-item-ref><a class=\"text-neutral-500 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 lg:px-2 [&amp;.active]:text-black/90 dark:[&amp;.active]:text-neutral-400\" href=\"#\" data-te-nav-link-ref><span class=\"text-[1.1rem] font-semibold\">All category</span></a></li><!-- Projects link --><li class=\"mb-4 lg:mb-0 lg:pr-2\" data-te-nav-item-ref><a class=\"text-neutral-500 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 lg:px-2 [&amp;.active]:text-black/90 dark:[&amp;.active]:text-neutral-400\" href=\"#\" data-te-nav-link-ref><span class=\"text-[1.1rem] font-semibold\">F.A.Q</span></a></li><li class=\"mb-4 pl-2 lg:mb-0 lg:pl-0 lg:pr-1\" data-te-nav-item-ref data-te-dropdown-ref><!-- Dropdown --><a class=\"flex items-center text-neutral-500 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-400 dark:focus:text-neutral-400 lg:px-2 [&amp;.active]:text-black/90 dark:[&amp;.active]:text-neutral-400\" href=\"#\" type=\"button\" id=\"dropdownMenuButton2\" data-te-dropdown-toggle-ref aria-expanded=\"false\"><span class=\"text-[1.1rem] font-semibold\">Dropdown</span> <span class=\"ml-1 w-2\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\" fill=\"currentColor\" class=\"h-5 w-5\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></span></a><ul class=\"absolute z-[1000] float-left m-0 hidden min-w-max list-none overflow-hidden rounded-lg border-none bg-white bg-clip-padding text-left text-base shadow-lg dark:bg-neutral-700 [&amp;[data-te-dropdown-show]]:block\" aria-labelledby=\"dropdownMenuButton2\" data-te-dropdown-menu-ref><li><a class=\"block w-full whitespace-nowrap bg-transparent px-4 py-2 text-sm font-normal text-neutral-700 hover:bg-neutral-100 active:text-neutral-800 active:no-underline disabled:pointer-events-none disabled:bg-transparent disabled:text-neutral-400 dark:text-neutral-200 dark:hover:bg-neutral-600\" href=\"#\" data-te-dropdown-item-ref>Action</a></li><li><a class=\"block w-full whitespace-nowrap bg-transparent px-4 py-2 text-sm font-normal text-neutral-700 hover:bg-neutral-100 active:text-neutral-800 active:no-underline disabled:pointer-events-none disabled:bg-transparent disabled:text-neutral-400 dark:text-neutral-200 dark:hover:bg-neutral-600\" href=\"#\" data-te-dropdown-item-ref>Another action</a></li><li><a class=\"block w-full whitespace-nowrap bg-transparent px-4 py-2 text-sm font-normal text-neutral-700 hover:bg-neutral-100 active:text-neutral-800 active:no-underline disabled:pointer-events-none disabled:bg-transparent disabled:text-neutral-400 dark:text-neutral-200 dark:hover:bg-neutral-600\" href=\"#\" data-te-dropdown-item-ref>Something else here</a></li></ul></li></ul></div><!-- Right elements --><div class=\"relative flex items-center\"><button id=\"theme-switcher\" class=\"mr-4 text-neutral-600 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 [&amp;.active]:text-black/90 dark:[&amp;.active]:text-neutral-400\" data-theme=\"light\"><div class=\"pointer-events-none\"><div class=\"inline-block text-center\" data-theme-icon=\"light\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"inline-block h-7 w-7\"><path id=\"path\" d=\"M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z\"></path></svg></div></div></button> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.Email != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Cart Icon --> <a class=\"mr-4 text-neutral-600 transition duration-200 hover:text-neutral-700 hover:ease-in-out focus:text-neutral-700 disabled:text-black/30 motion-reduce:transition-none dark:text-neutral-200 dark:hover:text-neutral-300 dark:focus:text-neutral-300 [&amp;.active]:text-black/90 dark:[&amp;.active]:text-neutral-400\" href=\"/account/cart\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-7 h-7\"><path d=\"M15.75 10.5V6a3.75 3.75 0 1 0-7.5 0v4.5m11.356-1.993 1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 0 1-1.12-1.243l1.264-12A1.125 1.125 0 0 1 5.513 7.5h12.974c.576 0 1.059.435 1.119 1.007ZM8.625 10.5a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm7.5 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Z\"></path></svg></a><form action=\"/account/logout\" method=\"post\"><button type=\"submit\" class=\"inline-flex items-center mr-4 rounded border-emerald-500 text-emerald-500 hover:border-emerald-600 hover:bg-emerald-400 hover:bg-opacity-10 hover:text-emerald-600 focus:border-emerald-700 focus:text-emerald-400 active:border-emerald-800 active:text-emerald-800 dark:border-emerald-300 dark:text-emerald-300 dark:hover:hover:text-white px-3 pb-2 pt-2.5 text-[0.9rem] font-bold uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0 flex-row-reverse\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6 ml-2\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M8.25 9V5.25A2.25 2.25 0 0 1 10.5 3h6a2.25 2.25 0 0 1 2.25 2.25v13.5A2.25 2.25 0 0 1 16.5 21h-6a2.25 2.25 0 0 1-2.25-2.25V15m-3 0-3-3m0 0 3-3m-3 3H15\"></path></svg> Logout</button></form><a href=\"/account\" role=\"button\" class=\"inline-flex items-center rounded bg-emerald-500 text-neutral-50 shadow-[0_4px_9px_-4px_rgba(51,45,45,0.7)] hover:bg-emerald-600 hover:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] focus:bg-emerald-800 focus:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] active:bg-emerald-700 active:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] px-3 pb-2 pt-2.5 text-[0.9rem] font-bold uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0\">Account</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center\"><a href=\"/login\" role=\"button\" class=\"inline-flex items-center mr-4 rounded border-emerald-500 text-emerald-500 hover:border-emerald-600 hover:bg-emerald-400 hover:bg-opacity-10 hover:text-emerald-600 focus:border-emerald-700 focus:text-emerald-400 active:border-emerald-800 active:text-emerald-800 dark:border-emerald-300 dark:text-emerald-300 dark:hover:hover:text-white px-3 pb-2 pt-2.5 text-[0.9rem] font-bold uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0 flex-row-reverse\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6 ml-2\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M18 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 1 1-6.75 0 3.375 3.375 0 0 1 6.75 0ZM3 19.235v-.11a6.375 6.375 0 0 1 12.75 0v.109A12.318 12.318 0 0 1 9.374 21c-2.331 0-4.512-.645-6.374-1.766Z\"></path></svg> Log In</a> <a href=\"/register\" role=\"button\" class=\"inline-flex items-center rounded bg-emerald-500 text-neutral-50 shadow-[0_4px_9px_-4px_rgba(51,45,45,0.7)] hover:bg-emerald-600 hover:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] focus:bg-emerald-800 focus:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] active:bg-emerald-700 active:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] px-3 pb-2 pt-2.5 text-[0.9rem] font-bold uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0\">Sign Up <svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6 ml-2\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15M12 9l-3 3m0 0 3 3m-3-3h12.75\"></path></svg></a></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<hr class=\"my-12 h-0.5 border-t-0 bg-neutral-500 dark:bg-neutral-50 opacity-100 dark:opacity-50\"><footer class=\"flex flex-col items-center text-center text-white\"><div class=\"container\"><div class=\" flex justify-center\"><a href=\"#!\" class=\"mr-9 text-neutral-800 dark:text-neutral-200\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-4 w-4\" fill=\"currentColor\" viewBox=\"0 0 24 24\"><path d=\"M9 8h-3v4h3v12h5v-12h3.642l.358-4h-4v-1.667c0-.955.192-1.333 1.115-1.333h2.885v-5h-3.808c-3.596 0-5.192 1.583-5.192 4.615v3.385z\"></path></svg></a> <a href=\"#!\" class=\"mr-9 text-neutral-800 dark:text-neutral-200\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-4 w-4\" fill=\"currentColor\" viewBox=\"0 0 24 24\"><path d=\"M24 4.557c-.883.392-1.832.656-2.828.775 1.017-.609 1.798-1.574 2.165-2.724-.951.564-2.005.974-3.127 1.195-.897-.957-2.178-1.555-3.594-1.555-3.179 0-5.515 2.966-4.797 6.045-4.091-.205-7.719-2.165-10.148-5.144-1.29 2.213-.669 5.108 1.523 6.574-.806-.026-1.566-.247-2.229-.616-.054 2.281 1.581 4.415 3.949 4.89-.693.188-1.452.232-2.224.084.626 1.956 2.444 3.379 4.6 3.419-2.07 1.623-4.678 2.348-7.29 2.04 2.179 1.397 4.768 2.212 7.548 2.212 9.142 0 14.307-7.721 13.995-14.646.962-.695 1.797-1.562 2.457-2.549z\"></path></svg></a> <a href=\"#!\" class=\"mr-9 text-neutral-800 dark:text-neutral-200\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-5 w-5\" fill=\"currentColor\" viewBox=\"0 0 24 24\"><path d=\"M7 11v2.4h3.97c-.16 1.029-1.2 3.02-3.97 3.02-2.39 0-4.34-1.979-4.34-4.42 0-2.44 1.95-4.42 4.34-4.42 1.36 0 2.27.58 2.79 1.08l1.9-1.83c-1.22-1.14-2.8-1.83-4.69-1.83-3.87 0-7 3.13-7 7s3.13 7 7 7c4.04 0 6.721-2.84 6.721-6.84 0-.46-.051-.81-.111-1.16h-6.61zm0 0 17 2h-3v3h-2v-3h-3v-2h3v-3h2v3h3v2z\" fill-rule=\"evenodd\" clip-rule=\"evenodd\"></path></svg></a> <a href=\"#!\" class=\"mr-9 text-neutral-800 dark:text-neutral-200\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-4 w-4\" fill=\"currentColor\" viewBox=\"0 0 24 24\"><path d=\"M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z\"></path></svg></a> <a href=\"#!\" class=\"mr-9 text-neutral-800 dark:text-neutral-200\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-4 w-4\" fill=\"currentColor\" viewBox=\"0 0 24 24\"><path d=\"M4.98 3.5c0 1.381-1.11 2.5-2.48 2.5s-2.48-1.119-2.48-2.5c0-1.38 1.11-2.5 2.48-2.5s2.48 1.12 2.48 2.5zm.02 4.5h-5v16h5v-16zm7.982 0h-4.968v16h4.969v-8.399c0-4.67 6.029-5.052 6.029 0v8.399h4.988v-10.131c0-7.88-8.922-7.593-11.018-3.714v-2.155z\"></path></svg></a> <a href=\"#!\" class=\"text-neutral-800 dark:text-neutral-200\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-4 w-4\" fill=\"currentColor\" viewBox=\"0 0 24 24\"><path d=\"M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z\"></path></svg></a></div></div><!--Copyright section--><div class=\"w-full p-4 text-center text-neutral-700  dark:text-neutral-200\">© 2023 GosBoost. All Rights Reserved.</div></footer><script src=\"/js/theme.js\"></script><script type=\"text/javascript\" src=\"/tailwind/node_modules/tw-elements/dist/js/tw-elements.umd.min.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
