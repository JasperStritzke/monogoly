import {defineStore} from "pinia";
import axios from "axios";
import useAlertStore from "./alert.store";
import {Currency} from "../util/currency";

export interface GameStore {
    gameId: string,
    players: Array<Player>,
    data: GameData,
}

export interface GameData {
    isLocked: boolean,
    currency: Currency,
    initialBalance: number,
    amountGo: number,
}

export interface Player {
    sendMessage(): boolean
}

type state = {
    player?: Player
    game?: GameStore
}


const useGameStore = defineStore("game", {
    state: () => ({
        player: undefined,
        game: undefined
    } as state),
    actions: {
        async joinGame(name: string, gameId: string, password: string): Promise<GameStore | undefined> {
            return axios.post("/game/join", {name, gameId, password})
                .then(({data: {game, you}}: { data: { you: Player, game: GameStore } }) => {
                    this.player = you
                    this.game = game

                    useAlertStore().dispatchSuccess(`Joined game #${game.gameId}.`)
                    return game
                })
                .catch(error => {
                    useAlertStore().dispatchError((error.response?.data as any).message)
                    return undefined
                })
        },
        async createGame(
            {
                name,
                currency,
                amountGo,
                initialBalance,
                password
            }: { name: string, currency: Currency, amountGo: number, initialBalance: number, password: string }
        ) {
            return axios.post("/game/create", {currency: currency.code, amountGo, initialBalance, password})
                .then(({data: {gameId, password}}: { data: { gameId: string, password: string } }) => {
                    useAlertStore().dispatchSuccess(`Successfully created game ${gameId}.`)

                    this.joinGame(name, gameId, password).catch()
                })
                .catch(error => useAlertStore().dispatchError((error.response?.data as any).message))
        }
    },
    getters: {
        isInGame(state): boolean {
            return !!state.game
        }
    }
})

export default useGameStore