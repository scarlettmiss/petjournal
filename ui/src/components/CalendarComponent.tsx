import React, {Component} from "react"
import {Calendar, dateFnsLocalizer, View, ViewsProps} from "react-big-calendar";
import {el, enUS} from "date-fns/locale";
import parse from "date-fns/parse";
import startOfWeek from "date-fns/startOfWeek";
import getDay from "date-fns/getDay";
import format from 'date-fns/format'
import CalendarEvent from "@/models/CalendarEvent";

interface CalendarComponentProps {
    events: CalendarEvent[]
    updateEvents?: (range: Date[] | { start: Date; end: Date }, view?: View) => void | undefined;
    views: ViewsProps
    onSelectEvent?: ((event: CalendarEvent, e: React.SyntheticEvent<HTMLElement>) => void) | undefined;
    defaultView?: View
}

interface CalendarComponentState {
}


export default class CalendarComponent extends Component<CalendarComponentProps, CalendarComponentState> {
    constructor(props: CalendarComponentProps) {
        super(props)
        this.state = {}
    }

    render() {
        const locales = {
            'en-US': enUS,
            'el-GR': el,
        }

        const localizer = dateFnsLocalizer({
            format,
            parse,
            startOfWeek,
            getDay,
            locales,
        })

        return (
            <Calendar
                style={{flex: 1}}
                views={this.props.views}
                defaultView={this.props.defaultView}
                localizer={localizer}
                events={this.props.events}
                startAccessor="start"
                endAccessor="end"
                onSelectEvent={this.props.onSelectEvent}
                onRangeChange={this.props.updateEvents}
                popup={true}
            />
        )
    }
}
