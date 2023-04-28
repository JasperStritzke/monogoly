import axios from "axios";

class ApiService {

    private static _cachedCurrencies?: Array<Currency>

    static async getCurrencies(): Promise<Array<Currency>> {
        return this._cachedCurrencies || await axios.get("/currencies")
            .then(({data}) => {
                console.log("Fetching...")
                const currencies = []

                for (let currency in data) {
                    currencies.push({
                        code: currency,
                        symbol: data[currency]
                    })
                }

                this._cachedCurrencies = currencies

                return currencies
            })
    }
}

export type Currency = {
    code: string,
    symbol: string
}

export default ApiService
