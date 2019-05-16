const axios = require('axios'),
      apiUrl = 'https://www.cau.ac.kr/ajax/FR_SCH_SVC/ScheduleListData.do',
      ics = require('ics'),
      querystring = require('querystring');

class CauCalendar {
    async getSchedules(year) {
        const apiResponse = await axios.post(apiUrl, querystring.stringify({
            'SCH_SITE_NO': 2,
            'SCH_YEAR': year
        }));

        return apiResponse.data.data.map(i => {
            return {
                start: [Number(i.START_Y), Number(i.START_M), Number(i.START_D)],
                end: [Number(i.END_Y), Number(i.END_M), Number(i.END_D)],
                title: i.SUBJECT
            }
        });
    }
    async createIcs(yearFrom, yearTo) {
        let events = [];
        for(var year = yearFrom; year <= yearTo; year++)
            events = events.concat(await this.getSchedules(year));
        
        let {error, value} = ics.createEvents(events);
        if (error)
            throw error;
        return value;
    }
}

module.exports = CauCalendar;