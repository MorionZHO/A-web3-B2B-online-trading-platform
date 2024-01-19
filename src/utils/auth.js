import {TOKEN_TIME_LIMIT} from "./constant";
import storage from "./storage";

export const setTokenTime = () => {
    storage.setItem("tokenTime", Date.now());
}

export const getTokenTIme = () => {
    return storage.getItem("tokenTime");
}

export const diffTokenTime = () => {
    const currentTime = Date.now();
    const tokenTime = getTokenTIme();
    return currentTime - tokenTime > TOKEN_TIME_LIMIT;
}

export const logout = () => {
    storage.clearItem("authToken");
    storage.clearItem("tokenTime");
    window.location.href='/login'
}