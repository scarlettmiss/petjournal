import {Event} from "react-big-calendar";
import * as React from "react";

export default class CalendarEvent implements Event {
    constructor(
        public allDay?: boolean | undefined,
        public title?: React.ReactNode | undefined,
        public start?: Date | undefined,
        public end?: Date | undefined,
        public resource?: any,
        public isEvent?: boolean) {
    }
}
