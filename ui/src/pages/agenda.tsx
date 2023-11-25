import React from "react"
import "react-phone-input-2/lib/material.css"
import TextUtils from "@/Utils/TextUtils"
import ErrorDto from "@/models/ErrorDto"
import {withRouter} from "next/router"
import dynamic from "next/dynamic"
import ErrorMessage from "@/components/ErrorMessage"
import {WithRouterProps} from "next/dist/client/with-router"
import BaseComponent from "@/components/BaseComponent"
import {recordCreationHandler, recordHandler, recordsCreationHandler, recordUpdateHandler} from "@/pages/api/record"
import {Record} from "@/models/record/Record"
import CalendarEvent from "@/models/CalendarEvent"
import CalendarComponent from "@/components/CalendarComponent"
import RecordDialog, {Mode} from "@/components/RecordDialog"
import RecordCreationViewModel from "@/viewmodels/record/RecordCreationViewModel"
import UpdateRecordViewModel from "@/viewmodels/record/UpdateRecordViewModel"
import {RecordType} from "@/enums/RecordType"

interface AgendaProps extends WithRouterProps {}

interface AgendaState {
    token?: string
    serverError?: string
    records: Record[]
    events: CalendarEvent[]
    loading: boolean
    selectedRecord?: Record
}

const ProtectedPage = dynamic(() => import("@/components/ProtectedPage"), {
    ssr: false,
})

class Agenda extends BaseComponent<AgendaProps, AgendaState> {
    private recordDialogRef: RecordDialog | null = null

    constructor(props: AgendaProps) {
        super(props)
        this.state = {
            records: [],
            events: [],
            loading: true,
        }
    }

    private getAllRecords = async (token?: string) => {
        const resp = await recordHandler(token)
        const data: Record[] = await resp.json()
        if (resp.ok) {
            this.setState({records: data}, () => this.updateEvents())
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else if ((data as ErrorDto).error) {
            throw new Error((data as ErrorDto).error)
        }
    }

    private initPage = (token?: string) => {
        this.getAllRecords(token).catch((e) => {
            console.error("Could not get User", e)
            this.setState({serverError: e.message})
        })
        this.setState({token, loading: true})
    }

    private updateEvents = () => {
        const events = this.state.records
            .filter((it) => it.recordType !== RecordType.WEIGHT && it.recordType !== RecordType.TEMPERATURE)
            .flatMap((r) => {
                const events: CalendarEvent[] = []
                const d = new Date(r.date!)
                events.push(new CalendarEvent(true, r.name!, d, d, r, true))
                return events
            })
        this.setState({events, loading: false})
    }

    private onViewRecordSelected = (record: Record) => {
        this.recordDialogRef?.setViewOnly()
        this.recordDialogRef?.setDataForUpdate(record, record.pet)
        this.recordDialogRef?.show()
    }

    private updateRecord = async (vm: UpdateRecordViewModel, r?: Record) => {
        const id = r?.id
        if (!id) {
            throw Error("no record to update")
        }

        const petId = r?.pet?.id
        if (!petId) {
            throw Error("Pet Id was not defined")
        }

        const resp = await recordUpdateHandler(vm, id, petId, this.state.token)
        const response: Record = await resp.json()
        if (resp.ok) {
            const records = this.state.records
            const index = records.findIndex((it) => it.id === id)
            records[index] = response
            this.setState({records: records}, () => this.updateEvents())
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            throw Error((response as ErrorDto).error)
        }
    }

    private onUpdateRecord = async (vm: UpdateRecordViewModel) => {
        this.recordDialogRef?.hide()
        this.updateRecord(vm, this.state.selectedRecord)
    }

    private onEventSelected = async (e: CalendarEvent) => {
        const record = e.resource
        this.setState({selectedRecord: record})

        this.onViewRecordSelected(record)
    }

    render() {
        if (TextUtils.isNotEmpty(this.state.serverError)) {
            return <ErrorMessage message={this.state.serverError} />
        }

        return (
            <ProtectedPage init={this.initPage} key={"agenda"} className={"relative"}>
                <div className={"p-8 h-full"}>
                    <CalendarComponent
                        events={this.state.events}
                        views={{month: true, week: true, day: true, agenda: true}}
                        onSelectEvent={this.onEventSelected}
                    />
                </div>
                <RecordDialog
                    ref={(ref) => (this.recordDialogRef = ref)}
                    onDismiss={() => this.recordDialogRef?.hide()}
                    onUpdate={this.onUpdateRecord}
                    mode={Mode.UPDATE}
                />
                {this.state.loading && (
                    <div role="status" className={"absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2"}>
                        <svg
                            aria-hidden="true"
                            className="w-20 h-20 animate-spin dark:text-indigo-300 fill-indigo-600"
                            viewBox="0 0 100 101"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                        >
                            <path
                                d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                                fill="currentColor"
                            />
                            <path
                                d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                                fill="currentFill"
                            />
                        </svg>
                        <span className="sr-only">Loading...</span>
                    </div>
                )}
            </ProtectedPage>
        )
    }
}

export default withRouter(Agenda)
