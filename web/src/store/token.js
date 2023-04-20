import storage from "."

const USER_TOKEN = 'userToken'

export const userToken = () => storage.getItem(USER_TOKEN)
export const setUserToken = (token) => storage.setItem(USER_TOKEN, token, 1000 * 60 * 60 * 24 )
export const removeUserToken = () => storage.removeItem(USER_TOKEN)