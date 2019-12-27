const CauCalendar = require('./cauacal'),
      initDatabase = require('./database'),
      config = require('./config.json');

(async () => {
    initDatabase(config.database);
    const cauCal = new CauCalendar();
    const yearFrom = 2004, yearTo = (new Date()).getFullYear() + 1;
    let schedules = [];
    for (let year = yearFrom; year <= yearTo; year++) {
        schedules = schedules.concat(await cauCal.getSchedules(year));
    }
    schedules = schedules.map(i => {
        i.start = new Date(i.start[0], i.start[1], i.start[2]);
        i.end = new Date(i.end[0], i.end[1], i.end[2]);
        return i;
    });
    await sequelize.transaction(async t => {
        await Event.destroy({
            where: {},
            truncate: true
        }, {transaction: t});
        await Event.bulkCreate(schedules, {transaction: t});
    });
})();