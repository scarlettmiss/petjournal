import {ReactElement} from "react"
import styles from "./button.module.css"

interface ButtonProps {
    title: string
    width?: "full" | "auto"
    variant?: "primary" | "secondary" | "third" | "group"
    type?: "submit" | "reset" | "button" | undefined
    onClick?: () => void
    className?: string
    icon?: ReactElement
    selected?: boolean
    disabled?: boolean
}

export default function Button(props: ButtonProps) {
    const widthStyle = props.width === "full" ? styles.btnFull : ""
    const variantStyle = (): string => {
        if (props.variant === "primary") {
            return props.disabled ? styles.btnPrimaryDisabled : styles.btnPrimary
        }
        if (props.variant === "secondary") {
            return props.disabled ? styles.btnSecondaryDisabled : styles.btnSecondary
        }
        if (props.variant === "group") {
            const groupSelectedStyle = props.selected ? styles.btnPrimary : styles.btnSecondary
            return `${styles.btnGroup} ${groupSelectedStyle}`
        }
        return ""
    }

    const onClick = (e: any) => {
        e.preventDefault()
        props.onClick && props.onClick()
    }

    return (
        <button
            key={props.title}
            type={props.type}
            className={`${styles.btn} ${widthStyle} ${variantStyle()} ${props.className ?? ""}`}
            onClick={onClick}
            disabled={props.disabled}
        >
            {props.title}
        </button>
    )
}
