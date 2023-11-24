import {Component} from "react"
import {Cookies} from "react-cookie"

abstract class BaseComponent<Props, State> extends Component<Props, State> {
    protected cookies: Cookies = new Cookies()

    protected logout = (callback?: () => void) => {
        this.cookies.remove("token")
        if (callback !== undefined) {
            callback()
        }
    }
}

export default BaseComponent
