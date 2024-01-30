/**
 * storage 封装
 */


export default {
    setItem(key, val) {
        let storage = this.getStorage();
        storage[key] = val;
        window.localStorage.setItem('web3trade', JSON.stringify(storage))
    },
    getItem(key) {
        return this.getStorage()[key];
    },
    getStorage() {
        return JSON.parse(window.localStorage.getItem('web3trade') || "{}");
    },
    clearItem(key) {
        let storage = this.getStorage();
        delete storage[key];
        window.localStorage.setItem('web3trade', JSON.stringify(storage));
    },
    clearAll() {
        window.localStorage.setItem('web3trade', "{}");
    }
}