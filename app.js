'use strict';

const Koa = require('koa'),
      fs = require('fs'),
      path = require('path'),
      Router = require('koa-router'),
      koaStatic = require('koa-static'),
      CauCalendar = require('./cauacal');
      //webdav = require('webdav-server').v2;

const config = require('./config.json'),
      app = new Koa(),
      //webdavServer = new webdav.WebDAVServer(),
      router = new Router(),
      cauCal = new CauCalendar();

router.get('/cau.ics', async (ctx, next) => {
    ctx.set('Content-Type', 'text/calendar; charset=utf-8');
    let {from, to} = ctx.request.query;
    const fromDefault = 2004, toDefault = (new Date()).getFullYear();
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
    
    // return
    let ics = await cauCal.createIcs(from, to),
        icsExtPos = ics.indexOf('\n', ics.indexOf('\n') + 1);
    ics = ics.substring(0, icsExtPos) + 'X-WR-CALNAME:중앙대학교 학사일정\nX-WR-CALDESC:caucalendar.online에서 제공하는 중앙대학교 학사일정' + ics.substring(icsExtPos)
    ctx.body = ics;
    next();
});
//webdavServer.setFileSystem('/', (success) => webdavServer.start())
app.use(koaStatic(path.join(__dirname, 'static')));
app.use(router.routes());
app.use(router.allowedMethods());
app.listen(config.port);
//app.use(webdav.extensions.express('/webdav', webdavServer));