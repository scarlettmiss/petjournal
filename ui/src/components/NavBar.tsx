import React, {ReactElement} from "react"
import Link from "next/link"
import Image from "next/image"
import Page from "@/models/page"
import styles from "./navbar.module.css"
import Avatar from "@/components/Avatar"
import {UserIcon} from "@heroicons/react/20/solid"
import {withRouter} from "next/router"
import {ArrowLeftOnRectangleIcon} from "@heroicons/react/24/outline"
import {WithRouterProps} from "next/dist/client/with-router"
import BaseComponent from "@/components/BaseComponent"

interface NavBarProps extends WithRouterProps {
    hideAllPages?: boolean
    pages?: Page[]
    buttons?: ReactElement
}

interface NavBarState {}

class NavBar extends BaseComponent<NavBarProps, NavBarState> {
    constructor(props: NavBarProps) {
        super(props)
        this.state = {}
    }

    private get pages() {
        const pages = new Array<Page>()
        if (this.props.hideAllPages) {
            return pages
        }
        // pages.push(new Page('/', "Pets"))
        pages.push(new Page("/agenda", "Agenda"))
        if (this.props.pages) {
            pages.push(...this.props.pages)
        }
        return pages
    }

    render() {
        return (
            <nav className={styles.navbar}>
                <div className={styles.container}>
                    <div className={styles.rightSideContainer}>
                        <Link href="/" className={styles.logoContainer}>
                            <Image src="/logo.png" width={33} height={35} className={styles.logo} alt="Pet Journal Logo" />
                            <span className={styles.logoTitle}>Pet Journal</span>
                        </Link>
                    </div>
                    <div className={styles.leftSideContainer}>
                        <div className={styles.pagesContainer}>
                            {this.pages.map((page) => (
                                <Link href={page.href} key={page.title} className={styles.pageContainer}>
                                    <span
                                        className={`${styles.pageLink} ${
                                            this.props.router.pathname === page.href ? styles.pageSelected : ""
                                        }`}
                                    >
                                        {page.title}
                                    </span>
                                </Link>
                            ))}
                        </div>
                        {!this.props.hideAllPages && (
                            <div className={"flex flex-row gap-2 items-center"}>
                                <Avatar
                                    icon={<UserIcon className={"flex p-1"} />}
                                    onCLick={() => this.props.router.push("/account")}
                                    className={"hover:bg-indigo-200 hover:text-indigo-700 h-[30px] w-[30px]"}
                                />

                                <ArrowLeftOnRectangleIcon
                                    className={
                                        "h-10 w-10 text-slate-300 p-2 rounded-full hover:bg-gray-600"
                                    }
                                    onClick={() => this.logout(() => this.props.router.replace("/auth/login"))}
                                />
                            </div>
                        )}
                        {this.props.buttons && <div className={styles.buttonsContainer}>{this.props.buttons}</div>}
                    </div>
                </div>
            </nav>
        )
    }
}

export default withRouter(NavBar)
