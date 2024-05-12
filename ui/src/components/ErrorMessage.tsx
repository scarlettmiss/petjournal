import React from "react"
import TextUtils from "@/Utils/TextUtils"

interface ErrorMessageProps {
    message?: string
    className?: string
}
export default function ErrorMessage(props: ErrorMessageProps) {
    return TextUtils.isNotEmpty(props.message) ? (
        <div
            className={`p-4 mb-4 text-sm rounded-lg bg-slate-700 text-red-400 ${
                props.className ? props.className : ""
            }`}
            role="alert"
        >
            <span className="font-medium">Error :</span> {props.message}
        </div>
    ) : (
        <></>
    )
}
