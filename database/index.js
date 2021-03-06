const Sequelize = require('sequelize');

module.exports = async connUrl => {
    global.sequelize = new Sequelize(connUrl);
    global.Event = sequelize.define('event', {
        title: {
            type: Sequelize.STRING,
            allowNull: false
        },
        start: {
            type: Sequelize.DATE,
            allowNull: false
        },
        end: {
            type: Sequelize.DATE,
            allowNull: false
        },
        uid: {
            type: Sequelize.STRING,
            allowNull: false
        }
    }, {
        timestamps: false,
        paranoid: false
    });

    await sequelize.sync();
};