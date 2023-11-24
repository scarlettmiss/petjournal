import React from "react"
import Image from "next/image"

interface AvatarProps {
    avatar?: string
    avatarTitle?: string
    icon?: React.ReactElement
    onCLick?: () => void
    textStyle?: string
    className?: string
}

export default function Avatar(props: AvatarProps) {
    return props.avatar ? (
        <Image
            unoptimized
            src={props.avatar}
            alt="Avatar"
            width={90}
            height={90}
            style={{objectFit: "cover"}}
            className={`rounded-full aspect-square ${props.className ?? ""}`}
            onClick={props.onCLick}
        />
    ) : (
        <div
            className={`aspect-square flex items-center justify-center bg-indigo-100 text-indigo-600 rounded-full text-center ${
                props.textStyle ?? "text-3xl"
            } font-bold ring-2 ring-indigo-500 ${props.className ?? ""}`}
            onClick={props.onCLick}
        >
            {props.icon ? props.icon : props.avatarTitle}
        </div>
    )
}
