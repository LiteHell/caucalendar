module.exports = () => {
    let now = new Date();
    return {
        fromDefault: 2004,
        toDefault: now.getMonth() == 11 ? now.getFullYear() + 1 : now.getFullYear()
    };
}