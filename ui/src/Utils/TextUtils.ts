export default class TextUtils {
    public static isEmpty(text?: string | null) {
        return !text || text.trim() === ""
    }

    public static isNotEmpty(text?: string | null) {
        return !TextUtils.isEmpty(text)
    }

    public static valueOrEmpty(text: string | null | undefined, defaultV: string): string {
        return TextUtils.isEmpty(text) ? defaultV : text!
    }

    public static isEmailValid(email: string): boolean {
        if (!email.trim()) {
            return false
        }

        // Regular expression pattern for email validation
        const pattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/

        // Match the email against the regular expression
        return pattern.test(email)
    }

    public static PasswordError(password: string): string {
        let error = ""
        if (password.length < 8) {
            error = "Password should be at least 8 characters long"
        }

        let done = /[a-z]+/.test(password)
        if (!done) {
            error = "Password should contain at least one lowercase character"
        }

        done = /[A-Z]+/.test(password)
        if (!done) {
            error = "Password should contain at least one uppercase character"
        }

        done = /[0-9]+/.test(password)
        if (!done) {
            error = "Password should contain at least one digit"
        }

        done = /[!@#$%^&*.?-]+/.test(password)
        if (!done) {
            error = "Password should contain at least one special character"
        }

        return error
    }
}
