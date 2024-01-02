// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func FullCalendar() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{\n        calendar: null,\n        events: [\n            {\n                id: 1,\n                title: &#39;Build my secret project 🛠&#39;,\n                start: &#39;2022-06-05&#39;,\n                end: &#39;2022-06-08&#39;,\n            },\n            {\n                id: 2,\n                title: &#39;Launch 🚀&#39;,\n                start: &#39;2022-06-23&#39;,\n            },\n        ],\n        newEventTitle: null,\n        newEventStart: null,\n        newEventEnd: null,\n        init() {\n            this.calendar = new FullCalendar.Calendar(this.$refs.calendar, {\n                events: (info, success) =&gt; success(this.events),\n                initialDate: &#39;2022-06-01&#39;,\n                initialView: &#39;dayGridMonth&#39;,\n                selectable: true,\n                unselectAuto: false,\n                editable: true,\n                select: (info) =&gt; {\n                    this.newEventStart = info.startStr\n                    this.newEventEnd = info.endStr\n                },\n                eventClick: (info) =&gt; {\n                    if (confirm(&#39;Are you sure you want to remove this event?&#39;)) {\n                        const index = this.getEventIndex(info)\n\n                        this.events.splice(index, 1)\n\n                        this.calendar.refetchEvents()\n                    }\n                },\n                eventChange: (info) =&gt; {\n                    const index = this.getEventIndex(info)\n\n                    this.events[index].start = info.event.startStr\n                    this.events[index].end = info.event.endStr\n                },\n            })\n\n            this.calendar.render()\n        },\n        getEventIndex(info) {\n            return this.events.findIndex((event) =&gt; event.id == info.event.id)\n        },\n        addEvent() {\n            if (! this.newEventTitle || ! this.newEventStart) {\n                return alert(&#39;Please choose a title and start date for the event!&#39;)\n            }\n\n            let event = {\n                id: Date.now(),\n                title: this.newEventTitle,\n                start: this.newEventStart,\n                end: this.newEventEnd,\n            }\n\n            this.events.push(event)\n            this.calendar.refetchEvents()\n\n            this.newEventTitle = null\n            this.newEventStart = null\n            this.newEventEnd = null\n\n            this.calendar.unselect()\n        },\n    }\"><div x-ref=\"calendar\"></div><h2 class=\"mt-8 font-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := `List of events`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><ul class=\"mt-1 list-inside list-disc text-sm text-gray-500\"><template x-for=\"event in events\" :key=\"event.id\"><li x-text=\"`${event.title}: ${event.start} ${(event.end ? &#39; through &#39; + event.end : &#39;&#39;)}`\"></li></template></ul><form x-on:submit.prevent=\"addEvent\" class=\"mt-8 max-w-md\"><h2 class=\"font-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := `Adding a new event`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p class=\"mt-1 text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := `Select a date or date range on the calendar and enter a title to create a new event.`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-4 flex items-center gap-2\"><label for=\"new-event-title\" class=\"sr-only\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var5 := `Event Title`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label> <input id=\"new-event-title\" type=\"text\" x-model=\"newEventTitle\" class=\"mt-1 w-full rounded-md border border-gray-200 px-3 py-2.5\" placeholder=\"Event title\"> <button type=\"submit\" class=\"shrink-0 rounded-md bg-white px-5 py-2.5 shadow\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var6 := `Submit`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var6)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div></form><h2 class=\"mt-8 font-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var7 := `Moving events and changing duration`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var7)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p class=\"mt-1 text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var8 := `Drag an event to move it to a different day. Drag the right edge to change the duration.`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var8)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><h2 class=\"mt-8 font-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var9 := `Deleting an event`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var9)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p class=\"mt-1 text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var10 := `Click an event and confirm to remove it.`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var10)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}