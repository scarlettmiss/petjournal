import TextUtils from "@/Utils/TextUtils"
import {Step} from "@/components/RecordDialog"
import {startOfDay} from "date-fns"

export default class RecordCreationViewModel {
    private _recordType: string = ""
    private _name: string = ""
    private _date: string = ""
    private _lot: string = ""
    private _result: string = ""
    private _description: string = ""
    private _notes: string = ""
    private _hasNextDate: boolean = false
    private _nextDate: string = ""

    public nameError: string = ""
    public dateError: string = ""
    public recordTypeError: string = ""
    public nextDateError: string = ""
    public resultError: string = ""

    public hasNameError: boolean = false
    public hasDateError: boolean = false
    public hasRecordTypeError: boolean = false
    public hasNextDateError: boolean = false
    public hasResultError: boolean = false

    private _valid: boolean = false

    public get stringify() {
        return JSON.stringify(this.fields)
    }

    validate(step: Step) {
        this._valid = false
        switch (step) {
            case Step.LINEDATA:
                this.checkDate()
                this.checkNumberResult()
                this.checkType()
                this._valid =
                    TextUtils.isEmpty(this.dateError) && TextUtils.isEmpty(this.resultError) && TextUtils.isEmpty(this.recordTypeError)
                break
            case Step.BASE:
                this.checkName()
                this.checkDate()
                this.checkType()
                this._valid =
                    TextUtils.isEmpty(this.nameError) && TextUtils.isEmpty(this.dateError) && TextUtils.isEmpty(this.recordTypeError)
                break
            case Step.EXTRA_INFO:
                this._valid = true
                break
            case Step.NEXT_DATE:
                this.checkNextDate()
                this._valid = TextUtils.isEmpty(this.nextDateError)
                break
        }
    }

    get recordType(): string {
        return this._recordType
    }

    set recordType(value: string) {
        this._recordType = value
    }

    public get isValid(): boolean {
        return this._valid
    }

    get hasNextDate(): boolean {
        return this._hasNextDate
    }

    set hasNextDate(value: boolean) {
        this._hasNextDate = value
    }

    public checkDate = () => {
        const isEmpty = TextUtils.isEmpty(this._date)
        this.dateError = isEmpty ? "The date is required" : ""
        this.hasDateError = isEmpty
    }

    public checkNextDate = () => {
        const isEmpty = TextUtils.isEmpty(this._nextDate)
        this.nextDateError = isEmpty ? "The date is required" : ""
        this.hasNextDateError = isEmpty
    }

    public checkType = () => {
        const isEmpty = TextUtils.isEmpty(this._recordType)
        this.recordTypeError = isEmpty ? "The record type is required" : ""
        this.hasRecordTypeError = isEmpty
    }

    public checkName = () => {
        const isEmpty = TextUtils.isEmpty(this._name)
        this.nameError = isEmpty ? "Name is required" : ""
        this.hasNameError = isEmpty
    }

    public checkNumberResult = () => {
        const isNotFinite = TextUtils.isEmpty(this._result) || !Number.isFinite(Number(this._result))
        this.resultError = isNotFinite ? "The value is required" : ""
        this.hasResultError = isNotFinite
    }

    get name(): string {
        return this._name
    }

    set name(value: string) {
        this._name = value
    }

    get nextDate(): string {
        return this._nextDate
    }

    set nextDate(value: string) {
        this._nextDate = value
    }

    get notes(): string {
        return this._notes
    }

    set notes(value: string) {
        this._notes = value
    }

    get description(): string {
        return this._description
    }

    set description(value: string) {
        this._description = value
    }

    get result(): string {
        return this._result
    }

    set result(value: string) {
        this._result = value
    }

    get lot(): string {
        return this._lot
    }

    set lot(value: string) {
        this._lot = value
    }

    get date(): string {
        return this._date
    }

    set date(value: string) {
        this._date = value
    }

    public get fields() {
        return {
            recordType: this._recordType,
            name: this._name,
            date: new Date(this._date! + "T00:00").getTime(),
            lot: this._lot,
            result: this._result,
            description: this._description,
            notes: this._notes,
            nextDate: this.hasNextDate ? new Date(this._nextDate! + "T00:00").getTime() : 0,
        }
    }
}
