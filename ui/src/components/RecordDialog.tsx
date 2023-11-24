import React, {Component} from "react"
import {XMarkIcon} from "@heroicons/react/20/solid"
import Button from "@/components/Button"
import RecordCreationViewModel from "@/viewmodels/record/RecordCreationViewModel"
import {Record} from "@/models/record/Record"
import TextInput from "@/components/TextInput"
import {addDays, format, isAfter, isBefore} from "date-fns"
import {RecordTypeUtils} from "@/enums/RecordType"
import styles from "@/components/textInput.module.css"
import UpdateRecordViewModel from "@/viewmodels/record/UpdateRecordViewModel"

interface RecordDialogProps {
    data?: Record
    onDismiss: () => void
    onCreate?: (vm: RecordCreationViewModel) => void
    onUpdate?: (vm: UpdateRecordViewModel) => void
    mode?: Mode
}

interface RecordDialogState {
    title: string
    vm: RecordCreationViewModel | UpdateRecordViewModel
    show: boolean
    mode: Mode
    step: Step
}

export enum Mode {
    CREATE = "CREATE",
    UPDATE = "UPDATE",
}

export enum Step {
    BASE = "BASE",
    EXTRA_INFO = "EXTRA_INFO",
    NEXT_DATE = "NEXT_DATE",
    LINEDATA = "LINEDATA"
}

export default class RecordDialog extends Component<RecordDialogProps, RecordDialogState> {
    constructor(props: RecordDialogProps) {
        super(props)
        this.state = {
            title: "Create Record",
            vm: new RecordCreationViewModel(),
            show: false,
            mode: props.mode ?? Mode.CREATE,
            step: Step.BASE,
        }
    }

    public setDataForCreation = (data: Record) => {
        const vm = new RecordCreationViewModel()
        vm.name = data.name ?? ""
        vm.recordType = data.recordType ?? ""
        vm.name = data.name ?? ""
        vm.lot = data.lot ?? ""
        vm.result = data.result ?? ""
        vm.description = data.description ?? ""
        vm.notes = data.notes ?? ""
        this.setCreate()
        this.setState(
            {
                step: Step.BASE,
                vm,
                title: "Create Record",
            },
            () => setTimeout(() => document?.getElementById("date")?.focus(), 1)
        )
    }

    public setDataForCreationWithDate = (data: Record, date: Date) => {
        const vm = new RecordCreationViewModel()
        vm.name = data.name ?? ""
        vm.recordType = data.recordType ?? ""
        vm.name = data.name ?? ""
        vm.lot = data.lot ?? ""
        vm.result = data.result ?? ""
        vm.description = data.description ?? ""
        vm.notes = data.notes ?? ""
        vm.date = format(date, "yyyy-MM-dd")
        this.setCreate()
        this.setState(
            {
                step: Step.BASE,
                vm,
                title: "Create Record",
            },
            () => setTimeout(() => document?.getElementById("date")?.focus(), 1)
        )
    }

    public setDataForUpdate = (data: Record) => {
        const vm = new UpdateRecordViewModel()
        vm.name = data.name ?? ""
        vm.date = data.date ? format(new Date(data.date), "yyyy-MM-dd") : ""
        vm.recordType = data.recordType ?? ""
        vm.name = data.name ?? ""
        vm.lot = data.lot ?? ""
        vm.result = data.result ?? ""
        vm.description = data.description ?? ""
        vm.notes = data.notes ?? ""
        vm.hasNextDate = data.nextDate !== undefined
        vm.nextDate = data.nextDate ? format(new Date(data.nextDate), "yyyy-MM-dd") : ""
        this.setState({step: Step.BASE, vm, title: "Update Record"})
    }

    public hide = () => {
        this.setState({show: false, vm: new RecordCreationViewModel()})
    }

    public show = () => {
        this.setState({show: true})
    }

    public setCreate = () => {
        this.setState({mode: Mode.CREATE, step: Step.BASE})
    }

    public setUpdate = () => {
        this.setState({mode: Mode.UPDATE})
    }

    private onNameChange = (value: string) => {
        const vm = this.state.vm
        vm.name = value
        this.setState({vm})
    }

    private onLotChange = (value: string) => {
        const vm = this.state.vm
        vm.lot = value
        this.setState({vm})
    }

    private onResultChange = (value: string) => {
        const vm = this.state.vm
        vm.result = value
        this.setState({vm})
    }

    private onDescriptionChange = (value: string) => {
        const vm = this.state.vm
        vm.description = value
        this.setState({vm})
    }

    private onNotesChange = (value: string) => {
        const vm = this.state.vm
        vm.notes = value
        this.setState({vm})
    }

    private onDateChange = (value: string) => {
        const vm = this.state.vm
        vm.date = isAfter(new Date(value), new Date()) ? format(new Date(), "yyyy-MM-dd").toString() : value
        this.setState({vm})
    }

    private onRecordTypeChange = (event: any) => {
        const vm = this.state.vm
        vm.recordType = event.currentTarget.value
        this.setState({vm})
    }

    private onHasNextDateChange = () => {
        const vm = this.state.vm
        vm.hasNextDate = !vm.hasNextDate
        this.setState({vm})
    }

    private onNextDateChange = (value: string) => {
        const vm = this.state.vm
        const minDate = addDays(new Date(), 1)
        vm.nextDate = isBefore(new Date(value), minDate) ? format(minDate, "yyyy-MM-dd").toString() : value
        this.setState({vm})
    }

    private checkName = () => {
        this.state.vm.checkName()
        this.forceUpdate()
    }

    private checkDate = () => {
        this.state.vm.checkDate()
        this.forceUpdate()
    }

    private checkNextDate = () => {
        this.state.vm.checkNextDate()
        this.forceUpdate()
    }

    private checkType = () => {
        this.state.vm.checkType()
        this.forceUpdate()
    }

    private get primaryButtonTitle() {
        if (this.state.step === Step.BASE) {
            return "Add Extra Information"
        }

        if (this.state.vm.hasNextDate && this.state.step === Step.EXTRA_INFO) {
            return "Schedule Next Date"
        }

        switch (this.state.mode) {
            case Mode.CREATE:
                return "Create Record"

            case Mode.UPDATE:
                return "Update Record"
        }
    }

    private get secondaryButtonTitle() {
        if (this.state.step !== Step.BASE) {
            return "Back"
        }

        return "Cancel"
    }

    private primaryButtonPress = (): void => {
        this.state.vm.validate(this.state.step)
        if (!this.state.vm.isValid) {
            this.forceUpdate()
            return
        }

        if (this.state.step === Step.BASE) {
            this.setState({step: Step.EXTRA_INFO})
            return
        } else if (this.state.vm.hasNextDate && this.state.step === Step.EXTRA_INFO) {
            this.setState({step: Step.NEXT_DATE})
            return
        }

        switch (this.state.mode) {
            case Mode.CREATE:
                return this.props.onCreate!(this.state.vm as RecordCreationViewModel)
            case Mode.UPDATE:
                return this.props.onUpdate!(this.state.vm as UpdateRecordViewModel)
        }
    }

    private secondButtonPress = () => {
        if (this.state.step === Step.BASE) {
            this.props.onDismiss()
            return
        } else if (this.state.step == Step.EXTRA_INFO) {
            this.setState({step: Step.BASE})
            return
        } else if (this.state.step == Step.NEXT_DATE) {
            this.setState({step: Step.EXTRA_INFO})
            return
        }


    }

    private get formBody() {
        const vm = this.state.vm
        return (
            <>
                {this.state.step === Step.BASE && (
                    <>
                        <TextInput
                            autoFocus={false}
                            id="name"
                            name="name"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Name"
                            value={vm.name}
                            onInput={this.onNameChange}
                            onBlur={this.checkName}
                            hasError={vm.hasNameError}
                            errorMessage={vm.nameError}
                            showLabel={true}
                        />

                        <div>
                            <label htmlFor={"recordType"} className={styles.label}>
                                Record Type
                            </label>
                            <select
                                value={vm.recordType}
                                id="recordType"
                                name="recordType"
                                onChange={this.onRecordTypeChange}
                                onBlur={this.checkType}
                                className={`${styles.textInput} ${styles.textInputFull} ${
                                    vm.hasRecordTypeError && styles.errorHighlight
                                } rounded-md capitalize`}
                            >
                                <option value="">Choose a Record Type</option>
                                {RecordTypeUtils.getAll().map((g) => {
                                    const name = RecordTypeUtils.getTitle(g)
                                    return (
                                        <option key={name} value={g} className={"capitalize"}>
                                            {name}
                                        </option>
                                    )
                                })}
                            </select>
                            <span
                                className={styles.errorMessage}>{vm.hasRecordTypeError ? vm.recordTypeError : ""}</span>
                        </div>

                        <TextInput
                            id="date"
                            name="date"
                            type="date"
                            max={format(new Date(), "yyyy-MM-dd")}
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Administration Date"
                            value={vm.date}
                            onInput={this.onDateChange}
                            onBlur={this.checkDate}
                            hasError={vm.hasDateError}
                            errorMessage={vm.dateError}
                            showLabel={true}
                        />
                    </>
                )}
                {this.state.step === Step.EXTRA_INFO && (
                    <>
                        <TextInput
                            id="lot"
                            name="lot"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Lot Number"
                            value={vm.lot}
                            onInput={this.onLotChange}
                            showLabel
                        />
                        <TextInput
                            id="result"
                            name="result"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Result"
                            value={vm.result}
                            onInput={this.onResultChange}
                            showLabel
                        />
                        <TextInput
                            id="description"
                            name="description"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Description"
                            value={vm.description}
                            onInput={this.onDescriptionChange}
                            showLabel
                        />
                        <TextInput
                            id="notes"
                            name="notes"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Notes"
                            value={vm.notes}
                            onInput={this.onNotesChange}
                            showLabel
                        />
                        <div>
                            <input
                                id="link-checkbox"
                                type="checkbox"
                                value={""}
                                checked={vm.hasNextDate}
                                onChange={this.onHasNextDateChange}
                                className="w-4 h-4 text-indigo-600 rounded focus:ring-indigo-600 ring-offset-gray-800 focus:ring-2 bg-gray-700 border-gray-600"
                            />
                            <label htmlFor="link-checkbox"
                                   className="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">
                                Is recurring
                            </label>
                        </div>
                    </>
                )}
                {this.state.step === Step.NEXT_DATE &&
                    <TextInput
                        id="date"
                        name="date"
                        type="date"
                        autoComplete="off"
                        required
                        width={"full"}
                        min={format(addDays(new Date(), 1), "yyyy-MM-dd")}
                        classNames="rounded-md"
                        placeholder="Schedule Next Occurance"
                        value={vm.nextDate}
                        onInput={this.onNextDateChange}
                        onBlur={this.checkNextDate}
                        hasError={vm.hasNextDateError}
                        errorMessage={vm.dateError}
                        showLabel={true}
                    />}
            </>
        )
    }

    render() {
        return (
            <div id="defaultModal" tabIndex={-1}
                 className={`fixed flex grow ${this.state?.show ? "" : "hidden"} z-50 h-screen w-full`}>
                <div className="relative self-center mx-auto w-full max-w-2xl max-h-full">
                    {/*Modal content*/}
                    <div className="relative bg-white rounded-lg shadow dark:bg-gray-700">
                        {/*Modal header*/}
                        <div className="flex items-start justify-between p-4 border-b rounded-t dark:border-gray-600">
                            <h3 className="text-xl font-semibold text-gray-900 dark:text-white text-center">{this.state.title}</h3>
                            <button
                                type="button"
                                className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ml-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
                                onClick={this.props.onDismiss}
                            >
                                <XMarkIcon className={"h-6 w-6"}/>
                                <span className="sr-only">Close modal</span>
                            </button>
                        </div>
                        {/*Modal body */}

                        <div className="px-12 py-4 space-y-4 max-h-96 overflow-y-auto justify-center">
                            <div className="shadow-sm px-4 py-4">
                                <form className="space-y-1 mb-3" method="POST" onSubmit={this.props.onDismiss}>
                                    {this.formBody}
                                </form>
                            </div>
                        </div>
                        {/*footer*/}
                        <div
                            className="flex items-center p-6 justify-end space-x-2 border-t border-gray-200 rounded-b dark:border-gray-600">
                            <Button
                                className={"right-0"}
                                title={this.secondaryButtonTitle}
                                variant={"secondary"}
                                type={"button"}
                                onClick={this.secondButtonPress}
                            />
                            <Button
                                className={"right-0"}
                                title={this.primaryButtonTitle}
                                variant={"primary"}
                                type={"button"}
                                onClick={this.primaryButtonPress}
                            />
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}
