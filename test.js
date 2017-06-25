
Array.prototype.filter = function (fn) {
    let res = []
    for (let i = 0; i < this.length; i++) {
        if (fn(this[i], i, this)) {
            res.push(arr[i])
        }
    }
    return res
}