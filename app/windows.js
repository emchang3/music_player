const { BrowserWindow } = require('electron');
const url = require('url');
const path = require('path');

const createMainWindow = () => {
  return new Promise(function(resolve, reject) {
    let win = new BrowserWindow({ width: 500, height: 500, titleBarStyle: 'hidden' });
    // win.setResizable(false);

    win.loadURL(url.format({
      pathname: path.join(__dirname, '..', 'views', 'index.html'),
      protocol: 'file:',
      slashes: true
    }));

    win.on('closed', () => {
      console.log('--- Closed: Main.');
      win = null;
    });

    win.webContents.on('did-finish-load', () => {
      console.log('--- Loaded: Main.');
      resolve(win);
    });
  });
}

module.exports = { createMainWindow: createMainWindow }