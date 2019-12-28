'use strict';

const Koa = require('koa'),
      fs = require('fs'),
      ics = require('ics'),
      path = require('path'),
      Router = require('koa-router'),
      koaStatic = require('koa-static'),
      Op = (require('sequelize')).Op,
      initDatabase = require('./database');
      //webdav = require('webdav-server').v2;

const config = require('./config.json'),
      app = new Koa(),
      //webdavServer = new webdav.WebDAVServer(),
      router = new Router();

const dateToICSDate = (dateObj) => {
    return [dateObj.getFullYear(), dateObj.getMonth() + 1, dateObj.getDate()];
}

initDatabase(config.database);

router.get('/cau.ics', async (ctx, next) => {
    ctx.set('Content-Type', 'text/calendar; charset=utf-8');
    let {from, to} = ctx.request.query;
    const fromDefault = 2004, toDefault = (new Date()).getFullYear() + 1;
    // set from, to
    if (from)
        from = Number(from);
    else
        from = fromDefault;
    if (to)
        to = Number(to)
    else
        to = toDefault;

    // swap if required
    if (from > to) {
        let tmp = from;
        from = to;
        to = tmp;
    }

    // validate
    if (isNaN(from) || !isFinite(from) || from < 2004)
        from = fromDefault;
    if (isNaN(to) || !isFinite(to) || to > (new Date()).getFullYear())
        to = toDefault; 
    
    // query database
    let events = await Event.findAll({
        attributes: ['title', 'start', 'end', 'uid'],
        where: {
            start: {
                [Op.gte]: new Date(from, 1, 1)
            },
            end: {
                [Op.lte]: new Date(to, 12, 31)
            }
        }
    });
    events = events.map(i => {
        return {
            start: dateToICSDate(i.start),
            end: dateToICSDate(i.end),
            uid: i.uid,
            title: i.title
        };
    })

    // create ics
    let {error: icsError, value: icsResult} = ics.createEvents(events);
    if (icsError)
        throw icsError;
        console.log(icsResult);

    // return
    let icsExtPos = icsResult.indexOf('\n', icsResult.indexOf('\n') + 1);
    icsResult = icsResult.substring(0, icsExtPos) + 'X-WR-CALNAME:중앙대학교 학사일정\nX-WR-CALDESC:caucalendar.online에서 제공하는 중앙대학교 학사일정' + icsResult.substring(icsExtPos)
    ctx.body = icsResult;
    next();
});
//webdavServer.setFileSystem('/', (success) => webdavServer.start())
app.use(koaStatic(path.join(__dirname, 'static')));
app.use(router.routes());
app.use(router.allowedMethods());
app.listen(config.port);
//app.use(webdav.extensions.express('/webdav', webdavServer));