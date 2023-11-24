import React, {Component} from "react"
import styles from "./deleteModal.module.css"
import {XMarkIcon} from "@heroicons/react/20/solid"
import {ExclamationTriangleIcon} from "@heroicons/react/24/outline"

interface DeleteModalProps {
    message: string
    onDelete: () => void
    onCancel: () => void
}

interface DeleteModalState {
 show: boolean
}

export default class DeleteModal extends Component<DeleteModalProps, DeleteModalState> {
    constructor(props: DeleteModalProps) {
        super(props)
        this.state = {
            show: false
        }
    }

    public hide = () => {
        this.setState({show: false})
    }

    public show = () => {
        this.setState({show: true})
    }

    render() {
        return (
            <div id="defaultModal" tabIndex={-1} className={`fixed flex grow ${this.state?.show ? "" : "hidden"} z-50 h-screen w-full`}>
                <div className={styles.modal}>
                    {/*Modal content*/}
                    <div className={styles.content}>
                        <button type="button" className={styles.iconButton} onClick={this.props.onCancel}>
                            <XMarkIcon className="w-6 h-6" />
                            <span className="sr-only">Close modal</span>
                        </button>
                        <div className="p-6 text-center">
                            <ExclamationTriangleIcon className="mx-auto mb-4 w-20 h-20 text-gray-200" />
                            <h3 className={styles.message}>{this.props.message}</h3>
                            <button onClick={this.props.onDelete} type="button" className={styles.btnDelete}>
                                {"Yes, I'm sure"}
                            </button>
                            <button onClick={this.props.onCancel} type="button" className={styles.btnCancel}>
                                {"No, cancel"}
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}
