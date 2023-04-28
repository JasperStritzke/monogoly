import {defineStore} from "pinia";

export type Alert = {
    id?: number
    type: "success" | "warning" | "error" | "info"
    message: string,
}

let incrementalAlertId = 0
const useAlertStore = defineStore('alerts', {
    state: () => ({
        alerts: []
    } as { alerts: Array<Alert> }),
    actions: {
        dispatchAlert(alert: Alert) {
            alert.id ||= incrementalAlertId++
            this.alerts.push(alert)
        },
        dispatchError(message: string) {
            this.dispatchAlert({
                type: "error",
                message,
            })
        },
        dispatchSuccess(message: string) {
            this.dispatchAlert({
                type: "success",
                message,
            })
        },
        dispatchWarning(message: string) {
            this.dispatchAlert({
                type: "warning",
                message,
            })
        },
        dispatchInfo(message: string) {
            this.dispatchAlert({
                type: "info",
                message,
            })
        },
        removeAlert(alert: Alert) {
            this.alerts = this.alerts.filter(a => a !== alert)
        }
    }
})

export default useAlertStore