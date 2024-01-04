import React from "react"
import styles from "./textInput.module.css"
import Input from "react-phone-input-2"

interface PhoneInputProps {
    autoFormat?: boolean
    countryCodeEditable?: boolean
    preferredCountries?: string[]
    value?: string
    id: string
    name: string
    required?: boolean
    autoFocus?: boolean
    country?: string
    width: "full" | "auto"
    classnames?: string
    onChange: (value: string) => void
    disabled?: boolean
}

export default function PhoneInput(props: PhoneInputProps) {
    const widthStyle = props.width === "full" ? styles.phoneInputFull : ""

    return (
        <div className="w-full shadow-sm">
            <label htmlFor={"phone"} className={styles.label}>
                Phone Number
            </label>
            <Input
                autoFormat={props.autoFormat}
                countryCodeEditable={props.countryCodeEditable}
                disableDropdown={props.disabled}
                preferredCountries={props.preferredCountries}
                inputProps={{
                    id: props.id,
                    name: props.name,
                    className: `${styles.textInput} ${styles.phoneInput} ${widthStyle} ${props.classnames ?? ""}`,
                    required: props.required,
                    autoFocus: props.autoFocus,
                    disabled: props.disabled,
                }}
                specialLabel={""}
                country={props.country}
                value={props.value}
                onChange={props.onChange}
            />
        </div>
    )
}
