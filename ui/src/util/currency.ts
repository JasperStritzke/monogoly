export type Currency = {
    symbol: string
    code: string
}
const currencies: Array<Currency> = [
    {symbol: "€", code: "EUR"},
    {symbol: "$", code: "USD"},
    {symbol: "$", code: "CAD"},
    {symbol: "£", code: "GBP"},
    {symbol: "kr", code: "NOK"},
    {symbol: "kr", code: "DKK"},
    {symbol: "¥", code: "YEN"},
    {symbol: "₹", code: "INR"},
    {symbol: "₿", code: "BTC"},
    {symbol: "&#9829;", code: "HRT"},
]

export default currencies