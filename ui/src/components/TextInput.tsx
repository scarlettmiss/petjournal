import React, {FormEvent, HTMLInputTypeAttribute} from "react"
import styles from "./textInput.module.css"

interface ButtonProps {
    id: string
    name: string
    icon?: React.ReactElement
    type?: HTMLInputTypeAttribute
    required?: boolean
    autoComplete: string
    classNames?: string
    placeholder: string
    value?: string
    width?: "full" | "auto"
    onInput: (value: string) => void
    onBlur?: () => void
    hasError?: boolean
    errorMessage?: string
    disabled?: boolean
    autoFocus?: boolean
    showLabel?: boolean
    max?: string
    min?: string
}

export default function TextInput(props: ButtonProps) {
    const widthStyle = props.width === "full" ? styles.textInputFull : ""

    const onInput = (event: FormEvent<HTMLInputElement>) => {
        const value = event.currentTarget.value
        props.onInput(value)
    }

    return (
        <div className={styles.textInputContainer}>
            <label htmlFor={props.id} className={props.showLabel ? styles.label : "sr-only"}>
                {props.placeholder}
            </label>
            {props.icon !== undefined && <div className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                {props.icon}
            </div>}
            <input
                required={props.required}
                id={props.id}
                name={props.name}
                type={props.type}
                max={props.max}
                min={props.min}
                autoComplete={props.autoComplete}
                className={`${styles.textInput} ${widthStyle} ${props.classNames ?? ""} ${props.hasError && styles.errorHighlight}`}
                placeholder={props.placeholder.replace("*", "")}
                value={props.value}
                onInput={onInput}
                disabled={props.disabled}
                onBlur={props.onBlur}
                autoFocus={props.autoFocus}
                style={{colorScheme: "dark"}}
            />
            <span className={styles.errorMessage}>{props.hasError ? props.errorMessage : ""}</span>
        </div>
    )
}
