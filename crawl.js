const CauCalendar = require('./cauacal');

let makeKSTString = (dateNumbers) => {
    return `${dateNumbers[0]}-${dateNumbers[1].toString().padStart(2, '0')}-${dateNumbers[2].toString().padStart(2, '0')}T12:00:00+09:00`;
}

module.exports = async () => {
    const cauCal = new CauCalendar();
    const {fromDefault: yearFrom, toDefault: yearTo} = require('./yearDefaults')();
    let schedules = [];
    for (let year = yearFrom; year <= yearTo; year++) {
        schedules = schedules.concat(await cauCal.getSchedules(year));
    }
    schedules = schedules.map(i => {
        i.start = new Date(makeKSTString(i.start));
        i.end = new Date(makeKSTString(i.end));
        return i;
    });
    await sequelize.transaction(async t => {
        await Event.destroy({
            where: {},
            truncate: true
        }, {transaction: t});
        await Event.bulkCreate(schedules, {transaction: t});
    });
};