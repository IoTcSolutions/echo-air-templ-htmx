// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Carousel() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"https://unpkg.com/smoothscroll-polyfill@0.4.4/dist/smoothscroll.js\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := ``
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</script><div x-data=\"{\n        skip: 3,\n        atBeginning: false,\n        atEnd: false,\n        next() {\n            this.to((current, offset) =&gt; current + (offset * this.skip))\n        },\n        prev() {\n            this.to((current, offset) =&gt; current - (offset * this.skip))\n        },\n        to(strategy) {\n            let slider = this.$refs.slider\n            let current = slider.scrollLeft\n            let offset = slider.firstElementChild.getBoundingClientRect().width\n            slider.scrollTo({ left: strategy(current, offset), behavior: &#39;smooth&#39; })\n        },\n        focusableWhenVisible: {\n            &#39;x-intersect:enter&#39;() {\n                this.$el.removeAttribute(&#39;tabindex&#39;)\n            },\n            &#39;x-intersect:leave&#39;() {\n                this.$el.setAttribute(&#39;tabindex&#39;, &#39;-1&#39;)\n            },\n        },\n        disableNextAndPreviousButtons: {\n            &#39;x-intersect:enter.threshold.05&#39;() {\n                let slideEls = this.$el.parentElement.children\n\n                // If this is the first slide.\n                if (slideEls[0] === this.$el) {\n                    this.atBeginning = true\n                // If this is the last slide.\n                } else if (slideEls[slideEls.length-1] === this.$el) {\n                    this.atEnd = true\n                }\n            },\n            &#39;x-intersect:leave.threshold.05&#39;() {\n                let slideEls = this.$el.parentElement.children\n\n                // If this is the first slide.\n                if (slideEls[0] === this.$el) {\n                    this.atBeginning = false\n                // If this is the last slide.\n                } else if (slideEls[slideEls.length-1] === this.$el) {\n                    this.atEnd = false\n                }\n            },\n        },\n    }\" class=\"flex w-full flex-col\"><div x-on:keydown.right=\"next\" x-on:keydown.left=\"prev\" tabindex=\"0\" role=\"region\" aria-labelledby=\"carousel-label\" class=\"flex space-x-6\"><h2 id=\"carousel-label\" class=\"sr-only\" hidden>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := `Carousel`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><!--")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := ` Prev Button `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("--><button x-on:click=\"prev\" class=\"text-6xl\" :aria-disabled=\"atBeginning\" :tabindex=\"atEnd ? -1 : 0\" :class=\"{ &#39;opacity-50 cursor-not-allowed&#39;: atBeginning }\"><span aria-hidden=\"true\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-12 w-12 text-gray-600\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"3\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M15 19l-7-7 7-7\"></path></svg></span> <span class=\"sr-only\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var5 := `Skip to previous slide page`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></button> <span id=\"carousel-content-label\" class=\"sr-only\" hidden>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var6 := `Carousel`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var6)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span><ul x-ref=\"slider\" tabindex=\"0\" role=\"listbox\" aria-labelledby=\"carousel-content-label\" class=\"flex w-full snap-x snap-mandatory overflow-x-scroll\"><li x-bind=\"disableNextAndPreviousButtons\" class=\"flex w-1/3 shrink-0 snap-start flex-col items-center justify-center p-2\" role=\"option\"><img class=\"mt-2 w-full\" src=\"https://picsum.photos/400/200?random=1\" alt=\"placeholder image\"> <button x-bind=\"focusableWhenVisible\" class=\"px-4 py-2 text-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var7 := `Do Something`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var7)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></li><li x-bind=\"disableNextAndPreviousButtons\" class=\"flex w-1/3 shrink-0 snap-start flex-col items-center justify-center p-2\" role=\"option\"><img class=\"mt-2 w-full\" src=\"https://picsum.photos/400/200?random=2\" alt=\"placeholder image\"> <button x-bind=\"focusableWhenVisible\" class=\"px-4 py-2 text-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var8 := `Do Something`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var8)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></li><li x-bind=\"disableNextAndPreviousButtons\" class=\"flex w-1/3 shrink-0 snap-start flex-col items-center justify-center p-2\" role=\"option\"><img class=\"mt-2 w-full\" src=\"https://picsum.photos/400/200?random=3\" alt=\"placeholder image\"> <button x-bind=\"focusableWhenVisible\" class=\"px-4 py-2 text-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var9 := `Do Something`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var9)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></li><li x-bind=\"disableNextAndPreviousButtons\" class=\"flex w-1/3 shrink-0 snap-start flex-col items-center justify-center p-2\" role=\"option\"><img class=\"mt-2 w-full\" src=\"https://picsum.photos/400/200?random=4\" alt=\"placeholder image\"> <button x-bind=\"focusableWhenVisible\" class=\"px-4 py-2 text-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var10 := `Do Something`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var10)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></li><li x-bind=\"disableNextAndPreviousButtons\" class=\"flex w-1/3 shrink-0 snap-start flex-col items-center justify-center p-2\" role=\"option\"><img class=\"mt-2 w-full\" src=\"https://picsum.photos/400/200?random=5\" alt=\"placeholder image\"> <button x-bind=\"focusableWhenVisible\" class=\"px-4 py-2 text-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var11 := `Do Something`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var11)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></li><li x-bind=\"disableNextAndPreviousButtons\" class=\"flex w-1/3 shrink-0 snap-start flex-col items-center justify-center p-2\" role=\"option\"><img class=\"mt-2 w-full\" src=\"https://picsum.photos/400/200?random=6\" alt=\"placeholder image\"> <button x-bind=\"focusableWhenVisible\" class=\"px-4 py-2 text-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var12 := `Do Something`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var12)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></li></ul><!--")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var13 := ` Next Button `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var13)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("--><button x-on:click=\"next\" class=\"text-6xl\" :aria-disabled=\"atEnd\" :tabindex=\"atEnd ? -1 : 0\" :class=\"{ &#39;opacity-50 cursor-not-allowed&#39;: atEnd }\"><span aria-hidden=\"true\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-12 w-12 text-gray-600\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"3\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M9 5l7 7-7 7\"></path></svg></span> <span class=\"sr-only\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var14 := `Skip to next slide page`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var14)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></button></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
