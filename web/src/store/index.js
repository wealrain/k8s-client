const DEFAULT_CACHE_TIME = 1000 * 60 * 60 * 24 * 7  // 7天



class Storage {

    constructor() {
        this.storage = localStorage
    }
    
    getKey(key) {
        return `__${key}__`.toUpperCase()
    }

    setItem(key, value,expire) {
        const data = JSON.stringify({
            value,
            expire: new Date().getTime() + (expire || DEFAULT_CACHE_TIME),
        })
        console.log('set item', data)
        this.storage.setItem(this.getKey(key), data)
    }
    
    getItem(key) {
        const data = this.storage.getItem(this.getKey(key))
        if (!data) {
            return null
        }
        console.log('get item', data)
        const { value, expire } = JSON.parse(data)
        if (new Date().getTime() > expire) {
            console.log('item expired')
            this.removeItem(key)
            return null
        }
        return value
    }
    
    removeItem(key) {
        this.storage.removeItem(this.getKey(key))
    }
    
    clear() {
       this.storage.clear()
    }
    
}

const storage = new Storage()


export default storage
