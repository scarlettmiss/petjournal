import React from "react"
import NavBar from "@/components/NavBar"
import TextUtils from "@/Utils/TextUtils"
import BaseComponent from "@/components/BaseComponent"
import {WithRouterProps} from "next/dist/client/with-router"
import {withRouter} from "next/router"

interface ProtectedPageProps extends WithRouterProps {
    children: any
    hideNav?: boolean
    className?: string
    init?: (token?: string) => void
}

interface ProtectedPageState {
    token?: string
}

class ProtectedPage extends BaseComponent<ProtectedPageProps, ProtectedPageState> {
    private initInterval?: NodeJS.Timer

    constructor(props: ProtectedPageProps) {
        super(props)
        this.state = {}
    }

    componentDidMount() {
        const t = this.cookies.get("token")

        if (t !== undefined) {
            this.setState({token: t})
            this.initInterval = setInterval(() => {
                if (this.props.router.isReady) {
                    this.props.init && this.props.init(t)
                    clearInterval(this.initInterval)
                }
            }, 1)
        } else {
            this.logout(() => this.props.router.replace("/auth/login"))
        }
    }

    render() {
        return TextUtils.isNotEmpty(this.state.token) ? (
            <div className={`flex flex-col h-screen bg-slate-700`}>
                {!this.props.hideNav && <NavBar/>}
                <span
                    className={`flex flex-col grow overflow-y-auto ${this.props.className}`}>{this.props.children}</span>
            </div>
        ) : (
            <></>
        )
    }
}

export default withRouter(ProtectedPage)
