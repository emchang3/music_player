const { app } = require('electron');
console.log(app);

const { createMainWindow } = require('./app/windows');

app.setName('Music Player');

app.on('ready', createMainWindow);

app.on('window-all-closed', () => app.quit());