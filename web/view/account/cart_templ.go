// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package account

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"FiberShop/internal/models"
	"FiberShop/web/view/layout"
	"fmt"
	"strconv"
)

func incrementQuantity(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_incrementQuantity_fcaa`,
		Function: `function __templ_incrementQuantity_fcaa(id){var input = document.getElementById('quantityInput_' + id);
    input.value = parseInt(input.value, 10) + 1;
    updateTotalPrice(id)
}`,
		Call:       templ.SafeScript(`__templ_incrementQuantity_fcaa`, id),
		CallInline: templ.SafeScriptInline(`__templ_incrementQuantity_fcaa`, id),
	}
}

func decrementQuantity(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_decrementQuantity_552a`,
		Function: `function __templ_decrementQuantity_552a(id){var input = document.getElementById('quantityInput_' + id);
    var newValue = parseInt(input.value, 10) - 1;
    input.value = newValue >= 1 ? newValue : 1;
    updateTotalPrice(id)
}`,
		Call:       templ.SafeScript(`__templ_decrementQuantity_552a`, id),
		CallInline: templ.SafeScriptInline(`__templ_decrementQuantity_552a`, id),
	}
}

func updateTotalPrice(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_updateTotalPrice_e79e`,
		Function: `function __templ_updateTotalPrice_e79e(id){updateTotalPrice(id)
}`,
		Call:       templ.SafeScript(`__templ_updateTotalPrice_e79e`, id),
		CallInline: templ.SafeScriptInline(`__templ_updateTotalPrice_e79e`, id),
	}
}

func Cart(data layout.Data, cartItems []models.CartItem, totalCartPrice float64) templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n            @layer utilities {\n                input[type=\"number\"]::-webkit-inner-spin-button,\n                input[type=\"number\"]::-webkit-outer-spin-button {\n                    -webkit-appearance: none;\n                    margin: 0;\n                }\n            }\n        </style> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(cartItems) < 1 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-center mt-5 font-bold  text-emerald-500 text-[3.75rem]\">Your Cart is empty</h1>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-center mt-5 font-bold  text-emerald-500 text-[3.75rem]\">Your Cart</h1><section class=\"flex justify-center \"><div class=\"xl:mx-20 justify-center flex-1 px-1 py-6 lg:py-4\"><div class=\"flex flex-wrap\"><div class=\"w-full lg:w-9/12\"><div class=\"px-10\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, item := range cartItems {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var3 string
					templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("item" + strconv.Itoa(item.ID))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 48, Col: 70}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"relative flex flex-wrap items-center pb-8 mb-8 -mx-4 border-b border-gray-200 dark:border-gray-700 xl:justify-between border-opacity-40\"><div class=\"w-full mb-2 lg:mb-0 h-96 md:h-44 md:w-44\"><img src=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var4 string
					templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(item.Product.ImageURL)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 51, Col: 71}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"\" class=\"object-cover w-full h-full\"></div><div class=\"w-full px-4 md:w-auto xl:mb-0\"><a class=\"block mb-3 text-xl font-medium dark:text-white hover:underline\" href=\"#\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var5 string
					templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(item.Product.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 56, Col: 62}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a><div class=\"flex mb-2 flex-wrap\"><p class=\"mr-4 text-sm font-medium\"><span class=\"dark:text-white\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var6 string
					templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(item.Product.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 59, Col: 103}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span><!--                                    <span class=\"ml-2 text-gray-600 dark:text-gray-300\">yellow</span>--></p><!--                                <p class=\"text-sm font-medium dark:text-white\">--><!--                                    <span>Size:</span>--><!--                                    <span class=\"ml-2 text-gray-600 dark:text-gray-300\">38</span>--><!--                                </p>--></div></div><div class=\"mt-6 mb-6 w-auto xl:mb-0 xl:mt-0\"><div class=\"flex items-center\"><div class=\"inline-flex items-center px-2 font-semibold border border-gray-300 rounded-md dark:bg-gray-800 dark:border-gray-700\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, decrementQuantity(strconv.Itoa(item.ID)))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button onclick=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var7 templ.ComponentScript = decrementQuantity(strconv.Itoa(item.ID))
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var7.Call)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-2 pr-2 border-r border-gray-300 dark:border-gray-600 dark:text-white hover:text-gray-400\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" fill=\"currentColor\" class=\"bi bi-dash\" viewBox=\"0 0 16 16\"><path d=\"M4 8a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7A.5.5 0 0 1 4 8z\"></path></svg></button> ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, updateTotalPrice(strconv.Itoa(item.ID)))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"number\" class=\"w-10 px-2 py-4 text-center border-0 rounded-md dark:bg-gray-800 bg-gray-50 dark:text-white\" placeholder=\"1\" value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var8 string
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(int(item.Quantity)))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 76, Col: 224}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" id=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("quantityInput_%d", item.ID))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 76, Col: 272}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" minlength=\"1\" min=\"1\" data-price=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var10 string
					templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%f", item.Product.Price))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 76, Col: 347}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" oninput=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var11 templ.ComponentScript = updateTotalPrice(strconv.Itoa(item.ID))
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var11.Call)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, incrementQuantity(strconv.Itoa(item.ID)))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button onclick=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var12 templ.ComponentScript = incrementQuantity(strconv.Itoa(item.ID))
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var12.Call)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-2 pl-2 border-l border-gray-300 dark:border-gray-600 dark:text-white hover:text-gray-400\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" fill=\"currentColor\" class=\"bi bi-plus\" viewBox=\"0 0 16 16\"><path d=\"M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4z\"></path></svg></button></div></div></div><div class=\"px-4 w-auto\"><span class=\"text-xl font-semibold text-emerald-500\"><span class=\"\">$</span> <span id=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var13 string
					templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("totalPrice_%d", item.ID))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 88, Col: 92}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var14 string
					templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%.2f", item.Product.Price*float64(item.Quantity)))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 88, Col: 161}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></span></div><button id=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var15 string
					templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("remove" + strconv.Itoa(item.ID)))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 91, Col: 91}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#alert\" hx-post=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var16 string
					templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/account/cart/delete?id=%d", item.ID))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 91, Col: 172}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block items-center rounded bg-emerald-500 text-neutral-50 shadow-[0_4px_9px_-4px_rgba(51,45,45,0.7)] hover:bg-emerald-600 hover:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] focus:bg-emerald-800 focus:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] active:bg-emerald-700 active:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] px-3 pb-2 pt-2.5 text-[0.9rem] font-bold uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0\">Remove</button></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"w-full lg:w-3/12\"><div class=\"px-6 mb-14\"><div class=\"mb-10\"><span class=\"mb-6 text-3xl font-bold dark:text-white\">Apply Coupon</span><form hx-post=\"/coupon/discount\" hx-target=\"#alert\"><input type=\"text\" id=\"code\" name=\"code\" class=\"flex-1 w-full px-8 py-4 mt-4 font-normal placeholder-gray-400 border dark:bg-gray-800 rounded-xl dark:border-gray-700 dark:placeholder-gray-500 md:flex-none md:mr-6 dark:text-gray-400\" placeholder=\"COUPON123\"> <button id=\"couponBtn\" class=\"block w-full text-center mt-4 rounded bg-emerald-500 text-neutral-50 shadow-[0_4px_9px_-4px_rgba(51,45,45,0.7)] hover:bg-emerald-600 hover:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] focus:bg-emerald-800 focus:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] active:bg-emerald-700 active:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] px-3 pb-2 pt-2.5 text-[0.9rem] font-bold uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0\" href=\"#\">Apply</button></form><div id=\"couponResponse\"></div></div><div><h2 class=\"mb-6 text-3xl font-bold dark:text-white\">Cart totals</h2><div class=\"flex items-center justify-between px-5 py-4 mb-3 font-medium leading-8 bg-gray-100 bg-opacity-50 border dark:text-white dark:bg-gray-800 dark:border-gray-800 rounded-xl\"><span>Subtotal</span> <span class=\"flex dark:text-gray-400 items-center text-xl\"><span class=\"mr-2 text-base\">$</span> <span id=\"subTotalCartPrice\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var17 string
				templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%.2f", totalCartPrice))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 120, Col: 110}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></span></div><div class=\"flex items-center justify-between px-5 py-4 mb-3 font-medium leading-8 bg-gray-100 bg-opacity-50 border dark:text-white dark:bg-gray-800 dark:border-gray-800 rounded-xl\"><span>Coupon</span> <span class=\"flex dark:text-gray-400 items-center text-xl\"><span class=\"mr-2 text-base\">$</span> <span id=\"couponValue\">0</span></span></div><div class=\"flex items-center justify-between px-5 py-4 mb-6 font-medium leading-8 bg-gray-100 border dark:text-white dark:bg-gray-800 dark:border-gray-800 rounded-xl\"><span>Total</span> <span class=\"flex items-center text-xl text-emerald-500\"><span class=\"mr-2 text-base\">$</span> <span id=\"totalCartPrice\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var18 string
				templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%.2f", totalCartPrice))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/view/account/cart.templ`, Line: 136, Col: 107}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></span></div><a class=\"block text-center mt-4 rounded bg-emerald-500 text-neutral-50 shadow-[0_4px_9px_-4px_rgba(51,45,45,0.7)] hover:bg-emerald-600 hover:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] focus:bg-emerald-800 focus:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] active:bg-emerald-700 active:shadow-[0_8px_9px_-4px_rgba(51,45,45,0.2),0_4px_18px_0_rgba(51,45,45,0.1)] px-3 pb-2 pt-2.5 text-[0.9rem] font-bold uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0\" href=\"#\">Checkout</a></div></div></div></div></div></section>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = Base().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <script src=\"/js/cart.js\"></script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base(data).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"alert\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
